package usecase

import (
	"regexp"
	"strconv"

	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/internal/repository"
	"github.com/kmdavidds/mager-spot-api/model"
	"github.com/kmdavidds/mager-spot-api/pkg/mt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type IInvoiceUsecase interface {
	Purchase(invoice entity.Invoice) (string, error)
	Verify(notificationPayload map[string]interface{})
}

type InvoiceUsecase struct {
	ir repository.IInvoiceRepository
	r  repository.Repository
}

func NewInvoiceUsecase(invoiceRepository repository.IInvoiceRepository, repository repository.Repository) IInvoiceUsecase {
	return &InvoiceUsecase{
		ir: invoiceRepository,
		r:  repository,
	}
}

func convertToINT64(str string) int64 {
	reg, _ := regexp.Compile("[^0-9]+")
	cleanStr := reg.ReplaceAllString(str, "")
	cleanInt, _ := strconv.Atoi(cleanStr)
	return int64(cleanInt)
}

func (iu *InvoiceUsecase) Purchase(invoice entity.Invoice) (string, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: invoice.ID.String(),
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID: invoice.PostID.String(),
			},
		},
	}

	switch invoice.Category {
	case "apartment-post":
		post, err := iu.r.ApartmentPostRepository.GetApartmentPost(model.ApartmentPostKey{ID: invoice.PostID})
		if err != nil {
			return "", err
		}
		req.TransactionDetails.GrossAmt = convertToINT64(post.Price) * convertToINT64(invoice.Amount)
		(*req.Items)[0].Name = post.Title
		(*req.Items)[0].Price = convertToINT64(post.Price)
		(*req.Items)[0].Qty = int32(convertToINT64(invoice.Amount))
		(*req.Items)[0].Category = invoice.Category
		(*req.Items)[0].MerchantName = post.User.Username
		invoice.SellerID = post.User.ID
	case "food-post":
		post, err := iu.r.FoodPostRepository.GetFoodPost(model.FoodPostKey{ID: invoice.PostID})
		if err != nil {
			return "", err
		}
		req.TransactionDetails.GrossAmt = convertToINT64(post.Price) * convertToINT64(invoice.Amount)
		(*req.Items)[0].Name = post.Title
		(*req.Items)[0].Price = convertToINT64(post.Price)
		(*req.Items)[0].Qty = int32(convertToINT64(invoice.Amount))
		(*req.Items)[0].Category = invoice.Category
		(*req.Items)[0].MerchantName = post.User.Username
		invoice.SellerID = post.User.ID
	case "product-post":
		post, err := iu.r.ProductPostRepository.GetProductPost(model.ProductPostKey{ID: invoice.PostID})
		if err != nil {
			return "", err
		}
		req.TransactionDetails.GrossAmt = convertToINT64(post.Price) * convertToINT64(invoice.Amount)
		(*req.Items)[0].Name = post.Title
		(*req.Items)[0].Price = convertToINT64(post.Price)
		(*req.Items)[0].Qty = int32(convertToINT64(invoice.Amount))
		(*req.Items)[0].Category = invoice.Category
		(*req.Items)[0].MerchantName = post.User.Username
		invoice.SellerID = post.User.ID
	case "shuttle-post":
		post, err := iu.r.ShuttlePostRepository.GetShuttlePost(model.ShuttlePostKey{ID: invoice.PostID})
		if err != nil {
			return "", err
		}
		req.TransactionDetails.GrossAmt = convertToINT64(post.Price) * convertToINT64(invoice.Amount)
		(*req.Items)[0].Name = post.Title
		(*req.Items)[0].Price = convertToINT64(post.Price)
		(*req.Items)[0].Qty = int32(convertToINT64(invoice.Amount))
		(*req.Items)[0].Category = invoice.Category
		(*req.Items)[0].MerchantName = post.User.Username
		invoice.SellerID = post.User.ID
	}

	pajak := midtrans.ItemDetails{
		ID:    "pajak_pembeli",
		Price: int64(float64(req.TransactionDetails.GrossAmt) * float64(0.05)),
		Qty:   1,
		Name:  "Pajak Sebesar 5%",
	}
	*req.Items = append(*req.Items, pajak)

	invoice.OriginalPrice = req.TransactionDetails.GrossAmt

	req.TransactionDetails.GrossAmt = int64(float64(req.TransactionDetails.GrossAmt) * float64(1.05))

	paymentLink, midtransErr := snap.CreateTransactionUrl(req)
	if midtransErr != nil {
		return "", midtransErr
	}

	invoice.PaymentLink = paymentLink

	err := iu.ir.CreateInvoice(invoice)
	if err != nil {
		return "", err
	}

	return paymentLink, nil
}

func (iu *InvoiceUsecase) Verify(notificationPayload map[string]interface{}) {
	transactionStatus := notificationPayload["transaction_status"]
	fraudStatus := notificationPayload["fraud_status"]
	orderID := notificationPayload["order_id"]

	switch transactionStatus {
	case mt.StatusCapture:
		switch fraudStatus {
		case mt.StatusChallenge:
			tx, err := iu.ir.UpdateInvoiceStatus("challenge", orderID.(string))
			if err != nil {
				tx.Rollback()
				return
			}
			tx.Commit()
		case mt.StatusAccept:
			tx, err := iu.ir.UpdateInvoiceStatus("success", orderID.(string))
			if err != nil {
				tx.Rollback()
				return 
			}
			invoice, err := iu.ir.GetInvoice(model.InvoiceParam{ID: orderID.(string)})
			if err != nil {
				tx.Rollback()
				return 
			}
			tx, err = iu.ir.AddBalance(tx, invoice)
			if err != nil {
				tx.Rollback()
				return 
			}
			tx.Commit()
		}
	case mt.StatusSettlement:
		tx, err := iu.ir.UpdateInvoiceStatus("success", orderID.(string))
			if err != nil {
				tx.Rollback()
				return
			}
			invoice, err := iu.ir.GetInvoice(model.InvoiceParam{ID: orderID.(string)})
			if err != nil {
				tx.Rollback()
				return 
			}
			tx, err = iu.ir.AddBalance(tx, invoice)
			if err != nil {
				tx.Rollback()
				return 
			}
			tx.Commit()
	case mt.StatusDeny:
	case mt.StatusCancel, mt.StatusExpire:
		tx, err := iu.ir.UpdateInvoiceStatus("failure", orderID.(string))
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	case mt.StatusPending:
		tx, err := iu.ir.UpdateInvoiceStatus("pending", orderID.(string))
		if err != nil {
			tx.Rollback()
			return 
		}
		tx.Commit()
	}
}

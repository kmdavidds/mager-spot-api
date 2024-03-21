package usecase

import (
	"regexp"
	"strconv"

	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/internal/repository"
	"github.com/kmdavidds/mager-spot-api/model"
	mt "github.com/kmdavidds/mager-spot-api/pkg/midtrans_extras"
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
	}

	switch invoice.Category {
	case "apartment-post":
		post, err := iu.r.ApartmentPostRepository.GetApartmentPost(model.ApartmentPostKey{ID: invoice.PostID})
		if err != nil {
			return "", err
		}
		req.TransactionDetails.GrossAmt = convertToINT64(post.Price)
	case "food-post":
		post, err := iu.r.FoodPostRepository.GetFoodPost(model.FoodPostKey{ID: invoice.PostID})
		if err != nil {
			return "", err
		}
		req.TransactionDetails.GrossAmt = convertToINT64(post.Price)
	case "product-post":
		post, err := iu.r.ProductPostRepository.GetProductPost(model.ProductPostKey{ID: invoice.PostID})
		if err != nil {
			return "", err
		}
		req.TransactionDetails.GrossAmt = convertToINT64(post.Price)
	case "shuttle-post":
		post, err := iu.r.ShuttlePostRepository.GetShuttlePost(model.ShuttlePostKey{ID: invoice.PostID})
		if err != nil {
			return "", err
		}
		req.TransactionDetails.GrossAmt = convertToINT64(post.Price)
	}

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
			iu.ir.UpdateInvoiceStatus("challenge", orderID.(string))
		case mt.StatusAccept:
			iu.ir.UpdateInvoiceStatus("success", orderID.(string))
		}
	case mt.StatusSettlement:
		iu.ir.UpdateInvoiceStatus("success", orderID.(string))
	case mt.StatusDeny:
	case mt.StatusCancel, mt.StatusExpire:
		iu.ir.UpdateInvoiceStatus("failure", orderID.(string))
	case mt.StatusPending:
		iu.ir.UpdateInvoiceStatus("pending", orderID.(string))
	}
}

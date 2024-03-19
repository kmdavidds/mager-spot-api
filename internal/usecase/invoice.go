package usecase

import (
	"strconv"

	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/internal/repository"
	"github.com/kmdavidds/mager-spot-api/model"
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
		parsedPrice, err := strconv.Atoi(post.Price)
		if err != nil {
			return "", err
		}
		req.TransactionDetails.GrossAmt = int64(parsedPrice)
	case "food-post":
		post, err := iu.r.FoodPostRepository.GetFoodPost(model.FoodPostKey{ID: invoice.PostID})
		if err != nil {
			return "", err
		}
		parsedPrice, err := strconv.Atoi(post.Price)
		if err != nil {
			return "", err
		}
		req.TransactionDetails.GrossAmt = int64(parsedPrice)
	case "product-post":
		post, err := iu.r.ProductPostRepository.GetProductPost(model.ProductPostKey{ID: invoice.PostID})
		if err != nil {
			return "", err
		}
		parsedPrice, err := strconv.Atoi(post.Price)
		if err != nil {
			return "", err
		}
		req.TransactionDetails.GrossAmt = int64(parsedPrice)
	case "shuttle-post":
		post, err := iu.r.ShuttlePostRepository.GetShuttlePost(model.ShuttlePostKey{ID: invoice.PostID})
		if err != nil {
			return "", err
		}
		parsedPrice, err := strconv.Atoi(post.Price)
		if err != nil {
			return "", err
		}
		req.TransactionDetails.GrossAmt = int64(parsedPrice)
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

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			iu.ir.UpdateInvoiceStatus("challenge", orderID.(string))
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
		} else if fraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			iu.ir.UpdateInvoiceStatus("success", orderID.(string))
		}
	} else if transactionStatus == "settlement" {
		// TODO set transaction status on your databaase to 'success'
		iu.ir.UpdateInvoiceStatus("success", orderID.(string))
	} else if transactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		// TODO set transaction status on your databaase to 'failure'
		iu.ir.UpdateInvoiceStatus("failure", orderID.(string))
	} else if transactionStatus == "pending" {
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		iu.ir.UpdateInvoiceStatus("pending", orderID.(string))
	}
}

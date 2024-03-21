package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
)

func (r *Rest) Purchase(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "failed get login user",
		})
		return
	}

	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse post id",
			"error":   err,
		})
		return
	}

	category := ctx.Param("category")
	amount := ctx.Query("amount")

	invoice := entity.Invoice{
		ID:       uuid.New(),
		UserID:   user.(entity.User).ID,
		PostID:   parsedId,
		Category: category,
		Status:   "pending",
		Amount: amount,
	}

	paymentLink, err := r.usecase.InvoiceUsecase.Purchase(invoice)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to purchase",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"paymentLink": paymentLink,
	})
}

func (r *Rest) Verify(ctx *gin.Context) {
	var notificationPayload map[string]interface{}

	err := ctx.ShouldBind(&notificationPayload)
	if err != nil {
		return
	}

	_, exists := notificationPayload["order_id"].(string)
	if !exists {
		return
	}

	r.usecase.InvoiceUsecase.Verify(notificationPayload)
}

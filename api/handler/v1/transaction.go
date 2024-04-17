package v1

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/new-task/api/handler/v1/http"
	"github.com/new-task/api/models"
)

// create 		Transaction
// @ID 			create transaction
// @Router		/transaction [POST]
// @Summary		create item
// @Tags        Transactions
// @Description	Here transaction can create
// @Accept 		json
// @Produce		json
// @Param 		body body models.AddTransaction true "TransactionBody"
// @Success 	201 {object} http.Response{data=models.Transaction} "Transaction data"
// @Response 	400 {object} http.Response{data=string} "Bad Request"
// @Failure 	500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) CreateT(c *gin.Context) {
	var body models.AddTransaction
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Task().CreateT(&body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.Created, res)
}

// create 		Transaction
// @ID 			get transaction
// @Router		/transaction [GET]
// @Summary		get item
// @Tags        Transactions
// @Description	Here transaction can get
// @Accept 		json
// @Produce		json
// @Param 		search query string false "SEARCH_KEY"
// @Param 		limit query int false "LIMIT"
// @Param 		page query int false "PAGE_NUMBER"
// @Success 	200 {object} http.Response{data=models.List} "TransactionViews"
// @Response 	400 {object} http.Response{data=string} "Bad Request"
// @Failure 	500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetT(c *gin.Context) {
	limit, page := h.CheckLimitOffset(c.Query("limit"), c.Query("page"))
	res, err := h.Storage.Task().GetT(&models.TrangetReq{
		SearchKey: c.Query("search"),
		Limit:     limit,
		Page:      page,
	})
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, res)
}

// update Transaction
// @ID 			update transaction
// @Router		/transaction [PUT]
// @Summary		UPDATE Transaction
// @Tags        Transactions
// @Description	Here Transaction can update
// @Accept 		json
// @Produce		json
// @Param 		body body models.TransactionUpdate true "TransactionBody"
// @Success 	200 {object} http.Response{data=models.TransactionUpdate} "Item data"
// @Response 	400 {object} http.Response{data=string} "Bad Request"
// @Failure 	500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) UpdateT(c *gin.Context) {
	var body models.TransactionUpdate
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Task().UpdateT(&body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, res)
}

// delete Transaction
// @ID 			delete transaction
// @Router		/transaction [DELETE]
// @Summary		Delete Transaction
// @Tags        Transactions
// @Description	Here Transaction can delete
// @Accept 		json
// @Produce		json
// @Param 		id query int true "TransactionID"
// @Success 	200 {object} http.Response{data=string} "Transaction data"
// @Response 	400 {object} http.Response{data=string} "Bad Request"
// @Failure 	500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) DeleteT(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	err = h.Storage.Task().DeleteT(id)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, "Deleted Successfully")
}

func (h *handlerV1) CheckLimitOffset(limitStr, pageStr string) (int, int) {
	var limit, page int
	var err error
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Println("error parse limit: ", err.Error())
		}
	} else {
		limit = 10
	}
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			log.Println("error parse page: ", err.Error())
		}
	} else {
		page = 1
	}

	return limit, page
}

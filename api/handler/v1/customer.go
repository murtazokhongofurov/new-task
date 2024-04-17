package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/new-task/api/handler/v1/http"
	"github.com/new-task/api/models"
)

// register Customer
// @ID 			register customer
// @Router		/customer [POST]
// @Summary		register customer
// @Tags        Customer
// @Description	Here customer can register
// @Accept 		json
// @Produce		json
// @Param 		body body models.Cregister true "CustomerBody"
// @Success 	201 {object} http.Response{data=models.Customer} "Customer data"
// @Response 	400 {object} http.Response{data=string} "Bad Request"
// @Failure 	500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) CreateC(c *gin.Context) {
	var body models.Cregister
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Task().CreateC(&body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.Created, res)
}

// Get Customer
// @ID 			get customer
// @Router		/customer [GET]
// @Summary		get customer
// @Tags        Customer
// @Description	Here customer can get
// @Accept 		json
// @Produce		json
// @Param 		id  query int true "CustomerID"
// @Success 	200 {object} http.Response{data=models.Customer} "Customer data"
// @Response 	400 {object} http.Response{data=string} "Bad Request"
// @Failure 	500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetC(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Task().GetC(id)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, res)
}

// update Customer
// @ID 			update customer
// @Router		/customer [PUT]
// @Summary		UPDATE customer
// @Tags        Customer
// @Description	Here customer can update
// @Accept 		json
// @Produce		json
// @Param 		body body models.Customer true "CustomerBody"
// @Success 	200 {object} http.Response{data=models.Customer} "Customer data"
// @Response 	400 {object} http.Response{data=string} "Bad Request"
// @Failure 	500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) UpdateC(c *gin.Context) {
	var body models.Customer
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Task().UpdateC(&body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, res)
}

// delete Customer
// @ID 			delete customer
// @Router		/customer [DELETE]
// @Summary		register customer
// @Tags        Customer
// @Description	Here customer can register
// @Accept 		json
// @Produce		json
// @Param 		id query int true "CustomerID"
// @Success 	200 {object} http.Response{data=string} "Customer data"
// @Response 	400 {object} http.Response{data=string} "Bad Request"
// @Failure 	500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) DeleteC(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		h.handleResponse(c, http.BadRequest, err)
		return
	}
	err = h.Storage.Task().DeleteC(id)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, "Deleted successfully")
}

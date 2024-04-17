package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/new-task/api/handler/v1/http"
	"github.com/new-task/api/models"
)

// add Item
// @ID 			add item
// @Router		/item  [POST]
// @Summary		add items
// @Tags        Item
// @Description	Here item can add
// @Accept 		json
// @Produce		json
// @Param 		body body models.ItemCreate true "Add ItemBody"
// @Success 	201 {object} http.Response{data=models.Item} "Customer data"
// @Response 	400 {object} http.Response{data=string} "Bad Request"
// @Failure 	500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) CreateP(c *gin.Context) {
	var body models.ItemCreate
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Task().CreateP(&body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.Created, res)
}

// Get Item
// @ID 			get item
// @Router		/item [GET]
// @Summary		register item
// @Tags        Item
// @Description	Here item can get
// @Accept 		json
// @Produce		json
// @Param 		id  query int true "ItemID"
// @Success 	200 {object} http.Response{data=models.Item} "Item data"
// @Response 	400 {object} http.Response{data=string} "Bad Request"
// @Failure 	500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetP(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Task().GetP(id)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, res)
}

// update Item
// @ID 			update item
// @Router		/item [PUT]
// @Summary		UPDATE item
// @Tags        Item
// @Description	Here item can update
// @Accept 		json
// @Produce		json
// @Param 		body body models.ItemUpdate true "Update ItemBody"
// @Success 	200 {object} http.Response{data=models.Item} "Item data"
// @Response 	400 {object} http.Response{data=string} "Bad Request"
// @Failure 	500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) UpdateP(c *gin.Context) {
	var body models.ItemUpdate
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Task().UpdateP(&body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, res)
}

// delete item
// @ID 			delete item
// @Router		/item [DELETE]
// @Summary		register customer
// @Tags        Item
// @Description	Here item can register
// @Accept 		json
// @Produce		json
// @Param 		id query int true "ItemID"
// @Success 	200 {object} http.Response{data=string} "Customer data"
// @Response 	400 {object} http.Response{data=string} "Bad Request"
// @Failure 	500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) DeleteP(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		h.handleResponse(c, http.BadRequest, err)
		return
	}
	err = h.Storage.Task().DeleteP(id)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, "Deleted successfully")
}

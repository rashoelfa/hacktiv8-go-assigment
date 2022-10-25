package controllers

import (
	"net/http"
	"time"

	"assigment2/models"

	"github.com/gin-gonic/gin"
)

type OrderRequestBody struct {
	OrderedAt    time.Time     `json:"orderedAt"`
	CustomerName string        `json:"customerName"`
	Items        []ItemsDetail `json:"items"`
}

type ItemsDetail struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}

func (h handler) AddOrder(ctx *gin.Context) {
	body := OrderRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var orders models.Orders
	var items []models.Items

	orders.Customer_name = body.CustomerName
	orders.Ordered_at = time.Now()

	if result := h.DB.Create(&orders); result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	var lastID models.Orders
	h.DB.Last(&lastID)

	itemsCollection := []ItemsDetail{}

	for i := 0; i < len(body.Items); i++ {
		items = append(items, models.Items{
			Item_code:   body.Items[i].ItemCode,
			Description: body.Items[i].Description,
			Quantity:    body.Items[i].Quantity,
			Order_id:    lastID.Order_id,
		})

		itemsCollection = append(itemsCollection, ItemsDetail{
			ItemCode:    items[i].Item_code,
			Description: items[i].Description,
			Quantity:    items[i].Quantity,
		})
	}

	if result := h.DB.Create(&items); result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	result := OrderRequestBody{
		orders.Ordered_at,
		body.CustomerName,
		itemsCollection,
	}

	ctx.JSON(http.StatusCreated, &result)
}

func (h handler) GetOrders(ctx *gin.Context) {
	var orders []models.Orders

	if result := h.DB.Find(&orders); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &orders)
}

func (h handler) GetOrderById(ctx *gin.Context) {
	id := ctx.Param("id")

	var order models.Orders
	var items []models.Items
	var response OrderRequestBody
	var itemsCollection []ItemsDetail

	if result := h.DB.First(&order, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	result := h.DB.Model(&items).Select("items.*").Joins("join orders on orders.order_id = items.order_id").Where("items.order_id = ?", id).Scan(&items)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	for i := 0; i < int(result.RowsAffected); i++ {
		itemsCollection = append(itemsCollection, ItemsDetail{
			ItemCode:    items[i].Item_code,
			Description: items[i].Description,
			Quantity:    items[i].Quantity,
		})
	}

	response = OrderRequestBody{
		order.Ordered_at,
		order.Customer_name,
		itemsCollection,
	}

	ctx.JSON(http.StatusOK, &response)
}

func (h handler) UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	body := OrderRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var orders models.Orders
	var items []models.Items

	if result := h.DB.First(&orders, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	orders.Customer_name = body.CustomerName

	if result := h.DB.Save(&orders); result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	var lastID models.Orders
	h.DB.First(&lastID, id)

	itemsCollection := []ItemsDetail{}

	if res := h.DB.Where("order_id = ?", lastID.Order_id).Order("item_id").Find(&items); res.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, res.Error)
		return
	}

	for i := 0; i < len(body.Items); i++ {
		items[i].Item_code = body.Items[i].ItemCode
		items[i].Description = body.Items[i].Description
		items[i].Quantity = body.Items[i].Quantity

		itemsCollection = append(itemsCollection, ItemsDetail{
			ItemCode:    items[i].Item_code,
			Description: items[i].Description,
			Quantity:    items[i].Quantity,
		})

		if result := h.DB.Where("order_id = ?", lastID.Order_id).Updates(&items[i]); result.Error != nil {
			ctx.AbortWithError(http.StatusInternalServerError, result.Error)
			return
		}
	}

	result := OrderRequestBody{
		orders.Ordered_at,
		body.CustomerName,
		itemsCollection,
	}

	ctx.JSON(http.StatusCreated, &result)
}

func (h handler) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	var order models.Orders

	if result := h.DB.First(&order, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&order)

	ctx.Status(http.StatusOK)
}

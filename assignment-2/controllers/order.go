package controllers

import (
	"assignment-2/models"
	"assignment-2/resources"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// PostOrder godoc
// @Summary Create new order
// @Description Create new order
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {array} Order
// @Router /orders [post]
func PostOrder(ctx *gin.Context) {
	db := models.GetDb()
	order := models.Order{}

	if err := ctx.ShouldBind(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	newOrder := models.Order{
		CustomerName: order.CustomerName,
		OrderedAt:    order.OrderedAt,
		Items:        order.Items,
	}

	if err := db.Create(&newOrder).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	addedItems := make([]resources.Item, len(newOrder.Items))

	for i, _ := range newOrder.Items {
		addedItems[i] = resources.Item{
			Description: newOrder.Items[i].Description,
			ItemID:      newOrder.Items[i].ID,
			Quantity:    newOrder.Items[i].Quantity,
		}
	}

	addedOrder := resources.Order{
		CustomerName: newOrder.CustomerName,
		Items:        addedItems,
		ID:           newOrder.ID,
		OrderedAt:    newOrder.OrderedAt,
	}

	ctx.JSON(http.StatusOK, resources.Response{Message: "add order success", Data: addedOrder})
}

// GetOrders godoc
// @Summary Get details of all orders
// @Description Get details of all orders
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {array} Order
// @Router /orders [get]
func GetOrders(ctx *gin.Context) {
	db := models.GetDb()
	orders := []models.Order{}

	if err := ctx.ShouldBind(&orders); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	if err := db.Model(&models.Order{}).Preload("Items").Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	ordersData := make([]resources.Order, len(orders))

	for i, _ := range orders {
		itemsData := make([]resources.Item, len(orders[i].Items))

		for j, _ := range orders[i].Items {
			itemsData[j] = resources.Item{
				Description: orders[i].Items[j].Description,
				ItemID:      orders[i].Items[j].ID,
				Quantity:    orders[i].Items[j].Quantity,
			}

			ordersData[i] = resources.Order{
				CustomerName: orders[i].CustomerName,
				Items:        itemsData,
				ID:           orders[i].ID,
				OrderedAt:    orders[i].OrderedAt,
			}
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"data": ordersData})
}

// PutOrder godoc
// @Summary Update order
// @Description Update order
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {array} Order
// @Router /orders [put]
func PutOrder(ctx *gin.Context) {
	db := models.GetDb()
	order := models.Order{}
	item := models.Item{}

	if err := db.Where("order_id = ?", ctx.Param("orderId")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})

		return
	}

	db.Unscoped().Where("order_id = ?", order.ID).Delete(item)

	if err := db.Unscoped().Where("order_id = ?", order.ID).Delete(item).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	if err := ctx.ShouldBind(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	if err := db.Save(order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	updatedItem := make([]resources.Item, len(order.Items))

	for i, _ := range order.Items {
		updatedItem[i] = resources.Item{
			Description: order.Items[i].Description,
			ItemID:      order.Items[i].ID,
			Quantity:    order.Items[i].Quantity,
		}
	}

	updatedOrder := resources.Order{
		CustomerName: order.CustomerName,
		Items:        updatedItem,
		ID:           order.ID,
		OrderedAt:    order.OrderedAt,
	}

	ctx.JSON(http.StatusOK, resources.Response{Message: "update order success", Data: updatedOrder})
}

// DeleteOrder godoc
// @Summary Delete order
// @Description Delete order
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 string
// @Router /orders [delete]
func DeleteOrder(ctx *gin.Context) {
	db := models.GetDb()
	order := models.Order{}

	if err := db.Where("order_id = ?", ctx.Param("orderId")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})

		return
	}

	if err := db.Select(clause.Associations).Delete(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "delete order success"})
}

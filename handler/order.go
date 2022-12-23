package handlers

import (
	"article/db"
	"article/model"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateOrder(c echo.Context) error {
	order := new(model.Order)
	err := c.Bind(order)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid order format")
	}

	// get the user id from header
	userID := c.Request().Header.Get("UserID")
	uid, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse user ID")
	}

	// check if the user is active
	user, err := db.GetUser(uint64(uid))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if !user.Active {
		return c.JSON(http.StatusBadRequest, "Inactive user!")
	}

	// check if there is a promotion and wether it is already used or not
	promotion, err := db.GetPromotion(order.PromotionID)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	if promotion.Used {
		return c.JSON(http.StatusBadRequest, "Promotion code already used")

	}

	// check if the item is available
	for _, item := range order.Items {
		itm, err := db.GetItem(uint64(item.ID))
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		if !itm.Available {
			return c.JSON(http.StatusBadRequest, "Item not available!")
		}
	}

	order.UserID = uint64(uid)
	order.Status = model.OrderStatusReview

	order, err = db.CreateOrder(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create order")
	}

	return c.JSON(http.StatusOK, order)
}

func UpdateOrder(c echo.Context) error {
	order := new(model.Order)
	err := c.Bind(order)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid order format")
	}

	// get the user id from header
	userID := c.Request().Header.Get("UserID")
	uid, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse user ID")
	}

	// Check if order belongs to the user
	orderBelongs, err := orderBelongsToUser(uint64(uid), uint64(order.ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to parse user ID")
	}

	if !orderBelongs {
		return c.NoContent(http.StatusInternalServerError)
	}

	order, err = db.UpdateOrder(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update order")
	}

	return c.JSON(http.StatusOK, order)
}

func GetOrders(c echo.Context) error {
	// get the user id from header
	userID := c.Request().Header.Get("UserID")
	uid, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse user ID")
	}

	// check if the user is authorized to access orders
	user, err := db.GetUser(uint64(uid))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// if user is client then fetch the client's orders only
	if user.Type == model.UserTypeClient {

		orders, err := db.GetOrderByUser(uint64(user.ID))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(http.StatusNotFound, "No orders found")
			}
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, orders)
	}

	items, err := db.GetOrders(uint64(uid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch items")
	}

	return c.JSON(http.StatusOK, items)
}

func GetOrder(c echo.Context) error {
	id := c.Param("id")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse order id")
	}

	// get the user id from header
	userID := c.Request().Header.Get("UserID")
	uid, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse user ID")
	}

	// Check if order belongs to the user
	orderBelongs, err := orderBelongsToUser(uint64(uid), uint64(orderID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to parse user ID")
	}

	if !orderBelongs {
		return c.NoContent(http.StatusInternalServerError)
	}

	order, err := db.GetOrder(uint64(orderID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch order")
	}

	return c.JSON(http.StatusOK, order)
}

func DeleteOrder(c echo.Context) error {
	id := c.Param("id")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse order id")
	}

	// get the user id from header
	userID := c.Request().Header.Get("UserID")
	uid, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse user ID")
	}

	// Check if order belongs to the user
	orderBelongs, err := orderBelongsToUser(uint64(uid), uint64(orderID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to parse user ID")
	}

	if !orderBelongs {
		return c.NoContent(http.StatusInternalServerError)
	}

	err = db.DeleteOrder(uint64(orderID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete order")
	}

	return c.NoContent(http.StatusOK)
}

// check if order belongs to the users
func orderBelongsToUser(userID, orderID uint64) (bool, error) {
	// Check if order belongs to the user
	ord, err := db.GetOrder(orderID)
	if err != nil {
		return false, err
	}

	if ord.UserID == uint64(userID) {
		return true, err
	}

	return false, nil
}

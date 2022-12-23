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

func CreateItem(c echo.Context) error {
	item := new(model.Item)
	err := c.Bind(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid item format")
	}

	item, err = db.CreateItem(item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create item")
	}

	return c.JSON(http.StatusOK, item)
}

func UpdateItem(c echo.Context) error {
	item := new(model.Item)
	err := c.Bind(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid item format")
	}

	item, err = db.UpdateItem(item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update item")
	}

	return c.JSON(http.StatusOK, item)
}

func GetItems(c echo.Context) error {
	availability := c.QueryParam("available")
	name := c.QueryParam("name")
	items, err := db.GetItems(name, availability)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch items")
	}

	return c.JSON(http.StatusOK, items)
}

func GetItem(c echo.Context) error {
	id := c.Param("id")
	itemID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse item id")
	}

	item, err := db.GetItem(uint64(itemID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch item")
	}

	return c.JSON(http.StatusOK, item)
}

func DeleteItem(c echo.Context) error {
	id := c.Param("id")
	itemID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse item id")
	}

	// Check if the item belongs to an order
	_, err = db.GetItemOrders(uint64(itemID))

	if err == nil {
		return c.JSON(http.StatusBadRequest, "Item is associated with order")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.NoContent(http.StatusInternalServerError)
	}

	err = db.DeleteItem(uint64(itemID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete item")
	}

	return c.NoContent(http.StatusOK)
}

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

func CreatePromotion(c echo.Context) error {
	promotion := new(model.Promotion)
	err := c.Bind(promotion)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid promotion format")
	}

	promotion, err = db.CreatePromotion(promotion)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create promotion")
	}

	return c.JSON(http.StatusOK, promotion)
}

func UpdatePromotion(c echo.Context) error {
	promotion := new(model.Promotion)
	err := c.Bind(promotion)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid promotion format")
	}

	promotion, err = db.UpdatePromotion(promotion)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update promotion")
	}

	return c.JSON(http.StatusOK, promotion)
}

func GetPromotions(c echo.Context) error {
	items, err := db.GetPromotions()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch items")
	}

	return c.JSON(http.StatusOK, items)
}

func GetPromotion(c echo.Context) error {
	id := c.Param("id")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse promotion id")
	}

	promotion, err := db.GetPromotion(uint64(orderID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch promotion")
	}

	return c.JSON(http.StatusOK, promotion)
}

func DeletePromotion(c echo.Context) error {
	id := c.Param("id")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse promotion id")
	}

	_, err = db.GetOrderByPromotion(uint64(orderID))

	if err == nil {
		return c.JSON(http.StatusBadRequest, "Promotion is associated with order")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.NoContent(http.StatusInternalServerError)
	}

	err = db.DeletePromotion(uint64(orderID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete promotion")
	}

	return c.NoContent(http.StatusOK)
}

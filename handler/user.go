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

func CreateUser(c echo.Context) error {
	user := new(model.User)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user format")
	}

	user, err = db.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	user := new(model.User)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user format")
	}

	user, err = db.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update user")
	}

	return c.JSON(http.StatusOK, user)
}

func GetUsers(c echo.Context) error {
	items, err := db.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch items")
	}

	return c.JSON(http.StatusOK, items)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse user id")
	}

	// get the user id from header
	usrID := c.Request().Header.Get("UserID")
	uid, err := strconv.Atoi(usrID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse user ID")
	}

	// check if the user is authorized to fetch user
	user, err := db.GetUser(uint64(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch user")
	}

	if user.Type == model.UserTypeClient && user.ID != uint(uid) {
		return c.NoContent(http.StatusUnauthorized)
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse user id")
	}

	// get order by user id
	_, err = db.GetOrderByUser(uint64(userID))

	if err == nil {
		return c.JSON(http.StatusBadRequest, "User is associated with order")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.NoContent(http.StatusInternalServerError)
	}

	err = db.DeleteUser(uint64(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete user")
	}

	return c.NoContent(http.StatusOK)
}

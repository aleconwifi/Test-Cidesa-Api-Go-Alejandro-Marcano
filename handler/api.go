package handlers

import (
	"article/db"
	"article/model"
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func HandleRequest() {

	e := echo.New()
	api := e.Group("/api")
	items := api.Group("/items")
	promotions := api.Group("/promotions")
	orders := api.Group("/orders")
	users := api.Group("/users")

	// Items
	items.POST("", AuthMiddleware([]string{model.UserTypeAdmin}, CreateItem))
	items.GET("", AuthMiddleware([]string{model.UserTypeAdmin}, GetItems))
	items.GET("/:id", AuthMiddleware([]string{model.UserTypeAdmin}, GetItem))
	items.PATCH("/:id", AuthMiddleware([]string{model.UserTypeAdmin}, UpdateItem))
	items.DELETE("/:", AuthMiddleware([]string{model.UserTypeAdmin}, DeleteItem))

	// Promotions
	promotions.POST("", AuthMiddleware([]string{model.UserTypeAdmin}, CreatePromotion))
	promotions.GET("", AuthMiddleware([]string{model.UserTypeAdmin}, GetPromotions))
	promotions.GET("/:id", AuthMiddleware([]string{model.UserTypeAdmin}, GetPromotion))
	promotions.PATCH("/:id", AuthMiddleware([]string{model.UserTypeAdmin}, UpdatePromotion))
	promotions.DELETE("/:", AuthMiddleware([]string{model.UserTypeAdmin}, DeletePromotion))

	// Orders
	orders.POST("", AuthMiddleware([]string{model.UserTypeAdmin, model.UserTypeClient}, CreateOrder))
	orders.GET("", AuthMiddleware([]string{model.UserTypeAdmin, model.UserTypeClient}, GetOrders))
	orders.GET("/:id", AuthMiddleware([]string{model.UserTypeAdmin}, GetOrder))
	orders.PATCH("/:id", AuthMiddleware([]string{model.UserTypeAdmin, model.UserTypeClient}, UpdateOrder))
	orders.DELETE("/:", AuthMiddleware([]string{model.UserTypeAdmin, model.UserTypeClient}, DeleteOrder))

	// Users
	users.POST("", AuthMiddleware([]string{model.UserTypeAdmin}, CreateUser))
	users.GET("", AuthMiddleware([]string{model.UserTypeAdmin}, GetUsers))
	users.GET("/:id", AuthMiddleware([]string{model.UserTypeAdmin, model.UserTypeClient}, GetUser))
	users.PATCH("/:id", AuthMiddleware([]string{model.UserTypeAdmin, model.UserTypeClient}, UpdateUser))
	users.DELETE("/:", AuthMiddleware([]string{model.UserTypeAdmin}, DeleteUser))

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func AuthMiddleware(authorizedUserTypes []string, next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse the header
		userID, err := base64.StdEncoding.DecodeString(c.Request().Header.Get("X-USERID"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid user id in header")
		}

		id, err := strconv.Atoi(string(userID))
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, "Invalid user id in header")
		}

		// get the user and check  the user type
		user, err := db.GetUser(uint64(id))
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		var authorized bool
		if len(authorizedUserTypes) == 0 {
			authorized = true
		} else {
			authorized = false
			// check the user type
			for _, authorizedUserType := range authorizedUserTypes {
				if authorizedUserType == user.Type {
					authorized = true
					break
				}
			}

			if !authorized {
				return c.NoContent(http.StatusUnauthorized)
			}
		}

		c.Request().Header.Set("UserID", string(userID))

		return next(c)
	}
}

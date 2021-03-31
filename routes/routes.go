package routes

import (
	"os"
	"retailStore/controllers"
	"retailStore/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	middlewares.LogMiddlewares(e)

	e.POST("/register", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUserController)

	e.GET("/items", controllers.GetItemController)
	e.GET("/items/:id", controllers.GetItemWIthParamsController)

	e.GET("/couriers", controllers.GetCouriersController)
	e.GET("/couriers/:id", controllers.GetCourierByIdController)
	e.DELETE("/couriers/:id", controllers.DeleteCourierByIdController)
	e.PUT("/couriers/:id", controllers.UpdateCourierByIdController)
	e.POST("/couriers", controllers.CreateCourierController)

	eJWT := e.Group("") 
	eJWT.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eJWT.POST("/address", controllers.CreateAddressController)
	eJWT.GET("/address", controllers.GetAddressController)
	eJWT.GET("/address/:id", controllers.GetAddressByIdController)

	eJWT.GET("/users", controllers.GetUserDetailController)
	eJWT.PUT("/users", controllers. UpdateUserDetailController)
	
	eJWT.GET("/shoppingcarts", controllers.GetShoppingCartController)
	eJWT.POST("/shoppingcarts", controllers.PostItemToShoppingCartController)
	eJWT.DELETE("/shoppingcarts", controllers.DeleteItemFromShoppingCartController)

	eJWT.GET("/orders", controllers.GetOrderController)
	eJWT.POST("/orders", controllers.PostOrderController)
	eJWT.DELETE("/orders", controllers.DeleteOrderController)

	eAdmin := eJWT.Group("")
	eAdmin.POST("/items", controllers.PostItemController)
	
	return e
}

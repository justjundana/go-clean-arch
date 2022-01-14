package routes

import (
	_newsController "alta/cleanarch/controllers/news"

	echo "github.com/labstack/echo/v4"
)

type ControllerList struct {
	NewsController _newsController.NewsController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	news := e.Group("news")
	news.GET("/list", cl.NewsController.GetAll)
	news.GET("/list/:id", cl.NewsController.GetByID)
	// news.POST("/store", cl.NewsController.Store, middlewareApp.RoleValidation("admin"))
	// news.PUT("/update", cl.NewsController.Update)
}

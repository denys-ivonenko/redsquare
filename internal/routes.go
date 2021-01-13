package internal

import (
	"github.com/denys-ivonenko/redsquare/internal/controllers"
	"github.com/labstack/echo/v4"
)

//EchoRouter ....
func EchoRouter() *echo.Echo {
	e := echo.New()
	e.File("/", "../../web/observer.html")
	e.File("/observer", "../../web/observer.html")
	e.File("/admin", "../../web/admin.html")
	e.POST("/move-square", controllers.MoveSquare)
	return e
}

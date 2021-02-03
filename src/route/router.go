package route

import (
	"fmt"
	// "github.com/eurie-inc/echo-sample/api"
	// "github.com/eurie-inc/echo-sample/db"
	// "github.com/eurie-inc/echo-sample/handler"
	// myMw "github.com/eurie-inc/echo-sample/middleware"
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
	// lib "../lib"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Debug()

	// request path에서 마지막에 '/' 있는 경우 제거
	e.Pre(middleware.RemoveTrailingSlash())
	// 처리 도중 문제 발생 시 서버를 죽이지 않고 계속 수행하도록 설정
  e.Use(middleware.Recover())
	// Set Bundle MiddleWare
	e.Use(echoMw.Logger())
	// e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	// e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)
	e.HTTPErrorHandler = customHTTPErrorHandler

	// Set Custom MiddleWare
	// e.Use(myMw.TransactionHandler(db.Init()))

	// Routes
	e.GET("/", handler)
	authRoute(e)
	gWWW := e.Group("/www")
	gWWWRoute(gWWW)

	return e
}

func authRoute(e *echo.Echo) {
	// e.POST("/login", wwwwDefault)
	e.GET("/login", handler)
}

func gWWWRoute(e *echo.Group) {
	grA := e.Group("/groupa")
	grA.POST("/", handler)
	grA.GET("/", handler)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("%d.jsp", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}

func handler(c echo.Context) error {
	return c.String(http.StatusOK, c.Request().RequestURI)
}

	// data, err := json.MarshalIndent(e.Routes(), "", "  ")
	// if err != nil {
	// 	return err
	// }
	// ioutil.WriteFile("routes.json", data, 0644)
package route

import (
	"fmt"
	"net/http"
	// "github.com/eurie-inc/echo-sample/api"
	// "github.com/eurie-inc/echo-sample/db"
	// "github.com/eurie-inc/echo-sample/handler"
	// myMw "github.com/eurie-inc/echo-sample/middleware"
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"

	"auth"
	"lib"
)

var db *lib.DB

func Init() *echo.Echo {
	e := echo.New()

	// e.Debug()
	db = lib.NewConnection()
	// request path에서 마지막에 '/' 있는 경우 제거
	// e.Pre(middleware.RemoveTrailingSlash())
	// 처리 도중 문제 발생 시 서버를 죽이지 않고 계속 수행하도록 설정
  e.Use(echoMw.Recover())
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

func Router() *echo.Echo {
	e := echo.New()

	// e.Debug()

	// e.Static("/static", "public")
	e.Static("/static", "public")

	// e.Use(echoMw.StaticWithConfig(echoMw.StaticConfig{
	// 	Root:   "public",
	// 	Browse: true,
	// }))

	// request path에서 마지막에 '/' 있는 경우 제거
	// e.Pre(echoMw.RemoveTrailingSlash())
	// 처리 도중 문제 발생 시 서버를 죽이지 않고 계속 수행하도록 설정
  e.Use(echoMw.Recover())
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
	fmt.Println("111111111")
	e.POST("/signin", signinHandler)
}

func gWWWRoute(e *echo.Group) {
	grA := e.Group("/aaa")
	grA.POST("/", handler)
	grA.GET("/", handler)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("/static/%d.html", code)
	
	errorPage = "404.html"
	fmt.Println("ERROR PAGE :::::::::::::::", errorPage)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}

	// c.Logger().Error(err)
}

func handler(c echo.Context) error {
	return c.String(http.StatusOK, c.Request().RequestURI)
	// return c.File("/static/404.html")
}

func signinHandler(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	fmt.Println("====================", email, "===", password, "===================")

	if bln, _ := auth.SignIn(db, email, password); bln {
		str := c.Request().RequestURI + " OK \n" + email + "\n" + password
		return c.String(http.StatusOK, str)
	} else {
		return c.String(http.StatusOK, c.Request().RequestURI)
	}

	// return c.String(http.StatusOK, c.Request().RequestURI)
	// return c.File("/static/404.html")
}
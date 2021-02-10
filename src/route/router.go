package route

import (
	"fmt"
	"net/http"
	"strconv"
	"errors"

	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"

	"auth"
	"lib"
)

// var db *lib.DB

func Router() *echo.Echo {
	e := echo.New()

	// e.Debug()
	// e.Static("/static", "public")

	// request path에서 마지막에 '/' 있는 경우 제거
	e.Pre(echoMw.RemoveTrailingSlash())
	// 처리 도중 문제 발생 시 서버를 죽이지 않고 계속 수행하도록 설정
  e.Use(echoMw.Recover())

	e.Use(echoMw.LoggerWithConfig(echoMw.LoggerConfig{
		Format: "[${time_rfc3339_nano}]id:${id}:remote_ip:${remote_ip},host:${host},method=${method},uri=${uri},status=${status},latency:${latency}\n${error}",
	}))

	e.Use(dbContext(lib.NewConnection()))

	// e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	// e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)
	e.HTTPErrorHandler = customHTTPErrorHandler

	// Routes
	e.GET("/", func (c echo.Context) error {
		// c.Response().Before(func() {
		// 	println("before response")
		// })
		// c.Response().After(func() {
		// 	println("after response")
		// })

		return c.String(http.StatusOK, "Hello, World!")
	})
	authRoute(e)
	// gWWW := e.Group("/www")
	// gWWWRoute(gWWW)

	return e
}

func dbContext(db *lib.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			c.Set("db", db)
			return next(c)
		}
	}
}

func authRoute(e *echo.Echo) {
	e.POST("/signin", signinHandler)
	e.GET("/register", registerHandler)
}

func handler(c echo.Context) error {
	fmt.Println("==================handler")
	return c.String(http.StatusOK, c.Request().RequestURI)
	// return c.File("/static/404.html")
}
func signinHandler(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if bln, err := auth.SignIn(c.Get("db").(*lib.DB), email, password); err != nil {
		return err
	} else {
		return c.String(http.StatusOK, strconv.FormatBool(bln))
	}
}

func registerHandler(c echo.Context) error {
	email := c.QueryParam("email")
	password := c.QueryParam("password")
	if len(email) == 0 || len(password) == 0 {
		return errors.New("You have to input email or password")
	}
	fmt.Println("== email: ", email, "=== password: ", password, "=")

	if n, err := auth.Register(c.Get("db").(*lib.DB), email, password); err != nil {
		return err
	} else {
		return c.String(http.StatusOK, strconv.FormatInt(n, 10))
	}
}

// func gWWWRoute(e *echo.Group) {
// 	grA := e.Group("/aaa")
// 	grA.POST("/", handler)
// 	grA.GET("/", handler)
// }

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("%d.html", code)
	errorPage = "404.html"

	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	} 
	c.Logger().Error(err)
}



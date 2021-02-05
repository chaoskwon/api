package main

import (
	"net/http"

	"github.com/Sirupsen/logrus"

	// "database/sql"
	// "time"
	// "fmt"
	// _ "github.com/go-sql-driver/mysql"

	lib "chaoskwon/api/lib"

	"path"
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

var db *lib.DB

func init() {
	db = lib.NewConnection()
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	defer db.Close()

	// router := route.Init()
	router := Router()
	// Start server
	router.Start(":1323")
}



// func (self *DB) GetAFiled(query as string, fld as interface{}) (interface{}, error) {
// 	if err := db.QueryRow(query).Scan(&fld); err != nil {
// 		return nil, err;
// 	}

// 	fmt.Println("value:::::::::::::", fld)
// 	return fld, nil
// }


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
	e.GET("/login", handler)
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
	errorPage := fmt.Sprintf("%d.html", code)
	
	errorPage = path.Join("static", "404.html")
	fmt.Println("ERROR PAGE :::::::::::::::", errorPage)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}

	// c.Logger().Error(err)
}

func handler(c echo.Context) error {
	// return c.String(http.StatusOK, c.Request().RequestURI)
	return c.File("/static/404.html")
}
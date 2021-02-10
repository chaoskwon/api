module route

go 1.15

replace lib => ../lib

replace auth => ../auth

require (
	auth v0.0.0-00010101000000-000000000000
	github.com/labstack/echo/v4 v4.1.17
	lib v0.0.0-00010101000000-000000000000
)

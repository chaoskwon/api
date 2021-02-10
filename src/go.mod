module main

go 1.15

replace lib => ./lib

replace route => ./route

replace auth => ./auth

require (
	github.com/go-sql-driver/mysql v1.5.0
	route v0.0.0-00010101000000-000000000000 // indirect
)

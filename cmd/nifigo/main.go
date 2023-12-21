package main

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "nifi-go"
	// Version is the version of the compiled software.
	Version = "0.0.1"
	//
	Build = "2023-01-01 00:00:00"
)

// @title nifi-go API
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	app := wireApp()

	app.Start(":1234")
}

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

func main() {
	app := wireApp()

	app.Start(":1234")
}

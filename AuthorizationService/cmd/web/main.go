package main

import (
	"gitlab.com/AlimKenzza/authorization/pkg/data"
	"gitlab.com/AlimKenzza/authorization/router"
)

func init() {
	data.Connect()
}

func main() {
	r := router.SetupRouter()
	// Listen and Serve in 0.0.0.0:8081
	r.Run(":8081")
}

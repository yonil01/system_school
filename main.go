package main

import (
	"fmt"
	"foro-hotel/api"
	"foro-hotel/internal/env"
	"os"
	"strconv"
)

func main() {
	c := env.NewConfiguration()

	port := os.Getenv("PORT")

	puerto, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("Error fatal in asignature port")
	}

	api.Start(puerto, c.App.ServiceName, c.App.LoggerHttp, c.App.AllowedDomains)

}

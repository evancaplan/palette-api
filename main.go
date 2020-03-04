package main

import (
	"github.com/rs/cors"
	"log"
	"net/http"
	"palatte-api/internal/httptransport"
	"palatte-api/internal/service"
)

func main() {
	calcService, err := service.NewCalculatorService()
	if err != nil {
		log.Fatalln(err)
	}
	handler := httptransport.NewHandler(calcService)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	h := c.Handler(handler)

	err = http.ListenAndServe(":8080", h)
	log.Fatalln(err)
}

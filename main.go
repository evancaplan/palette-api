package palatte_api

import (
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

	err = http.ListenAndServe(":8080", handler)
	log.Fatalln(err)
}

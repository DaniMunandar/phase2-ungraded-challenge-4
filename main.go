package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"

	"ungraded-challenge-4/handler"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/phase2_ngc4")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Buat instance dari CriminalReportHandler
	criminalReportHandler := handler.NewCriminalReportHandler(db)

	// Membuat router menggunakan httprouter
	router := httprouter.New()

	// Menentukan rute-rute untuk laporan kejahatan
	router.GET("/reports/:id", criminalReportHandler.GetCriminalReport)
	router.GET("/reports", criminalReportHandler.GetAllCriminalReports)
	router.POST("/reports", criminalReportHandler.CreateCriminalReport)
	router.PUT("/reports/:id", criminalReportHandler.UpdateCriminalReport)
	router.DELETE("/reports/:id", criminalReportHandler.DeleteCriminalReport)

	port := 8084
	fmt.Printf("Server is running on port %d...\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		log.Fatal(err)
	}
}

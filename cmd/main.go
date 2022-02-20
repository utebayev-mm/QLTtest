package main

import (
	"fmt"
	"net/http"
	"qltTestApi/controller"
	"qltTestApi/model/sql"
	"qltTestApi/utils"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// db.Create()
	r := Router()
	Port := ":8080"
	fmt.Println("server is listening on the port", Port)
	http.ListenAndServe(Port, r)
}

func Router() *mux.Router {
	r := mux.NewRouter()

	db := sql.InitDb()

	repo := utils.NewRepo(db)
	transaction := controller.NewTransaction(repo)
	category := controller.NewCategory(repo)

	r.HandleFunc("/", controller.MainPage).Methods("GET")
	r.HandleFunc("/addatransaction/", transaction.AddATransaction).Methods("GET", "POST")
	r.HandleFunc("/addacategory", category.AddACategory).Methods("GET", "POST")
	r.HandleFunc("/deleteatransaction/", transaction.DeleteTransactionByID).Methods("GET", "POST")
	r.HandleFunc("/deleteacategory", category.DeleteCategoryByID).Methods("GET", "POST")
	r.HandleFunc("/updatetransaction/{id:[0-9]+}", transaction.UpdateTransaction).Methods("GET", "POST")
	r.HandleFunc("/deletetransaction/{id:[0-9]+}", transaction.DeleteTransaction).Methods("GET", "POST")
	r.HandleFunc("/updatecategory/{id:[0-9]+}", category.UpdateCategory).Methods("GET", "POST")
	r.HandleFunc("/deletecategory/{id:[0-9]+}", category.DeleteCategory).Methods("GET", "POST")
	r.HandleFunc("/browsetransactions", transaction.BrowseTransactions).Methods("GET")
	r.HandleFunc("/browsecategories", category.BrowseCategories).Methods("GET")

	staticFileDirectory := http.Dir("../static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	return r
}

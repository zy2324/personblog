package main

import (
	"net/http"
	//"io/ioutil"
	//"encoding/json"
	//"fmt"
	"backend/model"
	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"

)

func main() {
	http.HandleFunc("/blog/add/article", model.AddArticle)
	http.HandleFunc("/blog/get/titles", model.GetTitles)
	http.ListenAndServe(":6000",nil)
}

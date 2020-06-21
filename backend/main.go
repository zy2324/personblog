package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"backend/model"
)

func add(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body) 
	if err != nil {
		return
	}

	var article Article

	if err := json.Unmarshal(body, &article); err != nil {
		return
	}

	fmt.Println(article)

	if article.Keyword == "yiqikanshijie"{
		addArticle(article)
	} else {
		fmt.Println("keyword wrong")
		w.Write([]byte("keyword wrong"))
	}
}

func main() {
	http.HandleFunc("/add/article", add)
	http.ListenAndServe(":6000",nil)
}

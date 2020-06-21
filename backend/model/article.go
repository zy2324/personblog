import model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Title string `json:"title"`
	Words string `json:words`
	Keyword string `json:keyword`
}

var db *sql.DB

func init()(success bool, db *sql.DB) {
	var isopen bool
	host := "127.0.0.1"
	user := ""
	pwd := ""
	database := "articl"

	db, err = sql.Open("mysql", user+":"+pwd+"@tcp("+host+":3306)/"+database+"?charset=utf8")
	if err != nil {
		fmt.Println("open db failed")
		return
	} 

	err = db.Ping()
	if err != nil {
		fmt.Println("ping failed")
		return
	}

	fmt.Println("connect mysql ok")
}

func addArticl(data Article) {
	res, err := db.exec("INSERT INTO article(title, words) VALUES(?,?)", data.Title, data.Words)
	if err != nil {
		fmt.Println("insert article failed")
	}
	fmt.Println("result:", res)
}
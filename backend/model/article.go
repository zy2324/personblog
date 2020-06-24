package model

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

func init() {
	host := "127.0.0.1"
	user := "root"
	pwd := ""
	database := "blog"

	db, _  = sql.Open("mysql", user+":"+pwd+"@tcp("+host+":3306)/"+database+"?charset=utf8")

	err := db.Ping()
	if err != nil {
		fmt.Println("ping failed")
	}

	fmt.Println("connect mysql ok")
	fmt.Println(db)
}

func insertArticle(data Article) {
	fmt.Println(db)
	res, err := db.Exec("INSERT INTO article(title, words) VALUES(?,?)", data.Title, data.Words)
	if err != nil {
		fmt.Println("insert article failed")
	}
	fmt.Println("result:", res)
}

func selectTitles() []string {
	var res []string
	rows, err := db.Query("SELECT title from article")
	if err != nil {
		fmt.Println("select titles failed")
	}
	defer rows.Close()

	for rows.Next() {
		var t string
		err := rows.Scan(&t)
		if err != nil {
			fmt.Println("scan title failed")
			return
		}
		res = append(res, t)
	}

	return res
}

func getLastArticle() (string, string) {
	var title, words string
	err := db.QueryRow("SELECT title, words from article order by id DESC limit 1").Scan(&title, &words)
	if err != nil {
		fmt.Println("select or scan last word failed")
		return
	}
	return title, words
}
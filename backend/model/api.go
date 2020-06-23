package model

import (
        "net/http"
        "io/ioutil"
        "encoding/json"
        "fmt"
)

func Addapi(w http.ResponseWriter, r *http.Request) {
        body, err := ioutil.ReadAll(r.Body) 
        if err != nil {
                return
        }

        var article Article

        if err := json.Unmarshal(body, &article); err != nil {
                return
        }

        fmt.Println(article, db)

        if article.Keyword == ""{
                addArticle(article)
                w.Write([]byte("add ok"))
        } else {
                fmt.Println("keyword wrong")
                w.Write([]byte("keyword wrong"))
        }
}

package model

import (
        "net/http"
        "io/ioutil"
        "encoding/json"
        "fmt"
)

func AddArticle(w http.ResponseWriter, r *http.Request) {
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
                insertArticle(article)
                w.Write([]byte("add ok"))
        } else {
                fmt.Println("keyword wrong")
                w.Write([]byte("keyword wrong"))
        }
}

func GetTitles(w http.ResponseWriter, r *http.Request) {
        titles := selectTitles()
        datas := make(map[string]([]string))

        datas["titles"] = titles
        js, err := json.Marshal(datas)
        if err != nil {
                return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(js)
 }

 func GetYizhou(w http.ResponseWriter, r *http.Request) {
        title, words := getLastArticle()
        datas := make(map[string]string)

        datas["yizhoutitle"] = title
        datas["yizhouwords"] = words

        js, err := json.Marshal(datas)
        if err != nil {
                return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(js)
 }

 func GetVisit(w http.ResponseWriter, r *http.Request) {
         res, err := getVisit()
         if err != nil {
                w.Write([]byte("get visit failed"))
                return
         }

         fs, err := json.Marshal(res)
         if err != nil {
                 return 
         }

         w.Header().Set("Content-Type", "application/json")
         w.Write(js)
 }
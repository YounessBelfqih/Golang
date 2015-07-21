package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"google.golang.org/appengine"
	"time"
)


func init() {
	type User struct {
		Name  string
		Email string
	}
	tpl, err := template.ParseFiles("templates/index.gohtml", "templates/tweets.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {


		username := req.FormValue("name")
		email := req.FormValue("email")
		fmt.Println("Name: ", username)
		fmt.Println("[]byte(name):", []byte(username))
		fmt.Println("email: ", email)

		err= tpl.Execute(res, User{
			Name: username,
			Email: email,
		})
		if err!=nil {
			http.Error(res, err.Error(), 500)
		}})
	http.HandleFunc("/tweets", handleTweets)
}

func handleTweets(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	tpl := template.Must(template.ParseFiles("templates/tweets.gohtml"))
	compose := req.FormValue("compose")
	fmt.Println("tweet: ", compose)
	if compose != "" {
		makeTweet(ctx, &Tweet{
			Text: compose,
			Username: "????",
			Time: time.Now(),
		})
	}

	tweets, err := getTweets(ctx)
	if err !=nil {
		http.Error(res, err.Error(), 500)
	}

	err = tpl.Execute(res, tweets)
	if err!=nil {
		http.Error(res, err.Error(), 500)
	}
}
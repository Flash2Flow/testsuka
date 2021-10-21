package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/go-session/session"
)

type UserFull struct {
	Login   		  string
	Email 			  string
	Password 		  string
	Developer   	  int
	Ban				  int
	Group       	  int
	Undesirable 	  int
	UserKey			  int
}

func reg(page http.ResponseWriter, req *http.Request) {

	login := req.FormValue("login")
	email := req.FormValue("email")
	password := req.FormValue("password")
	repassword := req.FormValue("repassword")

	store, err := session.Start(context.Background(), page, req)
	if err != nil {
		fmt.Fprint(page, err)
		return
	}


	active, ok := store.Get("active_login")

	if ok {
		Active := fmt.Sprintf("User logged: %v", active)
		log.Println(Active)
		http.Redirect(page, req, "/home/", 302)
	}else{
		http.Redirect(page, req, "/", 302)
	}

	if login == "" {
		fmt.Fprintf(page, "Login cant be nil")
		http.Redirect(page, req, "/", 302)
	}else{
		if email == ""{
			fmt.Fprintf(page, "Email cant be nil")
			http.Redirect(page, req, "/", 302)
		}else{
			if password == "" {
				fmt.Fprintf(page, "Password cant be nil")
				http.Redirect(page, req, "/", 302)
			}else{
				if password != repassword {
					fmt.Fprintf(page, "Password != repassword")
					http.Redirect(page, req, "/", 302)
				}else{
					users := fmt.Sprintf("/users/user_%s.json", login)
					f, err := os.Create(users)
					if err != nil {
						panic(err)
					}
					f.Close()

					md5_userkey := rand.Intn(9999999999)

					md5_password := GetMD5Hash(password)
					//reg
					c := UserFull{
						Login: login,
						Password: md5_password,
						Email: email,
						Developer: 0,
						Ban: 0,
						Group: 0,
						Undesirable: 0,
						UserKey: md5_userkey,
					}
					dat, err := json.Marshal(c)
					if err != nil {
						return
					}
					users_read := fmt.Sprintf("user_%s.json", login)
					err = ioutil.WriteFile(users_read, dat, 0644)
					if err != nil {
						return
					}

					store.Set("active_login", login)
					err = store.Save()
					if err != nil {
						fmt.Fprint(page, err)
						return
					}
					http.Redirect(page, req, "/home/", 302)
					auth := fmt.Sprintf("User auth: %s", login)
					log.Println(auth)
				}
			}
		}
	}

}
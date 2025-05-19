package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>TEST Shel website</h1><a href='/register'>Регистрация</a> | <a href='/login'>Вход</a>")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templates.ExecuteTemplate(w, "register.html", nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if userExists(username) {
		fmt.Fprintln(w, "the user exists!")
		return
	}

	f, _ := os.OpenFile("users.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	fmt.Fprintf(f, ">>> username: %s >>> password: %s\n", username, password)

	fmt.Fprintln(w, "Sueceful registred! <a href='/login'>enter</a>")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templates.ExecuteTemplate(w, "login.html", nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if validateUser(username, password) {
		fmt.Fprintf(w, "Welcome!, %s!", username)
	} else {
		fmt.Fprintln(w, "Something is not right")
	}
}

func userExists(username string) bool {
	f, err := os.Open("users.txt")
	if err != nil {
		return false
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ">>> username: "+username+" ") {
			return true
		}
	}
	return false
}

func validateUser(username, password string) bool {
	f, err := os.Open("users.txt")
	if err != nil {
		return false
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ">>> username: "+username+" ") && strings.Contains(line, ">>> password: "+password) {
			return true
		}
	}
	return false
}

//Comment

//<|════════════════════════════════════════════════════════════════════════════════════════════════════════════|>
// >>> Я использовал мини шаблон а именно ввод register и login а далее уже сам по гайдам с ютуба

// >>> Templates папка где лежит login.html,register.html

// func registerHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		templates.ExecuteTemplate(w, "register.html", nil)
// 		return
// 	}

// 	username := r.FormValue("username")
// 	password := r.FormValue("password")

// 	if userExists(username) {
// 		fmt.Fprintln(w, "the user exists!")
// 		return
// 	}

// 	f, _ := os.OpenFile("users.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	defer f.Close()
// 	fmt.Fprintf(f, ">>> username: %s >>> password: %s\n", username, password)

// 	fmt.Fprintln(w, "Sueceful registred! <a href='/login'>enter</a>")
// }

// >>> Данные строчки кода которые уже и выводят в user.txt логи имя и пороль
//fmt.Fprintf(f, ">>> username: %s >>> password: %s\n", username, password) >>> стиль вывода

//templates.ExecuteTemplate(w, "register.html", nil) >>> берет лог из данного файла
//templates.ExecuteTemplate(w, "login.html", nil) >>> берет лог из данного файла
//<|════════════════════════════════════════════════════════════════════════════════════════════════════════════|>

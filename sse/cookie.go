package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func GetURL() string {
	return fmt.Sprintf("http://localhost:%d/", PORT)
}

func main2() {
	http.Handle("/", http.FileServer(http.Dir("client")))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("client"))))
	http.HandleFunc("/random", random)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}

func secret(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("exampleCookie")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	w.Write([]byte(cookie.Value))
}

func login(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "exampleCookie",
		Value:    "Hello world 2!",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	http.SetCookie(w, &cookie)

	http.Redirect(w, r, GetURL(), http.StatusSeeOther)
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("exampleCookie")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	cookie.Value = ""
	http.SetCookie(w, cookie)

	http.Redirect(w, r, GetURL(), http.StatusSeeOther)
}

func rand2(w http.ResponseWriter, r *http.Request) {
	/*session, _ := store.Get(r, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}*/

	/*cookie, err := r.Cookie("exampleCookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		if len(cookie.Value) > 0 {
			start(w)
		}
	}*/
}

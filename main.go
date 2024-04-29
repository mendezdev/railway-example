package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("starting....")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/appetizers", func(w http.ResponseWriter, r *http.Request) {
		url := "https://seanallen-course-backend.herokuapp.com/swiftui-fundamentals/appetizers"
		res, err := http.Get(url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			msg := fmt.Sprintf("error sending request: %s", err.Error())
			w.Write([]byte(msg))
			return
		}

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			msg := fmt.Sprintf("error reading body response: %s", err.Error())
			w.Write([]byte(msg))
			return
		}

		fmt.Printf("RESPONSE: %s\n", string(resBody))

		w.Header().Add("Content-Type", "application/json")

		w.Write(resBody)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe("0.0.0.0:"+port, r)
}

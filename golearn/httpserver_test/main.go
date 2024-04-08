package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	// g := gin.Default()

	// g.GET("/", func(ctx *gin.Context) {

	// 	ctx.JSON(200, Result{
	// 		Message: "Sucess",
	// 	})
	// })

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello world"))
	// })
	// http.Handle("/Test", Myhandler{})
	// http.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
	// 	defer r.Body.Close()
	// 	bytes, err := io.ReadAll(r.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(string(bytes))
	// })

	// http.ListenAndServe(":8080", nil)

	var handler Myhandler

	s := http.Server{
		Addr:           ":8080",
		Handler:        handler,
		WriteTimeout:   1 * time.Second,
		ReadTimeout:    1 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}

type Result struct {
	Message string
}

type Myhandler int

func (m Myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	val := r.Header.Get("AccessToken")

	fmt.Printf("val: %v\n", val)

	w.Write([]byte("hello world1"))
	//<-time.After(2 * time.Second)
}

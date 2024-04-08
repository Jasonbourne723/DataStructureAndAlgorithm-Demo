package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err.Error())
	}

	//reader := bufio.NewReader(resp.Body)

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("b: %v\n", string(b))

	client := http.Client{}

	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	req.Header.Add("AccessToken", "asdfsafasdfsf")

	result, err := client.Do(req)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	defer result.Body.Close()

	bb, err := io.ReadAll(result.Body)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("bb: %v\n", string(bb))

	var buf bytes.Buffer

	json.NewEncoder(&buf).Encode(struct{ Name string }{"jason"})

	http.Post("http://localhost:8080/go", "application/json", &buf)

}

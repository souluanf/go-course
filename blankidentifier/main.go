package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, _ := http.Get("https://google.com.br")
	body, _ := io.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", body)
}

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	bd := make([]byte, 999999)
	res.Body.Read(bd)
	fmt.Println(string(bd))
}

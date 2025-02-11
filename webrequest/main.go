package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.freecodecamp.org/learn/back-end-development-and-apis/"

func main() {
	fmt.Println("hii webrequest testing")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Our response type is : %T", response)

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(data)

	fmt.Println(content)
}

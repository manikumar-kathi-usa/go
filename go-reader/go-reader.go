package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	for i := 0; i < 100; i++ {
		resp, err := http.Get(fmt.Sprint("https://jsonplaceholder.typicode.com/todos/1", i))
		if err != nil {
			panic(err)
		}
		buff := make([]byte, 100)

		n, err := resp.Body.Read(buff)
		if n > 0 {
			fmt.Println(string(buff[:n]))
		}

		if err == io.EOF {
			log.Println("clsoign connection")
			resp.Body.Close()
		}

	}

}

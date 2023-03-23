package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("http://10.129.8.44:3003/ping-server.php/ping/8.8.8.8`id`/3")

	if err != nil {
		fmt.Println(err)
	}

	// defer res.Body.Close()

	// if error != nil {
	// 	fmt.Println(error.Error())
	// }

	body, errRead := io.ReadAll(res.Body)
	defer res.Body.Close()

	if errRead != nil {
		fmt.Println(errRead)
	}

	fmt.Println(res.Header)
	fmt.Println(string(body))
	// fmt.Println(body)

}

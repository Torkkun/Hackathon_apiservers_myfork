package main

import (
	"app/handler"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("=== START ===")
	handler.Router.Run(":8080")

	//定期実行用　別でやったほうがよさそうだったらのちに変更するべきかも
	go Regular()
}

func Regular() {
	t := time.NewTicker(time.Minute)
	defer t.Stop()
	for range t.C {
		url := "http://localhost:8080/tojiuru/execution"
		res, err := http.Get(url)
		if err != nil {
			log.Fatalln(err.Error())
		} else {
			log.Println("[*] "+res.Status, res.Body)
		}

	}
}

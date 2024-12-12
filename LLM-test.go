package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Answer string `json:"answer"`
}

func main() {
	client := &http.Client{}
	var data = strings.NewReader(`{"inputs":{"target_language":"Chinese","query":"hello world"}}`)
	req, err := http.NewRequest("POST", "http://localhost/api/completion-messages", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "dify.black.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7")
	req.Header.Set("authorization", "Bearer XXX")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("origin", "http://localhost")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response Response
	json.Unmarshal(bodyText, &response)

	fmt.Println(response.Answer)
}
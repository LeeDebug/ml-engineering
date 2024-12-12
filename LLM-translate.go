package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func query(input string) (result string, err error) {
	client := &http.Client{}

	type Request struct {
		Inputs struct {
			TargetLanguage string `json:"target_language"`
			Query          string `json:"query"`
		} `json:"inputs"`
		ResponseMode string `json:"response_mode,omitempty"`
	}

	type Response struct {
		Answer string `json:"answer"`
	}

	var payload Request
	payload.Inputs.TargetLanguage = "Chinese"
	payload.Inputs.Query = input

	payloadBuf, err := json.Marshal(payload)
	if err != nil {
		return result, err
	}

	var data = strings.NewReader(string(payloadBuf))
	req, err := http.NewRequest("POST", "http://localhost/api/completion-messages", data)
	if err != nil {
		return result, err
	}

	req.Header.Set("authority", "dify.black.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
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
		return result, err
	}

	var response Response
	json.Unmarshal(bodyText, &response)
	return response.Answer, nil
}

func visit(paths *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			*paths = append(*paths, path)
		}
		return nil
	}
}

func main() {
	var mdFiles []string
	root := "./" // ../ml-engineering

	err := filepath.Walk(root, visit(&mdFiles))
	if err != nil {
		panic(err)
	}

	for _, filePath := range mdFiles {
		fmt.Println("开始处理文件", filePath)

		buf, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("读取文件出错: %s\n", err)
			return
		}

		translated, err := query(string(buf))
		if err != nil {
			fmt.Printf("调用模型翻译出错: %s\n", err)
			return
		}

		err = os.WriteFile(filePath, []byte(translated), 0644)
		if err != nil {
			fmt.Printf("保存文件出错: %s\n", err)
			return
		}
		fmt.Println("搞定！")
	}
}
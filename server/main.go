package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Member struct {
	Name     string
	Id       string
	Homecity string
}

func main() {
	jsonFile, err := os.Open("test.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var member Member

	json.Unmarshal([]byte(byteValue), &member)
	fmt.Println(member.Id)

	request := fmt.Sprintf(`{"id": "%s"}`, member.Id)

	jsonBody := []byte(request)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/word", bodyReader)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)

	respBody, _ := io.ReadAll(resp.Body)

	fmt.Println(string(respBody))
	resp.Body.Close()
}

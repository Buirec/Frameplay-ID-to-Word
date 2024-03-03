package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Member struct {
	Name     string
	Id       string
	Homecity string
}

func getId(input string) string {
	jsonFile, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var member Member

	json.Unmarshal([]byte(byteValue), &member)
	fmt.Println(member.Id)

	if _, err := strconv.Atoi(member.Id); err != nil {
		fmt.Println("This is not a valid ID")
		return "false"
	}

	return member.Id

}

func postId(id string) string {
	request := fmt.Sprintf(`{"id": "%s"}`, id)

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
	resp.Body.Close()

	return string(respBody)

}

func main() {
	var input string

	fmt.Print("Input the json file name: ")
	fmt.Scan(&input)

	id := getId(input)

	if id == "false" {
		return
	}

	fmt.Println(postId(id))

}

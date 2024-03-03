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
	// Code to handle file input
	jsonFile, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// Converts the file into a byte array
	byteValue, _ := io.ReadAll(jsonFile)

	var member Member

	// Parses the byte array into Go readable JSON data
	json.Unmarshal([]byte(byteValue), &member)

	// Validate the JSON has the correct fields
	if member.Id == "" || member.Name == "" || member.Homecity == "" {
		fmt.Println("This JSON file has the incorrect structure")
		return ""
	}

	fmt.Println(member.Id)

	// Checks if the ID gotten from the data is valid
	if _, err := strconv.Atoi(member.Id); err != nil {
		fmt.Println("This is not a valid ID")
		return ""
	}

	return member.Id

}

func postId(id string) string {
	// Creates a JSON packet
	request := fmt.Sprintf(`{"id": "%s"}`, id)

	jsonBody := []byte(request)
	bodyReader := bytes.NewReader(jsonBody)

	// Generates a request for this packet
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/word", bodyReader)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	// Sends the request to our locally run API and parses the result
	resp, err := client.Do(req)

	respBody, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	return string(respBody)

}

func main() {
	id := getId("data.json")

	if id == "" {
		return
	}

	fmt.Println(postId(id))

}

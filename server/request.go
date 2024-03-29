package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Config struct {
	Server  string
	Timeout time.Duration
}

func PostId(id string) string {
	// Checks if the ID given is valid
	if _, err := strconv.Atoi(id); err != nil {
		fmt.Println("This is not a valid ID")
		return ""
	}

	// Creates a JSON packet
	request := fmt.Sprintf(`{"id": "%s"}`, id)

	jsonBody := []byte(request)
	bodyReader := bytes.NewReader(jsonBody)

	var config Config

	// TODO: Given a real commerical application, validate the string here to ensure people just can't put whatever into the input

	// This gets our config file which should be protected
	jsonFile, err := content.ReadFile("private/config.json")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	err = json.Unmarshal([]byte(jsonFile), &config)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// Generates a request for this packet
	req, err := http.NewRequest(http.MethodPost, config.Server+"word", bodyReader)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: config.Timeout * time.Second,
	}

	// Sends the request to our locally run API and parses the result
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	respBody, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	return string(respBody)

}

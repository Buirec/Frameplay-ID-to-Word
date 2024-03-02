package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Member struct {
	Name     string
	Id       string
	Homecity string
}

func idToWord(id string) string {
	return "Hello"
}

func main() {
	memberJson := `{"name": "John Doe","id": "12345","homecity": "Melbourne"}`

	var member Member

	json.Unmarshal([]byte(memberJson), &member)
	fmt.Println(member.Id)

	word := idToWord(member.Id)

	resp, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + word)
	if err != nil {
		fmt.Println(err)
	}

	respBody, _ := io.ReadAll(resp.Body)
	var body []map[string]string
	json.Unmarshal(respBody, &body)

	fmt.Println(body[0]["word"])
	resp.Body.Close()
}

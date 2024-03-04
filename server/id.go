package server

import (
	"embed"
	"encoding/json"
	"fmt"
	"strconv"
)

//go:embed private
var content embed.FS

type Member struct {
	Name     string
	Id       string
	Homecity string
}

func GetId(input string) string {
	var member Member

	// TODO: Given a real commerical application, validate the string here to ensure people just can't put whatever into the input

	// Code to handle file input
	jsonFile, err := content.ReadFile("private/" + input)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// Parses the byte array into Go readable JSON data
	json.Unmarshal([]byte(jsonFile), &member)

	// Validate the JSON has the correct fields
	if member.Id == "" || member.Name == "" || member.Homecity == "" {
		fmt.Println("This JSON file has the incorrect structure")
		return ""
	}

	// This is for the STDOUT as covered by the spec
	fmt.Println(member.Id)

	// Checks if the ID gotten from the data is valid
	if _, err := strconv.Atoi(member.Id); err != nil {
		fmt.Println("This is not a valid ID")
		return ""
	}

	return member.Id

}

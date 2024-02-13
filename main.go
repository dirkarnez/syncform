package main

import (
	"fmt"
	"regexp"
	"time"

	"encoding/json"

	"github.com/invopop/jsonschema"
)

type SampleUser struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Friends     []int          `json:"friends"`
	Tags        map[string]any `json:"tags"`
	BirthDate   time.Time      `json:"birth_date"`
	YearOfBirth string         `json:"year_of_birth"`
	Metadata    any            `json:"metadata"`
	FavColor    string         `json:"fav_color"`
}

type User struct {
	Name string `json:"name"`
}

func main() {
	user := &SampleUser{}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	m1 := regexp.MustCompile(`"([^"]+)"\s*:\s*`)

	fmt.Println(m1.ReplaceAllString(string(b), "$1:"))

	s := jsonschema.Reflect(&SampleUser{})

	for pair := s.Definitions["SampleUser"].Properties.Oldest(); pair != nil; pair = pair.Next() {
		memberType := pair.Value.Type
		fmt.Printf("%s %s\n", pair.Key, memberType)
	} // prints:

}

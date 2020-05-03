package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Position is a GitHub job
type Position struct {
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	URL         string      `json:"url"`
	CreatedAt   string      `json:"created_at"`
	Company     string      `json:"company"`
	CompanyURL  string      `json:"company_url"`
	Location    string      `json:"location"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	HowToApply  string      `json:"how_to_apply"`
	CompanyLogo interface{} `json:"company_logo"`
}

func main() {

	url := "https://jobs.github.com/positions.json?description=developer&location=london"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", "Postman")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var positions []Position
	err = json.Unmarshal(body, &positions)

	fmt.Println(positions)
}

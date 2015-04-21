package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	NS      = os.Getenv("SIMPLEDB_NS")
	API_URL = "http://simpledb2.herokuapp.com/"
)

func main() {
	if len(NS) == 0 {
		fmt.Println("ERROR: unable to find SIMPLEDB_NS env variable")
		os.Exit(-1)
	}

	cmd := "list"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "list":
		list()
		break
	case "add":
		add(strings.Join(os.Args[2:], " "))
		break
	}

}

func list() {
	//currently we just show the tasks
	resp, _ := http.Get(API_URL + NS + "/tasks")

	p := []payload{}
	json.NewDecoder(resp.Body).Decode(&p)
	resp.Body.Close()

	for i, taskPayload := range p {
		t := task{}
		json.Unmarshal([]byte(taskPayload.Data), &t)

		fmt.Printf("%v. %s\n", i, t.Title[0])
	}

}

type task struct {
	Title []string `json:"title"`
}

type payload struct {
	Data string `json:"Data"`
}

func add(title string) {
	//currently we just show the tasks
	values := url.Values{}
	values.Add("title", title)
	resp, _ := http.PostForm(API_URL+NS+"/tasks", values)
	_ = resp
}

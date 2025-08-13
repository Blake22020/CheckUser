package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name      string `json:"login"`
	Bio       string `json:"bio"`
	Repos     int    `json:"public_repos"`
	Followers int    `json:"followers"`
}

func main() {
	var name string
	fmt.Scanln(&name)

	resp, err := http.Get("https://api.github.com/users/" + name)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		panic(err)
	}
	fmt.Printf("%s\n%s\nПубличных репозиториев: %d\nПодписчиков:%d\n", user.Name, user.Bio, user.Repos, user.Followers)
}

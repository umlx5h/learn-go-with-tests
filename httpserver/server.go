package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var (
	score = map[string]int{
		"Floyd":  10,
		"Pepper": 20,
	}
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, GetPlayerScore(name))
}

func GetPlayerScore(name string) string {
	return strconv.Itoa(score[name])
}

package main

import (
	"fmt"
	"net/http"
)

func playerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, 20)
}
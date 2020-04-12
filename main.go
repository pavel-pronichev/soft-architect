package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RespOk struct {
	Status string `json:"status"`
}

func healthCheck(w http.ResponseWriter, f *http.Request) {
	b, err := json.Marshal(RespOk{"OK"})
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}

	fmt.Fprint(w, string(b))
}

func main() {
	http.HandleFunc("/health/", healthCheck)
	http.ListenAndServe("0.0.0.0:8000", nil)
}

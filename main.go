package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ranupratapsingh/snowflake_id_k8s/snowflake"
)

func main() {
	http.HandleFunc("/generate", generateHandler)
	http.HandleFunc("/generate_base36", generateBase36Handler)
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	id := snowflake.GenerateInt64()
	w.Header().Set("Content-Type", "application/json")
	res := map[string]int64{"id": id}
	json.NewEncoder(w).Encode(res)
}

func generateBase36Handler(w http.ResponseWriter, r *http.Request) {
	s36 := snowflake.GenerateBase36()
	w.Header().Set("Content-Type", "application/json")
	res := map[string]string{"id": s36}
	json.NewEncoder(w).Encode(res)
}

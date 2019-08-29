package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// CheckRequest is a non-exported type checkRequest from
// "github.com/concourse/concourse/atc/resource"
type CheckRequest struct {
	Source  map[string]interface{} `json:"source"`
	Version map[string]string      `json:"version"`
}

type CheckResponse []map[string]string

func main() {
	log.Printf("check\n")

	var request CheckRequest
	if err := json.NewDecoder(os.Stdin).Decode(&request); err != nil {
		log.Fatalf("failed to unmarshal check request; %s", err)
	}
	log.Printf("req:%+v\n", request)

	response := CheckResponse{{"ref": "123"}}
	output, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal check response; %s", err)
	}

	fmt.Printf(string(output))
}

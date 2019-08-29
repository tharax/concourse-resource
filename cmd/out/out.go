package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// PutRequest is a non-exported type putRequest from
// "github.com/concourse/concourse/atc/resource"
type PutRequest struct {
	Source map[string]interface{} `json:"source"`
	Params map[string]interface{} `json:"params,omitempty"`
}

type PutResponse struct {
	Version  map[string]string        `json:"version"`
	Metadata []map[string]interface{} `json:"metadata,omitempty"`
}

func main() {
	log.Printf("out\n")

	var request PutRequest
	if err := json.NewDecoder(os.Stdin).Decode(&request); err != nil {
		log.Fatalf("failed to unmarshal get request; %s", err)
	}
	log.Printf("req:%+v\n", request)

	response := PutResponse{
		Version: map[string]string{"ver": "789"},
		Metadata: []map[string]interface{}{
			{"ref": "10,11,12"},
		},
	}
	output, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal get response; %s", err)
	}

	fmt.Printf(string(output))

}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// GetRequest is a non-exported type getRequest from
// "github.com/concourse/concourse/atc/resource"
type GetRequest struct {
	Source  map[string]interface{} `json:"source"`
	Params  map[string]interface{} `json:"params,omitempty"`
	Version map[string]string      `json:"version,omitempty"`
}

type GetResponse struct {
	Version  map[string]string        `json:"version"`
	Metadata []map[string]interface{} `json:"metadata,omitempty"`
}

func main() {
	log.Printf("in\n")

	var request GetRequest
	if err := json.NewDecoder(os.Stdin).Decode(&request); err != nil {
		log.Fatalf("failed to unmarshal get request; %s", err)
	}
	log.Printf("req:%+v\n", request)

	response := GetResponse{
		Version: request.Version,
		Metadata: []map[string]interface{}{
			{"ref": "456"},
		},
	}
	output, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal get response; %s", err)
	}

	fmt.Printf(string(output))
}

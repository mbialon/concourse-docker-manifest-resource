package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mbialon/concourse-docker-manifest-resource/pkg/docker/manifest"
)

type Request struct {
	Source  *Source  `json:"source"`
	Version *Version `json:"version"`
}

type Source struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type Version struct {
	Digest string `json:"digest"`
}

func main() {
	if err := os.Chdir(os.Args[1]); err != nil {
		log.Fatalf("cannot change dir: %v", err)
	}
	var request Request
	if err := json.NewDecoder(os.Stdin).Decode(&request); err != nil {
		log.Fatalf("cannot decode input: %v", err)
	}
	manifestList := fmt.Sprintf("%s:%s@%s",
		strings.TrimSpace(request.Source.Repository),
		strings.TrimSpace(request.Source.Tag),
		request.Version.Digest)
	if err := manifest.Inspect(manifestList); err != nil {
		log.Fatalf("cannot inspect manifest: %v", err)
	}
	output := map[string]interface{}{
		"version": request.Version,
	}
	if err := json.NewEncoder(os.Stdout).Encode(output); err != nil {
		log.Fatalf("cannot encode output: %v", err)
	}
}

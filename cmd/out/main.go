package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Request struct {
	Source *Source `json:"source"`
	Params *Params `json:"params"`
}

type Source struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type Params struct {
	Arch    string `json:"arch"`
	OS      string `json:"os"`
	TagFile string `json:"tag_file"`
}

func main() {
	if err := os.Chdir(os.Args[1]); err != nil {
		log.Fatalf("cannot change dir: %v", err)
	}
	var request Request
	if err := json.NewDecoder(os.Stdin).Decode(&request); err != nil {
		log.Fatalf("cannot decode input: %v", err)
	}
	b, err := ioutil.ReadFile(request.Params.TagFile)
	if err != nil {
		log.Fatalf("cannot read tag file: %v", err)
	}
	tag := string(b)
	manifest := strings.TrimSpace(request.Source.Repository) + ":" + strings.TrimSpace(tag)
	manifestList := strings.TrimSpace(request.Source.Repository) + ":" + strings.TrimSpace(request.Source.Tag)
	if err := dockerLogin(request.Source); err != nil {
		log.Fatalf("cannot login to docker hub: %v", err)
	}
	if err := createManifest(manifestList, manifest); err != nil {
		log.Fatalf("cannot create manifest: %v", err)
	}
	if err := annotateManifest(manifestList, manifest, request.Params); err != nil {
		log.Fatalf("cannot annotate manifest: %v", err)
	}
	digest, err := pushManifest(manifestList)
	if err != nil {
		log.Fatalf("cannot push manifest: %v", err)
	}
	output := map[string]interface{}{
		"version": map[string]string{
			"digest": digest,
		},
	}
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(output); err != nil {
		log.Fatalf("cannot encode output: %v", err)
	}
}

func dockerLogin(source *Source) error {
	cmd := exec.Command("docker", "login", "-u", source.Username, "-p", source.Password)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func createManifest(manifestList, manifest string) error {
	cmd := exec.Command("docker", "manifest", "create", "--amend", manifestList, manifest)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func annotateManifest(manifestList, manifest string, params *Params) error {
	cmd := exec.Command("docker", "manifest", "annotate", "--arch", params.Arch, "--os", params.OS, manifestList, manifest)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func pushManifest(manifestList string) (string, error) {
	cmd := exec.Command("docker", "manifest", "push", manifestList)
	cmd.Stderr = os.Stderr
	b, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}

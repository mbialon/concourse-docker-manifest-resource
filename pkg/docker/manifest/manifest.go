package manifest

import (
	"os"
	"os/exec"
	"strings"
)

type Annotation struct {
	Manifest     string
	Architecture string
	OS           string
}

func Create(manifestList string, manifests []string) error {
	args := append([]string{"manifest", "create", "--amend", manifestList}, manifests...)
	cmd := exec.Command("docker", args...)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Annotate(manifestList string, annotations []Annotation) error {
	for _, a := range annotations {
		cmd := exec.Command("docker", "manifest", "annotate", "--arch", a.Architecture, "--os", a.OS, manifestList, a.Manifest)
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}

func Push(manifestList string) (string, error) {
	cmd := exec.Command("docker", "manifest", "push", manifestList)
	cmd.Stderr = os.Stderr
	b, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}

func Inspect(manifestList string) error {
	cmd := exec.Command("docker", "manifest", "inspect", manifestList)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

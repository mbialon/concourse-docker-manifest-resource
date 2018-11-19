package docker

import (
	"os"
	"os/exec"
)

func Login(username, password string) error {
	cmd := exec.Command("docker", "login", "-u", username, "-p", password)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

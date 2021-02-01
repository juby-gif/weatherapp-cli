package utils

import (
	"os"
	"os/exec"
)

func GetScreenCleared() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

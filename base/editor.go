package base

import (
	"os"
	"log"
	"strings"
	"syscall"
	"os/exec"
	"io/ioutil"
)

func GetEditorText() string {
	editorFile, err := ioutil.TempFile("", "mypass")
	if err != nil {
		log.Fatal(err)
	}
	defer syscall.Unlink(editorFile.Name())
	cmd := exec.Command("vi", editorFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err_file := cmd.Start()
	if err_file != nil {
		log.Fatal(err)
	}
	err_file = cmd.Wait()
	if err_file != nil {
		log.Fatal(err)
	}
	data, err_file := ioutil.ReadFile(editorFile.Name())
	if err_file != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(data))
}

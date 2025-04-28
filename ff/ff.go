package ff

import (
	"errors"
	"os"
	"os/exec"
)

func FileExists(filepath string) bool {
	if _, err := os.Open(filepath); os.IsNotExist(err) {
		return false
	}
	return true
}

func WriteIfNotExist(filepath string, body []byte) error {
	if FileExists(filepath) {
		return errors.New("File already exists!")
	}
    if err := os.WriteFile(filepath, body, 0666); err != nil {
		return err
	}
	return nil
}

func CreateAndWrite(name string, body []byte) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = file.Write(body)
	if err != nil {
		return err
	}
	return nil
}

func OpenInEditor(editor string, filepath string) error {
	cmd := exec.Command(editor, filepath)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Clone(repository string) (string, error) {
	tmp, err := os.MkdirTemp(os.TempDir(), "git")
	if err != nil {
		return "", err
	}
	tmp = fmt.Sprintf("%s/tmp", tmp)
	cmd := exec.Command("git", "clone", repository, tmp)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		os.RemoveAll(tmp)
		return "", err
	}
	return tmp, nil
}

func UpdateImage(path string, key string, tag string) error {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	yaml, err := Unmarshal(yamlFile)
	if err != nil {
		return err
	}
	originalTag := yaml.Get(key)
	originalTagValue, ok := originalTag.(string)
	if !ok {
		return fmt.Errorf("image path does not have a string value")
	}
	imageWithoutTag := strings.Split(originalTagValue, ":")[0]
	yaml.Set(key, fmt.Sprintf("%s:%s", imageWithoutTag, tag))
	marshalledYaml, err := yaml.Marshall()
	if err != nil {
		return err
	}
	return os.WriteFile(path, marshalledYaml, os.ModePerm)
}

func StageAll(path string) error {
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Commit(path string) error {
	cmd := exec.Command("git", "commit", "-m", "\"auto image update\"")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Push(path string, repository string, username string, token string) error {
	url := strings.Split(repository, "@")[1]
	cmd := exec.Command("git", "push", fmt.Sprintf("https://%s:%s@%s", username, token, url), "master")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Clear(path string) error {
	return os.RemoveAll(path)
}

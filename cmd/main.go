package main

import (
	"argocdimageupdater/internal"
	"fmt"
	"os"
)

func main() {
	repository := os.Getenv("CI_REPOSITORY_URL")
	fmt.Println("CI_REPOSITORY_URL", repository)
	user := os.Getenv("CI_USER")
	fmt.Println("CI_USER", user)
	token := os.Getenv("CI_TOKEN")
	fmt.Println("CI_TOKEN", token)
	tag := os.Getenv("CI_COMMIT_TAG")
	fmt.Println("CI_COMMIT_TAG", tag)
	filePath := os.Getenv("CI_IMAGE_FILE")
	fmt.Println("CI_IMAGE_FILE", filePath)
	keyPath := os.Getenv("CI_IMAGE_KEY")
	fmt.Println("CI_IMAGE_KEY", keyPath)

	path, err := internal.Clone(repository)
	if err != nil {
		panic(err)
	}
	err = internal.UpdateImage(fmt.Sprintf("%s/%s", path, filePath), keyPath, tag)
	if err != nil {
		panic(err)
	}
	err = internal.Commit(path)
	if err != nil {
		panic(err)
	}
	err = internal.Push(path, repository, user, token)
	if err != nil {
		panic(err)
	}
}

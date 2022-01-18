package service

import (
	"fmt"
	"io/ioutil"

	"path/filepath"
)

// creating of default dockerfile
func createDefaultDockerfile(path, image string) error {
	d1 := []byte(image)
	err := ioutil.WriteFile(filepath.Join(path, "Dockerfile"), d1, 0644)
	if err != nil {
		return fmt.Errorf("unable to create dockerfile: %v", err)
	}
	return nil
}

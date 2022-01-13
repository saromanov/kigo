package service

import (
	"fmt"
	"io/ioutil"
)


// creating of default dockerfile
func createDefaultDockerfile(image string) error {
	d1 := []byte(image)
	err := ioutil.WriteFile("Dockerfile", d1, 0644)
	if err != nil {
		return fmt.Errorf("unable to create dockerfile: %v", err)
	}
	return nil
}
package service

import (
	"context"
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Generate defines interface for generation service
type Generate interface {
	Run(ctx context.Context) error 
}
// service provides generation of the service
type service struct {
	cfg Config
}

// New provides initialization of the service generation
func New(cfg Config) *service {
	return &service {
		cfg: cfg,
	}
}


// Run starts generation of the service files
// Also, it starts to generate suppotted files like Dockerfile
func (s *service) Run(ctx context.Context) error {
	if s.cfg.Name == "" {
		return errors.New("name of the service is not defined")
	}

	cmdPath, err := createDir(s.cfg.Name, "cmd")
	if err != nil {
		return err
	}
	_, err = createDir(s.cfg.Name, "internal")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath.Join(cmdPath, "main.go"), []byte("package main\n\nfunc main() {\n\n}\n"), 0644); err != nil {
		return err
	}

	if err := createDefaultDockerfile(cmdPath, s.cfg.BaseImage); err != nil {
		return err
	}
	return nil
}

func createDir(dirname, name string) (string, error) {
	mainPath := filepath.Join(dirname, name)
	if err := os.MkdirAll(mainPath, fs.ModeDir); err != nil {
		return "", err
	}
	return mainPath, nil
}
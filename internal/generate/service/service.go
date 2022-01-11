package service

import (
	"context"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

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


// Run starts generation of the service
func (s *service) Run(ctx context.Context) error {
	if s.cfg.Name == "" {
		return errors.New("name of the service is not defined")
	}
	mainPath := filepath.Join(s.cfg.Name, "cmd")
	if err := os.MkdirAll(mainPath, fs.ModeDir); err != nil {
		return err
	}
	if _, err := os.Create(filepath.Join(mainPath, "main.go")); err != nil {
		return err
	}
	return nil
}
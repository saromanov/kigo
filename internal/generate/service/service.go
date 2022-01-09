package service

import (
	"context"
	"io/fs"
	"os"
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

func (s *service) Run(ctx context.Context) error {
	if err := os.Mkdir(s.cfg.Name, fs.ModeDir); err != nil {
		return err
	}
	return nil
}
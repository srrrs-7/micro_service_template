package domain

import (
	"context"
	"template/pkg/sftp/domain/repository"
	"template/util/env"
)

type Service struct {
	env  *env.Env
	sftp repository.SftpHandler
}

func NewService(
	env *env.Env,
	sftp repository.SftpHandler,
) *Service {
	return &Service{
		env:  env,
		sftp: sftp,
	}
}

func (s Service) Upload(ctx context.Context) error {
	return nil
}

func (s Service) Download(ctx context.Context) error {
	return nil
}

package awsS3

import (
	"context"
	"template/pkg/sftp/domain/repository"
	"template/util/env"

	"github.com/aws/aws-sdk-go/service/s3"
)

type Repository struct {
	env    *env.Env
	client *s3.S3
}

func NewRepository(
	env *env.Env,
	client *s3.S3,
) repository.SftpHandler {
	return &Repository{
		env:    env,
		client: client,
	}
}

func (r *Repository) Upload(ctx context.Context) error {
	return nil
}
func (r *Repository) Download(ctx context.Context) error {
	return nil
}

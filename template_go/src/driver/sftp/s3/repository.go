package awsS3

import (
	"template/pkg/domain"
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
) *Repository {
	return &Repository{
		env:    env,
		client: client,
	}
}

func (r *Repository) Upload() ([]*domain.Entity, error) {
	return nil, nil
}
func (r *Repository) Download() ([]*domain.Entity, error) {
	return nil, nil
}

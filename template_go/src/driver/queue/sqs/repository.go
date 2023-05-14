package awsSqs

import (
	"template/pkg/domain"
	"template/util/env"

	"github.com/aws/aws-sdk-go/service/sqs"
)

type Repository struct {
	env    *env.Env
	client *sqs.SQS
}

func NewRepository(
	env *env.Env,
	client *sqs.SQS,
) *Repository {
	// SQSクライアントを作成
	return &Repository{
		env:    env,
		client: client,
	}
}

func (r *Repository) SendMessage() ([]*domain.Entity, error) {
	return nil, nil
}
func (r *Repository) ReceiveMessage() ([]*domain.Entity, error) {
	return nil, nil
}
func (r *Repository) DeleteMessage() ([]*domain.Entity, error) {
	return nil, nil
}
func (r *Repository) CreateQueue() ([]*domain.Entity, error) {
	return nil, nil
}
func (r *Repository) DeleteQueue() ([]*domain.Entity, error) {
	return nil, nil
}
func (r *Repository) GetQueueAttributes() ([]*domain.Entity, error) {
	return nil, nil
}
func (r *Repository) SetQueueAttributes() ([]*domain.Entity, error) {
	return nil, nil
}

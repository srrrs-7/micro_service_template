package main

import (
	"template/driver/api"
	"template/driver/api/container"
	"template/driver/db"
	"template/driver/db/table"
	"template/driver/queue"
	awsSqs "template/driver/queue/sqs"
	"template/driver/sftp"
	awsS3 "template/driver/sftp/s3"
	"template/pkg/domain"
	"template/util/aws"
	"template/util/env"
)

func main() {
	env := env.SetEnv()
	// DB connection
	tx := db.NewDb(env)
	db.CloseDb(tx)
	// aws session
	sess := aws.NewAwsSession(env)
	sqs := queue.NewSqs(env, sess)
	s3 := sftp.NewS3(env, sess)
	// repository
	store := table.NewRepository(env, tx)
	server := container.NewRepository()
	queue := awsSqs.NewRepository(env, sqs)
	sftp := awsS3.NewRepository(env, s3)
	// service
	domainService := domain.NewService(store, server, queue, sftp)
	// router
	api.NewRouter(env, domainService)
}

package main

import (
	"flag"
	"template/driver"
	"template/driver/api"
	"template/driver/db"
	"template/driver/db/table"
	"template/driver/db/table2"
	"template/driver/queue"
	awsSqs "template/driver/queue/sqs"
	"template/driver/sftp"
	awsS3 "template/driver/sftp/s3"
	"template/pkg"
	apiService "template/pkg/api/domain"
	dbService "template/pkg/db/domain"
	db2Service "template/pkg/db/domain2"
	queueService "template/pkg/queue/domain"
	sftpService "template/pkg/sftp/domain"
	"template/util/aws"
	"template/util/env"
)

func init() {
	flag.Parse()
}

func main() {
	env := env.SetEnv()
	// api client
	api := api.NewClient(&api.Auth{
		AuthUrl: env.AuthUrl,
		Url:     env.Url,
		ID:      env.Id,
		Secret:  env.Secret,
		Token:   env.Token,
	})
	// DB connection
	tx := db.NewDb(env)
	db.CloseDb(tx)
	// aws session
	sess := aws.NewAwsSession(env)
	sqs := queue.NewSqs(env, sess)
	s3 := sftp.NewS3(env, sess)
	// repository
	client := api.NewRepository(api)
	store := table.NewRepository(env, tx)
	store2 := table2.NewRepository(env, tx)
	queue := awsSqs.NewRepository(env, sqs)
	sftp := awsS3.NewRepository(env, s3)
	// service
	apiService := apiService.NewService(client)
	dbService := dbService.NewService(store)
	db2Service := db2Service.NewService(store2)
	queueService := queueService.NewService(queue)
	sftpService := sftpService.NewService(sftp)

	// logic
	logicRepo := pkg.NewRepositories(
		apiService,
		dbService,
		db2Service,
		queueService,
		sftpService,
	)

	// new router
	driver.NewRouter(env, *logicRepo)
}

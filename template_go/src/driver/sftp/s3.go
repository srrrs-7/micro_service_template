package sftp

import (
	"template/util/env"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewS3(env *env.Env, sess *session.Session) *s3.S3 {
	return s3.New(sess)
}

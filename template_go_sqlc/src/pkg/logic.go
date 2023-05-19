package pkg

import (
	"net/http"
	api "template/pkg/api/domain/usecase"
	db "template/pkg/db/table/usecase"
	queue "template/pkg/queue/domain/usecase"
	sftp "template/pkg/sftp/domain/usecase"
)

type Repositories struct {
	api   api.Client
	db    db.Store
	queue queue.Queuer
	sftp  sftp.SftpHandler
}

func NewRepositories(
	api api.Client,
	db db.Store,
	queue queue.Queuer,
	sftp sftp.SftpHandler,
) *Repositories {
	return &Repositories{
		api:   api,
		db:    db,
		queue: queue,
		sftp:  sftp,
	}
}

func (repo *Repositories) MainLogic(w http.ResponseWriter, r *http.Request) {

}

func (repo *Repositories) SecondLogic(w http.ResponseWriter, r *http.Request) {

}

func (repo *Repositories) SubLogic(w http.ResponseWriter, r *http.Request) {

}

package domain

import (
	"net/http"
)

type Service struct {
	store  Store
	server Server
	queue  QueueHandler
	sftp   SftpHandler
}

func NewService(
	store Store,
	server Server,
	queue QueueHandler,
	sftp SftpHandler,
) *Service {
	return &Service{
		store:  store,
		server: server,
		queue:  queue,
		sftp:   sftp,
	}
}

func (s *Service) Logic(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello world"))
}

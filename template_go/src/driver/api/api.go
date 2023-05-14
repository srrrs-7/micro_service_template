lepackage api

import (
	"fmt"
	"net/http"
	"template/pkg/domain"
	"template/util/env"
	"template/util/log"

	"github.com/gorilla/mux"
)

func NewRouter(
	env *env.Env,
	domain *domain.Service,
	// domain domain.Service, 複数のサービスを受け付ける
	// domain domain.Service,
	// domain domain.Service,
) {
	r := mux.NewRouter()
	r.Use(log.NewLogging)

	r.HandleFunc("/domain/", domain.Logic).Methods("GET")

	fmt.Println("start server on port:", env.HttpPort)
	err := http.ListenAndServe(":"+env.HttpPort, r)
	if err != nil {
		panic(err)
	}
}

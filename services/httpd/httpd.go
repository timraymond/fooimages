package httpd

import "net/http"

type Request struct {
	Text string
}

var _ http.Handler = &Service{}

type Service struct {
	mux *http.ServeMux
}

func (s *Service) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.mux.ServeHTTP(rw, req)
}

func NewService() *Service {
	mux := http.NewServeMux()

	mux.HandleFunc("/images", func(rw http.ResponseWriter, req *http.Request) {
	})

	return &Service{mux}
}

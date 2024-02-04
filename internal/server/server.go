package server

import "time"

type HTTPServer struct {
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
}

func NewServer(s HTTPServer) *HTTPServer {
	return &HTTPServer{
		WriteTimeout: s.WriteTimeout,
		ReadTimeout:  s.ReadTimeout,
	}
}

package autoria

import (
	"net/http"
	"time"
)

type Service struct {
	client *http.Client
	apikey string
	debug  bool
}

type Opts struct {
	APIKey  string
	Debug   bool
	Timeout time.Duration
}

func New(opts Opts) *Service {
	return &Service{
		apikey: opts.APIKey,
		client: &http.Client{
			Timeout: opts.Timeout,
		},
		debug: opts.Debug,
	}
}

func (s *Service) GetApiKey() string {
	return s.apikey
}

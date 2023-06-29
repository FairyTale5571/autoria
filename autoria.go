package autoria

import (
	"net/http"
	"sync"
	"time"
)

type Provider interface {
	GetApiKey() string

	GetCategories() (Categories, error)
	GetBodyStyles(parentID int) ([]BaseWithParentID, error)
	GetBodyStylesWithGroups(parentID int) ([][]BaseWithParentID, error)
	GetMarksByCategory(categoryID int) (Marks, error)
	GetModelsByCategoryAndMarkID(categoryID, markID int) (Models, error)
	GetModelsByCategoryAndMarkIDWithGroups(categoryID, markID int) (Models, error)
	GetGenerationsByModelID(modelID int) ([]Generations, error)
}

type service struct {
	apikey     string
	client     *http.Client
	mutex      sync.Mutex
	maxRetries int
	debug      bool
}

type Opts struct {
	APIKey     string
	MaxRetries int
	Timeout    time.Duration
	Debug      bool
}

func New(opts Opts) Provider {
	return &service{
		apikey: opts.APIKey,
		client: &http.Client{
			Timeout: opts.Timeout,
		},
		mutex:      sync.Mutex{},
		maxRetries: opts.MaxRetries,
		debug:      opts.Debug,
	}
}

func (s *service) GetApiKey() string {
	return s.apikey
}

package localtunnel

import (
	"net/http"
	"time"
)

// DefaultBaseURL is the default value for the Options.BaseURL
var DefaultBaseURL = "https://localtunnel.me"

// DefaultMaxConnections is the default value for Options.MaxConnections
var DefaultMaxConnections = 10

// DefaultLogger is the default value for Options.Log, it does nothing.
var DefaultLogger Logger = logger{}

// Logger is implemented by built-in log.Logger as well as logrus.Entry
type Logger interface {
	Println(v ...interface{})
}

// Options for connecting to a localtunnel server
type Options struct {
	Subdomain      string
	BaseURL        string
	MaxConnections int
	Log            Logger
	HTTP           *http.Client
}

func (o *Options) setDefaults() {
	if o.BaseURL == "" {
		o.BaseURL = DefaultBaseURL
	}
	if o.MaxConnections == 0 {
		o.MaxConnections = DefaultMaxConnections
	}
	if o.Log == nil {
		o.Log = DefaultLogger
	}
	if o.HTTP == nil {
		o.HTTP = http.Client{Timeout: 30 * time.Second}
	}
}

type logger struct{}

func (logger) Println(...interface{}) {}

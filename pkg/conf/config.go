package conf

import (
	"go.opencensus.io/trace"
	"github.com/FirYuen/scaffold/pkg/log"
	"github.com/FirYuen/scaffold/pkg/storage/orm"
	"time"
)

type Config struct {
	App    AppConfig
	Http   ServerConfig
	Grpc   ServerConfig
	Web    WebConfig


	// component config
	Logger  log.Config
	MySQL   orm.Config
	Trace   trace.Config

}
// AppConfig app config
type AppConfig struct {
	Name              string
	Version           string
	Mode              string
	PprofPort         string
	URL               string
	JwtSecret         string
	JwtTimeout        int
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
}

// ServerConfig server config
type ServerConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
// WebConfig web config
type WebConfig struct {
	Name   string
	Domain string
	Secret string
	Static string
}
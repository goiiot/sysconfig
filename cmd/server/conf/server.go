package conf

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

const (
	cfgServerWebRoot         = "server.web.home"
	cfgServerWebHTTPListen   = "server.web.http.listen"
	cfgServerWebHTTPSListen  = "server.web.https.listen"
	cfgServerWebHTTPSTLSCert = "server.web.https.tls_cert"
	cfgServerWebHTTPSTLSKey  = "server.web.https.tls_key"
)

var serverFlags = []Flag{
	stringFlag(&config.Server.Web.Root, cfgServerWebRoot, "", "", "server web root"),
	stringFlag(&config.Server.Web.HTTP.Listen, cfgServerWebHTTPListen, "", ":8080", "server web http listen address"),
	stringFlag(&config.Server.Web.HTTPS.Listen, cfgServerWebHTTPSListen, "", ":8443", "server web https listen address"),
	stringFlag(&config.Server.Web.HTTPS.TLSCert, cfgServerWebHTTPSTLSCert, "", "", "server web tls cert"),
	stringFlag(&config.Server.Web.HTTPS.TLSKey, cfgServerWebHTTPSTLSKey, "", "", "server web tls key"),
}

// ServerWebHttpsConfig server.web.https
type ServerWebHttpsConfig struct {
	Enabled bool   `yaml:"enabled"`
	Listen  string `yaml:"listen"`
	TLSCert string `yaml:"tls_cert"`
	TLSKey  string `yaml:"tls_key"`
}

// ServerWebHttpConfig server.web.http
type ServerWebHttpConfig struct {
	Enabled bool   `yaml:"enabled"`
	Listen  string `yaml:"listen"`
}

type ServerAuthUser struct {
	Username string   `yaml:"username" form:"username" json:"username" binding:"required"`
	Password string   `yaml:"password" form:"password" json:"password" binding:"required"`
	Services []string `yaml:"services" form:"-" json:"-" binding:"-"`

	capReg *regexp.Regexp
}

func (u *ServerAuthUser) initCapability() {
	if u == nil || u.Services == nil {
		return
	}

	for i, v := range u.Services {
		if v == "all" {
			u.Services[i] = ".*"
		}
	}

	u.capReg = regexp.MustCompile(fmt.Sprintf(`^https?://.*?/api/v\d*/(%s)`, strings.Join(u.Services, "|")))
}

func (u *ServerAuthUser) CapableOf(url string) bool {
	if u == nil || u.capReg == nil {
		return false
	}

	ok := u.capReg.MatchString(url)
	return ok
}

// ServerAuthConfig server.auth
type ServerAuthConfig struct {
	Enabled        bool             `yaml:"enabled"`
	SecretKey      string           `yaml:"secret_key"`
	SessionTimeout time.Duration    `yaml:"session_timeout"`
	Users          []ServerAuthUser `yaml:"users"`
}

// ServerWebConfig server.web
type ServerWebConfig struct {
	Root     string               `yaml:"root"`
	Username string               `yaml:"username"`
	Password string               `yaml:"password"`
	HTTP     ServerWebHttpConfig  `yaml:"http"`
	HTTPS    ServerWebHttpsConfig `yaml:"https"`
}

// ServerConfig match server config block in yaml
type ServerConfig struct {
	Web  ServerWebConfig  `yaml:"web"`
	Auth ServerAuthConfig `yaml:"auth"`
}

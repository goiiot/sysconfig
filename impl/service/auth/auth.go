package auth

import (
	"strings"
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"github.com/goiiot/sysconfig/cmd/server/conf"
	"github.com/goiiot/sysconfig/impl/log"
	"github.com/goiiot/sysconfig/impl/service/utils"
	JWT "gopkg.in/dgrijalva/jwt-go.v3"
)

const identityKey = "id"

// read only after initialization
var users map[string]*conf.ServerAuthUser

func createAuthMiddleware(config *conf.ServerAuthConfig) *jwt.GinJWTMiddleware {
	users = make(map[string]*conf.ServerAuthUser)
	for i, v := range config.Users {
		users[v.Username] = &config.Users[i]
	}

	return &jwt.GinJWTMiddleware{
		Realm:            "test jwt",
		SigningAlgorithm: "HS256",
		Key:              []byte(config.SecretKey),
		Timeout:          config.SessionTimeout,
		TokenLookup:      "cookie:token",
		Authenticator: func(c *gin.Context) (interface{}, error) {
			user := new(conf.ServerAuthUser)
			if err := c.ShouldBind(user); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			if u, ok := users[user.Username]; ok && u.Password == user.Password {
				return u, nil
			}
			return nil, jwt.ErrForbidden
		},
		IdentityHandler: func(claims JWT.MapClaims) interface{} {
			if username, ok := claims["id"].(string); ok {
				return users[username]
			}
			return nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if data == nil {
				return false
			}
			if user, ok := data.(*conf.ServerAuthUser); ok {
				scheme := "http://"
				if c.Request.TLS != nil {
					scheme = "https://"
				}
				return user.CapableOf(strings.Join([]string{scheme, c.Request.Host, c.Request.RequestURI}, ""))
			}
			return false
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if data == nil {
				return jwt.MapClaims{}
			}

			if u, ok := data.(*conf.ServerAuthUser); ok {
				return jwt.MapClaims{
					identityKey: u.Username,
				}
			}
			return jwt.MapClaims{}
		},
		Unauthorized: func(c *gin.Context, code int, reason string) {
			utils.RespErrJSON(c, code, code, reason)
			log.D("unauthorized user", zap.String("reason", reason))
		},
		LoginResponse: func(c *gin.Context, code int, token string, until time.Time) {
			c.SetCookie("token", token, int(config.SessionTimeout.Seconds()), "/", strings.SplitN(c.Request.Host, ":", 2)[0], false, false)
			utils.RespOkJSON(c)
			log.D("user login success", zap.String("token", token))
		},
		RefreshResponse: func(c *gin.Context, code int, token string, until time.Time) {
			c.SetCookie("token", token, int(config.SessionTimeout.Seconds()), "/", strings.SplitN(c.Request.Host, ":", 2)[0], false, false)
			utils.RespOkJSON(c)
			log.D("user login refresh success", zap.String("token", token))
		},
	}
}

func InitMiddlewareAuth(v1 *gin.RouterGroup, config *conf.ServerAuthConfig) {
	if !config.Enabled {
		return
	}

	m := createAuthMiddleware(config)
	if err := m.MiddlewareInit(); err != nil {
		log.F("init auth middleware failed", false, zap.Error(err))
	}

	v1.POST("/auth", m.LoginHandler)
	v1.Use(m.MiddlewareFunc())
}

package restecho

import (
	"net/http"
	"time"

	"github.com/kiriwill/events-api/pkg/repository/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(svc *mysql.MysqlRepository, authTokenSecret string, jwtDuration time.Duration) *echo.Echo {
	e := echo.New()
	e.Validator = NewValidator()
	e.Use(middleware.Logger())

	auth := NewAuth(authTokenSecret, jwtDuration)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	v1 := e.Group("v1")
	user := v1.Group("/user")
	user.POST("/", PostUserHandler(svc))
	user.POST("/signin", SignInHandler(svc, auth))

	events := v1.Group("/events", auth.CustomJwtMiddleware())
	events.POST("/", PostEventHandler(svc))
	events.GET("/", SearchEventHandler(svc))

	metrics := events.Group("/metrics", auth.CustomJwtMiddleware())
	metrics.GET("/", GetMetricsHandler(svc))

	return e
}

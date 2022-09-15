package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type LogMiddleware struct {
	*echo.Echo
}

func NewLogMiddleware(echo *echo.Echo) *LogMiddleware {
	return &LogMiddleware{Echo: echo}
}

func (m LogMiddleware) LogMiddleWare() {
	m.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339_nano}]: {"id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out},"data": ${query}}` + "\n",
		CustomTimeFormat: "2006-01-02T15:04:05.00000",
	}))
}

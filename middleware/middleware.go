package middleware

import (
	"fmt"
	"net/http"
	"os"

	"test-echo/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

var secret = os.Getenv("ACCESS_SECRET")

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	//SigningKey: []byte(secret),
	SigningKey: []byte("test-apk"),
})

func MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		config.MakeLogEntry(c).Info("incoming request")
		return next(c)
	}
}

func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	config.MakeLogEntry(c).Error(report.Message)
	c.HTML(report.Code, report.Message.(string))
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.WriteHeader(code)
}

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		writer := statusRecorder{w, 0}
		handler.ServeHTTP(&writer, req)
		if writer.statusCode == 0 {
			log.Errorf("Expected status code to be set but it wasn't")
		}
	})
}

var Logger = middleware.Logger()

var Recover = middleware.Recover()

var CORSWithConfig = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: []string{"http://localhost:4200", "*"},
	//AllowOrigins: []string{"*"},
	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	//AllowHeaders: []string{"*"},
	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
})

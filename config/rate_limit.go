package config

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
	"time"
)

func GetRateLimiterConfig() middleware.RateLimiterConfig {
	limit, err := strconv.ParseInt(os.Getenv("RATE_LIMIT_ALLOW"), 10, 64)
	if err != nil {
		limit = 10
	}

	return middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: limit, Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string,err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}

}

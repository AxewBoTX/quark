package handlers

import (
	"github.com/labstack/echo/v4"

	"quark/client/lib"
)

func SessionManagerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session_cookie, session_cookie_get_err := c.Cookie(lib.SESSION_COOKIE_NAME)
		if session_cookie_get_err != nil {
			lib.InfoWithColor("INFO", "0", lib.COLOR_BLUE, "Session Cookie Not Detected")
			return next(c)
		}
		lib.InfoWithColor(
			"INFO",
			"0",
			lib.COLOR_BLUE,
			"Session Cookie Detected",
			"Cookie",
			session_cookie,
		)
		return next(c)
	}
}

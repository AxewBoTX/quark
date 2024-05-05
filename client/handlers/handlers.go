package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"quark/client/lib"
	"quark/client/web"
	"quark/client/web/routes"
)

// index(/) route handler
func IndexHandler(c echo.Context) error {
	TemplateRenderState := c.Get("TemplateRenderState")
	if TemplateRenderState == true {
		return web.RenderTemplTemplate(c, http.StatusOK, routes.Home_Page())
	} else {
		return nil
	}
}

// (/register) route handler
func RegisterHandler(c echo.Context) error {
	TemplateRenderState := c.Get("TemplateRenderState")
	if TemplateRenderState == true {
		return web.RenderTemplTemplate(c, http.StatusOK, routes.Register_Page())
	} else {
		return nil
	}
}

// (/login) route handler
func LoginHandler(c echo.Context) error {
	TemplateRenderState := c.Get("TemplateRenderState")
	if TemplateRenderState == true {
		return web.RenderTemplTemplate(c, http.StatusOK, routes.Login_Page())
	} else {
		return nil
	}
}

// (/chat) route handler
func ChatHandler(c echo.Context) error {
	TemplateRenderState := c.Get("TemplateRenderState")
	if TemplateRenderState == true {
		c.SetRequest(
			c.Request().
				WithContext(context.WithValue(c.Request().Context(), "realtime_server_addr", "ws://"+lib.SERVER_HOST+lib.SERVER_PORT+"/ws")),
		)
		return web.RenderTemplTemplate(c, http.StatusOK, routes.Chat_Page())
	} else {
		return nil
	}
}

package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"quark/client/web"
)

// index(/) route handler
func IndexHandler(c echo.Context) error {
	TemplateRenderState := c.Get("TemplateRenderState")
	if TemplateRenderState == true {
		return web.RenderTemplTemplate(c, http.StatusOK)
	} else {
		return nil
	}
}

// (/register) route handler
func RegisterHandler(c echo.Context) error {
	TemplateRenderState := c.Get("TemplateRenderState")
	if TemplateRenderState == true {
		return web.RenderTemplTemplate(c, http.StatusOK)
	} else {
		return nil
	}
}

// (/login) route handler
func LoginHandler(c echo.Context) error {
	TemplateRenderState := c.Get("TemplateRenderState")
	if TemplateRenderState == true {
		return web.RenderTemplTemplate(c, http.StatusOK)
	} else {
		return nil
	}
}

// (/chat) route handler
func ChatHandler(c echo.Context) error {
	TemplateRenderState := c.Get("TemplateRenderState")
	if TemplateRenderState == true {
		// lib.InfoWithColor(
		// 	"INFO",
		// 	"0",
		// 	lib.COLOR_BLUE,
		// 	"Session Data",
		// 	"Value",
		// 	c.Get("session-data"),
		// )
		return web.RenderTemplTemplate(c, http.StatusOK)
	} else {
		return nil
	}
}

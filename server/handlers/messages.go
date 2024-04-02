package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	MessageListFetchQuery string = `SELECT * FROM %s;`
	MessageFetchQuery     string = `SELECT * FROM %s WHERE id = ? LIMIT 1;`
	MessageInsertQuery    string = `INSERT INTO %s (id,user_id,body,created) VALUE(?,?,?,?);`
	MessageDeleteQuery    string = `DELETE FROM %s WHERE id = ?;`
)

func Messages(router *echo.Group, DB *sql.DB) {
	// (/messages/) route GET request handler
	router.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Message List Fetch")
	})

	// (/messages/:messageID) route GET request handler
	router.GET("/:messageID", func(c echo.Context) error {
		messageID := c.Param("messageID")
		return c.JSON(http.StatusOK, fmt.Sprintf("Message Fetch With ID: %s", messageID))
	})

	// (/messages/) route POST request handler
	router.POST("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Message Insert")
	})

	// (/messages/:messageID) route DELETE request handler
	router.DELETE("/:messageID", func(c echo.Context) error {
		messageID := c.Param("messageID")
		return c.JSON(http.StatusOK, fmt.Sprintf("Message Delete With ID: %s", messageID))
	})
}

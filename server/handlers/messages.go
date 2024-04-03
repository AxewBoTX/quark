package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"quark/server/lib"
)

var (
	MessageListFetchQuery string = `SELECT * FROM %s;`
	MessageFetchQuery     string = `SELECT * FROM %s WHERE id = ? LIMIT 1;`
	MessageInsertQuery    string = `INSERT INTO %s (id,user_id,body,created) VALUES(?,?,?,?);`
	MessageDeleteQuery    string = `DELETE FROM %s WHERE id = ?;`
)

func Messages(router *echo.Group, DB *sql.DB) {
	// (/messages/) route GET request handler
	router.GET("/", func(c echo.Context) error {
		rows, rows_fetch_err := DB.Query(fmt.Sprintf(MessageListFetchQuery, lib.MESSAGE_TABLE_NAME))
		if rows_fetch_err != nil {
			lib.ErrorWithColor(
				"ERROR",
				"0",
				lib.COLOR_RED,
				"Failed To Fetch Database rows",
				"Error",
				rows_fetch_err,
			)
			return c.String(http.StatusInternalServerError, "Database Rows Fetch Error")
		}
		defer func() {
			rows.Close()
		}()

		var messages []lib.Message
		for rows.Next() {
			var message lib.Message
			if row_scan_err := rows.Scan(&message.ID, &message.UserID, &message.Body, &message.Created); row_scan_err != nil {
				lib.ErrorWithColor(
					"ERROR",
					"0",
					lib.COLOR_RED,
					"Failed To Scan Database Row",
					"Error",
					row_scan_err,
				)
				return c.String(http.StatusInternalServerError, "Database Row Scan Error")
			}
			messages = append(messages, message)
		}

		return c.JSON(http.StatusOK, messages)
	})

	// (/messages/:messageID) route GET request handler
	router.GET("/:messageID", func(c echo.Context) error {
		messageID := c.Param("messageID")
		var message lib.Message
		if row_fetch_err := DB.QueryRow(fmt.Sprintf(MessageFetchQuery, lib.MESSAGE_TABLE_NAME), messageID).Scan(
			&message.ID, &message.UserID, &message.Body, &message.Created,
		); row_fetch_err != nil {
			lib.ErrorWithColor(
				"ERROR",
				"0",
				lib.COLOR_RED,
				"Failed To Fetch Database Row",
				"Error",
				row_fetch_err,
			)
			return c.String(http.StatusInternalServerError, "Database Row Fetch Error")
		}
		return c.JSON(http.StatusOK, message)
	})

	// (/messages/) route POST request handler
	router.POST("/", func(c echo.Context) error {
		current_time := time.Now().Format(time.RFC3339)
		var req_message lib.Message
		if req_message_bind_err := c.Bind(&req_message); req_message_bind_err != nil {
			lib.ErrorWithColor(
				"ERROR",
				"0",
				lib.COLOR_RED,
				"Failed To Bind Request Data",
				"Error",
				req_message_bind_err,
			)
			return c.String(http.StatusInternalServerError, "Request Data Bind Error")
		}
		req_message.ID = uuid.New().String()
		req_message.Created = current_time

		if _, row_create_err := DB.Exec(fmt.Sprintf(MessageInsertQuery, lib.MESSAGE_TABLE_NAME),
			req_message.ID, req_message.UserID, req_message.Body, current_time,
		); row_create_err != nil {
			lib.ErrorWithColor(
				"ERROR",
				"0",
				lib.COLOR_RED,
				"Failed To Create Database Row",
				"Error",
				row_create_err,
			)
			return c.String(http.StatusInternalServerError, "Database Row Create Error")
		}

		return c.JSON(http.StatusOK, req_message)
	})

	// (/messages/:messageID) route DELETE request handler
	router.DELETE("/:messageID", func(c echo.Context) error {
		messageID := strings.TrimSpace(c.Param("messageID"))
		if len(messageID) == 0 || messageID == "" {
			lib.ErrorWithColor("ERROR", "0", lib.COLOR_RED, "Request path parameter is nil")
			return c.String(http.StatusBadRequest, "Request path parameter is nil")
		}
		if _, message_delete_err := DB.Exec(fmt.Sprintf(MessageDeleteQuery, lib.MESSAGE_TABLE_NAME), messageID); message_delete_err != nil {
			lib.ErrorWithColor(
				"ERROR",
				"0",
				lib.COLOR_RED,
				"Failed To Delete Database Row",
				"Error",
				message_delete_err,
			)
			return c.String(http.StatusInternalServerError, "Database Row Delete Error")
		}
		return c.String(http.StatusOK, "SUCCESS")
	})
}

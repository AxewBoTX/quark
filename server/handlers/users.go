package handlers

import (
	"database/sql"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"quark/server/lib"
)

func Users(router *echo.Group, DB *sql.DB) {
	// (/users/) route GET request handler
	router.GET("/", func(c echo.Context) error {
		rows, rows_fetch_err := DB.Query("SELECT * FROM usr_base")
		if rows_fetch_err != nil {
			log.Error("Failed To Fetch usr_base rows", "Error", rows_fetch_err)
			return c.String(http.StatusInternalServerError, "Database Rows Fetch Error")
		}
		defer func() {
			rows.Close()
		}()

		var users []lib.User
		for rows.Next() {
			var user lib.User
			if row_scan_err := rows.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.UserAuthToken); row_scan_err != nil {
				log.Error("Failed To Scan usr_base Row", "Error", row_scan_err)
				return c.String(http.StatusInternalServerError, "Database Row Scan Error")
			}
			users = append(users, user)
		}

		return c.JSON(http.StatusOK, users)
	})

	// (/users/:userID) route GET request handler
	router.GET("/:userID", func(c echo.Context) error {
		userID := c.Param("userID")
		var user lib.User
		if row_fetch_err := DB.QueryRow("SELECT * FROM usr_base WHERE id = ? LIMIT 1", userID).Scan(
			&user.ID, &user.Username, &user.PasswordHash, &user.UserAuthToken,
		); row_fetch_err != nil {
			log.Error("Failed To Fetch Database Row", "Error", row_fetch_err)
			return c.String(http.StatusInternalServerError, "Database Row Fetch Error")
		}
		return c.JSON(http.StatusOK, user)
	})

	// (/users/) route POST request handler
	router.POST("/", func(c echo.Context) error {
		var req_user lib.User
		if req_user_bind_err := c.Bind(&req_user); req_user_bind_err != nil {
			log.Error("Failed To Bind Request Data", "Error", req_user_bind_err)
			return c.String(http.StatusInternalServerError, "Request Data Bind Error")
		}
		req_user.ID = uuid.New().String()

		if _, row_create_err := DB.Exec(
			`INSERT INTO usr_base (id,username,passwordHash,userAuthToken) VALUES(?,?,?,?)`,
			req_user.ID, req_user.Username, req_user.PasswordHash, req_user.UserAuthToken,
		); row_create_err != nil {
			log.Error("Failed To Create Database Row", row_create_err)
			return c.String(http.StatusInternalServerError, "Database Row Create Error")
		}

		return c.JSON(http.StatusOK, req_user)
	})

	// (/users/:userID) route PATCH request handler
	router.PATCH("/:userID", func(c echo.Context) error {
		return c.String(http.StatusOK, "Update User with ID "+c.Param("userID"))
	})

	// (/users/:userID) route DELETE request handler
	router.DELETE("/:userID", func(c echo.Context) error {
		return c.String(http.StatusOK, "DELETE User with ID "+c.Param("userID"))
	})
}

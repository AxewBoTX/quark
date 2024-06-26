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

// database queries
var (
	UserListFetchQuery            string = `SELECT * FROM %s;`
	UserFetchQuery_Username       string = `SELECT * FROM %s WHERE username = ? LIMIT 1;`
	UserFetchQuery_ID             string = `SELECT * FROM %s WHERE id = ? LIMIT 1;`
	UserFetchQuerry_UserAuthToken string = `SELECT * FROM %s WHERE userAuthToken = ? LIMIT 1;`
	UserInsertQuery               string = `INSERT INTO %s (id,username,passwordHash,userAuthToken,created,updated) VALUES(?,?,?,?,?,?);`
	UserDeleteQuery               string = `DELETE FROM %s WHERE id = ?;`
)

func Users(router *echo.Group, DB *sql.DB) {
	// (/users/) route GET request handler
	router.GET("/", func(c echo.Context) error {
		// fetch user rows from database
		rows, rows_fetch_err := DB.Query(fmt.Sprintf(UserListFetchQuery, lib.USER_TABLE_NAME))
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

		// scan the rows into an array
		var users []lib.User
		for rows.Next() {
			var user lib.User
			if row_scan_err := rows.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.UserAuthToken, &user.Created, &user.Updated); row_scan_err != nil {
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
			users = append(users, user)
		}

		return c.JSON(http.StatusOK, users)
	})

	// (/users/token/:token) route GET request handler
	router.GET("/token/:token", func(c echo.Context) error {
		token := c.Param("token")
		if len(token) == 0 || token == "" {
			lib.ErrorWithColor("ERROR", "0", lib.COLOR_RED, "Request path parameter is nil")
			return c.String(http.StatusBadRequest, "Token must not be nil")
		}
		var user lib.User
		// fetch database row and scan into struct
		if row_fetch_err := DB.QueryRow(fmt.Sprintf(UserFetchQuerry_UserAuthToken, lib.USER_TABLE_NAME), token).Scan(
			&user.ID, &user.Username, &user.PasswordHash, &user.UserAuthToken, &user.Created, &user.Updated,
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
		user.PasswordHash = ""
		return c.JSON(http.StatusOK, user)
	})

	// (/users/:username) route GET request handler
	router.GET("/:username", func(c echo.Context) error {
		username := c.Param("username")
		if len(username) == 0 || username == "" {
			lib.ErrorWithColor("ERROR", "0", lib.COLOR_RED, "Request path parameter is nil")
			return c.String(http.StatusBadRequest, "Username must not be nil")
		}
		var user lib.User
		// fetch database row and scan into struct
		if row_fetch_err := DB.QueryRow(fmt.Sprintf(UserFetchQuery_Username, lib.USER_TABLE_NAME), username).Scan(
			&user.ID, &user.Username, &user.PasswordHash, &user.UserAuthToken, &user.Created, &user.Updated,
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
		return c.JSON(http.StatusOK, user)
	})

	// (/users/) route POST request handler
	router.POST("/", func(c echo.Context) error {
		current_time := time.Now().Unix()
		var req_user lib.User

		// bind request data to struct
		if req_user_bind_err := c.Bind(&req_user); req_user_bind_err != nil {
			lib.ErrorWithColor(
				"ERROR",
				"0",
				lib.COLOR_RED,
				"Failed To Bind Request Data",
				"Error",
				req_user_bind_err,
			)
			return c.String(http.StatusInternalServerError, "Request Data Bind Error")
		}
		req_user.ID = uuid.New().String()
		req_user.Created = current_time
		req_user.Updated = current_time

		// insert database row
		if _, row_create_err := DB.Exec(fmt.Sprintf(UserInsertQuery, lib.USER_TABLE_NAME),
			req_user.ID, req_user.Username, req_user.PasswordHash, req_user.UserAuthToken, current_time, current_time,
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

		return c.JSON(http.StatusOK, req_user)
	})

	// (/users/:userID) route PATCH request handler
	router.PATCH("/:userID", func(c echo.Context) error {
		UserID := c.Param("userID")
		if len(UserID) == 0 || UserID == "" {
			lib.ErrorWithColor("ERROR", "0", lib.COLOR_RED, "Request path parameter is nil")
			return c.String(http.StatusBadRequest, "UserID must not be nil")
		}

		// bind request data into a struct
		var req_user lib.User
		if req_user_bind_err := c.Bind(&req_user); req_user_bind_err != nil {
			lib.ErrorWithColor(
				"ERROR",
				"0",
				lib.COLOR_RED,
				"Failed To Bind Request Data",
				"Error",
				req_user_bind_err,
			)
			return c.String(http.StatusInternalServerError, "Request Data Bind Error")
		}
		req_user.ID = UserID

		// update database row
		if _, user_update_err := DB.Exec(lib.GenerateSQLUpdateQuery(req_user)); user_update_err != nil {
			lib.ErrorWithColor(
				"ERROR",
				"0",
				lib.COLOR_RED,
				"Failed To Update Databse Row",
				"Error",
				user_update_err,
			)
			return c.String(http.StatusInternalServerError, "Database Row Update Error")
		}

		// fetch updated database row
		var up_user lib.User
		if user_fetch_err := DB.QueryRow(fmt.Sprintf(UserFetchQuery_ID, lib.USER_TABLE_NAME), UserID).
			Scan(
				&up_user.ID,
				&up_user.Username,
				&up_user.PasswordHash,
				&up_user.UserAuthToken,
				&up_user.Created,
				&up_user.Updated,
			); user_fetch_err != nil {
			lib.ErrorWithColor(
				"ERROR",
				"0",
				lib.COLOR_RED,
				"Failed To Fetch Database Row",
				"Error",
				user_fetch_err,
			)
			return c.String(http.StatusInternalServerError, "Database Row Fetch Error")
		}

		return c.JSON(http.StatusOK, up_user)
	})

	// (/users/:userID) route DELETE request handler
	router.DELETE("/:userID", func(c echo.Context) error {
		UserID := strings.TrimSpace(c.Param("userID"))
		if len(UserID) == 0 || UserID == "" {
			lib.ErrorWithColor("ERROR", "0", lib.COLOR_RED, "Request path parameter is nil")
			return c.String(http.StatusBadRequest, "Request path parameter is nil")
		}

		// delete database row
		if _, user_delete_err := DB.Exec(fmt.Sprintf(UserDeleteQuery, lib.USER_TABLE_NAME), UserID); user_delete_err != nil {
			lib.ErrorWithColor(
				"ERROR",
				"0",
				lib.COLOR_RED,
				"Failed To Delete Database Row",
				"Error",
				user_delete_err,
			)
			return c.String(http.StatusInternalServerError, "Database Row Delete Error")
		}
		return c.String(http.StatusOK, "SUCCESS")
	})
}

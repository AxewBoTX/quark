package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"quark/client/lib"
)

// (/auth/register) route handler
func AuthRegisterHandler(c echo.Context) error {
	auth_client := resty.New()

	username := strings.TrimSpace(c.FormValue("username"))
	password := strings.TrimSpace(c.FormValue("password"))
	var passwordHash string
	userAuthToken := uuid.New().String()
	userID := uuid.New().String()
	password_confirm := strings.TrimSpace(c.FormValue("passwordConfirm"))

	// match password and passwordConfirm and then hash password
	if password == password_confirm {
		if hash, pass_hash_err := lib.HashString(password); pass_hash_err != nil {
			lib.ErrorWithColor(
				"ERROR",
				"0",
				lib.COLOR_RED,
				"Failed To Hash String",
				"Error",
				pass_hash_err,
			)
			return c.JSON(http.StatusInternalServerError, "FAIL_RGS")
		} else {
			passwordHash = hash
		}
	} else {
		lib.ErrorWithColor(
			"ERROR",
			"0",
			lib.COLOR_RED,
			"Bad Request",
		)
		return c.JSON(http.StatusBadRequest, "FAIL_RGS")
	}

	// post register data to the server
	res, user_register_err := auth_client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"id":            userID,
			"username":      username,
			"passwordHash":  passwordHash,
			"userAuthToken": userAuthToken,
		}).
		Post(lib.SERVER_HOST + lib.SERVER_PORT + "/users/")
	if user_register_err != nil {
		lib.ErrorWithColor(
			"ERROR",
			"0",
			lib.COLOR_RED,
			"Failed To Send Data To Server",
			"Error",
			user_register_err,
		)
		return c.JSON(http.StatusInternalServerError, "FAIL_RGS")
	}

	// handle server response
	var user lib.User
	if resp_decode_err := json.Unmarshal(res.Body(), &user); resp_decode_err != nil {
		lib.ErrorWithColor(
			"ERROR",
			"0",
			lib.COLOR_RED,
			"Failed To Decode Server Response",
			"Error",
			resp_decode_err,
		)
		return c.JSON(http.StatusInternalServerError, "FAIL_RGS")
	}
	return c.JSON(http.StatusOK, "PASS_RGS")
}

// (/auth/login) route handler
func AuthLoginHandler(c echo.Context) error {
	auth_client := resty.New()
	username := strings.TrimSpace(c.FormValue("username"))
	password := strings.TrimSpace(c.FormValue("password"))

	// fetch user data from server
	res, user_login_err := auth_client.R().
		Get(lib.SERVER_HOST + lib.SERVER_PORT + "/users/" + username)
	if user_login_err != nil {
		lib.ErrorWithColor(
			"ERROR",
			"0",
			lib.COLOR_RED,
			"Failed To Get Data From Server",
			"Error",
			user_login_err,
		)
		return c.JSON(http.StatusInternalServerError, "FAIL_LGN")
	}

	// handle server response
	var user lib.User
	if resp_decode_err := json.Unmarshal(res.Body(), &user); resp_decode_err != nil {
		lib.ErrorWithColor(
			"ERROR",
			"0",
			lib.COLOR_RED,
			"Failed To Decode Server Response",
			"Error",
			resp_decode_err,
		)
		return c.JSON(http.StatusInternalServerError, "FAIL_LGN")
	}

	// verify password && update UserAuthToken
	if lib.CheckStringHash(password, user.PasswordHash) == true {
		res, user_auth_token_update_err := auth_client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(map[string]interface{}{
				"userAuthToken": uuid.New().String(),
			}).
			Patch(lib.SERVER_HOST + lib.SERVER_PORT + "/users/" + user.ID)
		if user_auth_token_update_err != nil {
			lib.ErrorWithColor(
				"ERROR",
				"0",
				lib.COLOR_RED,
				"Failed To Send Data To Server",
				"Error",
				user_auth_token_update_err,
			)
			return c.JSON(http.StatusInternalServerError, "FAIL_LGN")
		}

		// handle server response
		var up_user lib.User
		if resp_decode_err := json.Unmarshal(res.Body(), &up_user); resp_decode_err != nil {
			lib.ErrorWithColor(
				"ERROR",
				"0",
				lib.COLOR_RED,
				"Failed To Decode Server Response",
				"Error",
				resp_decode_err,
			)
			return c.JSON(http.StatusInternalServerError, "FAIL_LGN")
		}

		// set session cookie
		c.SetCookie(&http.Cookie{
			Name:    "usr_session",
			Value:   up_user.UserAuthToken,
			Path:    "/",
			Expires: time.Now().Add(90 * 24 * time.Hour),
		})
	} else {
		lib.ErrorWithColor(
			"ERROR",
			"0",
			lib.COLOR_RED,
			"Wrong Password",
		)
		return c.JSON(http.StatusInternalServerError, "FAIL_LGN_PASS")
	}

	return c.JSON(http.StatusOK, "PASS_LGN")
}

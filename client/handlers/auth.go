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

	pass_toast := `<div id="registerToast" role="alert" class="alert alert-success max-w-[300px] fixed top-5 right-5">
		<svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
		<span>Your account was successfully registered</span>
	</div>`
	fail_toast := `<div id="registerToast" role="alert" class="alert alert-error max-w-[300px] fixed top-5 right-5">
		<svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
		<span>Registration Failed</span>
	</div>`

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
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status_code": "FAIL_RGS",
				"toast":       fail_toast,
			})
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
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status_code": "FAIL_RGS",
			"toast":       fail_toast,
		})
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
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status_code": "FAIL_RGS",
			"toast":       fail_toast,
		})
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
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status_code": "FAIL_RGS",
			"toast":       fail_toast,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status_code": "PASS_RGS",
		"toast":       pass_toast,
	})
}

// (/auth/login) route handler
func AuthLoginHandler(c echo.Context) error {
	auth_client := resty.New()
	username := strings.TrimSpace(c.FormValue("username"))
	password := strings.TrimSpace(c.FormValue("password"))

	fail_toast := `<div id="loginToast" role="alert" class="alert alert-error max-w-[300px] fixed top-5 right-5">
		<svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
		<span>Failed to Login</span>
	</div>`
	fail_toast_pass := `<div id="loginToast" role="alert" class="alert alert-error max-w-[300px] fixed top-5 right-5">
		<svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
		<span>Wrong Password</span>
	</div>`

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
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status_code": "FAIL_LGN",
		})
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
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status_code": "FAIL_LGN",
			"toast":       fail_toast,
		})
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
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status_code": "FAIL_LGN",
				"toast":       fail_toast,
			})
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
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status_code": "FAIL_LGN",
				"toast":       fail_toast,
			})
		}

		// set session cookie
		c.SetCookie(&http.Cookie{
			Name:    lib.SESSION_COOKIE_NAME,
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
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status_code": "FAIL_LGN_PASS",
			"toast":       fail_toast_pass,
		})
	}
	c.Response().Header().Set("HX-Redirect", "/chat")
	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"status_code": "PASS_LGN",
	})
}

// (/auth/logout) route handler
func AuthLogoutHandler(c echo.Context) error {
	// remove the session cookie
	c.SetCookie(&http.Cookie{
		Name:   lib.SESSION_COOKIE_NAME,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	// redirect user to /login route
	c.Response().Header().Set("HX-Redirect", "/login")
	return c.JSON(http.StatusOK, "PASS_LGO")
}

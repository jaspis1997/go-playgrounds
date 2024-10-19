package line

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

type (
	WebEngine  = *gin.Engine
	WebContext = *gin.Context
)

func NewServer() WebEngine {
	return initWebRoutes(gin.Default())
}

func initWebRoutes(e WebEngine) WebEngine {
	e.LoadHTMLFiles("templates/signup.html")
	e.Static("/public", "./public")

	e.GET("/", func(ctx WebContext) {
		ctx.HTML(http.StatusOK, "signup.html", "")
	})
	e.GET("/auth", func(ctx WebContext) {
		const format_url = "https://access.line.me/oauth2/v2.1/authorize?response_type=%s&client_id=%s&redirect_uri=%s&state=%s&scope=%s&bot_prompt=%s"
		var (
			response_type = "code"
			client_id     = os.Getenv(Env_LineClientId)
			redirect_uri  = url.QueryEscape(fmt.Sprintf("http://localhost:8009/callback?token=%s", "example_token"))
			state         = "example12345abcde"
			scope         = "profile"
			bot_prompt    = "aggressive"
		)

		url := fmt.Sprintf(
			format_url,
			response_type,
			client_id,
			redirect_uri,
			state,
			scope,
			bot_prompt,
		)
		log.Print(url)
		ctx.Redirect(http.StatusFound, url)
		// ctx.SetCookie("sessions", "example", 0, "/", "", false, true)
		// ctx.String(http.StatusOK, "%s", url)
	})
	mypage := func(ctx WebContext) {
		ctx.String(http.StatusOK, "username:%s", "username")
	}

	e.GET("/callback", func(ctx WebContext) {
		// ctx.SetCookie("sessions", "example", 0, "/", "", false, true)
		// t, _ := ctx.Cookie("sessions")
		t := ctx.Query("token")
		ctx.String(http.StatusOK, "token:%s", t)
	})
	loginCheck := func(ctx WebContext) {

	}
	e.GET("/mypage", loginCheck, mypage)

	return e
}

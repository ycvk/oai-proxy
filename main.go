package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
	"gopkg.in/yaml.v3"
	"net/http"
	"net/url"
	"slices"
)

var (
	//go:embed config.yml
	configByte []byte
	cfg        Config
)

type Config struct {
	AccessToken  string   `yaml:"access_token"`
	RefreshToken string   `yaml:"refresh_token"`
	SitePassword string   `yaml:"site_password"`
	AllowUsers   []string `yaml:"allow_users"`
}

const (
	refreshURL  = "https://token.oaifree.com/api/auth/refresh"
	registerURL = "https://chat.oaifree.com/token/register"
	oauthURL    = "https://new.oaifree.com/api/auth/oauth_token"
	proxyURL    = "new.oaifree.com"
	// 默认值
	siteLimit         = ""
	expiresIn         = "0"
	gpt35Limit        = "-1"
	gpt4Limit         = "-1"
	showConversations = "false"
	resetLimit        = "false"
)

func init() {
	err := yaml.Unmarshal(configByte, &cfg)
	if err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		token := cfg.AccessToken
		if isTokenExpired(token) {
			err := refreshAccessToken(&cfg)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		}

		sitePassword := c.PostForm("site_password")
		if sitePassword != cfg.SitePassword {
			c.String(http.StatusUnauthorized, "密码错误")
			return
		}

		uniqueName := c.PostForm("unique_name")
		if !slices.Contains(cfg.AllowUsers, uniqueName) {
			c.String(http.StatusForbidden, "未授权的用户")
			return
		}

		accessToken := cfg.AccessToken

		formData := url.Values{
			"unique_name":        {uniqueName},
			"access_token":       {accessToken},
			"site_limit":         {siteLimit},
			"expires_in":         {expiresIn},
			"gpt35_limit":        {gpt35Limit},
			"gpt4_limit":         {gpt4Limit},
			"show_conversations": {showConversations},
			"reset_limit":        {resetLimit},
		}

		var respData map[string]any
		resp, err := req.R().SetFormDataFromValues(formData).SetSuccessResult(&respData).Post(registerURL)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error registering token")
			return
		}
		defer resp.Body.Close()

		tokenKey, ok := respData["token_key"].(string)
		if !ok {
			tokenKey = "未找到 Share_token"
		}

		loginResp := map[string]string{}

		post, err := req.R().SetBody(map[string]string{
			"share_token": tokenKey,
		}).SetSuccessResult(&loginResp).Post(oauthURL)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer post.Body.Close()

		c.Redirect(http.StatusMovedPermanently, loginResp["login_url"])
	})

	r.Run("0.0.0.0:48881")
}

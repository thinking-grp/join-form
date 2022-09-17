package main

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
	"github.com/disgoorg/snowflake/v2"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tidwall/gjson"

	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/", getRoot)
	e.POST("/join", postJoin)

	e.Logger.Fatal(e.Start(":3040"))
}

func getRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Running thinking join form")
}

type JoinParam struct {
	Name         string `json:"name"`
	TwitterID    string `json:"twitter_id"`
	GitHubID     string `json:"github_id"`
	Introduction string `json:"introduction"`
	Site         string `json:"site"`
}

func postJoin(c echo.Context) error {
	param := new(JoinParam)
	if err := c.Bind(param); err != nil {
		return err
	}

	if param.Name == "" {
		return c.String(http.StatusBadRequest, "name is required")
	}
	if param.TwitterID == "" {
		return c.String(http.StatusBadRequest, "twitter_id is required")
	}
	if param.GitHubID == "" {
		return c.String(http.StatusBadRequest, "github_id is required")
	}
	if param.Introduction == "" {
		return c.String(http.StatusBadRequest, "introduction is required")
	}
	if param.Site == "" {
		return c.String(http.StatusBadRequest, "site is required")
	}

	var token string = gjson.Get(readConfig(), "token").String()
	var channelID int64 = gjson.Get(readConfig(), "id").Int()

	client := webhook.New(snowflake.ID(channelID), token)

	embeds := discord.NewEmbedBuilder().
		SetTitle("参加申請").
		SetDescription("アイコン画像が取得できていない場合はGitHubIDが間違いである可能性があります").
		AddField("名前", param.Name, false).
		AddField("自己紹介", param.Introduction, false).
		AddField("Twitter", "[@"+param.TwitterID+"](https://twitter.com/"+param.TwitterID+")", false).
		AddField("GitHub", "[@"+param.GitHubID+"](https://github.com/"+param.GitHubID+")", false).
		AddField("サイト", param.Site, false).
		SetImage("https://github.com/" + param.GitHubID + ".png").
		Build()

	_, err := client.CreateEmbeds([]discord.Embed{embeds})

	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}

	return c.String(http.StatusOK, "Done")
}

func readConfig() string {
	var configPath string = filepath.Join("src", "/config.json")

	f, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	json, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

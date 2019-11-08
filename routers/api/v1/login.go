package v1

import (
	"encoding/json"
	"fmt"
	"github.com/nguyen930411/go-gin-cms/pkg/app"
	"github.com/nguyen930411/go-gin-cms/pkg/e"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const redirectUri = "http://localhost:8000/azure/callback"
const client_id = "b035a92e-8247-4da2-b5b8-f982a015d43f"
const client_secret = "cM?4d[TDtyKsyUQ4[7L?5cs?630@M3xK"

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
	}
	return string(b)
}

func GetLogin(c *gin.Context) {
	authUrl := `https://login.microsoftonline.com/common/oauth2/v2.0/authorize?
client_id=` + client_id + `
&response_type=code
&redirect_uri=` + url.QueryEscape(redirectUri) + `
&response_mode=query
&scope=openid%20offline_access%20https%3A%2F%2Fgraph.microsoft.com%2Fuser.read
&state=` + RandomString(16)
	c.Redirect(http.StatusMovedPermanently, authUrl)
	//c.Abort()
	appG := app.Gin{C: c}

	data := make(map[string]interface{})
	data["lists"] = "test list"
	data["total"] = "test total"

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetCallback(c *gin.Context) {
	//appG := app.Gin{C: c}

	tokenUrl := "https://login.microsoftonline.com/common/oauth2/v2.0/token"

	form := url.Values{}
	form.Add("client_id", client_id)
	form.Add("scope", "https://graph.microsoft.com/user.read")
	form.Add("code", c.Query("code"))
	form.Add("redirect_uri", redirectUri)
	form.Add("grant_type", "authorization_code")
	form.Add("client_secret", client_secret)
	req, err := http.NewRequest("POST", tokenUrl, strings.NewReader(form.Encode()))
	req.PostForm = form
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	var userJson map[string]interface{}
	json.Unmarshal([]byte(body), &userJson)
	fmt.Printf("Results: %v\n", userJson)
	reqUser, err := http.NewRequest("GET", tokenUrl, nil)
	reqUser.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	reqUser.Header.Set("Authorization", "Bearer " + userJson["access_token"].(string))
	respUser, err := client.Do(req)
	defer respUser.Body.Close()
	bodyUser, _ := ioutil.ReadAll(respUser.Body)
	fmt.Println("response Body:", string(bodyUser))

	//data := make(map[string]interface{})
	defer c.Abort()
	//data["data"] = c.Query("code")

	//appG.Response(http.StatusOK, e.SUCCESS, data)
}

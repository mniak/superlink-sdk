package superlink

import (
	"context"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type Client struct {
	resty       *resty.Client
	authConfig  clientcredentials.Config
	token       *oauth2.Token
	AccessToken string
}

func NewClient(baseurl, clientid, clientsecret string) *Client {
	return &Client{
		resty: resty.New().SetHostURL(baseurl),
		authConfig: clientcredentials.Config{
			ClientID:     clientid,
			ClientSecret: clientsecret,
			TokenURL:     fmt.Sprintf("%s/api/public/v2/token", baseurl),
			AuthStyle:    oauth2.AuthStyleAutoDetect,
		},
	}
}

func (c *Client) getToken() (*oauth2.Token, error) {
	if c.token == nil || c.token.Expiry.Before(time.Now().Add(5*time.Second)) {
		ctx := context.Background()
		token, err := c.authConfig.Token(ctx)
		c.token = token
		return c.token, err
	}
	return c.token, nil
}

func (c *Client) SetDebug(d bool) *Client {
	c.resty.SetDebug(d)
	return c
}

func (c *Client) CreateLink(link Link) (result LinkCreated, err error) {
	token, err := c.getToken()
	if err != nil {
		return
	}
	resp, err := c.resty.R().
		SetAuthToken(token.AccessToken).
		SetResult(&LinkCreated{}).
		SetBody(link).
		Post("api/public/v1/products")

	if err != nil {
		return
	}
	return resp.Result().(LinkCreated), nil
}

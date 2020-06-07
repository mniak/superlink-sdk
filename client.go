package superlink

import "github.com/go-resty/resty/v2"

type client struct {
	resty       *resty.Client
	AccessToken string
}

func NewClient(baseurl string) client {
	return client{
		resty: resty.New().SetHostURL(baseurl),
	}
}

func (c *client) CreateLink() {
	c.resty.R().
		SetAuthToken(c.AccessToken).
		Post("api/public/v1/products")
}

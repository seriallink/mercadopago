package mp

import "net/http"

const oauthTokenPath = "/oauth/token"

type Client struct {
	ClientId     string
	ClientSecret string
	PublicKey    string
	AuthToken    AuthToken
	httpClient   *http.Client
}

type AuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	LiveMode     bool   `json:"live_mode"`
	UserId       int64  `json:"user_id"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_id"`
	Scope        string `json:"scope"`
}

func NewClient(id, secret, key, token string) *Client {

	return &Client{
		ClientId:     id,
		ClientSecret: secret,
		PublicKey:    key,
		AuthToken: AuthToken{
			AccessToken: token,
		},
		httpClient: http.DefaultClient,
	}

}

func (c *Client) Authorize() error {

	data := Params{
		"grant_type":    "client_credentials",
		"client_id":     c.ClientId,
		"client_secret": c.ClientSecret,
	}

	return c.Post(oauthTokenPath, data, nil, &c.AuthToken)

}

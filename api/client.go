package api

type Client struct {
	authorization Authorization
	baseURL       string
}

func NewClient(authorization Authorization) Client {
	return Client{
		authorization: authorization,
		baseURL:       EndpointURL,
	}
}

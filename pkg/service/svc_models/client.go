package svc_models

type Client struct {
	Fund  int
	Tasks []Task
}

func NewClient() *Client {
	return &Client{Fund: 10000}
}

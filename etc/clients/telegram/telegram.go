package telegram

import (
	"encoding/json"
	"etc/lib/e"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host    string
	baseUrl string
	client  http.Client
}

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

func New(host string, token string) *Client {
	return &Client{
		host:    host,
		baseUrl: newBaseUrl(token),
		client:  http.Client{},
	}
}

func newBaseUrl(token string) string {
	return fmt.Sprintf("bot%s", token)
}

func (c *Client) Updates(offset, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, e.Wrap(fmt.Sprintf("bad update data '%s'", getUpdatesMethod), err)
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, e.Wrap("Unmarshal error:", err)
	}

	return res.Result, nil
}

func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return e.Wrap("can't send message", err)
	}

	return nil
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.baseUrl, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, e.Wrap("can't create request", err)
	}

	req.URL.RawQuery = query.Encode()

	res, err := c.client.Do(req)
	if err != nil {
		return nil, e.Wrap("can't get response", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, e.Wrap("can't read response", err)
	}

	return body, nil
}

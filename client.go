package dingtalk

import (
	"time"
)

const (
	REFRESH_TOKEN_INTERVAL = 7000
)

type DTalkClient struct {
	CorpID       string
	CorpSecret   string
	accessToken  string
	getTokenTime int64
}

func NewDTalkClient(corpid, corpSecret string) (*DTalkClient, error) {
	client := &DTalkClient{
		CorpID:     corpid,
		CorpSecret: corpSecret,
	}

	_, err := client.GetAccessToken()
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			time.Sleep(REFRESH_TOKEN_INTERVAL * time.Second)
			if client.accessToken != "" {
				client.accessToken = ""
			}
		}
	}()

	return client, nil
}

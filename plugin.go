package main

import (
	"crypto/tls"
	"errors"
	"net/http"
	"strings"

	"github.com/0ghny/go-synochat"
)

type Plugin struct {
	Url     string
	Token   string
	Message string
	SkipSSL bool
}

func (p Plugin) Exec() (err error) {
	if strings.TrimSpace(p.Token) == "" || strings.TrimSpace(p.Url) == "" {
		return errors.New("missing synochat token or server url")
	}

	var client *synochat.Client
	if p.SkipSSL {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client, err = synochat.NewCustomClient(p.Url, &http.Client{Transport: tr})
	} else {
		client, err = synochat.NewClient(p.Url)
	}
	if err != nil {
		return err
	}

	err = client.SendMessage(&synochat.ChatMessage{Text: p.Message}, p.Token)
	if err != nil {
		return err
	}

	return nil
}

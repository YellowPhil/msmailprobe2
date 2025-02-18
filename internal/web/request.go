package web

import (
	"crypto/tls"
	b64 "encoding/base64"
	"net/http"
	"strings"
	"time"

	"github.com/corpix/uarand"
	"github.com/rs/zerolog/log"
	"github.com/yellowphil/msmailprobe2/internal/errors"
)

func WebRequest(url string) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   time.Duration(3 * time.Second),
		Transport: tr,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", uarand.GetRandom())

	resp, err := client.Do(req)
	if err != nil {
		log.Debug().Err(err).Stack()
		return err
	}

	switch resp.StatusCode {
	case 401:
		return &errors.ErrUnauthorized{}
	case 200:
		return nil
	default:
		return &errors.ErrSomethingWrong{StatusCode: resp.StatusCode}
	}
}

func NTLMRequest(url string) []byte {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	timeout := time.Duration(3 * time.Second)

	client := &http.Client{
		Timeout:   timeout,
		Transport: tr,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", uarand.GetRandom())
	req.Header.Set("Authorization", "NTLM TlRMTVNTUAABAAAAB4IIogAAAAAAAAAAAAAAAAAAAAAGAbEdAAAADw==")
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err)
		return []byte("")
	}

	ntlmResponse := resp.Header.Get("WWW-Authenticate")
	log.Debug().Msgf("NTLM response: %s", ntlmResponse)

	data := strings.Split(ntlmResponse, " ")
	base64DecodedResp, err := b64.StdEncoding.DecodeString(data[1])
	if err != nil {
		log.Debug().Err(err).Stack()
	}
	return base64DecodedResp
}

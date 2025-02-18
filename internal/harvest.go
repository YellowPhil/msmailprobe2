package internal

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/yellowphil/msmailprobe2/internal/errors"
	"github.com/yellowphil/msmailprobe2/internal/web"
)

func HarvestInternalDomain(hostUrl string) string {
	log.Info().Msg("Attempting to harvest internal domain:")

	var (
		err          error
		urlToHarvest string
		urls         = []string{"/ews", "/autodiscover/autodiscover.xml", "/rpc", "/mapi", "/oab"}
	)
	for _, url := range urls {
		targetUrl := fmt.Sprintf("%s%s", hostUrl, url)
		err = web.WebRequest(targetUrl)
		if _, ok := err.(*errors.ErrUnauthorized); ok {
			urlToHarvest = targetUrl
		}
	}
	if urlToHarvest == "" {
		log.Fatal().Msg("Unable to resolve host provided to harvest internal domain name.")
	}
	log.Debug().Msgf("URL that will be used for harvesting internal domain is %s", urlToHarvest)

	response := web.NTLMRequest(urlToHarvest)
	if len(response) == 0 {
		log.Fatal().Msg("Unable to parse NTLM response for internal domain name")
	}

	var (
		continueAppending     bool
		internalDomainBuilder strings.Builder
	)

	for _, decimalValue := range response {
		switch decimalValue {
		case 0:
			continue
		case 2:
			continueAppending = false
		case 15:
			continueAppending = true
		default:
			if continueAppending {
				internalDomainBuilder.WriteByte(decimalValue)
			}
		}
	}

	internalDomain := internalDomainBuilder.String()

	return internalDomain
}

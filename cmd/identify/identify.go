package identify

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/yellowphil/msmailprobe2/internal"
	weberrors "github.com/yellowphil/msmailprobe2/internal/errors"
	"github.com/yellowphil/msmailprobe2/internal/web"
)

var (
	Target      string
	IdentifyCmd = &cobra.Command{
		Use:     "identify",
		Short:   "identify enumeration endpoints",
		Long:    "~~Identify Command~~",
		Example: "identify -t mail.target.com",
		Run:     run,
	}
)

func run(cmd *cobra.Command, args []string) {
	if !strings.HasPrefix(Target, "https://") && !strings.HasPrefix(Target, "http://") {
		Target = fmt.Sprintf("https://%s", Target)
	}
	internalDomain := internal.HarvestInternalDomain(Target)
	log.Info().Msgf("[+] Found internal domain: %s", internalDomain)
	urlEnum(Target)
}

func urlEnum(target string) {
	// amogus.sex  => amogus-sex.mail.protection.outlook.com"
	domainSlice := strings.Split(target, ".")
	o365Domain := fmt.Sprintf("%s-%s.mail.protection.outlook.com",
		domainSlice[len(domainSlice)-2], domainSlice[len(domainSlice)-1])
	if addr, err := net.LookupIP(o365Domain); err != nil || addr == nil {
		log.Warn().Msg("[-] Domain is not using o365 resources.")
	} else {
		log.Info().Msg("[+] Domain is using o365 resources.")
	}

	timeBasedEndpoints := []string{"/Microsoft-Server-ActiveSync", "/autodiscover/autodiscover.xml", "/owa"}
	log.Info().Msg("Identifying endpoints vulnerable to time-based enumeration:")
	enumEndpointFound := false
	sprayEndpointFound := false

	for _, url := range timeBasedEndpoints {
		targetUrl := fmt.Sprintf("%s%s", target, url)

		if err := web.WebRequest(targetUrl); err == nil || errors.Is(err, &weberrors.ErrUnauthorized{}) {
			log.Info().Msgf("[+] %s", targetUrl)
			enumEndpointFound = true
		} else {
			log.Debug().Msgf("[-] %s\t%v", targetUrl, err)
		}

	}
	if !enumEndpointFound {
		log.Warn().Msg("[-] No Exchange endpoints vulnerable to time-based enumeration discovered.")
	}
	log.Info().Msg("Identifying exposed Exchange endpoints for potential spraying")

	SprayEndpoints401 := []string{"/rpc", "/oab", "/ews", "/mapi"}
	SprayEndpoints200 := []string{"/ecp", "/owa"}

	for _, url := range SprayEndpoints401 {
		targetUrl := fmt.Sprintf("%s%s", target, url)
		if err := web.WebRequest(targetUrl); errors.Is(err, &weberrors.ErrUnauthorized{}) {
			log.Info().Msgf("[+] %s", targetUrl)
			sprayEndpointFound = true
		} else {
			log.Debug().Err(err)
			log.Debug().Msgf("[-] %s\t%v", targetUrl, err)
		}
	}
	for _, url := range SprayEndpoints200 {
		targetUrl := fmt.Sprintf("%s%s", target, url)
		if err := web.WebRequest(targetUrl); err == nil {
			log.Info().Msgf("[+] %s", targetUrl)
			sprayEndpointFound = true
		} else {
			log.Debug().Msgf("[-] %s\t%v", targetUrl, err)
		}

	}
	if !sprayEndpointFound {
		log.Warn().Msg("[-] No onprem Exchange services identified.")
	}
}

func init() {
	IdentifyCmd.Flags().StringVarP(&Target, "target", "t", "", "target host")
	IdentifyCmd.MarkFlagRequired("target")
}

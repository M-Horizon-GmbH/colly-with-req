package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/M-Horizon-GmbH/colly-with-req"
	"github.com/M-Horizon-GmbH/colly-with-req/proxy"
)

func main() {
	c := colly.NewCollector()
	c.ImpersonateChrome()

	// PROXY need this format http://username:password@host:port
	// Set the PROXY environment variable to a comma-separated list of proxy URLs
	// Like this: export PROXY="http://localhost:8080,socks5://localhost:1080"
	proxiesEnv := os.Getenv("PROXY")
	if proxiesEnv == "" {
		log.Fatal("PROXY environment variable not set")
	}

	proxyList := parseProxyList(proxiesEnv)
	if len(proxyList) == 0 {
		log.Fatal("no valid proxy URLs found in PROXY")
	}
	fmt.Println("proxyList:", proxyList)

	rp, err := proxy.RoundRobinProxySwitcher(proxyList...)
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	c.OnHTML("element-selector", func(e *colly.HTMLElement) {
		log.Println(e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println(string(r.Body))
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error on %s: %s", r.Request.URL, err)
	})

	c.Visit("https://www.howsmyssl.com/")
	c.Visit("https://ifconfig.me/ip")
	c.Visit("https://ifconfig.me/ip")
	c.Visit("https://ifconfig.me/ip")
	c.Wait()
}

// parseProxyList splittet am Komma, trimmt Leerraum
// und verwirft leere Einträge.
func parseProxyList(env string) []string {
	parts := strings.Split(env, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}

package extensions

import (
	"github.com/M-Horizon-GmbH/colly-with-req"
)

// URLLengthFilter filters out requests with URLs longer than URLLengthLimit
func URLLengthFilter(c *colly.Collector, URLLengthLimit int) {
	c.OnRequest(func(r *colly.Request) {
		if len(r.URL.String()) > URLLengthLimit {
			r.Abort()
		}
	})
}

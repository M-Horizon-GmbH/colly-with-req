# colly-with-req

An advanced, high-performance scraping toolkit built for [tradebys.com](https://tradebys.com), combining the elegance of Colly with the power of the `req` HTTP client for robust, browser-like requests.

## 🚀 Features

* **Seamless Integration**: Replaced Colly’s default `http.Client` with the `req/v3` client.
* **Browser Fingerprinting**: Use `ImpersonateChrome()` or `ImpersonateFirefox()` to mimic real browser TLS fingerprints, HTTP/2 settings, header order, and more.
* **Anti-Crawler Bypass**: Effortlessly bypass fingerprint-based blocking with built‑in HTTP fingerprint impersonation.
* **Rich HTTP Client Capabilities**: Leverage retries, timeouts, automatic decoding, debug dumps, and middleware from `req`.

* **Scalable & Configurable**: Full Colly feature set (concurrency, delays, caching, robots.txt) plus `req`’s extensibility.

## ⚙️ Installation

Add to your Go module:

```bash
go get go get github.com/M-Horizon-GmbH/colly-with-req
```

## 📦 Usage

```go
package main

import (
 "log"

 "github.com/M-Horizon-GmbH/colly-with-req"
)

func main() {
 c := colly.NewCollector()
 c.ImpersonateChrome()

 c.OnHTML(".label", func(e *colly.HTMLElement) {
  log.Println(e.Text)
 })

 c.OnRequest(func(r *colly.Request) {
  log.Println("Visiting", r.URL)
 })

 c.Visit("https://www.howsmyssl.com/")
 c.Wait()
}
```

### Impersonation Example

```go
reqClient := req.C().ImpersonateFirefox()
```

Switch seamlessly between different browser fingerprints to adapt to site-specific anti-crawler defenses.

## 🔧 Configuration

* **TLS Fingerprint**: `ImpersonateChrome()`, `ImpersonateFirefox()`
* **Retries & Backoff**: `reqClient.SetRetry(...)`
* **Debugging**: `reqClient.DevMode()`, `EnableDump()` on requests
* **Middleware**: Add request/response middleware via `reqClient.OnBeforeRequest(...)` and `OnAfterResponse(...)`

*For full `req` docs, visit [https://req.cool/docs/tutorial/http-fingerprint/](https://req.cool/docs/tutorial/http-fingerprint/)*

## 📖 Why `req` + Colly?

Colly is legendary for its crawler-friendly API and concurrency controls, but its default HTTP client can be easily fingerprinted by sophisticated anti-bot systems. By integrating `req/v3`:

1. **True Browser Emulation**: Realistic TLS handshake, HTTP/2 frames, header ordering.
2. **Built-In HTTP3 & HTTP2 Support**: Auto-negotiation or forced protocol selection.
3. **Rich Debug & Middleware**: Simplify troubleshooting, logging, and custom logic.

## 🙏 Credits

This project builds upon the outstanding work of:

* [Colly](https://github.com/gocolly/colly) — Lightning fast and elegant scraping framework.
* [req/v3](https://github.com/imroc/req) — Simple Go HTTP client with “Black Magic”.

All credits go to the Colly and `req` maintainers and contributors.

## 🔗 Used at Tradebys

We use this integrated package at [tradebys.com](https://tradebys.com) to track product prices across multiple marketplaces, empowering sellers with real‑time market positioning insights without wasting marketing budget.

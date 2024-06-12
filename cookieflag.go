// Package cookieflag provides a Caddy module that modifies various flags (Secure, HttpOnly, ...) in Set-Cookie headers.
// It allows users to customize these flags based on their needs.
package cookieflag

import (
	"net/http"
	"strings"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(CookieFlag{})
	httpcaddyfile.RegisterHandlerDirective("cookieflag", parseCaddyfile)
}

// CookieFlag is a middleware that modifies various flags (Secure, HttpOnly, ...) in Set-Cookie headers..
//
// Syntax:
//
//	cookieflag [<matcher>] [(+|-)<field>] {
//		+<field>
//		-<field>
//	}
type CookieFlag struct {
	// The list of cookie flags to be modified.
	//
	// Prepend the flag name with a `+` to add that flag ot with `-` to remove it
	//
	// Flag | Description
	// ------------|-------------
	// secure | The `Secure` flag
	// httponly | The `HttpOnly` flag
	Flags []string `json:"flags,omitempty"`
}

// CaddyModule returns the Caddy module information.
func (CookieFlag) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.cookieflag",
		New: func() caddy.Module { return new(CookieFlag) },
	}
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (cf CookieFlag) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	rr := &responseRewriter{ResponseWriter: w, flags: cf.Flags}
	return next.ServeHTTP(rr, r)
}

// responseRewriter is an http.ResponseWriter that can modify Set-Cookie headers.
type responseRewriter struct {
	http.ResponseWriter
	flags []string
}

func (rw *responseRewriter) WriteHeader(statusCode int) {
	headers := rw.Header()["Set-Cookie"]
	for i, header := range headers {
		headers[i] = modifySetCookieHeader(header, rw.flags)
	}
	rw.ResponseWriter.WriteHeader(statusCode)
}

// modifySetCookieHeader modifies the Set-Cookie header based on the provided flags.
func modifySetCookieHeader(header string, flags []string) string {
	// Check which flags are specified and remove them if necessary
	for _, flag := range flags {
		switch flag {
		case "+secure":
			header = strings.ReplaceAll(header, "; Secure", "")
			header += "; Secure"
		case "-secure":
			header = strings.ReplaceAll(header, "; Secure", "")
		case "+httponly":
			header = strings.ReplaceAll(header, "; HttpOnly", "")
			header += "; HttpOnly"
		case "-httponly":
			header = strings.ReplaceAll(header, "; HttpOnly", "")
		}
	}

	return strings.TrimSpace(header)
}

// parseCaddyfile sets up the handler from Caddyfile tokens. Syntax:
//
//	cookieflag [+secure|-secure] [+httponly|-httponly]
//
// or:
//
//	cookieflag {
//	  +secure
//	  -httponly
//	}
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var cf CookieFlag
	err := cf.UnmarshalCaddyfile(h.Dispenser)
	return cf, err
}

// UnmarshalCaddyfile sets up the handler from Caddyfile tokens.
func (cf *CookieFlag) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		// Check if it is a block
		if d.NextArg() {
			// Single line configuration
			cf.Flags = append(cf.Flags, d.Val())
		} else {
			// Block configuration
			for nesting := d.Nesting(); d.NextBlock(nesting); {
				cf.Flags = append(cf.Flags, d.Val())
			}
		}
	}
	return nil
}

// Interface guards
var (
	_ caddyhttp.MiddlewareHandler = (*CookieFlag)(nil)
	_ caddy.Module                = (*CookieFlag)(nil)
	_ caddyfile.Unmarshaler       = (*CookieFlag)(nil)
)

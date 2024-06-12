# caddy-cookieflag

![Build and test](https://github.com/teodorescuserban/caddy-cookieflag/actions/workflows/test.yml/badge.svg)
[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/teodorescuserban/caddy-cookieflag)
[![Go Report Card](https://goreportcard.com/badge/github.com/teodorescuserban/caddy-cookieflag?)](https://goreportcard.com/report/github.com/teodorescuserban/caddy-cookieflag)

This is a caddy plugin. Needs >= 2.8.0.
Adds or removes the "Secure" and "HttpOnly" flags on you cookies set by the upstream.

## Usage

### One line usage

```caddyfile
:8881 {

    cookieflag +secure
    cookieflag -httponly

    reverse_proxy localhost:8889 {
        header_down -Server
    }
}

:8889 {
    header Content-Type "text/html; charset=utf-8"
    header +Set-Cookie "samesite-ex1=aaa; SameSite=Lax"
    header +Set-Cookie "max-age-ex1=bbb; Max-Age=0"
    header +Set-Cookie "secure-ex1=ccc; Secure"
    header +Set-Cookie "httponly-ex1=ddd; HttpOnly"
    header +Set-Cookie "path-ex1=eee; Path=/"
    header +Set-Cookie "haveitall-ex1=fff; HttpOnly; Secure; Path=/; Max-Age=0; SameSite=Lax; haveitall-ex2=ggg"
    respond "{host}{uri}"
}
```

### Block usage

```caddyfile
:8881 {

    cookieflag {
        +secure
        -httponly
    }

    reverse_proxy localhost:8889 {
        header_down -Server
    }
}

:8889 {
    header Content-Type "text/html; charset=utf-8"
    header +Set-Cookie "samesite-ex1=aaa; SameSite=Lax"
    header +Set-Cookie "max-age-ex1=bbb; Max-Age=0"
    header +Set-Cookie "secure-ex1=ccc; Secure"
    header +Set-Cookie "httponly-ex1=ddd; HttpOnly"
    header +Set-Cookie "path-ex1=eee; Path=/"
    header +Set-Cookie "haveitall-ex1=fff; HttpOnly; Secure; Path=/; Max-Age=0; SameSite=Lax; haveitall-ex2=ggg"
    respond "{host}{uri}"
}
```

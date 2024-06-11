# caddy-cookieflag

![Build and test](https://github.com/teodorescuserban/caddy-cookieflag/actions/workflows/test.yml/badge.svg)
[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/teodorescuserban/caddy-cookieflag)
[![Go Report Card](https://goreportcard.com/badge/github.com/teodorescuserban/caddy-cookieflag)](https://goreportcard.com/report/github.com/teodorescuserban/caddy-cookieflag)

This is a caddy plugin. Works with caddy 2.
Adds or removes the "Secure" and "HttpOnly" flags on you cookies set by the upstream.

## Usage

### Set the module order

You will need to specify the execution order of this module in your caddyfile. This is done in the global options block.

```caddyfile
{
    ...
    order argsort before reverse_proxy
    ...
}
```

### One line usage

Once the order has been set in the global options block, use `argsort lowecase` in any server block

```caddyfile
# Add this block in top-level settings:
{
    ...
    order cookieflag before reverse_proxy
    ...
}

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

### Forward the normalized request to an upstream

Once the order has been set in the global options block, you ensure query arguments sorting for an upstream server

```caddyfile
{
    order argsort before header
}

:8882 {
    argsort
    reverse_proxy localhost:8883
}

:8883 {
    header Content-Type "text/html; charset=utf-8"
    respond "Hello."
}
```

### Block usage

Once the order has been set in the global options block, use `argsort lowecase` in any server block

```caddyfile
# Add this block in top-level settings:
{
    ...
    order cookieflag before reverse_proxy
    ...
}

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

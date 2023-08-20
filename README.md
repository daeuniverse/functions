# Functions

## Intro

A serverless functions for internal interaction with [@daeuniverse](https://github.com/daeuniverse).

## Prerequisites

Export `GITHUB_TOKEN` as environment variable

```bash
export GITHUB_TOKEN=<token>
```

## How to use

```bash
go run main.go
```

## API Comsumption

- healthcheck - `/GET localhost:5888/api/v2/health`
- fetch geodata(geosite.dat) - `/GET localhost:5888/api/v2/geodata?file=geosite.dat`
- fetch geodata(geoip.dat) - `/GET localhost:5888/api/v2/geodata?file=geoip.dat`
- trigger daed-pick-build - `/POST localhost:5888/api/v2/workflow?name=daed-pick-build&daed=main&wing=origin/main&dae=origin/main`

## References

- [go-chi-v5](https://pkg.go.dev/github.com/go-chi/chi/v5)
- [go-resty](https://pkg.go.dev/github.com/go-resty/resty/v2)
- [go-github-sdk](https://pkg.go.dev/github.com/google/go-github/v52/github)
- [How to pipe http.Response to http.ResponseWriter](https://stackoverflow.com/questions/28891531/piping-http-response-to-http-responsewriter)
- [Golang communicating the name of the file served to the browser/client](https://stackoverflow.com/questions/44510661/golang-communicating-the-name-of-the-file-served-to-the-browser-client)

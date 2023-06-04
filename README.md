# Functions

## Intro

TBD.

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

- healthcheck - `/GET localhost:3000/api/health`
- fetch geodata(geosite.dat) - `/GET localhost:3000/api/geodata?file=geosite.dat`
- fetch geodata(geoip.dat) - `/GET localhost:3000/api/geodata?file=geoip.dat`

## References

- [How to pipe http.Response to http.ResponseWriter](https://stackoverflow.com/questions/28891531/piping-http-response-to-http-responsewriter)
- [Golang communicating the name of the file served to the browser/client](https://stackoverflow.com/questions/44510661/golang-communicating-the-name-of-the-file-served-to-the-browser-client)

# Building

``` bash
 go build -o simple-k8s-webserver  cmd/main.go
 ```
# Running

The webserver uses the [labstack/echo](https://echo.labstack.com/) framework. The server `simple-k8s-webserver` listens on two ports,
`8080` and `9360` for the main webserver and prometheus endpoint respectively for the following routes

- 8080
  - `/` returns `200: Hello, World`
  - `/livez` returns `200` when ready
  - `/readyz` returns `200` when ready
  - `/panic` generates a `500` server side error
  - any other routes return `404`

- 9360
  - `/metrics` returns the default prometheus metrics harvested by the [echo
	prometheus middleware](https://echo.labstack.com/middleware/prometheus/)
  - any other routes return `404`

## Logging

Logs are provided by the [echo middleware logging](https://echo.labstack.com/middleware/logger/) component and
write JSON structured logs to STDOUT.



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

## Kubernetes

The `.kustomize` directory contains a simple

- Pod autoscaler
- Deployment
- Service

Which are intended to launch the app in a single
node cluster, and expose the ports on that node.
Port 8080 is exposed as 30090 and 9360 as 30090

Run `kubectl get service` to inspect which ports have
been exposed externaly,

e.g.

```
 kubectl get services
NAME                   TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)                         AGE
kubernetes             ClusterIP   10.96.0.1     <none>        443/TCP                         7d6h
simple-k8s-webserver   NodePort    10.99.4.187   <none>        8080:30080/TCP,9360:30090/TCP   67m
```



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
e.g.
```
{"time":"2023-05-14T22:35:20.889216444Z","id":"","remote_ip":"192.168.65.4","host":"localhost:30080","method":"GET","uri":"/","user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36","status":200,"error":"","latency":31165,"latency_human":"31.165µs","bytes_in":0,"bytes_out":13}
```

## Kubernetes

The `.kustomize` directory contains a simple

- Pod autoscaler
- Deployment
- Service

Which are intended to launch the app in a single
node cluster, and expose the ports on that node.
Port 8080 is exposed as 30080 and 9360 as 30090

Run `kubectl get service` to inspect which ports have
been exposed externaly,

e.g.

```
 kubectl get services
NAME                   TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)                         AGE
kubernetes             ClusterIP   10.96.0.1     <none>        443/TCP                         7d6h
simple-k8s-webserver   NodePort    10.99.4.187   <none>        8080:30080/TCP,9360:30090/TCP   67m
```

For example:

<img width="674" alt="image" src="https://github.com/tobywan/simple-k8s-webserver/assets/17815047/0c4315d9-5265-43a4-946a-fa7b374c1271">

<img width="674" alt="image" src="https://github.com/tobywan/simple-k8s-webserver/assets/17815047/f67d58a9-d226-46bb-bb90-a92d0ecf1e6d">

<img width="674" alt="image" src="https://github.com/tobywan/simple-k8s-webserver/assets/17815047/5b24e493-af5d-4f45-9379-f554dd214150">

<img width="674" alt="image" src="https://github.com/tobywan/simple-k8s-webserver/assets/17815047/bd4af241-6d2a-43dc-8de8-8dfc1717cea1">

<img width="674" alt="image" src="https://github.com/tobywan/simple-k8s-webserver/assets/17815047/5137538c-72d8-4dfe-84d4-a778a3e9a14f">

<img width="674" alt="image" src="https://github.com/tobywan/simple-k8s-webserver/assets/17815047/cb16db7a-fc3f-4ee5-8ff6-256d929b6967">




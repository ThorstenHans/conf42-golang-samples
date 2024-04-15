# Hello Conf42 sample

This sample is based on the `http-go` template provided by Spin CLI (`spin templates list`).

Started from a simple "Hello World" we added the `Router` provided by Spin SDK for Go and registered the following endpoint:

`GET /:kind` -> Return a greeting using the route parameter `:kind` and construct a HTTP response using `Content-Type: application/json`.

## Build and Run the sample locally

To build the sample, you can run `spin build` in the current folder.


```bash
# build the Spin App
spin build
Building component hello-conf42 with `tinygo build -target=wasi -gc=leaking -no-debug -o main.wasm main.go`
Finished building all Spin components
```

To run the sample locally use `spin up`:

```bash
# run the Spin App locally
spin up
Logging component stdio to ".spin/logs/"

Serving http://127.0.0.1:3000
Available Routes:
  hello-conf42: http://127.0.0.1:3000 (wildcard)
```

You can now `curl` the endpoint as shown below:

```bash
# Invoke the Endpoint
curl -iX GET http://localhost:3000/GoLang

HTTP/1.1 200 OK
content-type: application/json
content-length: 35
date: Mon, 15 Apr 2024 14:39:44 GMT

{"message":"Hello Conf42 GoLang!"}
```
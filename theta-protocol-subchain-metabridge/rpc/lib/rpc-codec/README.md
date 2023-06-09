# JSON-RPC 2.0 [![GoDoc](https://godoc.org/github.com/powerman/rpc-codec/jsonrpc2?status.svg)](http://godoc.org/github.com/powerman/rpc-codec/jsonrpc2) [![CircleCI](https://circleci.com/gh/powerman/rpc-codec.svg?style=svg)](https://circleci.com/gh/powerman/rpc-codec) [![Coverage Status](https://coveralls.io/repos/powerman/rpc-codec/badge.svg?branch=master&service=github)](https://coveralls.io/github/powerman/rpc-codec?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/powerman/rpc-codec)](https://goreportcard.com/report/github.com/powerman/rpc-codec)

jsonrpc2 is a codec for net/rpc.

Implements [JSON-RPC 2.0](http://www.jsonrpc.org/specification) and
[JSON-RPC 2.0 Transport: HTTP](http://www.simple-is-better.org/json-rpc/transport_http.html)
specifications with following limitations:

- Client: Batch Request not supported.
- HTTP Client&Server: Pipelined Requests/Responses not supported.
- HTTP Client&Server: GET Request not supported.


## Installation

```sh
go get github.com/powerman/rpc-codec/...
```

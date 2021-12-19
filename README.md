# gap

[![Build Status](https://github.com/ectobit/gap/workflows/check/badge.svg)](https://github.com/ectobit/gap/actions)
[![Go Reference](https://pkg.go.dev/badge/go.ectobit.com/gap.svg)](https://pkg.go.dev/go.ectobit.com/gap)

<!-- [![Go Report](https://goreportcard.com/badge/go.ectobit.com/gap)](https://goreportcard.com/report/go.ectobit.com/gap) -->

[![License](https://img.shields.io/badge/license-BSD--2--Clause--Patent-orange.svg)](https://github.com/ectobit/gap/blob/main/LICENSE)

Custom generic HTTP handler providing automatic JSON decoding/encoding of HTTP request/response to your concrete types. `gap.Wrap` allows to use these custom handlers just as idiomatic Go't HTTP handler functions.

:warning: :construction: This project is unstable and still under heavy construction. It uses Go 1.18beta1.

## Problem description

Standard HTTP handler functions in Go have signature `func(http.ResponseWriter, *http.Request)` which doesn't allow an easy RESTful API implementation. On each request you have to parse the request body using JSON decoder, process the request and return the response body by using JSON encoder. This means lot of code repetitions or at least not elegant and clean code.

## Solution

Using benefits of Go's generics, it's possible now to have generic requests and responses and hide the JSON encoding and decoding in a wrapper which provides the standard HTTP handler functions at the end. This way you can focus to your business logic and not to boring parts like encoding, decoding and many errors that may happen. This projects provides HTTP handler function with signature `func(*gap.Request[I]) *gap.Response[O]`, where I is your request body type and O is response body type. For GET requests request body type is simply `struct{}`.

## Contribution

- `make test` runs unit tests
- `make test-cov` displays test coverage

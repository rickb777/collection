# collection

[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg?style=flat-square)](https://godoc.org/github.com/rickb777/collection)
[![Build Status](https://travis-ci.org/rickb777/collection.svg?branch=master)](https://travis-ci.org/rickb777/collection)
[![Issues](https://img.shields.io/github/issues/rickb777/collection.svg)](https://github.com/rickb777/collection/issues)

A suite of general-purpose collections for Go.

This consists of regular types generated automatically from templates using
[*Runtemplate v3*](https://github.com/rickb777/runtemplate/blob/master/v3/README.md).

Three categories of collections are provided:

 * simple lists, sets & maps - these are wrappers around the equivalent built-in Go language features
 * shared lists, sets & maps - these all use structs that have mutex locking for each method
 * immutable lists, sets & maps - these all have structs containing the data but do not have any mutation methods.

Currently, the following built-in types are supported:

 * `string`
 * `int`
 * `uint`
 * `int64`
 * `uint64`

## Tests

Note that extensive tests are present in the [code generator](https://github.com/rickb777/runtemplate/tree/master/v3
/builtintest). Because all the Go source files here are auto-generated, no additional tests are needed here.

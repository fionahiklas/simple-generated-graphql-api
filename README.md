# simple-generated-graphql-api

## Overview

Simple test of GraphQL generation in Go

Based on this [tutorial](https://www.howtographql.com/graphql-go/1-getting-started/) and the [gqlgen](https://github.com/99designs/gqlgen) tool.


## Setup

* From the [gqlgen github](https://github.com/99designs/gqlgen) README

```
go mod init simple-generated-graphql-api
printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go
go mod tidy
```

* Running the initial generation

```
go run github.com/99designs/gqlgen init
```



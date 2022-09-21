# simple-generated-graphql-api

## Overview

Simple test of GraphQL generation in Go

Based on this [tutorial](https://www.howtographql.com/graphql-go/1-getting-started/) and the [gqlgen](https://github.com/99designs/gqlgen) tool.


## Notes

### Setup

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

* This created alot of generated code that was then committed

### Changing the Schema

* Completely changed the schema to represent a different system with burglar alarms in homes
* Ran this command

```
go run github.com/99designs/gqlgen generate
```

* This generated the following errors

```
validation failed: packages.Load: /Users/fiona/wd/fionahiklas/simple-generated-graphql-api/graph/schema.resolvers.go:43:72: NewTodo not declared by package model
/Users/fiona/wd/fionahiklas/simple-generated-graphql-api/graph/schema.resolvers.go:43:89: Todo not declared by package model
/Users/fiona/wd/fionahiklas/simple-generated-graphql-api/graph/schema.resolvers.go:46:62: Todo not declared by package model
```

* The code that isn't relevant (and causes the errors) is highlighted in `schema.resolvers.go`
* Deleted this and then ran the `generate` step again 
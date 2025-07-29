---
id: Organizing a Go module
aliases: []
tags: []
---

[Ref](https://go.dev/doc/modules/layout)

Go projects can include packages, command line programs, or the combination of two.

A Go package can be split into multiple files within the same directory:

```
project-root-directory/
  go.mod
  modname.go
  modname_test.go
  auth.go
  auth_test.go
  hash.go
  hash_test.go
```

Larger packages or commands may be better spliting off some functionalities into supporting packages.

It is recommended placing such packages into a directory named `internal`. That way we don't expose the functionalities to other module.

```language
project-root-directory/
  -> Functionalities in internal/ will not be exposed
  internal/
    auth/
      auth.go
      auth_test.go
    hash/
      hash.go
      hash_test.go
  go.mod
  modname.go
  modname_test.go
```

A module can consist of multiple importable packages, and each package has its own directory.

```language
project-root-directory/
  go.mod
  modname.go
  modname_test.go
  auth/
    auth.go
    auth_test.go
    token/
      token.go
      token_test.go
  hash/
    hash.go
  internal/
    trace/
      trace.go
```

> The Ollama repository follows this structure, and this might be due to 1) The original creators prefer a flat structure or 2)

We can also have multiple commands:

```
project-root-directory/
  go.mod
  internal/
    ... shared internal packages
  prog1/
    main.go
  prog2/
    main.go
```

When building servers with Go, it is recommended to keep the server logic inside `internal/` and Go commands in `cmd/` (The most common approach)

```language
project-root-directory/
  go.mod
  internal/
    auth/
      ...
    metrics/
      ...
    model/
      ...
  cmd/
    api-server/
      main.go
    metrics-analyzer/
      main.go
    ...
  ... the project's other directories with non-Go code
```

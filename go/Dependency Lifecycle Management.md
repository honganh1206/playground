# Dependency Lifecycle Manamgement

Services like HTTP servers, database connections, and background workers have dependencies must be managed *EXPLICITLY*

While Go provides primitives like `context.Context` and `sync.WaitGroup` (See more in [Context](./Context.md)) it is *the programmer's responsibility* to ensure resources are constructed, started, and shutdown in the correct order

## The `github.com/jacoelho/component` library

A minimal framework to structure and orchestrate stateful services in Go

How? Components are registered with explicit dependencies and organized into levels in a directed acyclic graph (DAG - [See more](https://en.wikipedia.org/wiki/Directed_acyclic_graph))

What is better? Reducing manual writings of start/stop logic, enforce correct component ordering, and top of all centralize lifecycle management across different components

# Golang Docs

Golang notes, scripts, small projects, and solutions from several sources.

## Golang Concepts

- Data Types
  - Strings, Ints, & Booleans
  - Arrays & Slices
  - Maps, Structs
- Variables & Constants
- Formatted Output
- User Input: Get and Validate
- Pointers
- Scope Rules
- Control Flow
  - Loops
  - If-else & Switch
- Encapsulate logic
  - Functions
- Code organization
  - Packages
- Make our app faster..
  - Goroutines

## What is Go / Golang?

- Programming language developed at Google in 2007, Open-sourced in 2009
- Multicore processors, cloud infra, big networked computation clusters became common (infras became scalable, distributed, dynamic, more capacity)
- Doing multiple tasks at once (**Concurrency**) -- Developers need to write code to prevent conflicts (when tasks run in parallel), complex code, expensive & slow (therefore Golang comes). Programming languages with built in concurrency mechanisms: C++, Java, Golang. Without: Python, Nodejs
- Go was designed to run on multiple cores and built to support concurrency
- Concurrency in Go is cheap and easy
- Go is good for performant apps, run on scaled and distributed systems.

## Characteristics of Go

- Combines both **simplicity and readable** syntax of dynamically typed language like Python with the **Efficiency and safety** of a lower-level, statically typed language like C++
- For Server-side or Backend Language (Microservices, web apps, database services)
- Technologies written in golang: Docker, HashiCorp Vault, Kubernetes, CockroachDB, etc
- Simple Syntax: Easy to learn, read and write code
- Fast build time, start up and run
- Requires fewer resources
- Compiles into a single binary (machine code)
  - Faster than intepreted languages like Python
  - Consistent across different OS

## Structure of Go file

- All golang code must belong to a **package**
- The first statement in Go file must be "package ..."
- Go programs are organized into packages
- Go's standard library, provides different core packages to use
- A package is a collection of source files

## Golang CLI Commands

- Create new module
  ```sh
  go mod init <module_name>
  ```
- Compiling and running go files
  ```sh
  go run main.go
  ```

## Data Types in Go

- You need to tell Go compiler, the data type when declaring the variable
  - More robust, reduces the likelihood of errors
  - Helps developers to catch type mismatches sooner (at compile time)
- Type Inference: BUT, Go can infer the type when you assign a value

## Sources

- [Golang Full Course for Beginners](https://www.youtube.com/watch?v=yyUHQIec83I&t=1s)
- https://go.dev/doc/articles/wiki/
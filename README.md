# Golang Docs

Golang notes, scripts, small projects, and solutions from several sources.

## What is Go / Golang?

Go or also called Golang is a Programming Language developed by Google in 2007, Open-sourced in 2009. The people initially developed Go are Rob Pike, Roberrt Griesemer and Ken Thompson

## Why Golang was created?

- Multicore processors, cloud infra, big networked computation clusters became common (infras became scalable, distributed, dynamic, more capacity)
- Doing multiple tasks at once (**Concurrency**) -- Developers need to write code to prevent conflicts (when tasks run in parallel), complex code, expensive & slow (therefore Golang comes). Programming languages with built in concurrency mechanisms: C++, Java, Golang. Without: Python, Nodejs
- Go was designed to run on multiple cores and built to support concurrency
- Concurrency in Go is cheap and easy
- Go is good for performant apps, run on scaled and distributed systems.

## What Golang is used for?

- Server side scripting language, network based program, cross-platform app development, command line tools / utility
- Examples of technologies built using Golang is Docker. Big companies like Uber, BBC, Facebook, Apple and Google use Golang

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

## Advantages of using Golang

- Fun and easy to learn (very simple)
- Extremely fast runtime process and compilation time
- Concurrent programming (multi threads) support
- Garbage collection (automatic memory management) support
- Support in multiple OS platforms (Windows, Mac, Linux, Raspberry Pi, etc)
- Complete dependency and tooling
- Great community support. A lot of free and open source tools.

## Golang CLI Commands

- Create new module
  ```sh
  go mod init <module_name>
  ```
- Compiling go script
  ```sh
  go build -o <compiled_binary> <script.go>
  ```
- Run compiled go binary
  ```sh
  ./<compiled_binary>
  ```
- Running `.go` script without compilation file output
  ```sh
  go run <script.go>
  ```

# deepmodules

<div align="center">

![Figure 4.1](./figure41.png)

<p>
<i>
The best modules are those that provide powerful functionality yet have
simple interfaces. I use the term deep to describe such modules.
</i>
— John Ousterhout <i>A Philosophy of Software Design</i>
</p>

</div>

A blunt tool that analyze the "depth" (lines of code per export) of Go
packages. Identify opportunities to refactor shallow modules, optimize
your package architecture, and build a codebase that stands the test
of time.

## Introduction

Deep modules are a concept in software engineering and system design
that refers to high-quality, reusable components with simple interfaces
but complex internal implementations.

This approach to module design aims to manage complexity effectively
by encapsulating intricate logic and data structures within a clean,
easy-to-use exterior.

Deep modules effectively hide their internal workings, exposing only
what's necessary for other parts of the system to interact with them.

## Installation

```bash
go install github.com/olahol/deepmodules@latest
```

## Usage

```bash
deepmodules /path/to/go/repo

┌────────────────┬───────────────────────────────────┬─────┬───────┬──────┐
│PACKAGE         │DIR                                │LINES│EXPORTS│DEPTH │
├────────────────┼───────────────────────────────────┼─────┼───────┼──────┤
│create          │pkg/cmd/repo/create                │1279 │3      │426.33│
│create          │pkg/cmd/release/create             │839  │2      │419.50│
│label           │pkg/cmd/label                      │893  │3      │297.67│
│status          │pkg/cmd/pr/status                  │529  │2      │264.50│
│set             │pkg/cmd/variable/set               │473  │2      │236.50│
│link            │pkg/cmd/project/link               │233  │1      │233.00│
│view            │pkg/cmd/pr/view                    │466  │2      │233.00│
```

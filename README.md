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

A tool for analyzing module "depth" (lines of code per export) of Go
packages. Identify opportunities to refactor shallow modules, optimize
your package architecture, and build a codebase that stands the test
of time!

## Installation

```bash
go install github.com/olahol/deepmodules@latest
```

## Example

```bash
$ deepmodules /path/to/esbuild/repo

┌───────────┬────────────────────┬─────┬───────┬───────┐
│PACKAGE    │DIR                 │LINES│EXPORTS│DEPTH  │
├───────────┼────────────────────┼─────┼───────┼───────┤
│css_parser │internal/css_parser │9256 │8      │1157.00│
│js_parser  │internal/js_parser  │25274│23     │1098.87│
│js_printer │internal/js_printer │4923 │6      │820.50 │
│linker     │internal/linker     │7302 │13     │561.69 │
│css_printer│internal/css_printer│1141 │3      │380.33 │
│cli        │pkg/cli             │1779 │5      │355.80 │
│bundler    │internal/bundler    │3331 │11     │302.82 │
│runtime    │internal/runtime    │604  │2      │302.00 │
│resolver   │internal/resolver   │5503 │31     │177.52 │
```

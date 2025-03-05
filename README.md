# Comparison between shopify's go-lua and gopher-lua

This repository contains some tests and benchmarks to compare the two Lua implementations [go-lua](https://github.com/epikur-io/go-lua) and [gopher-lua](https://github.com/epikur-io/gopher-lua).

## Running benchmarks:

**Note:** Requires [taksfile](https://taskfile.dev/installation/) top be installed.

```bash
# overall benchmark
task benchmark

# detailed benchmark with mem/cpu profile for go-lua
task benchmark_golua

# detailed benchmark with mem/cpu profile for gopher-lua
task benchmark_gopherlua
```
# Comparison between shopify's go-lua and gopher-lua

This repository contains some tests and benchmarks to compare the two Lua implementations [go-lua](https://github.com/epikur-io/go-lua) and [gopher-lua](https://github.com/epikur-io/gopher-lua).

## Running benchmarks:

**Note:** Requires [taskfile](https://taskfile.dev/installation/) top be installed.

```bash
# overall benchmark
task benchmark

# detailed benchmark with mem/cpu profile for go-lua
task benchmark_golua

# detailed benchmark with mem/cpu profile for gopher-lua
task benchmark_gopherlua
```

## Results

### Gopher-Lua

```
goos: linux
goarch: amd64
pkg: github.com/epikur-io/compare-go-lua-vms/gopher-lua
cpu: Intel(R) Core(TM) i7-7820HQ CPU @ 2.90GHz
BenchmarkGopherLuaTables
BenchmarkGopherLuaTables/tables/append-table.lua
BenchmarkGopherLuaTables/tables/append-table.lua-8                 10000            324644 ns/op          267797 B/op        965 allocs/op
BenchmarkGopherLuaMist
BenchmarkGopherLuaMist/misc/debug.lua
BenchmarkGopherLuaMist/misc/debug.lua-8                            10000            241662 ns/op          288335 B/op       1164 allocs/op
BenchmarkGopherLuaMist/misc/fibonacci-iterative.lua
BenchmarkGopherLuaMist/misc/fibonacci-iterative.lua-8              10000            218149 ns/op          244602 B/op        971 allocs/op
BenchmarkGopherLuaMist/misc/fibonacci-recursive.lua
BenchmarkGopherLuaMist/misc/fibonacci-recursive.lua-8              10000            239228 ns/op          243428 B/op        951 allocs/op
BenchmarkGopherLuaMist/misc/fibonacci-tail-tail.lua
BenchmarkGopherLuaMist/misc/fibonacci-tail-tail.lua-8              10000            239596 ns/op          266148 B/op       1056 allocs/op
BenchmarkGopherLuaMist/misc/url-encode.lua
BenchmarkGopherLuaMist/misc/url-encode.lua-8                       10000            492859 ns/op          346134 B/op       3221 allocs/op
BenchmarkGopherLuaIssues
BenchmarkGopherLuaIssues/issues/gopher-lua-225.issue.lua
BenchmarkGopherLuaIssues/issues/gopher-lua-225.issue.lua-8         10000            421775 ns/op          314342 B/op       1624 allocs/op
```

### Shopify Go-Lua

```
goos: linux
goarch: amd64
pkg: github.com/epikur-io/compare-go-lua-vms/go-lua
cpu: Intel(R) Core(TM) i7-7820HQ CPU @ 2.90GHz
BenchmarkGoLuaTables
BenchmarkGoLuaTables/tables/append-table.lua
BenchmarkGoLuaTables/tables/append-table.lua-8                      10000            313257 ns/op           65955 B/op       2304 allocs/op
BenchmarkGoLuaMisc
BenchmarkGoLuaMisc/misc/fibonacci-iterative.lua
BenchmarkGoLuaMisc/misc/fibonacci-iterative.lua-8                   10000            121713 ns/op           35033 B/op        800 allocs/op
BenchmarkGoLuaMisc/misc/fibonacci-recursive.lua
BenchmarkGoLuaMisc/misc/fibonacci-recursive.lua-8                   10000            111947 ns/op           34321 B/op        783 allocs/op
BenchmarkGoLuaMisc/misc/fibonacci-tail-tail.lua
BenchmarkGoLuaMisc/misc/fibonacci-tail-tail.lua-8                   10000            113051 ns/op           35899 B/op        814 allocs/op
BenchmarkGoLuaIssues
BenchmarkGoLuaIssues/issues/gopher-lua-225.issue.lua
BenchmarkGoLuaIssues/issues/gopher-lua-225.issue.lua-8              10000            186958 ns/op           48228 B/op       1136 allocs/op
```
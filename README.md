Server benchmark
===

## Benchmark

- On OS X
- Core 8

## Ruby-unicorn

```
$ cd ruby-unicorn
$ unicorn --config-file unicorn.rb
```

```
$ ab http://localhost:4444 -n 1000 -c 100
# snip...
Requests per second:    4220.28 [#/sec] (mean)
```

## Golang-net-http


```
$ cd golang-net-http
$ go run server.go
```

```
$ ab http://localhost:5555 -n 1000 -c 100
# snip...
Requests per second:    19128.14 [#/sec] (mean)
```

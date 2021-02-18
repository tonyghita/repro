# repro

> A small reproduction for a Go 1.16 internal compiler error.

The issue is tracked at <https://github.com/golang/go/issues/44355>.

I wasn't able to figure out how to run this reproduction on play.golang.org, so I made this small repository instead.

I was able to `git bisect` the commit which introduced this error to <https://github.com/golang/go/commit/f2c0c2b90200b470c39a2db821b7c707604fe083>.

To reproduce the issue, run tests using *Go1.16*.

```sh
git clone https://github.com/tonyghita/repro.git
cd repro
go test ./...
```

You should see the output

```sh
# github.com/tonyghita/repro_test [github.com/tonyghita/repro.test]
./repro_test.go:10:13: internal compiler error: 'TestF.func1': Value live at entry. It shouldn't be. func TestF.func1, node ~R0, value v19

Please file a bug report including a short program that triggers the error.
https://golang.org/issue/new
FAIL	github.com/tonyghita/repro [build failed]
FAIL
```

Verify that the issue does not exist on Go 1.15.

```sh
go get golang.org/dl/go1.15
go1.15 download

go1.15 test ./...
```

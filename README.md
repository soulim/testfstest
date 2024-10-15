# testfsttest

A call to `fstest.DirFS` takes unexpectedly long time to complete if given directory contains at least one relatively big file.

## Experiment

A test with multiple subtests uses `fstest.TestFS` to check existance of files in given directory.
See [`main_test.go`](main_test.go)

## Setup

- Go version: go1.23.2 darwin/arm64
- `testdata/many-big` directory contains a set of files and one of them is big (4.4MB).
- `testdata/many-small` directory contains a set of file, but none of them is big.
- `testdata/one-<category>` directories contain only one file of each category: big, small, and empty.

## Results

```console
go test ./... -v
=== RUN   TestTestFS
=== RUN   TestTestFS/testdata/many-big/big.jpg
    main_test.go:32: TestFS() time: 4.008915917s
=== RUN   TestTestFS/testdata/many-big/empty.txt
    main_test.go:32: TestFS() time: 3.955715375s
=== RUN   TestTestFS/testdata/many-big/small.txt
    main_test.go:32: TestFS() time: 3.953422583s
=== RUN   TestTestFS/testdata/many-small/empty.txt
    main_test.go:32: TestFS() time: 2.408916ms
=== RUN   TestTestFS/testdata/many-small/small.txt
    main_test.go:32: TestFS() time: 721.459µs
=== RUN   TestTestFS/testdata/one-big/big.jpg
    main_test.go:32: TestFS() time: 3.968942875s
=== RUN   TestTestFS/testdata/one-empty/empty.txt
    main_test.go:32: TestFS() time: 243.958µs
=== RUN   TestTestFS/testdata/one-small/small.txt
    main_test.go:32: TestFS() time: 2.164458ms
--- PASS: TestTestFS (15.89s)
    --- PASS: TestTestFS/testdata/many-big/big.jpg (4.01s)
    --- PASS: TestTestFS/testdata/many-big/empty.txt (3.96s)
    --- PASS: TestTestFS/testdata/many-big/small.txt (3.95s)
    --- PASS: TestTestFS/testdata/many-small/empty.txt (0.00s)
    --- PASS: TestTestFS/testdata/many-small/small.txt (0.00s)
    --- PASS: TestTestFS/testdata/one-big/big.jpg (3.97s)
    --- PASS: TestTestFS/testdata/one-empty/empty.txt (0.00s)
    --- PASS: TestTestFS/testdata/one-small/small.txt (0.00s)
PASS
ok      github.com/soulim/testfstest    16.170s
```

- If a given directory contains at least one big file, then it takes significant time to test each file from this directory.
- Directories without relatively big files handled in milliseconds or microseconds.

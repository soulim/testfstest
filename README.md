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
    main_test.go:32: DirFS() time: 4.056086875s
=== RUN   TestTestFS/testdata/many-big/empty.txt
    main_test.go:32: DirFS() time: 3.944343834s
=== RUN   TestTestFS/testdata/many-big/small.txt
    main_test.go:32: DirFS() time: 3.940132208s
=== RUN   TestTestFS/testdata/many-small/empty.txt
    main_test.go:32: DirFS() time: 2.439167ms
=== RUN   TestTestFS/testdata/many-small/small.txt
    main_test.go:32: DirFS() time: 705.167µs
=== RUN   TestTestFS/testdata/one-big/big.jpg
    main_test.go:32: DirFS() time: 3.95575275s
=== RUN   TestTestFS/testdata/one-empty/empty.txt
    main_test.go:32: DirFS() time: 233.166µs
=== RUN   TestTestFS/testdata/one-small/small.txt
    main_test.go:32: DirFS() time: 1.989042ms
--- PASS: TestTestFS (15.90s)
    --- PASS: TestTestFS/testdata/many-big/big.jpg (4.06s)
    --- PASS: TestTestFS/testdata/many-big/empty.txt (3.94s)
    --- PASS: TestTestFS/testdata/many-big/small.txt (3.94s)
    --- PASS: TestTestFS/testdata/many-small/empty.txt (0.00s)
    --- PASS: TestTestFS/testdata/many-small/small.txt (0.00s)
    --- PASS: TestTestFS/testdata/one-big/big.jpg (3.96s)
    --- PASS: TestTestFS/testdata/one-empty/empty.txt (0.00s)
    --- PASS: TestTestFS/testdata/one-small/small.txt (0.00s)
PASS
ok      github.com/soulim/testfstest    16.124s
```

- If a given directory contains at least one big file, then it takes significant time to test each file from this directory.
- Directories without relatively big files handled in milliseconds or microseconds.

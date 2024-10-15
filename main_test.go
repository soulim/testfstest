package main_test

import (
	"io"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
	"time"
)

func TestTestFS(t *testing.T) {
	src := "testdata"

	dst, err := os.MkdirTemp(".", "testrun-*")
	if err != nil {
		t.Fatalf("cannot create a temporary directory, error: %v", err)
	}

	var ts0, ts1 time.Time

	ts0 = time.Now()
	copy(t, filepath.Join(src, "IMG_20240802_111620.jpg"), filepath.Join(dst, "IMG_20240802_111620.jpg"))
	ts1 = time.Now()
	t.Logf("copy() time: %v", ts1.Sub(ts0))

	ts0 = time.Now()
	err = fstest.TestFS(os.DirFS(dst), "IMG_20240802_111620.jpg")
	ts1 = time.Now()
	t.Logf("DirFS() time: %v", ts1.Sub(ts0))
	if err != nil {
		t.Fatal(err)
	}
}

func copy(t testing.TB, src string, dst string) {
	t.Helper()

	r, err := os.Open(src)
	if err != nil {
		t.Fatalf("cannot open a source file, error: %v", err)
	}

	w, err := os.Create(dst)
	if err != nil {
		t.Fatalf("cannot create a destinaton file, error: %v", err)
	}

	n, err := io.Copy(w, r)
	if err != nil {
		t.Fatalf("cannot copy source to destinaton, error: %v", err)
	}

	t.Logf("copied:\n\tsrc: %s\n\tdst: %s\n\tsize: %v\n", src, dst, n)
}

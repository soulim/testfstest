package main_test

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
	"time"
)

func TestTestFS(t *testing.T) {
	files := []string{
		filepath.Join("testdata", "many-big", "big.jpg"),
		filepath.Join("testdata", "many-big", "empty.txt"),
		filepath.Join("testdata", "many-big", "small.txt"),
		filepath.Join("testdata", "many-small", "empty.txt"),
		filepath.Join("testdata", "many-small", "small.txt"),
		filepath.Join("testdata", "one-big", "big.jpg"),
		filepath.Join("testdata", "one-empty", "empty.txt"),
		filepath.Join("testdata", "one-small", "small.txt"),
	}

	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			var ts0, ts1 time.Time

			dir := filepath.Dir(file)

			ts0 = time.Now()
			err := fstest.TestFS(os.DirFS(dir), filepath.Base(file))
			ts1 = time.Now()
			t.Logf("TestFS() time: %v", ts1.Sub(ts0))
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

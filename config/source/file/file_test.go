package file_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/Breeze0806/go-admin-core/config/source/file"
)

func TestConfig(t *testing.T) {
	data := []byte(`{"foo": "bar"}`)
	path := filepath.Join(os.TempDir(), fmt.Sprintf("file.%d", time.Now().UnixNano()))
	fh, err := os.Create(path)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		fh.Close()
		os.Remove(path)
	}()
	_, err = fh.Write(data)
	if err != nil {
		t.Error(err)
	}
}

func TestFile(t *testing.T) {
	data := []byte(`{"foo": "bar"}`)
	path := filepath.Join(os.TempDir(), fmt.Sprintf("file.%d", time.Now().UnixNano()))
	fh, err := os.Create(path)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		fh.Close()
		os.Remove(path)
	}()

	_, err = fh.Write(data)
	if err != nil {
		t.Error(err)
	}

	f := file.NewSource(file.WithPath(path))
	c, err := f.Read()
	if err != nil {
		t.Error(err)
	}
	if string(c.Data) != string(data) {
		t.Logf("%+v", c)
		t.Error("data from file does not match")
	}
}

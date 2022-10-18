package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"testing"
)

func TestMakeZipFile(t *testing.T) {
	t.Run("main test", func(t *testing.T) {
		buffer := new(bytes.Buffer)
		MakeZipFile(buffer)

		// zipを読み込むためのreader生成
		r, err := zip.NewReader(bytes.NewReader(buffer.Bytes()), int64(len(buffer.Bytes())))
		if err != nil {
			t.Error("error")
		}

		// outputの変数に読み込ませる
		output := bytes.Buffer{}
		rc, err := r.File[0].Open()
		t.Cleanup(func() { rc.Close() })

		output.ReadFrom(rc)
		actual := string(output.Bytes())
		expect := `id,name,address,birthday,gender
`
		if actual != expect {
			t.Error("\nactual： ", actual, "\nexpect： ", expect)
		}
		fmt.Println(actual)
		fmt.Println(len(actual))
		fmt.Println(len(expect))

	})
}

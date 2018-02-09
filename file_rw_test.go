package testGo

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func Test_file_w(t *testing.T) {
	f, err := os.Create("./test")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	defer f.Close()
	f.WriteString("hello world")
}

func Test_file_r(t *testing.T) {
	f, err := os.Open("./test")
	if err != nil {

		if strings.Contains(err.Error(), "no such file or directory") {
			f, err = os.Create("./test")
			if err != nil {
				t.Error(err)
				t.Fail()
			}
		} else {
			t.Error(err)
			t.Fail()
		}

	}
	defer f.Close()
	f.WriteString("hello world\n")

	reader := bufio.NewReader(f)
	line, _, err := reader.ReadLine()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if string(line) != "hello world" {
		t.Fail()
	}
}

package envLoader

import (
	"testing"
	"os"
	"fmt"
)

var cases = [][2]string{
	{"test", "yep"},
	{"second", "envSet"},
	{"third", "thisIsTheLast"},
}

func TestLoad(t *testing.T) {
	f, err := os.OpenFile(".env", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0775)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range cases {
		fmt.Fprintf(f, "%s,%s\n", v[0], v[1])
	}
	f.Close()

	if err:=Load();nil!=err{
		t.Fatal(err)
	}
	for _,kv:=range cases {
		if os.Getenv(kv[0]) != kv[1] {
			t.Fail()
		}
	}
	if err:=os.Remove(".env");nil!=err{
		t.Fail()
	}
}

func TestLoadCustomFilename(t *testing.T) {
	f, err := os.OpenFile(".env.file", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0775)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range cases {
		fmt.Fprintf(f, "%s,%s\n", v[0], v[1])
	}
	f.Close()

	if err:=Load(".env.file");nil!=err{
		t.Fatal(err)
	}
	for _,kv:=range cases {
		if os.Getenv(kv[0]) != kv[1] {
			t.Fail()
		}
	}
	if err:=os.Remove(".env.file");nil!=err{
		t.Fail()
	}
}

func TestLoadWithSpaces(t *testing.T) {
	f, err := os.OpenFile(".env.file", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0775)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range cases {
		fmt.Fprintf(f, "%s   ,   %s    \n", v[0], v[1])
		//put spaces around
	}
	f.Close()

	if err:=Load(".env.file");nil!=err{
		t.Fatal(err)
	}
	for _,kv:=range cases {
		if os.Getenv(kv[0]) != kv[1] {
			t.Fail()
		}
	}
	if err:=os.Remove(".env.file");nil!=err{
		t.Fail()
	}
}

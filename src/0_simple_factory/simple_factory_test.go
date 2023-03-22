package simple_factory

import "testing"

func TestImplA(t *testing.T) {
	api := NewApi(1)
	s := "Hello"
	res := api.operation(s)
	if res != ("ImplA s==" + s) {
		t.Fatal("ImplA test fail")
	}
}

func TestImplB(t *testing.T) {
	api := NewApi(2)
	s := "Hello"
	res := api.operation(s)
	if res != ("ImplB s==" + s) {
		t.Fatal("ImplB test fail")
	}
}

package simple_factory

import "testing"

func Test1(t *testing.T) {
	api := NewAPI(1)
	println(api.Say("zxc1"))
}

func Test2(t *testing.T) {
	api := NewAPI(2)
	println(api.Say("zxc2"))
}

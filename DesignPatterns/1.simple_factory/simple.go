package simple_factory

type API interface {
	Say(name string) string
}

type HiAPI struct {
}

func (h *HiAPI) Say(name string) string {
	return name + ":Hi"
}

type HelloAPI struct {
}

func (h *HelloAPI) Say(name string) string {
	return name + ":Hello"
}
func NewAPI(t int) API {
	if t == 1 {
		return &HiAPI{}
	} else if t == 2 {
		return &HelloAPI{}
	}
	return nil
}

package simple_factory

import "fmt"

type Api interface {
	operation(s string) string
}

type ImplA struct{}

type ImplB struct{}

func (a *ImplA) operation(s string) string {
	return fmt.Sprintf("ImplA s==%s", s)
}

func (b *ImplB) operation(s string) string {
	return fmt.Sprintf("ImplB s==%s", s)
}

// NewApi 相当于工厂
func NewApi(t int) Api {
	var api Api
	if t == 1 {
		api = &ImplA{}
	} else if t == 2 {
		api = &ImplB{}
	}
	return api
}

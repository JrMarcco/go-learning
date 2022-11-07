package demo

import (
	"fmt"
	"go_learning/simple_web/framework"
)

type ServiceProviderDemo struct {
}

var _ framework.ServiceProvider = new(ServiceProviderDemo)

func (s *ServiceProviderDemo) Register(framework.Container) framework.NewInstance {
	return NewServiceDemo
}

func (s *ServiceProviderDemo) Boot(framework.Container) error {
	fmt.Println("boot func")
	return nil
}

func (s *ServiceProviderDemo) Params(container framework.Container) []any {
	return []any{container}
}

func (s *ServiceProviderDemo) IsDefer() bool {
	return true
}

func (s *ServiceProviderDemo) Name() string {
	return Key
}

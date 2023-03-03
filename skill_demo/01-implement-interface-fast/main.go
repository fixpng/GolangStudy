package main

type someInterface interface {
	DoSomething()
	DoAnotherThing()
}

type someStruct struct {
}

func (s someStruct) DoSomething() {
	//TODO implement me
	panic("implement me")
}

func (s someStruct) DoAnotherThing() {
	//TODO implement me
	panic("implement me")
}

// 如何快速实现接口 - how to quick impl someInterface for someStruct

// 1. imple someInterface for someStruct{}

// 2.
var _ someInterface = (*someStruct)(nil)

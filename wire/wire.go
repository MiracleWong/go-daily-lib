//+build wireinject

package main

import "github.com/google/wire"

func InitializeEvent(name string) Event {
	wire.Build(NewEvent, NewGreeter, NewMessage) // <--- 传入provider函数
	return Event{} //返回值没有实际意义，只需符合函数签名即可
}
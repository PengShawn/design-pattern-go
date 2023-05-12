package main

// Device 接收者接口
type Device interface {
	on()
	off()
}

package main

// OnCommand 具体命令接口
type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

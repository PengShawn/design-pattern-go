package main

// OffCommand 具体命令接口
type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

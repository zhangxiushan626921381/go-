package main

import "fmt"

//1.接口的实现
type USBer interface {
	read()
	write()
}

//2.创建对象
type UsbDev struct { // 设备
	id     int
	name   string
	rspeed int
	wspeed int
}
type mobile struct { // 手机
	UsbDev
	call string
}
type USB struct { // U盘
	UsbDev
}

//3.实现方法
func (m *mobile) read() {
	fmt.Printf("%s正在读取数据速度为:%d\n", m.name, m.rspeed)
}
func (m *mobile) write() {
	fmt.Printf("%s在写入数据速度为：%d\n", m.name, m.wspeed)
}
func (u *USB) read() {
	fmt.Printf("%s正在读取数据速度为:%d\n", u.name, u.rspeed)
}
func (u *USB) write() {
	fmt.Printf("%s在写入数据速度为：%d\n", u.name, u.wspeed)
}

//4.多态的实现 将接口作为函数的参数来传递
func Usbdev(usb USBer) {
	usb.read()
	usb.write()
}
func main() {
	//接口类型
	var usb USBer
	usb = &mobile{UsbDev{101, "诺基亚", 5, 10}, "大象"}
	Usbdev(usb)
	usb = &USB{UsbDev{102, "U盘", 4, 6}}
	Usbdev(usb)
}

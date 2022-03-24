package main

func main() {
	initClient()
	sunname := "mychannel"
	msg := "nihao"
	addsub(sunname)
	post(sunname, msg)
	unsub(sunname)
}

package main

var ch =make(chan int)
var cb =make(chan int)

var i int = 0
func main(){
	for j :=0;j<=49;j++ {
		go func() {
			println(i)
			i++
			if i == 100 {
				cb <- 1
			}
			ch <- 1
		}()
		go func() {
			<-ch
			println(i)
			i++
			if i == 100 {
				cb <- 1
			}
		}()
	}
	<-cb
}
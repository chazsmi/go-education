package main

func test() {
	// START DECLARE OMIT
	// Declaring a channel type
	var mychannel chan int
	// Initializing a channel
	mychannel = make(chan int)
	newchannel := make(chan int)

	// Sending on a channel
	// The arrow inticates the direction of data
	mychannel <- 1
	// receiving on a channel
	value = <-mychannel
	// END DECLARE OMIT
}

func main() {

}

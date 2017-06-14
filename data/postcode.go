package data

func ListNswPostcodes() <-chan int {
	firstPc := 1001
	lastPc := 3707
	pc := make(chan int, lastPc-firstPc)

	go func() {
		for x := firstPc; x <= lastPc; x++ {
			pc <- x
		}
		close(pc)
	}()
	return pc

}

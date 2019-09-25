package main

func fanIn(channels ...chan int) chan int {
	out := make(chan int)
	done := make(chan struct{})
	for _, c := range channels {
		go func(in chan int) {
			for num := range in {
				out <- num
			}
			done <- struct{}{}
		}(c)
	}
	go func() {
		for i := 0; i < len(channels); i++ {
			<-done
		}
		close(out)
	}()
	return out
}

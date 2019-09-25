package main

func fanOut(in chan int) chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- square(i)
		}
		close(out)
	}()
	return out
}

func square(num int) int {
	return num * num
}

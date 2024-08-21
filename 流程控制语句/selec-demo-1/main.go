package main

func main() {
	ch := make(chan int, 1)
	ch <- 100

	select {
	case i := <-ch:
		println("case1 recv: ", i)
	case i := <-ch:
		println("case2 recv: ", i)
	}
}

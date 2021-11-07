package Account

func CountTo(max int) (<-chan int, func()) {
	done := make(chan struct{})
	ch := make(chan int)
	cancel := func() {
		close(done)
	}
	go func(max int) {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			default:
				ch <- i
			}

		}
		close(ch)
	}(max)
	return ch, cancel
}

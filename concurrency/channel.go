package concurrency

// Match reads from an input channel and returns a channel that only contains
// the bytes that match the target string.
func Match(input chan byte, target string) chan byte {
	channel := make(chan byte, len(target))
	go transfer(input, channel, target)
	return channel
}

func transfer(input, output chan byte, target string) {
	var i int
	for v := range input {
		if v == target[i] {
			i++
			if i == len(target)-1 {
				close(output)
				return
			}
		} else {
			i = 0
		}
		output <- v
	}
	close(output)
}

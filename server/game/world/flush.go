package world

func (w *World) Flush() []int {
	all := make([]int, 0, len(w.updates))
	w.updates <- -1 // mark stop
	i := 0
	for u := 0; u != -1; u = <-w.updates {
		all = append(all, u)
		i += 1
	}

	return all
}

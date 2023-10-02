package progress

func ProgressChan[T any](ch <-chan T, options ...func(*Progress)) <-chan T {
	progressChan := make(chan T)
	progress := New(options...)
	go func() {
		for value := range ch {
			progress.Inc()
			progressChan <- value
		}
		progress.Done()
		close(progressChan)
	}()
	return progressChan
}

# progress

progress is a simple Go library that provides a minimalistic clone of the tqdm library for Python. It allows you to add progress bars to your Go applications to track the progress of tasks.

## Installation

You can install progress using `go get`:

```shell
go get github.com/albekov/progress
```

## Usage

Here's a basic example of how to use progress:

```go
progress := progress.New(progress.WithTotal(100))
for i := 0; i < 100; i++ {
    t := rand.Intn(100) + 5
    time.Sleep(time.Duration(t) * time.Millisecond)
    progress.Inc()
}
progress.Done()
```

The result is:

![](https://github.com/albekov/progress/blob/master/simple.gif)

See the [examples](examples) directory for more examples.

## License

progress is licensed under the MIT license. See the [LICENSE](LICENSE) file for details.

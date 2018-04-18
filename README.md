# random-string-stream

golang implementation example of stream and random

## Usage

```go
for s := range randomString.Stream() {
    fmt.Println(s)
}
```

## Command-Line

```go
randomString -n 3
a0f
...
```

## Installation

```go
$ go get github.com/nabetama-training/random-string-stream
```

## Author

Mao Nabeta [@nabetama](https://twitter.com/nabetama)

# Recipe - Using the common I/O interfaces

- `io.Reader` 🔗 (https://pkg.go.dev/io#Reader)
- `io.Writer` 🔗 (https://pkg.go.dev/io#Writer)

The Go language provides a number of I/O interfaces that are used throughout the standard library. It is best practice to make use of these interfaces wherever possible rather than passing structures or other types directly.

```go
type Reader interface {  
        Read(p []byte) (n int, err error)  
}

type Writer interface {  
        Write(p []byte) (n int, err error)  
}
```

Go makes it easy to combine interfaces.

- `io.Seeker` 🔗 (https://pkg.go.dev/io#Seeker)
- `io.ReadSeeker` 🔗 (https://pkg.go.dev/io#ReadSeeker)
- `io.MultiWriter` 🔗 (https://pkg.go.dev/io#MultiWriter)
- `io.Seek` 🔗 (https://pkg.go.dev/io#Seeker.Seek)
- `io.CopyBuffer` 🔗 (https://pkg.go.dev/io#CopyBuffer)

```go
type Seeker interface {  
        Seek(offset int64, whence int) (int64, error)  
}

type ReadSeeker interface {  
        Reader  
        Seeker  
}
```

- `io.Pipe` 🔗 (https://pkg.go.dev/io#Pipe)

```go
func Pipe() (*PipeReader, *PipeWriter)
```

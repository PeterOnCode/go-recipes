# Recipe - Using the bytes and strings packages

- https://pkg.go.dev/bytes#Buffer
- https://pkg.go.dev/bytes#Buffer.Write
- https://pkg.go.dev/bytes#NewBuffer
- https://pkg.go.dev/bytes#NewBufferString

The bytes library (see https://pkg.go.dev/bytes) provides a number of convenience functions when working with data. 

- A `buffer`, for example (see https://pkg.go.dev/bytes#Buffer), is far more flexible than an array of bytes when working with stream-processing libraries or methods. Once you've created a `buffer`, it can be used to satisfy an `io.Reader` interface so that you can take advantage of `io` functions to manipulate the data. 
- For streaming applications, you'd probably want to use a `buffer` and a `scanner` (see https://pkg.go.dev/bufio#Scanner). The `bufio` package (see https://pkg.go.dev/bufio) comes in handy for these cases. 
- Sometimes, using an array or slice is more appropriate for smaller datasets or when you have a lot of memory on your machine.

Go provides a lot of flexibility in converting data between interfaces when using these basic types — it's relatively simple to convert between strings and bytes. When working with strings, the `strings` package (see https://pkg.go.dev/strings) provides a number of convenience functions to work with, search, and manipulate strings. 

In some cases, a good regular expression may be appropriate, but most of the time, the `strings` and `strconv` (see https://pkg.go.dev/strconv) packages are sufficient. 

The `strings` package allows you to make a string look like a title (see https://pkg.go.dev/strings#Title), split it into an array (see https://pkg.go.dev/strings#Split), or trim whitespace (see https://pkg.go.dev/strings#Split). It also provides a `Reader` interface of its own that can be used instead of the `bytes` package reader type (see https://pkg.go.dev/bytes#Reader).


# Arrays

Go's most basic collections is an array. When you define an array, you must specify what type of data it may contain and how big the array is in the following form:

`[<size>]<type>`

For example,

1. `[10]int` is an array of size 10 that contains `int`s, 
2. while `[5]string` is an array of size 5 that contains `string`s.

>  The key to making this an array is specifying the size. If your definition didn't have the size, it would seem like it works, but it would not be an array - it'd be a slice.

You can set element values to be any type, including pointers and arrays.

## Overview of Go Arrays

Arrays in Go are fundamental data structures that store a fixed-size sequential collection of elements of the same type. Unlike slices, which are more flexible, arrays have a predetermined length defined at the time of declaration. This characteristic leads to both advantages and limitations in their usage.

### Characteristics of Go Arrays

1. **Fixed Size**: The size of an array is part of its type and cannot be changed after its declaration. This allows for efficient memory allocation but limits flexibility[2][3].
   
2. **Contiguous Memory Allocation**: Arrays store elements in contiguous memory locations, which facilitates fast access since the address of each element can be calculated based on its index[1].

3. **Value Type**: In Go, arrays are treated as value types. When an array is passed to a function, a copy of the entire array is made, not just a reference to its first element. This means modifications within the function do not affect the original array[1][4].

4. **Initialization**: Arrays can be initialized either by specifying all values or allowing the compiler to infer the size. For example:
   ```go
   var a [5]int          // Declares an array of 5 integers
   b := [5]int{1, 2, 3, 4, 5} // Initializes an array with specific values
   c := [...]int{1, 2, 3} // Compiler infers size
   ```

### Declaring and Initializing Arrays

The syntax for declaring an array in Go is:
```go
var name [length]type
```
For instance:
```go
var nums [5]int // An array of five integers
```
You can also initialize arrays directly:
```go
arr := [5]int{1, 2, 3, 4, 5}
```

### Array Operations

Go provides several built-in functions for working with arrays:

- **Length**: Use `len(array)` to get the number of elements.
- **Accessing Elements**: Elements can be accessed using their index (e.g., `array`).
- **Looping Through Arrays**: You can use a `for` loop or `for range` to iterate through elements.

Example:
```go
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}
```

### Best Practices and Performance Considerations

- **Memory Efficiency**: Since arrays have a fixed size and are allocated in contiguous memory, they can be more memory-efficient compared to slices when the size is known beforehand.
- **Avoiding Large Arrays on Stack**: If an array exceeds a certain size (e.g., 10 MB as per Go 1.23), it will be allocated on the heap instead of the stack[1].

### Common Pitfalls

- **Copying Behavior**: Be cautious about the copying behavior when passing arrays to functions; this can lead to unexpected results if modifications are made within the function.
- **Zero Values**: Uninitialized elements in an array default to their zero values (e.g., `0` for integers) which can lead to confusion if not accounted for[4].

In summary, understanding how arrays work in Go is crucial for effective programming. They offer predictable performance and memory usage but come with limitations regarding flexibility and mutability compared to slices.

Citations:
[1] https://victoriametrics.com/blog/go-array/
[2] https://codedamn.com/news/go/arrays-in-go
[3] https://www.digitalocean.com/community/tutorials/understanding-arrays-and-slices-in-go
[4] https://dev.to/dawkaka/go-arrays-and-slices-a-deep-dive-dp8

## Resources

1. A Tour of Go - Arrays (see https://go.dev/tour/moretypes/6)
2. [Arrays in Golang](https://golangdocs.com/arrays-in-golang)
3. [Go by Example: Arrays](https://gobyexample.com/arrays)
4. [Go Array - programiz.com](https://www.programiz.com/golang/arrays)

## Source Code

- [basic\_arrays.go](<./basic_arrays.go>)
- [basic\_arrays\_test.go](<./basic_arrays_test.go>)

```shell
go test .
```
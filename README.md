# ReadMe of Go Recipes

Recipes to build modular, readable, and testable Golang applications across various domains

## Recipes

### Complex Types

#### ↪️ Arrays

Go's most basic collections is an array. When you define an array, you must specify what type of data it may contain and how big the array is 

- [arrays](<./arrays/README.md>)

### I/O and Filesystems

#### ↪️ Using the common I/O interfaces

The Go language provides a number of I/O interfaces that are used throughout the standard library. It is best practice to make use of these interfaces wherever possible rather than passing structures or other types directly.

- [interfaces](<./interfaces/README.md>)

#### ↪️ Using the bytes and strings packages

The `bytes` and `strings` packages have a number of useful helpers to work with and convert the data from `string` to `byte` types, and vice versa. They allow the creation of buffers that work with a number of common I/O interfaces.

- [bytestrings](./bytestrings/README.md)


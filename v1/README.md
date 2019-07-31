# V1 - It sort of looks like Python...

The first version of our CSV system will merely use some built-in Go packages "encoding/csv" (standard library package for reading and writing CSV), "fmt" (standard library package for formatting and printing text representations of Go objects), "os" (standard library for operating system functionality), "strconv" (standard library package for casting and converting string representations into other data types).

## Concepts 
 
1. error handling
2. range
3. defer

### error handling

Error handling is considered an explicit, standard practice in Go.  In a recent [Go Time Podcast](https://changelog.com/gotime/91) Dave Cheney speaks directly to the explicit nature. 
> I think maybe what you’re alluding to is the fact that in Go error handling seems to be very explicit. It’s not just the technicalities of exceptions versus explicit return values, it’s more – at least to me, it’s the tradition that we have of… People often say, you think about the unhappy or the sad part first, and because the error handling is part of the return value, not any kind of additional mechanism, it’s kind of like the thing that you have to think about first.

The pattern `if err != nil {...}` is a standard idiom found in Go - "try this, but if you get an error here's what you do!"  Nearly all function calls in Go will return an error if something goes wrong, and when it does `err` will not be `nil`.  This is sharply in contrast to nearly all other OO-languages which offer variants of the try/catch idiom.

Instead of considering errors to be out of the ordinary, Go lets authors keep errors at the forefront of their programs.  Errors should be dealt with in explicit and clear ways that prevent undefined behaviour in later parts of the program.

### range
Go offers its users a convenient way to iterate over several types: arrays/slices, maps, strings, and channels.  Using the `range` keyword will unpack two values for each iteration over the iterable, and it varies depending on the type.  For simplicity, when iterating over an array the left value will be the current array index and the right value will be the value of the element at that index position.  When `range`-ing over a map the left value will be the key and the right value will be the object.  It is not uncommon to use an underscore to discard an object if it is not needed.  In this case the `v1/main.go` discards the index.

[Official notes](https://github.com/golang/go/wiki/Range) and [Go By Example](https://gobyexample.com/range).

### defer

If you've ever wanted to add a callback function to a registery that will run after a function has returned then you've come to the right place!  In Go `defer` will call the deferred function once the function has returned to the caller.  Most commonly associated with resource cleanup, it can also be used for locking (similar to a context manager in Python) or use cases sometimes associated with function decorators (such as timing/profiling a function call).

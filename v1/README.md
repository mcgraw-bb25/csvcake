# V1 - It sort of looks like Python...

The first version of our CSV system will merely use some built-in Go packages "encoding/csv" (standard library package for reading and writing CSV), "fmt" (standard library package for formatting and printing text representations of Go objects), "os" (standard library for operating system functionality), "strconv" (standard library package for casting and converting string representations into other data types).

## Concepts 
 
1. error handling
2. range
3. defer

### error handling

Error handling is considered an explicit, standard practice in Go.  In a recent [Go Time Podcast](https://changelog.com/gotime/91) Dave Cheney speaks directly to the explicit nature. 
> I think maybe what you’re alluding to is the fact that in Go error handling seems to be very explicit. It’s not just the technicalities of exceptions versus explicit return values, it’s more – at least to me, it’s the tradition that we have of… People often say, you think about the unhappy or the sad part first, and because the error handling is part of the return value, not any kind of additional mechanism, it’s kind of like the thing that you have to think about first.

### range

### defer

It is
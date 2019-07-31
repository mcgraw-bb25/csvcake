# V2 - A Simple Go Version of DictReader

The second version of our CSV system hides the underlying IO and file system details from us, aiming to replace the DictReader from Python via the MapReader.  Our goal is to return a map which allows for accessing values by their column headers, with all values typed as strings

## Concepts 
 
1. make builtin
2. returning errors
3. packages

### make builtin

Before you use an array, map, or channel you must use the builtin `make()` to allocate and initialize the object.  If you attempt you use any of these objects before calling make, the runtime will panic.  [The compiler will not help you](https://play.golang.org/p/5hxG5gLb1fK).  

### returning errors

When creating functions it is common, if not expected, that the function will return more than one parameter.  Most commonly it will be a value and an error, but that certainly isn't a strict rule.  Some functions only return errors, and others will return two or more values and an error.  It is rare for an idiomatic function in Go to not return an error.  When a function runs to completion and does not cause an error, `nil` is returned in place of an error.

Some functions will return errors that have been returned by other functions, merely passing them up the call stack.  Other times you can use the standard library package "errors" and create a new error using errors.New() or fmt.Errorf().  Finally there is Dave Cheney's library [pkg/errors](https://github.com/pkg/errors) which is especially useful for wrapping errors and lengthy article from Dave on [errors in Go.](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully). 

### packages

Here we also see our first use of a package.  (Note: Packages are not modules)  We reference the package in a (fairly) explicit manner - even though it's in a subfolder from `main.go` we are still specifying the full path for the package.  Once imported we can then access the exported types, interfaces, and functions from our main package.
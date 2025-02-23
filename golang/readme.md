Go is a procedural programming language4

it is dynmicly type ( y:=0)

Go is a statically typed, concurrent, and garbage-collected programming 


language Statically Typed

This means that variable types are checked at compile time, not at runtime.
Example: In Go, if you declare a variable as an int, you cannot later assign a string to it.
go
Copy
Edit
var x int = 10  // x is an integer
x = "hello"      // ❌ This will cause a compilation error
Concurrent

Go has built-in support for handling multiple tasks at the same time using goroutines.
Concurrency allows Go programs to perform multiple operations in parallel, making it great for high-performance applications like web servers or real-time systems.
Example:
go
Copy
Edit
go func() { 
    fmt.Println("Hello from Goroutine!") 
}()
The go keyword starts a new goroutine, running the function independently.
Garbage-Collected

Go has automatic memory management, meaning the programmer does not need to manually free memory.
The Garbage Collector (GC) automatically detects and removes unused objects from memory, preventing memory leaks.
This makes Go efficient but can also introduce some slight overhead compared to manual memory management (like in C/C++).


Go is known for its support for concurrency, which is the ability to run multiple tasks simultaneously. Concurrency is achieved in Go through the use of Goroutines and Channels, which allow you to write code that can run multiple operations at the same time. This makes Go an ideal choice for building high-performance and scalable network services, as well as for solving complex computational problems.


Basic type: Numbers, strings, and booleans come under this category.
Aggregate type: Array and structs come under this category.
Reference type: Pointers, slices, maps, functions, and channels come under this category.
Interface type
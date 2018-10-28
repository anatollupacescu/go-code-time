# Pointers

1. pointer receivers vs value receivers
1. value types vs pointer types vs reference types
1. collection of pointers vs collection of values

#### General guidelines for receivers:

When choosing the receiver type - "If in doubt, use a pointer" - or Dave Cheney's advice: `you should prefer declaring methods on *T unless you have a strong reason to do otherwise.`

Strong reasons for a value receiver:
 1. avoid getting things changed by surprise via a pointer
 1. the source for that value is changing: an instance of `time.Time`
 1. implement the `prototype` pattern

Before using a value receiver we must ask ourselves: If the method does not mutate its receiver, does it need to be a method?

Using values as function call arguments also makes sense if structs are small and likely to stay that way `Point2D{x,y}`

Use pointer receiver:
1. to change the internal state of the object.
1. when referencing to atomic types like `sync.Mutex`.
1. to signal to user that this variable is not supposed to be copied: `*fs.File` or `*testing.T`
1. you need to re-slice or reallocate the slice or otherwise change the underlying backing data structure.

## General advices regarding pointers
Don't use them only to improve speed, unless there's an explicit requirement and benchmarking.
Don't use pointers to basic data types unless `null` has semantic significance, or the pointer is required by the used library, like in the case of serialization
Slices or maps of pointers is a corner case, should be only use when there's a strong case for it, like the need to change the state of the individual elements in the collection. They don't normally add any benefits but do add overhead because of increased chance of cache misses.
Because in `Go` every assignment is a copy - if by some requirement you **must** use an array (and not a slice) then as a workaround it makes sense to use an array of pointers.

#### Pointers as arguments:
If you pass them as arguments the expectation is still that the client will not change the state the pointer is referencing.
#### Bonus:
How large is large? Assume it's equivalent to passing all its elements as arguments to the method. If that feels too large, it's also too large for the receiver.

## Links
1. https://www.callicoder.com/golang-pointers/ (the basics)
1. https://www.ardanlabs.com/blog/2014/12/using-pointers-in-go.html
1. https://go101.org/article/pointer.html
1. https://dave.cheney.net/2018/07/12/slices-from-the-ground-up

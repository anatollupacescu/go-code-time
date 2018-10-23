# Pointers

1. pointer receivers vs value receivers
1. value types vs pointer types vs reference types
1. collection of pointers vs collection of values

## Links
1. https://www.ardanlabs.com/blog/2014/12/using-pointers-in-go.html
1. https://go101.org/article/pointer.html

General guidelines for receivers:

Value receivers will give you a copy, so use it when you don't intend to change its state.

Pointer receiver are used when you need to change it's state.

If the receiver is a map, slice, func or chan - don't use a pointer to them because they are implemented with pointers internally.
One particular case to use pointer to slices would be if you need to re-slice or reallocate the slice.

However you must use pointers when referencing to atomic types like sync.Mutex, you don't want a copy of that.

Use pointers to signal to user that this variable is not supposed to be copied: *fs.File

Use value receivers to:
 1. avoid getting things changed by surprise via a pointer 
 1. the source for that value is changing: an iterator or an instance of time.Time
 1. structs are small and likely to stay that way ```Point2D{x,y}``` or are reference types: map, slice

Pointers to basic data types: serialization

Slices or maps of pointers is a corner case, should be only use when there's a strong case for it.
    
Bonus:
When choosing the receiver type - "If in doubt, use a pointer". However 

How large is large? Assume it's equivalent to passing all its elements as arguments to the method. If that feels too large, it's also too large for the receiver.
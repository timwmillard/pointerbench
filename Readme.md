# Go Struct Pointers vs Struct Literals Benchmark

The ago old debate about pointers vs literals.  Go makes it easy to swap betwen pointers and literal values.  You don't need to think about stranded pointers from
variable that get allocated on the stack.  Go takes care of all this for you.

## But which one should you use?

Like any good answers for a computer engineer, it depends!!!

## Pointers

The argument for pointers is that you don't need to copy the entier struct when passing and returing from a function call.  You only need to copy the pointer value and that's it.

## Literal Values

The argument for literal values, is that allocations on the stack are much cheaper that heap allocations.  When allocating data on the heap, the program first needs to request
memeory from the operation system.  Then once it's finished using the memeory it needs to be freed by the garbarge collector.

All this comes at a cost.  When allocation memory
on the stack, the program simple move the stack pointer and the memory is allocated or disgared.  The other disadvantage about pointers and heap allocations it the memeory is
spread out everywhere.  This leads to more L2 cache misses.  Where as storing all the data together in a struct, when you program reads the first btye from memory, the CPU also
fetchs other close by memory and stores it in L1, L2 and L3 caches.  Subsequention read are now much faster.  But if the next peice of data was located in a different location
(like where a pointer could be pointer), the CPU needs to go back to main memeroy to fetch the data.

## Benchmarks

The only real way to tell if you program is faster or slower using a certain technique is to benchmark it.  So I've created some benchmarks on copying different sized structs
using literal values and pointers.  I would recommend benchmarking you own program before making any decision and use this as a guide only.

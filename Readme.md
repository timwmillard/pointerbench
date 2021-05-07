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

```
goos: darwin
goarch: amd64
pkg: pointerbench
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkToRecords-8             	   76118	     14859 ns/op	   81920 B/op	       1 allocs/op
BenchmarkToRecordsPointer-8      	    7795	    164921 ns/op	  161920 B/op	   10001 allocs/op
BenchmarkToRecord8s-8            	    5164	    194936 ns/op	  647171 B/op	       1 allocs/op
BenchmarkToRecord8sPointer-8     	    3480	    337193 ns/op	  721924 B/op	   10001 allocs/op
BenchmarkToRecord16s-8           	    2889	    353150 ns/op	 1286148 B/op	       1 allocs/op
BenchmarkToRecord16sPointer-8    	    2334	    484936 ns/op	 1361924 B/op	   10001 allocs/op
BenchmarkToRecord24s-8           	    2094	    486829 ns/op	 1925122 B/op	       1 allocs/op
BenchmarkToRecord24sPointer-8    	    1695	    635061 ns/op	 2001923 B/op	   10001 allocs/op
BenchmarkToRecord32s-8           	    1567	    736798 ns/op	 2564108 B/op	       1 allocs/op
BenchmarkToRecord32sPointer-8    	    1574	    761292 ns/op	 2641925 B/op	   10001 allocs/op
BenchmarkToRecord64s-8           	     772	   1522499 ns/op	 5120006 B/op	       1 allocs/op
BenchmarkToRecord64sPointer-8    	     914	   1278165 ns/op	 5201924 B/op	   10001 allocs/op
BenchmarkToRecord128s-8          	     369	   3156291 ns/op	10240005 B/op	       1 allocs/op
BenchmarkToRecord128sPointer-8   	     531	   2246717 ns/op	10321923 B/op	   10001 allocs/op
PASS
ok  	pointerbench	18.188s
```
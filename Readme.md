# Random algorithms

This repo is a dumping ground of random algorithms I'm experimenting with.  

### csharpmergesort
This is a merge sort written in c#.  It takes a parameter called size to specify the array length and prints the time taken to sort (and whether the array is sorted).
```
cd csharpmergesort
dotnet run --size 4000000
```

### gomergesort
This is a merge sort written in go.  It has a normal non-concurrent function and a concurrent function.  The concurrent version is ... interesting.  With unbounded concurrency and given a large enough array size, it will exhaust all system ram (my machine has 16GB of ram and 2GB swap).  The level of "concurrency" is controlled by `minConcurrentSize` constant in `gomergesort/main.go`.  Once an array is less than that size, the algorithm uses the non-concurrent version of the function.  It does result in somewhat of a speed up as demonstrated by the benchmarks.  I've written some go benchmarks that can be run of increasing input sizes to see the differences (there's also two insertion sort test cases to show how slow it is).
```
cd gomergesort/cmd
go run .
go test -bench=.
```

### knightstour
This is a brute force search of the knights tour in go.  The program takes three parameters:
1. `-s <int>` to specify the size of the board (only works with square boards)
2. `-x <int>` to specify the starting x location
3. `-y <int>` to specify the starting y location

It will search all possible paths and output the successful paths (where every space has been visited), the total number of paths searched, and the total execution time.

It will spin up gofuncs until the path size is greater than two.  I haven't done the math on how many gofuncs this creates, but It's gonna be a lot (in the thousands?) and will saturate the CPU on most systems.  As with `gomergesort`, if you unbound the concurrency, the program will spin up so many gofuncs that it will consume ALL system ram and get killed.

```
cd knightstour/cmd
go run . -s 5 -x 0 -y 0
```

The output of the above looks like this:
```
done size: 5    Start: H8       complete paths: 304     searched paths: 625,308 execution time: 331.322066ms
```

`knightstour/cmd/script.sh` takes a parameter specifying the size and runs the program for all possible board locations.  Its output looks like this:
```
done size: 5    Start: H8       complete paths: 304     searched paths: 625,308 execution time: 324.918731ms
done size: 5    Start: H7       complete paths: 0       searched paths: 727,156 execution time: 389.442825ms
done size: 5    Start: H6       complete paths: 56      searched paths: 595,892 execution time: 326.768057ms
done size: 5    Start: H5       complete paths: 0       searched paths: 727,156 execution time: 388.322748ms
.
.
.
```

Since it's a brute force algorithm and the problem space grows exponentially, jumping from a board size of 5 to a board size of 6 causes ONE iteration to take an hour.
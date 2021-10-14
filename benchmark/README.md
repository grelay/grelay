# **Benchmakr**

This benchmark was based on [circuit](https://github.com/cep21/circuit) library.

## Running

```shell
$ make bench
```

Result:

```shell
> make bench
cd benchmark && go test -v -benchmem -run=^$ -bench=. ./...
goos: darwin
goarch: amd64
pkg: github.com/grelay/grelay/v1/benchmark
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz

BenchmarkCiruits/gRelay/Default/passing/1-12              	  701451	      1491 ns/op	     507 B/op	      10 allocs/op
BenchmarkCiruits/gRelay/Default/passing/75-12             	 4510262	       413.0 ns/op	     548 B/op	       9 allocs/op
BenchmarkCiruits/gRelay/Default/failing/1-12              	  778444	      1547 ns/op	     543 B/op	      10 allocs/op
BenchmarkCiruits/gRelay/Default/failing/75-12             	 5439175	       556.9 ns/op	     536 B/op	       9 allocs/op

BenchmarkCiruits/cep21-circuit/Hystrix/passing/1-12       	 1402143	       843.4 ns/op	     208 B/op	       4 allocs/op
BenchmarkCiruits/cep21-circuit/Hystrix/passing/75-12      	 5468440	       225.8 ns/op	     208 B/op	       4 allocs/op
BenchmarkCiruits/cep21-circuit/Hystrix/failing/1-12       	 8452128	       142.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/Hystrix/failing/75-12      	18480543	        63.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/Minimal/passing/1-12       	 3945703	       301.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/Minimal/passing/75-12      	17082249	        67.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/Minimal/failing/1-12       	11417625	       103.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/Minimal/failing/75-12      	74411128	        14.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/UseGo/passing/1-12         	  869540	      1463 ns/op	     296 B/op	       6 allocs/op
BenchmarkCiruits/cep21-circuit/UseGo/passing/75-12        	 3430490	       298.0 ns/op	     296 B/op	       6 allocs/op
BenchmarkCiruits/cep21-circuit/UseGo/failing/1-12         	  742660	      1525 ns/op	     296 B/op	       6 allocs/op
BenchmarkCiruits/cep21-circuit/UseGo/failing/75-12        	 4815441	       265.7 ns/op	     296 B/op	       6 allocs/op

BenchmarkCiruits/GoHystrix/DefaultConfig/passing/1-12     	  208128	      5819 ns/op	    1244 B/op	      22 allocs/op
BenchmarkCiruits/GoHystrix/DefaultConfig/passing/75-12    	  529312	      3147 ns/op	    1390 B/op	      24 allocs/op
BenchmarkCiruits/GoHystrix/DefaultConfig/failing/1-12     	  147070	      7312 ns/op	    1273 B/op	      23 allocs/op
BenchmarkCiruits/GoHystrix/DefaultConfig/failing/75-12    	  928563	      1103 ns/op	    1248 B/op	      24 allocs/op

BenchmarkCiruits/rubyist/Threshold-10/passing/1-12        	  701619	      1450 ns/op	     378 B/op	       6 allocs/op
BenchmarkCiruits/rubyist/Threshold-10/passing/75-12       	 1592112	       752.2 ns/op	     341 B/op	       6 allocs/op
BenchmarkCiruits/rubyist/Threshold-10/failing/1-12        	 9217720	       115.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/rubyist/Threshold-10/failing/75-12       	 8941886	       133.4 ns/op	       0 B/op	       0 allocs/op

BenchmarkCiruits/gobreaker/Default/passing/1-12           	 5994787	       197.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/gobreaker/Default/passing/75-12          	 3628560	       330.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/gobreaker/Default/failing/1-12           	10850232	       104.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/gobreaker/Default/failing/75-12          	 6940129	       173.9 ns/op	       0 B/op	       0 allocs/op

BenchmarkCiruits/handy/Default/passing/1-12               	 1233432	       967.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/handy/Default/passing/75-12              	 1071403	      1121 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/handy/Default/failing/1-12               	 1125232	      1050 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/handy/Default/failing/75-12              	  951085	      1113 ns/op	       0 B/op	       0 allocs/op

BenchmarkCiruits/iand_circuit/Default/passing/1-12        	17755350	        64.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/iand_circuit/Default/passing/75-12       	 9525399	       123.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/iand_circuit/Default/failing/1-12        	65656539	        15.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/iand_circuit/Default/failing/75-12       	532295302	         2.303 ns/op	       0 B/op	       0 allocs/op

PASS
ok  	github.com/grelay/grelay/v1/benchmark	63.483s
```
# **Benchmakr**

This benchmark was based on [circuit](https://github.com/cep21/circuit) library.

gRelay implementation is equivalent with `cep21-circuit/UseGo`.

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

BenchmarkCiruits/gRelay

BenchmarkCiruits/gRelay/Default/passing/1-12              	  742922	      1475 ns/op	     523 B/op	      11 allocs/op
BenchmarkCiruits/gRelay/Default/passing/75-12             	 4388548	       334.9 ns/op	     563 B/op	      10 allocs/op
BenchmarkCiruits/gRelay/Default/failing/1-12              	  781750	      1522 ns/op	     559 B/op	      11 allocs/op
BenchmarkCiruits/gRelay/Default/failing/75-12             	 5531700	       378.0 ns/op	     559 B/op	      11 allocs/op

BenchmarkCiruits/cep21-circuit

BenchmarkCiruits/cep21-circuit/Hystrix/passing/1-12       	 1506835	       796.3 ns/op	     208 B/op	       4 allocs/op
BenchmarkCiruits/cep21-circuit/Hystrix/passing/75-12      	 5424374	       220.1 ns/op	     208 B/op	       4 allocs/op
BenchmarkCiruits/cep21-circuit/Hystrix/failing/1-12       	 8733158	       137.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/Hystrix/failing/75-12      	18249513	        66.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/Minimal/passing/1-12       	 4036084	       301.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/Minimal/passing/75-12      	17819472	        63.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/Minimal/failing/1-12       	11114912	       101.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/Minimal/failing/75-12      	78535341	        14.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/UseGo/passing/1-12         	  770818	      1363 ns/op	     296 B/op	       6 allocs/op
BenchmarkCiruits/cep21-circuit/UseGo/passing/75-12        	 5468922	       218.2 ns/op	     296 B/op	       6 allocs/op
BenchmarkCiruits/cep21-circuit/UseGo/failing/1-12         	  858354	      1411 ns/op	     296 B/op	       6 allocs/op
BenchmarkCiruits/cep21-circuit/UseGo/failing/75-12        	 5368263	       224.0 ns/op	     296 B/op	       6 allocs/op

BenchmarkCiruits/GoHystrix

BenchmarkCiruits/GoHystrix/DefaultConfig/passing/1-12     	  206504	      5325 ns/op	    1249 B/op	      22 allocs/op
BenchmarkCiruits/GoHystrix/DefaultConfig/passing/75-12    	  471795	      2726 ns/op	    1320 B/op	      24 allocs/op
BenchmarkCiruits/GoHystrix/DefaultConfig/failing/1-12     	  160808	      6953 ns/op	    1267 B/op	      23 allocs/op
BenchmarkCiruits/GoHystrix/DefaultConfig/failing/75-12    	 1174665	      1023 ns/op	    1247 B/op	      24 allocs/op

BenchmarkCiruits/rubyist

BenchmarkCiruits/rubyist/Threshold-10/passing/1-12        	  749019	      1444 ns/op	     378 B/op	       6 allocs/op
BenchmarkCiruits/rubyist/Threshold-10/passing/75-12       	 1597680	       730.3 ns/op	     344 B/op	       5 allocs/op
BenchmarkCiruits/rubyist/Threshold-10/failing/1-12        	 9552349	       111.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/rubyist/Threshold-10/failing/75-12       	 9117228	       130.7 ns/op	       0 B/op	       0 allocs/op

BenchmarkCiruits/gobreaker

BenchmarkCiruits/gobreaker/Default/passing/1-12           	 6170560	       191.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/gobreaker/Default/passing/75-12          	 3798652	       316.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/gobreaker/Default/failing/1-12           	12112557	       100.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/gobreaker/Default/failing/75-12          	 7305661	       162.5 ns/op	       0 B/op	       0 allocs/op

BenchmarkCiruits/handy

BenchmarkCiruits/handy/Default/passing/1-12               	 1288095	       936.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/handy/Default/passing/75-12              	  982584	      1110 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/handy/Default/failing/1-12               	  967164	      1035 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/handy/Default/failing/75-12              	 1000000	      1071 ns/op	       0 B/op	       0 allocs/op

BenchmarkCiruits/iand_circuit

BenchmarkCiruits/iand_circuit/Default/passing/1-12        	18967089	        61.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/iand_circuit/Default/passing/75-12       	 9826328	       122.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/iand_circuit/Default/failing/1-12        	69069673	        15.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/iand_circuit/Default/failing/75-12       	562519341	         2.130 ns/op	       0 B/op	       0 allocs/op

PASS
ok  	github.com/grelay/grelay/v1/benchmark	57.141s
```
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
BenchmarkCiruits
BenchmarkCiruits/cep21-circuit/Hystrix/passing/1-12       	 1438222	       809.1 ns/op	     208 B/op	       4 allocs/op
BenchmarkCiruits/cep21-circuit/Hystrix/passing/75-12      	 5423850	       231.7 ns/op	     208 B/op	       4 allocs/op
BenchmarkCiruits/cep21-circuit/Hystrix/failing/1-12       	 8839462	       137.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/Hystrix/failing/75-12      	20404459	        56.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/Minimal/passing/1-12       	 3923539	       303.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/Minimal/passing/75-12      	16548805	        71.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/Minimal/failing/1-12       	10558832	       102.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/Minimal/failing/75-12      	76433164	        14.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/cep21-circuit/UseGo/passing/1-12         	  792894	      1375 ns/op	     296 B/op	       6 allocs/op
BenchmarkCiruits/cep21-circuit/UseGo/passing/75-12        	 5771398	       215.3 ns/op	     296 B/op	       5 allocs/op
BenchmarkCiruits/cep21-circuit/UseGo/failing/1-12         	  817326	      1410 ns/op	     296 B/op	       6 allocs/op
BenchmarkCiruits/cep21-circuit/UseGo/failing/75-12        	 5632882	       221.7 ns/op	     296 B/op	       6 allocs/op

BenchmarkCiruits/GoHystrix/DefaultConfig/passing/1-12     	  206464	      5504 ns/op	    1245 B/op	      22 allocs/op
BenchmarkCiruits/GoHystrix/DefaultConfig/passing/75-12    	  444634	      2312 ns/op	    1291 B/op	      24 allocs/op
BenchmarkCiruits/GoHystrix/DefaultConfig/failing/1-12     	  157387	      7402 ns/op	    1265 B/op	      23 allocs/op
BenchmarkCiruits/GoHystrix/DefaultConfig/failing/75-12    	 1000000	      1110 ns/op	    1247 B/op	      24 allocs/op

BenchmarkCiruits/rubyist/Threshold-10/passing/1-12        	  801668	      1495 ns/op	     377 B/op	       6 allocs/op
BenchmarkCiruits/rubyist/Threshold-10/passing/75-12       	 1549971	       749.0 ns/op	     342 B/op	       6 allocs/op
BenchmarkCiruits/rubyist/Threshold-10/failing/1-12        	 9462261	       114.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/rubyist/Threshold-10/failing/75-12       	 7272706	       159.7 ns/op	       0 B/op	       0 allocs/op

BenchmarkCiruits/gobreaker/Default/passing/1-12           	 5946483	       200.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/gobreaker/Default/passing/75-12          	 3588241	       325.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/gobreaker/Default/failing/1-12           	11549697	       101.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/gobreaker/Default/failing/75-12          	 7034408	       174.7 ns/op	       0 B/op	       0 allocs/op

BenchmarkCiruits/handy/Default/passing/1-12               	 1208620	       981.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/handy/Default/passing/75-12              	  935779	      1128 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/handy/Default/failing/1-12               	 1000000	      1064 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/handy/Default/failing/75-12              	  989696	      1128 ns/op	       0 B/op	       0 allocs/op

BenchmarkCiruits/iand_circuit/Default/passing/1-12        	18387626	        63.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/iand_circuit/Default/passing/75-12       	 9865338	       120.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/iand_circuit/Default/failing/1-12        	69672188	        16.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkCiruits/iand_circuit/Default/failing/75-12       	456441234	         2.665 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/grelay/grelay/v1/benchmark	45.692s
```
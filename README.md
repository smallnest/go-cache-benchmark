## Go Cache Benchmarks

Benchmarks of in-memory cache libraries for Go with expiration/TTL support.


**Notice**: this project is based on [Xeoncross/go-cache-benchmark](https://github.com/Xeoncross/go-cache-benchmark) and added multiple-size-values benchmark.

## Run Test

    go get
    go test -bench=. -benchmem

## Results

```
goos: darwin
goarch: amd64
pkg: github.com/smallnest/go-cache-benchmark
BenchmarkCache/Kodingcache-64/Set-4     	 1000000	      1622 ns/op	     400 B/op	       3 allocs/op
BenchmarkCache/Kodingcache-64/Get-4     	 1752364	       618 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/HashicorpLRU-64/Set-4    	 2297848	       529 ns/op	     142 B/op	       6 allocs/op
BenchmarkCache/HashicorpLRU-64/Get-4    	 5118595	       243 ns/op	      40 B/op	       2 allocs/op
BenchmarkCache/Cache2Go-64/Set-4        	 1000000	      1549 ns/op	     346 B/op	       7 allocs/op
BenchmarkCache/Cache2Go-64/Get-4        	 3249121	       357 ns/op	      39 B/op	       2 allocs/op
BenchmarkCache/GoCache-64/Set-4         	 1000000	      1285 ns/op	     239 B/op	       3 allocs/op
BenchmarkCache/GoCache-64/Get-4         	 4058758	       284 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/Freecache-64/Set-4       	 2216034	       540 ns/op	      24 B/op	       1 allocs/op
BenchmarkCache/Freecache-64/Get-4       	 5191746	       230 ns/op	      24 B/op	       2 allocs/op
BenchmarkCache/BigCache-64/Set-4        	 1838832	       671 ns/op	      62 B/op	       2 allocs/op
BenchmarkCache/BigCache-64/Get-4        	 2493453	       518 ns/op	      82 B/op	       3 allocs/op
BenchmarkCache/Ristretto-64/Set-4       	 1209412	      1066 ns/op	     208 B/op	       5 allocs/op
BenchmarkCache/Ristretto-64/Get-4       	 2194956	       724 ns/op	      46 B/op	       3 allocs/op
BenchmarkCache/SyncMap-64/Set-4         	 1000000	      1057 ns/op	     218 B/op	       7 allocs/op
BenchmarkCache/SyncMap-64/Get-4         	 4655073	       249 ns/op	      23 B/op	       1 allocs/op

BenchmarkCache/Kodingcache-1024/Set-4   	 1000000	      1511 ns/op	     400 B/op	       3 allocs/op
BenchmarkCache/Kodingcache-1024/Get-4   	 2316610	       578 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/HashicorpLRU-1024/Set-4  	 2355328	       505 ns/op	     142 B/op	       6 allocs/op
BenchmarkCache/HashicorpLRU-1024/Get-4  	 5733535	       206 ns/op	      39 B/op	       2 allocs/op
BenchmarkCache/Cache2Go-1024/Set-4      	 2112726	      1522 ns/op	     282 B/op	       7 allocs/op
BenchmarkCache/Cache2Go-1024/Get-4      	 2513055	       444 ns/op	      39 B/op	       2 allocs/op
BenchmarkCache/GoCache-1024/Set-4       	 2252820	       736 ns/op	     170 B/op	       5 allocs/op
BenchmarkCache/GoCache-1024/Get-4       	 3257652	       346 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/Freecache-1024/Set-4     	 2258368	       526 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/Freecache-1024/Get-4     	 7151340	       172 ns/op	      24 B/op	       2 allocs/op
BenchmarkCache/BigCache-1024/Set-4      	 1705735	       745 ns/op	      40 B/op	       3 allocs/op
BenchmarkCache/BigCache-1024/Get-4      	 4658552	       309 ns/op	      86 B/op	       2 allocs/op
BenchmarkCache/Ristretto-1024/Set-4     	 1891932	       900 ns/op	     367 B/op	       5 allocs/op
BenchmarkCache/Ristretto-1024/Get-4     	 2362500	       528 ns/op	      47 B/op	       3 allocs/op
BenchmarkCache/SyncMap-1024/Set-4       	 1000000	      1040 ns/op	     249 B/op	       8 allocs/op
BenchmarkCache/SyncMap-1024/Get-4       	 4766206	       275 ns/op	      23 B/op	       1 allocs/op

BenchmarkCache/Kodingcache-10240/Set-4  	 1000000	      1082 ns/op	     400 B/op	       3 allocs/op
BenchmarkCache/Kodingcache-10240/Get-4  	 1968775	       567 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/HashicorpLRU-10240/Set-4 	 2414696	       501 ns/op	     142 B/op	       6 allocs/op
BenchmarkCache/HashicorpLRU-10240/Get-4 	 5767788	       211 ns/op	      39 B/op	       2 allocs/op
BenchmarkCache/Cache2Go-10240/Set-4     	 2043338	       576 ns/op	     223 B/op	       6 allocs/op
BenchmarkCache/Cache2Go-10240/Get-4     	 2699784	       414 ns/op	      39 B/op	       2 allocs/op
BenchmarkCache/GoCache-10240/Set-4      	 1827577	       678 ns/op	     184 B/op	       4 allocs/op
BenchmarkCache/GoCache-10240/Get-4      	 3087848	       361 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/Freecache-10240/Set-4    	 8182722	       146 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/Freecache-10240/Get-4    	 7280942	       176 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/BigCache-10240/Set-4     	  559812	      2109 ns/op	      79 B/op	       2 allocs/op
BenchmarkCache/BigCache-10240/Get-4     	 4213150	       264 ns/op	      91 B/op	       2 allocs/op
BenchmarkCache/Ristretto-10240/Set-4    	 1823619	       987 ns/op	     358 B/op	       5 allocs/op
BenchmarkCache/Ristretto-10240/Get-4    	 2379097	       499 ns/op	      47 B/op	       3 allocs/op
BenchmarkCache/SyncMap-10240/Set-4      	 2659936	       707 ns/op	     131 B/op	       6 allocs/op
BenchmarkCache/SyncMap-10240/Get-4      	 3405146	       320 ns/op	      23 B/op	       1 allocs/op

BenchmarkCache/Kodingcache-1048576/Set-4         	 1000000	      1049 ns/op	     400 B/op	       3 allocs/op
BenchmarkCache/Kodingcache-1048576/Get-4         	 2250337	       575 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/HashicorpLRU-1048576/Set-4        	 2433124	       509 ns/op	     142 B/op	       6 allocs/op
BenchmarkCache/HashicorpLRU-1048576/Get-4        	 5649486	       224 ns/op	      39 B/op	       2 allocs/op
BenchmarkCache/Cache2Go-1048576/Set-4            	 2206428	       592 ns/op	     223 B/op	       6 allocs/op
BenchmarkCache/Cache2Go-1048576/Get-4            	 2615947	       431 ns/op	      39 B/op	       2 allocs/op
BenchmarkCache/GoCache-1048576/Set-4             	 1819345	       670 ns/op	     185 B/op	       4 allocs/op
BenchmarkCache/GoCache-1048576/Get-4             	 3089766	       395 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/Freecache-1048576/Set-4           	 8273787	       159 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/Freecache-1048576/Get-4           	 7180322	       164 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/BigCache-1048576/Set-4            	    6937	    373832 ns/op	  260905 B/op	       5 allocs/op
BenchmarkCache/BigCache-1048576/Get-4            	 6896461	       188 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/Ristretto-1048576/Set-4           	 1000000	      1109 ns/op	     337 B/op	       5 allocs/op
BenchmarkCache/Ristretto-1048576/Get-4           	 2890052	       422 ns/op	      47 B/op	       3 allocs/op
BenchmarkCache/SyncMap-1048576/Set-4             	 2587011	       693 ns/op	     132 B/op	       6 allocs/op
BenchmarkCache/SyncMap-1048576/Get-4             	 3461686	       311 ns/op	      23 B/op	       1 allocs/op
```

![benchmark](benchmarks.png)

_Note: sync.Map does not support expiration and is only included for comparison_
_Note: hashicorp/golang-lru does not support expires/TTL and is only included for comparison_


## Warning

Please note that these libraries are benchmarked against storage of small strings. If you are storing large blobs the results _will_ differ and you are encouraged to benchmark with your custom payloads.

- https://github.com/dgraph-io/ristretto
- https://golang.org/pkg/sync/#Map
- https://github.com/coocood/freecache
- https://github.com/allegro/bigcache
- https://github.com/patrickmn/go-cache
- https://github.com/muesli/cache2go
- https://github.com/bluele/gcache

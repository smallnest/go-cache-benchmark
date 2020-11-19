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
BenchmarkCache/Kodingcache-64/Set-4     	 1000000	      1192 ns/op	     400 B/op	       3 allocs/op
BenchmarkCache/Kodingcache-64/Get-4     	 2290170	       583 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/HashicorpLRU-64/Set-4    	 2297430	       501 ns/op	     142 B/op	       6 allocs/op
BenchmarkCache/HashicorpLRU-64/Get-4    	 5814535	       206 ns/op	      40 B/op	       2 allocs/op
BenchmarkCache/Cache2Go-64/Set-4        	 1000000	      1240 ns/op	     346 B/op	       7 allocs/op
BenchmarkCache/Cache2Go-64/Get-4        	 3480003	       329 ns/op	      39 B/op	       2 allocs/op
BenchmarkCache/GoCache-64/Set-4         	 1000000	      1024 ns/op	     239 B/op	       3 allocs/op
BenchmarkCache/GoCache-64/Get-4         	 3349364	       302 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/Freecache-64/Set-4       	 2328043	       518 ns/op	      24 B/op	       1 allocs/op
BenchmarkCache/Freecache-64/Get-4       	 5724613	       211 ns/op	      24 B/op	       2 allocs/op
BenchmarkCache/BigCache-64/Set-4        	 2134650	       621 ns/op	      57 B/op	       2 allocs/op
BenchmarkCache/BigCache-64/Get-4        	 2870516	       406 ns/op	      83 B/op	       3 allocs/op
BenchmarkCache/Ristretto-64/Set-4       	 1252058	      1033 ns/op	     205 B/op	       5 allocs/op
BenchmarkCache/Ristretto-64/Get-4       	 2196135	       471 ns/op	      47 B/op	       3 allocs/op
BenchmarkCache/SyncMap-64/Set-4         	 2465827	       745 ns/op	     134 B/op	       6 allocs/op
BenchmarkCache/SyncMap-64/Get-4         	 3567628	       307 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/ConcurrentMap-64/Set-4   	 1986038	       539 ns/op	     121 B/op	       3 allocs/op
BenchmarkCache/ConcurrentMap-64/Get-4   	 3881145	       306 ns/op	      23 B/op	       1 allocs/op

BenchmarkCache/Kodingcache-1024/Set-4   	 1000000	      1104 ns/op	     400 B/op	       3 allocs/op
BenchmarkCache/Kodingcache-1024/Get-4   	 2167081	       639 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/HashicorpLRU-1024/Set-4  	 2345680	       522 ns/op	     143 B/op	       6 allocs/op
BenchmarkCache/HashicorpLRU-1024/Get-4  	 5468550	       213 ns/op	      39 B/op	       2 allocs/op
BenchmarkCache/Cache2Go-1024/Set-4      	 2216598	      1074 ns/op	     279 B/op	       7 allocs/op
BenchmarkCache/Cache2Go-1024/Get-4      	 2642286	       430 ns/op	      39 B/op	       2 allocs/op
BenchmarkCache/GoCache-1024/Set-4       	 2263203	       727 ns/op	     170 B/op	       5 allocs/op
BenchmarkCache/GoCache-1024/Get-4       	 3338199	       340 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/Freecache-1024/Set-4     	 2242026	       529 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/Freecache-1024/Get-4     	 6950230	       184 ns/op	      24 B/op	       2 allocs/op
BenchmarkCache/BigCache-1024/Set-4      	 1534689	       835 ns/op	      40 B/op	       3 allocs/op
BenchmarkCache/BigCache-1024/Get-4      	 4195605	       346 ns/op	      93 B/op	       2 allocs/op
BenchmarkCache/Ristretto-1024/Set-4     	 1808474	       889 ns/op	     342 B/op	       5 allocs/op
BenchmarkCache/Ristretto-1024/Get-4     	 2489349	       516 ns/op	      50 B/op	       3 allocs/op
BenchmarkCache/SyncMap-1024/Set-4       	 2732754	       698 ns/op	     131 B/op	       6 allocs/op
BenchmarkCache/SyncMap-1024/Get-4       	 3242845	       326 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/ConcurrentMap-1024/Set-4 	 2112396	       532 ns/op	     116 B/op	       3 allocs/op
BenchmarkCache/ConcurrentMap-1024/Get-4 	 4004212	       312 ns/op	      23 B/op	       1 allocs/op

BenchmarkCache/Kodingcache-10240/Set-4  	 1000000	      1050 ns/op	     400 B/op	       3 allocs/op
BenchmarkCache/Kodingcache-10240/Get-4  	 2051797	       585 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/HashicorpLRU-10240/Set-4 	 2322882	       517 ns/op	     142 B/op	       6 allocs/op
BenchmarkCache/HashicorpLRU-10240/Get-4 	 5475622	       218 ns/op	      39 B/op	       2 allocs/op
BenchmarkCache/Cache2Go-10240/Set-4     	 2196616	       616 ns/op	     223 B/op	       6 allocs/op
BenchmarkCache/Cache2Go-10240/Get-4     	 2430381	       461 ns/op	      39 B/op	       2 allocs/op
BenchmarkCache/GoCache-10240/Set-4      	 1789729	       652 ns/op	     187 B/op	       4 allocs/op
BenchmarkCache/GoCache-10240/Get-4      	 2886171	       392 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/Freecache-10240/Set-4    	 7992661	       151 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/Freecache-10240/Get-4    	 7124287	       168 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/BigCache-10240/Set-4     	  564171	      2305 ns/op	      78 B/op	       2 allocs/op
BenchmarkCache/BigCache-10240/Get-4     	 4267731	       278 ns/op	      90 B/op	       2 allocs/op
BenchmarkCache/Ristretto-10240/Set-4    	 1806175	       959 ns/op	     352 B/op	       5 allocs/op
BenchmarkCache/Ristretto-10240/Get-4    	 2410005	       501 ns/op	      47 B/op	       3 allocs/op
BenchmarkCache/SyncMap-10240/Set-4      	 2689389	       702 ns/op	     131 B/op	       6 allocs/op
BenchmarkCache/SyncMap-10240/Get-4      	 3484850	       316 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/ConcurrentMap-10240/Set-4         	 2216274	       509 ns/op	     113 B/op	       3 allocs/op
BenchmarkCache/ConcurrentMap-10240/Get-4         	 4133799	       305 ns/op	      23 B/op	       1 allocs/op

BenchmarkCache/Kodingcache-1048576/Set-4         	 1000000	      1136 ns/op	     400 B/op	       3 allocs/op
BenchmarkCache/Kodingcache-1048576/Get-4         	 2266717	       582 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/HashicorpLRU-1048576/Set-4        	 2394577	       507 ns/op	     142 B/op	       6 allocs/op
BenchmarkCache/HashicorpLRU-1048576/Get-4        	 5403476	       225 ns/op	      39 B/op	       2 allocs/op
BenchmarkCache/Cache2Go-1048576/Set-4            	 2098029	       578 ns/op	     223 B/op	       6 allocs/op
BenchmarkCache/Cache2Go-1048576/Get-4            	 2624553	       425 ns/op	      39 B/op	       2 allocs/op
BenchmarkCache/GoCache-1048576/Set-4             	 1805330	       654 ns/op	     186 B/op	       4 allocs/op
BenchmarkCache/GoCache-1048576/Get-4             	 2636514	       390 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/Freecache-1048576/Set-4           	 7822702	       181 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/Freecache-1048576/Get-4           	 6129246	       171 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/BigCache-1048576/Set-4            	    7149	    417176 ns/op	  253171 B/op	       5 allocs/op
BenchmarkCache/BigCache-1048576/Get-4            	 6234769	       186 ns/op	      38 B/op	       2 allocs/op
BenchmarkCache/Ristretto-1048576/Set-4           	  909507	      1228 ns/op	     364 B/op	       5 allocs/op
BenchmarkCache/Ristretto-1048576/Get-4           	 2002611	       621 ns/op	      47 B/op	       3 allocs/op
BenchmarkCache/SyncMap-1048576/Set-4             	 1000000	      1105 ns/op	     218 B/op	       7 allocs/op
BenchmarkCache/SyncMap-1048576/Get-4             	 3761005	       302 ns/op	      23 B/op	       1 allocs/op
BenchmarkCache/ConcurrentMap-1048576/Set-4       	 2110754	       579 ns/op	     117 B/op	       3 allocs/op
BenchmarkCache/ConcurrentMap-1048576/Get-4       	 3638185	       342 ns/op	      23 B/op	       1 allocs/op

BenchmarkConcurrentMap/Set-4                     	 2136591	       484 ns/op	     116 B/op	       3 allocs/op
BenchmarkConcurrentMap/Get-4                     	 3961012	       332 ns/op	      23 B/op	       1 allocs/op
```


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

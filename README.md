# Neighbor search

```
$ go test -run='^$' -benchmem -bench "."
goos: darwin
goarch: amd64
BenchmarkV1_0/tolerance=0-8                 1168           1026170 ns/op               0 B/op          0 allocs/op
BenchmarkV1_0/tolerance=1-8                 1134           1031805 ns/op               0 B/op          0 allocs/op
BenchmarkV1_0/tolerance=2-8                 1075           1022265 ns/op               0 B/op          0 allocs/op
BenchmarkV1_0/tolerance=4-8                 1057           1015113 ns/op               0 B/op          0 allocs/op
BenchmarkV1_0/tolerance=8-8                 1048           1021607 ns/op               0 B/op          0 allocs/op
BenchmarkV1_0/tolerance=16-8                1018           1025472 ns/op               0 B/op          0 allocs/op
BenchmarkV1_1/tolerance=0-8                 1497            701109 ns/op               0 B/op          0 allocs/op
BenchmarkV1_1/tolerance=1-8                 1492            696637 ns/op               0 B/op          0 allocs/op
BenchmarkV1_1/tolerance=2-8                 1479            690664 ns/op               0 B/op          0 allocs/op
BenchmarkV1_1/tolerance=4-8                 1573            688512 ns/op               0 B/op          0 allocs/op
BenchmarkV1_1/tolerance=8-8                 1542            688857 ns/op               0 B/op          0 allocs/op
BenchmarkV1_1/tolerance=16-8                1556            675998 ns/op               0 B/op          0 allocs/op
BenchmarkV2_0/tolerance=0-8               393519              2939 ns/op               0 B/op          0 allocs/op
BenchmarkV2_0/tolerance=1-8                24608             49682 ns/op               0 B/op          0 allocs/op
BenchmarkV2_0/tolerance=2-8                 2776            449351 ns/op               0 B/op          0 allocs/op
BenchmarkV2_0/tolerance=4-8                  172           6866302 ns/op               0 B/op          0 allocs/op
BenchmarkV2_0/tolerance=8-8                   20          84607615 ns/op               0 B/op          0 allocs/op
BenchmarkV2_0/tolerance=16-8                   6         189268543 ns/op               0 B/op          0 allocs/op
BenchmarkV2_1/tolerance=0-8               341040              3450 ns/op              47 B/op          5 allocs/op
BenchmarkV2_1/tolerance=1-8                24624             51168 ns/op            2461 B/op         12 allocs/op
BenchmarkV2_1/tolerance=2-8                 2986            392376 ns/op           25554 B/op         14 allocs/op
BenchmarkV2_1/tolerance=4-8                  204           6450680 ns/op          925620 B/op         27 allocs/op
BenchmarkV2_1/tolerance=8-8                   27          73891654 ns/op        14627411 B/op         39 allocs/op
BenchmarkV2_1/tolerance=16-8                   7         158200439 ns/op        38981972 B/op         44 allocs/op
BenchmarkV3_0/tolerance=0-8                 2073            578176 ns/op               0 B/op          0 allocs/op
BenchmarkV3_0/tolerance=1-8                 2186            544949 ns/op               0 B/op          0 allocs/op
BenchmarkV3_0/tolerance=2-8                 2205            545329 ns/op               0 B/op          0 allocs/op
BenchmarkV3_0/tolerance=4-8                 2217            541437 ns/op               0 B/op          0 allocs/op
BenchmarkV3_0/tolerance=8-8                 2234            541450 ns/op               0 B/op          0 allocs/op
BenchmarkV3_0/tolerance=16-8                2194            535396 ns/op               0 B/op          0 allocs/op
BenchmarkV4_0/tolerance=0-8                   66          19598465 ns/op             861 B/op          4 allocs/op
BenchmarkV4_0/tolerance=1-8                   64          20000143 ns/op             490 B/op          3 allocs/op
BenchmarkV4_0/tolerance=2-8                   57          19591752 ns/op             249 B/op          2 allocs/op
BenchmarkV4_0/tolerance=4-8                   63          19830432 ns/op              58 B/op          2 allocs/op
BenchmarkV4_0/tolerance=8-8                   55          19976722 ns/op              92 B/op          2 allocs/op
BenchmarkV4_0/tolerance=16-8                  66          19959907 ns/op             139 B/op          2 allocs/op
PASS
```

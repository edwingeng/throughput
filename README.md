# Overview
This is a simple throughput comparison of serveral socket implementations. I hope you can get a rough idea about their performance differences.

# Results
### go
```
2017/11/18 10:53:18 size: 128, n: 1000000
2017/11/18 10:53:24 n: 1000000, time: 5.489, throughput: 182186

2017/11/18 10:53:26 size: 128, n: 1000000, nodelay
2017/11/18 10:53:40 n: 1000000, time: 13.912, throughput: 71879

```

### go-zmq
```
2017/11/18 10:52:22 size: 128, n: 1000000, nodelay
2017/11/18 10:52:24 n: 1000000, time: 2.346, throughput: 426226
```

### c
```
2017/11/18 18:11:05 size: 128, n: 1000000
2017/11/18 18:11:07 n: 1000000, time: 0.990, throughput: 1010468

2017/11/18 18:11:09 size: 128, n: 1000000, nodelay
2017/11/18 18:11:15 n: 1000000, time: 2.811, throughput: 355721
```

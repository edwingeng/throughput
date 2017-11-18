# Overview
This is a simple throughput comparison of serveral TCP c/s implementations. I hope you can get a rough idea about their performance differences.

# Results
### go
Comparison group: A, B
```
2017/11/18 10:53:18 size: 128, n: 1000000
2017/11/18 10:53:24 n: 1000000, time: 5.489, throughput: 182186

2017/11/18 10:53:26 size: 128, n: 1000000, nodelay
2017/11/18 10:53:40 n: 1000000, time: 13.912, throughput: 71879

```

### go-zmq
Comparison group: A
```
2017/11/18 10:52:22 size: 128, n: 1000000, nodelay
2017/11/18 10:52:24 n: 1000000, time: 2.346, throughput: 426226
```

### c
Comparison group: A
```
2017/11/18 18:29:28 size: 128, n: 1000000
2017/11/18 18:29:32 n: 1000000, time: 1.813, throughput: 551612

2017/11/18 18:29:34 size: 128, n: 1000000, nodelay
2017/11/18 18:29:46 n: 1000000, time: 5.736, throughput: 174329
```

### go-one-off
Comparison group: B
```
2017/11/18 18:41:24 size: 128, n: 1000000
2017/11/18 18:41:27 n: 1000000, time: 2.928, throughput: 341507

2017/11/18 18:41:29 size: 128, n: 1000000, nodelay
2017/11/18 18:41:36 n: 1000000, time: 7.164, throughput: 139593
```

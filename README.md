# Overview
This is a simple throughput comparison of serveral TCP c/s implementations. I hope you can get a rough idea about their performance differences.

**These toys cannot measure the performance of receiving messages on the server side.**

# Comparison Group A

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
2017/11/18 18:29:28 size: 128, n: 1000000
2017/11/18 18:29:32 n: 1000000, time: 1.813, throughput: 551612

2017/11/18 18:29:34 size: 128, n: 1000000, nodelay
2017/11/18 18:29:46 n: 1000000, time: 5.736, throughput: 174329
```

# Comparison Group B

### go
```
2017/11/18 10:53:18 size: 128, n: 1000000
2017/11/18 10:53:24 n: 1000000, time: 5.489, throughput: 182186

2017/11/18 10:53:26 size: 128, n: 1000000, nodelay
2017/11/18 10:53:40 n: 1000000, time: 13.912, throughput: 71879
```

### go-one-off
```
2017/11/18 18:41:24 size: 128, n: 1000000
2017/11/18 18:41:27 n: 1000000, time: 2.928, throughput: 341507

2017/11/18 18:41:29 size: 128, n: 1000000, nodelay
2017/11/18 18:41:36 n: 1000000, time: 7.164, throughput: 139593
```

# Comparison Group C

### go-one-off
```
2017/11/19 10:55:40 size: 128, n: 1000000
2017/11/19 10:55:40 size: 128, n: 1000000
2017/11/19 10:55:40 size: 128, n: 1000000
2017/11/19 10:55:44 n: 1000000, time: 3.920, throughput: 255104
2017/11/19 10:55:44 n: 1000000, time: 4.000, throughput: 249980
2017/11/19 10:55:44 n: 1000000, time: 4.095, throughput: 244203

2017/11/19 10:55:46 size: 128, n: 1000000, nodelay
2017/11/19 10:55:46 size: 128, n: 1000000, nodelay
2017/11/19 10:55:46 size: 128, n: 1000000, nodelay
2017/11/19 10:56:12 n: 1000000, time: 26.019, throughput: 38433
2017/11/19 10:56:12 n: 1000000, time: 26.025, throughput: 38424
2017/11/19 10:56:12 n: 1000000, time: 26.073, throughput: 38354
```

### go-final
```
2017/11/19 12:03:30 size: 128, n: 1000000
2017/11/19 12:03:30 size: 128, n: 1000000
2017/11/19 12:03:30 size: 128, n: 1000000
2017/11/19 12:03:33 n: 1000000, time: 2.742, throughput: 364693
2017/11/19 12:03:33 n: 1000000, time: 2.790, throughput: 358405
2017/11/19 12:03:34 n: 1000000, time: 3.193, throughput: 313213

2017/11/19 12:03:35 listening on :8888
2017/11/19 12:03:36 size: 128, n: 1000000, nodelay
2017/11/19 12:03:36 size: 128, n: 1000000, nodelay
2017/11/19 12:03:36 size: 128, n: 1000000, nodelay
2017/11/19 12:04:02 n: 1000000, time: 26.682, throughput: 37478
2017/11/19 12:04:02 n: 1000000, time: 26.703, throughput: 37449
2017/11/19 12:04:02 n: 1000000, time: 26.719, throughput: 37426
```

# Conclusion
- Call IO functions as less as possible, no matter in what programming language.
- Sending messages in 'go' is about 60% percents slower than in 'c'.

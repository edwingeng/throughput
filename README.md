# Overview
This is a simple throughput comparison of serveral TCP c/s implementations. I hope you can get a rough idea about their performance differences.

**These toys do not measure the performance of receiving messages on the server side.**

# Comparison Group A

### go
```
2019/01/31 15:40:42 payload size: 128, n: 1000000
2019/01/31 15:40:48 n: 1000000, time: 6.279, throughput: 159257

2019/01/31 15:40:50 payload size: 128, n: 1000000, TCP_NODELAY
2019/01/31 15:41:02 n: 1000000, time: 11.878, throughput: 84189
```

### go-zmq
```
2019/01/31 15:41:59 payload size: 128, n: 1000000, TCP_NODELAY
2019/01/31 15:42:01 n: 1000000, time: 2.126, throughput: 470454
```

### c
```
2019/01/31 15:39:18 payload size: 128, n: 1000000
2019/01/31 15:39:23 n: 1000000, time: 3.229, throughput: 309655

2019/01/31 15:39:25 payload size: 128, n: 1000000, TCP_NODELAY
2019/01/31 15:39:35 n: 1000000, time: 6.176, throughput: 161907
```

# Comparison Group B

### go
```
2019/01/31 15:40:42 payload size: 128, n: 1000000
2019/01/31 15:40:48 n: 1000000, time: 6.279, throughput: 159257

2019/01/31 15:40:50 payload size: 128, n: 1000000, TCP_NODELAY
2019/01/31 15:41:02 n: 1000000, time: 11.878, throughput: 84189
```

### go-one-off
```
2019/01/31 15:44:35 payload size: 128, n: 1000000
2019/01/31 15:44:39 n: 1000000, time: 3.605, throughput: 277420

2019/01/31 15:44:41 payload size: 128, n: 1000000, TCP_NODELAY
2019/01/31 15:44:47 n: 1000000, time: 6.001, throughput: 166645
```

# Comparison Group C

### go-one-off
```
2019/01/31 15:45:17 payload size: 128, n: 1000000
2019/01/31 15:45:17 payload size: 128, n: 1000000
2019/01/31 15:45:17 payload size: 128, n: 1000000
2019/01/31 15:45:22 n: 1000000, time: 4.441, throughput: 225197
2019/01/31 15:45:22 n: 1000000, time: 4.454, throughput: 224516
2019/01/31 15:45:22 n: 1000000, time: 4.486, throughput: 222919

2019/01/31 15:45:24 payload size: 128, n: 1000000, TCP_NODELAY
2019/01/31 15:45:24 payload size: 128, n: 1000000, TCP_NODELAY
2019/01/31 15:45:24 payload size: 128, n: 1000000, TCP_NODELAY
2019/01/31 15:45:48 n: 1000000, time: 23.897, throughput: 41845
2019/01/31 15:45:48 n: 1000000, time: 23.903, throughput: 41835
2019/01/31 15:45:48 n: 1000000, time: 23.904, throughput: 41833
```

### go-final
```
2019/01/31 15:46:27 payload size: 128, n: 1000000
2019/01/31 15:46:27 payload size: 128, n: 1000000
2019/01/31 15:46:27 payload size: 128, n: 1000000
2019/01/31 15:46:30 n: 1000000, time: 3.725, throughput: 268454
2019/01/31 15:46:30 n: 1000000, time: 3.776, throughput: 264827
2019/01/31 15:46:31 n: 1000000, time: 3.850, throughput: 259754

2019/01/31 15:46:33 payload size: 128, n: 1000000, TCP_NODELAY
2019/01/31 15:46:33 payload size: 128, n: 1000000, TCP_NODELAY
2019/01/31 15:46:33 payload size: 128, n: 1000000, TCP_NODELAY
2019/01/31 15:46:56 n: 1000000, time: 23.920, throughput: 41806
2019/01/31 15:46:56 n: 1000000, time: 23.921, throughput: 41805
2019/01/31 15:46:56 n: 1000000, time: 23.924, throughput: 41798
```

# Conclusions
- Call IO functions as less as possible, no matter in what programming language.
- Sending messages in 'go' is about 50% percents slower than in 'c'.

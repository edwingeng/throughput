# Overview
This is a simple throughput comparison of the tcp socket in golang, zeromq in golang and tcp socket in c. I hope you can get a rough idea about their performance differences.

# Results
### go
```
size: 128, n: 1000000
total messages: 1000000, time: 5.427139791s, throughput: 184259

size: 128, n: 1000000, nodelay
total messages: 1000000, time: 13.930531631s, throughput: 71784
```

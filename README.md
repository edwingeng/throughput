# Overview
This is a simple throughput comparison of the tcp socket in golang, zeromq in golang and tcp socket in c. I hope you can get a rough idea about their performance differences.

# Results
### go
```
nodelay: false, size: 128, n: 1000000
total messages: 1000000, time: 5.519737641s, throughput: 181168

nodelay: true, size: 128, n: 1000000
total messages: 1000000, time: 14.062820876s, throughput: 71109
```

Go language tests
=================

### Install Go

Described in this repository: [Flaiers/the-go](https://github.com/Flaiers/the-go)

***

### Install project

```bash
git clone git@github.com:Flaiers/go-tests.git
```

***

### Install wrk

```bash
git clone git@github.com:wg/wrk.git ; \
cd wrk ; \
make -j 4 ; \
sudo cp wrk /usr/bin/
```

***

### Run

```bash
wrk -c200 -t4 -d120s http://127.0.0.1:8000
```

&nbsp;

net/http
--------
### Easy requests, no db

I run:

```bash
go run src/main.go
```

Result:

```
Running 2m test @ http://127.0.0.1:8000
  4 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.69ms    4.66ms 101.23ms   87.59%
    Req/Sec    19.18k     3.19k   33.45k    69.76%
  9162924 requests in 2.00m, 1.11GB read
Requests/sec:  76299.02
Transfer/sec:      9.46MB
```

&nbsp;

Gin
---
### Easy requests, no db

I run:

```bash
go get -u github.com/gin-gonic/gin ; \
go run src/main.go
```

Result:

```
Running 2m test @ http://127.0.0.1:8000
  4 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.93ms    5.21ms  85.09ms   87.62%
    Req/Sec    19.71k     3.74k   42.56k    69.34%
  9417494 requests in 2.00m, 1.14GB read
Requests/sec:  78417.86
Transfer/sec:      9.72MB
```

&nbsp;

Echo
----

I run:

```bash
go run src/main.go
```

Result:

```
Running 2m test @ http://127.0.0.1:8000
  4 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.95ms    4.99ms  99.07ms   88.04%
    Req/Sec    17.89k     3.11k   40.05k    72.45%
  8545856 requests in 2.00m, 1.03GB read
Requests/sec:  71180.59
Transfer/sec:      8.82MB
```

&nbsp;

Iris
----

I run:

```bash
go run src/main.go
```

Result:

```
Running 2m test @ http://127.0.0.1:8000
  4 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.03ms    5.19ms  97.29ms   87.47%
    Req/Sec    18.02k     3.53k   33.11k    70.70%
  8610067 requests in 2.00m, 1.04GB read
Requests/sec:  71693.01
Transfer/sec:      8.89MB
```

&nbsp;

Beego
------

I run:

```bash
go build -o core main.go ; \
./core
```

Result:

```
Running 2m test @ http://127.0.0.1:8000
  4 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.23ms    5.62ms  93.06ms   88.14%
    Req/Sec    18.16k     2.98k   45.97k    69.46%
  8672630 requests in 2.00m, 1.05GB read
Requests/sec:  72218.26
Transfer/sec:      8.95MB
```

Go language tests
=================

Use tests
---------
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
    Latency     3.96ms    5.12ms 128.98ms   87.54%
    Req/Sec    18.62k     4.48k   37.91k    68.42%
  8894634 requests in 2.00m, 1.08GB read
Requests/sec:  74077.34
Transfer/sec:      9.18MB
```

# Rece-LB

Building a HTTP load balancer and reverse proxy from scratch written in `Go`.

## Features

- Round Robin
- Config support using json

## Running it

Start multiple instances of `server.py`

```bash
# create a virtual env (optional)
python3 -m venv venv
source venv/bin/activate

# install flask
pip install flask

# make multiple servers
./run.sh
```

Start the loadbalancer

```bash
go run loadbalancer.go
```

Spam the loadbalancer with requests

```bash
for i in {1..20}; do curl 127.0.0.1:8080; done
```

To kill all the instances of the servers

```bash
./clear.sh
```

## Config

### port

The port to the run the load balancer on.

### hosts

A list of servers we want to forward requests to.

### algorithm

Algorithm for load balancing. Currently only `RoundRobin`.

## Sample Config

```json
{
  "port": 8080,
  "servers": [
    "http://localhost:5001",
    "http://localhost:5002",
    "http://localhost:5003",
    "http://localhost:5004",
    "http://localhost:5005"
  ],
  "algorithm": "RoundRobin"
}
```

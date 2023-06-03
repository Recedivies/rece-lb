# Rece-LB

Building a HTTP load balancer and reverse proxy from scratch written in `Go`.

## Features

- Random
- Round Robin
- Least Connection
- Passive health check at configurable intervals
- Active health check
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
go run main.go
```

Spam the loadbalancer with requests

```bash
for i in {1..20}; do curl -i 127.0.0.1:8080; done
```

To kill all the instances of the servers

```bash
./clear.sh
```

## Config

### port

The port to the run the load balancer on.

### servers

A list of servers we want to forward requests to.

### algorithm

Algorithm for load balancing. Can be `Random`, `RoundRobin`, or `LeastConnection` .

### health_check_type

Defines the type of health check you want to perform on your server. Can be `passive` or `active` (active health check is by default).

### health_check_interval

Time interval in seconds to perform health check on servers.

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
  "algorithm": "Random",
  "health_check_type": "passive",
  "health_check_interval": 5
}
```

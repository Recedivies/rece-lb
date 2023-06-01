#!/bin/bash

for i in {1..5}
do 
  python3 server.py "server-$i" "500$i" &
done
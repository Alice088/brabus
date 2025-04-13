#!/bin/bash

docker run --name nats --network nats --rm -p 4222:4222 -p 8222:8222 nats --http_port 8222

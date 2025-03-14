#!/bin/bash

# Start Docker Compose for the emulator
docker compose -f emulator/docker-compose.yaml up -d

# Start Azurite in the background
azurite &

# Run PostgreSQL control plane container
docker run --hostname=3482db53b646 \
  --name=pgcontrolplane \
  --env PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/lib/postgresql/14/bin \
  --env LANG=en_US.utf8 \
  --env PGDATA=/var/lib/postgresql/data \
  --env POSTGRES_PASSWORD=test \
  --env POSTGRES_USER=postgres \
  --volume $(pwd)/pgdata:/var/lib/postgresql/data \
  --network bridge \
  -p 5432:5432 \
  --restart=no \
  --runtime=runc \
  -d postgres:14

# Wait for containers to start
sleep 5

# Verify running containers
docker ps

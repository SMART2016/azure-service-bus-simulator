#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Install Azurite if not found
install_azurite() {
    echo "Checking Azurite installation..."
    if command_exists azurite; then
        echo "Azurite is already installed."
    else
        echo "Azurite is not installed. Installing now..."
        if [[ "$OSTYPE" == "darwin"* ]]; then
            # MacOS installation via Homebrew
            if ! command_exists brew; then
                echo "Homebrew not found. Please install Homebrew first."
                exit 1
            fi
            brew install --cask azurite
        else
            # Linux installation via NPM
            if ! command_exists npm; then
                echo "NPM not found. Installing Node.js and NPM..."
                sudo apt update && sudo apt install -y nodejs npm
            fi
            echo "Installing Azurite via NPM..."
            sudo npm install -g azurite
        fi
    fi
}

# Install and start Azurite
install_azurite
azurite &

# Start Docker Compose for the emulator
echo "Starting emulator with Docker Compose..."
docker compose -f emulator/docker-compose.yaml up -d

# Start PostgreSQL container
echo "Starting PostgreSQL container..."
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

# Wait for services to initialize
sleep 5

# Display running containers
echo "Running Docker containers:"
docker ps

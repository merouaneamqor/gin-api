#!/bin/bash

# The URL to the Go tar.gz file
GO_URL="https://go.dev/dl/go1.21.5.linux-amd64.tar.gz"

# The directory to which we'll install Go
GO_DIR="/usr/local"

# Remove any previous installation of Go
sudo rm -rf $GO_DIR/go

# Download the specified Go package
wget $GO_URL

# Extract the downloaded tar.gz file to the specified directory
sudo tar -C $GO_DIR -xzf go1.21.5.linux-amd64.tar.gz

# Remove the downloaded tar.gz file
rm go1.21.5.linux-amd64.tar.gz

# Set up environment variables
echo "export GOROOT=$GO_DIR/go" >> ~/.profile
echo "export PATH=\$PATH:\$GOROOT/bin" >> ~/.profile
echo "export PATH=$GOPATH\$bin:$GOROOT\$bin:$PATH" >> ~/.profile
# Apply the environment variables
source ~/.profile

# Verify the installation
go version

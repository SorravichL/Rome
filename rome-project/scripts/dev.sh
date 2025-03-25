#!/bin/bash

# Exit on any error
set -e

echo "Starting Go backend"
cd go-backend
go run main.go &
GO_PID=$!
cd ..

echo "Starting TypeScript backend"
cd ts-backend
npm run dev &
TS_PID=$!
cd ..

echo "Both services are running:"
echo "Go backend"
echo "TypeScript backend"

# Clean up on exit
trap "kill $GO_PID $TS_PID" EXIT
wait

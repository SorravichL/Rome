#!/bin/bash

# Exit on any error
set -e

echo "Starting Go backend on port 5001..."
cd go-backend
go run main.go &
GO_PID=$!
cd ..

echo "Starting TypeScript backend on port 5002..."
cd ts-backend
npm run dev &
TS_PID=$!
cd ..

echo "Both services are running:"
echo "Go backend:        http://localhost:5001"
echo "TypeScript backend: http://localhost:5002"

# Clean up on exit
trap "kill $GO_PID $TS_PID" EXIT
wait

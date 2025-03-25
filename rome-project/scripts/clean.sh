#!/bin/bash

# Exit on error
set -e

echo "🧹 Running TypeScript Linter..."
cd ts-backend
npm run lint

echo "🎨 Formatting TypeScript code..."
npm run format
cd ..

echo "🧹 Running Go Linter..."
cd go-backend
golangci-lint run

echo "🎨 Formatting Go code..."
go fmt ./...
cd ..

echo "✅ Codebase linted and formatted!"

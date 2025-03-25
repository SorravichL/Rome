#!/bin/bash

# Ensure output folders exist
mkdir -p ./go-backend/gen
mkdir -p ./ts-backend/gen

echo "🔁 Generating Go code..."
oapi-codegen -generate types,client,chi-server \
  -o ./go-backend/gen/api.gen.go \
  -package gen \
  ./shared/openapi/api.yaml

echo "🔁 Generating TypeScript code..."
npx openapi-typescript ./shared/openapi/api.yaml -o ./ts-backend/gen/api.ts

echo "✅ Code generation completed!"
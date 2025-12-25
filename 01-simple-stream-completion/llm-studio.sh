#!/bin/bash

# Script to pull all models for Docker Model Runner used in main.go
# This script reads models from docker.env and downloads them

ENV_FILE="llm-studio.env"

# Check if llm-studio.env exists
if [ ! -f "$ENV_FILE" ]; then
  echo "❌ Error: $ENV_FILE not found!"
  exit 1
fi

echo "🚀 Starting to pull Ollama models from $ENV_FILE..."
echo ""

# Source the environment file
source "$ENV_FILE"

lms get "$CHAT_MODEL"


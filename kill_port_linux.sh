#!/bin/bash

if [ -f .env ]; then
  # Read PORT from .env
  export $(grep -v '^#' .env | xargs)
  
  if [ -n "$PORT" ]; then
    echo "Found PORT=$PORT in .env"
    
    # Linux specific port check using lsof or netstat
    # Try lsof first (more reliable on linux)
    if command -v lsof >/dev/null 2>&1; then
      PID=$(lsof -t -i:$PORT)
    elif command -v netstat >/dev/null 2>&1; then
      PID=$(netstat -nlp | grep ":$PORT" | awk '{print $7}' | cut -d'/' -f1)
    else
        echo "Error: lsof or netstat not found. Cannot determine PID."
        exit 1
    fi

    if [ -n "$PID" ]; then
        echo "Killing process on port $PORT (PID: $PID)..."
        kill -9 $PID
        echo "Process killed."
    else
        echo "No process found on port $PORT."
    fi
  else
    echo "PORT not found in .env"
  fi
else
  echo ".env file not found."
fi

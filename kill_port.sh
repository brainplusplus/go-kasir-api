#!/bin/bash

if [ -f .env ]; then
  # Read PORT from .env
  export $(grep -v '^#' .env | xargs)
  
  if [ -n "$PORT" ]; then
    echo "Found PORT=$PORT in .env"
    
    # Find PID on port (filtering for LISTENING to avoid TIME_WAIT or other states)
    PIDs=$(netstat -ano | grep ":$PORT" | grep "LISTENING" | awk '{print $5}' | sort -u)
    
    if [ -n "$PIDs" ]; then
        for pid in $PIDs; do
            if [ "$pid" != "0" ]; then
                echo "Killing process on port $PORT (PID: $pid)..."
                taskkill //F //PID $pid
                echo "Process $pid killed."
            fi
        done
    else
        echo "No process found on port $PORT."
    fi
  else
    echo "PORT not found in .env"
  fi
else
  echo ".env file not found."
fi

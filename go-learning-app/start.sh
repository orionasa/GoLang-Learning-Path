#!/bin/bash
# This script starts the backend and frontend servers.

echo "Starting backend server..."
(cd backend && go run main.go) &
BACKEND_PID=$!

echo "Starting frontend server..."
(cd frontend && npm install && npm run dev) &
FRONTEND_PID=$!

# Wait for both processes to complete
wait $BACKEND_PID
wait $FRONTEND_PID

#!/bin/bash

# Define the URL
URL="http://localhost:8080/"

# Start time
SECONDS=0

# Loop to send 1000 requests
for i in {1..10000}
do
    curl -s $URL &
done

# Wait for all background jobs to complete
wait

# Calculate time taken
elapsed_time=$SECONDS  # Elapsed time in seconds

echo "Time taken: ${elapsed_time} seconds"

#!/bin/bash

# Clear the shared file upon each launch, 
# to make sure it's unique between runs
rm shared_volume/shared

docker compose up --build
rm shared_volume/shared
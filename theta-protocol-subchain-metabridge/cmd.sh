#!/bin/bash

# Enter the name of the Docker image you want to use
image_name="theta-subc-node"

# Get the container ID of the running container with the specified image name
container_id=$(docker ps -qf "ancestor=$image_name")

if [ -z "$container_id" ]
then
  echo "No running container with image name $image_name found"
else
  # Run the 'exec' command inside the container
  docker exec $container_id bash -c '"$@"' sh "$@"
fi
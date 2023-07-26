
get_version() {
  git describe --tags --abbrev=0
}

# Function to build the Docker image
build_image() {
  git rev-parse HEAD > ./REVISION
  ver=`get_version`
  echo "Building image with version $ver"
  docker build --platform linux/amd64 -t snowflake-id-gen:latest -f ./Dockerfile .
  docker tag snowflake-id-gen:latest snowflake-id-gen:$ver
}

# function for push image to docker hub
push_image() {
  ver=`get_version`
  docker tag snowflake-id-gen:$ver thecode/snowflake-id-gen:$ver
  docker push thecode/snowflake-id-gen:$ver
}

# function for help
print_help() {
  echo "Usage: $0 [build | push | run]"
  exit 1
}

# Parse the arguments and execute the corresponding actions
case "$1" in
  "build")
    build_image
    ;;
  "push")
    push_image
    ;;
  "run")
    run_container
    ;;
  *)
    print_help
    ;;
esac

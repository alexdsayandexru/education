docker run -it --rm \
  -p 8181:8080 \
  -v "$(pwd)/config.yml:/config.yml" \
  flipt/flipt:latest ./flipt --config /config.yml

  docker run -d \
      -p 8080:8080 \
      -p 9000:9000 \
      -v $HOME/flipt:/var/opt/flipt \
      docker.flipt.io/flipt/flipt:v2-beta

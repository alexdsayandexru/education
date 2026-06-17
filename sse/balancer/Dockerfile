FROM nginx

RUN \
  apt-get update && \
  apt-get install -y \
    vim

COPY nginx2.conf /etc/nginx/nginx.conf

COPY localhost.crt /etc/nginx/localhost.crt
COPY localhost.key /etc/nginx/localhost.key
1. Собрать образ
docker image build -t load_balancer_sse .

2. Создать контейнер
docker run --name proxy_nginx_tls -p 4433:443 -d proxy_nginx_tls
(docker run --name LOAD_BALANCER_SSE -p 8888:80 -d load_balancer_sse)


3. Выполнить запросы к балансировщику
curl https://localhost:4433/random --insecure

#while sleep 0.5; do curl https://localhost:4433; done

4. Создать сертификат
openssl req -x509 -sha256 -nodes -newkey rsa:2048 -days 365 -keyout localhost.key -out localhost.crt
openssl x509 -text -noout -in localhost.crt
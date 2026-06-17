Сборка образа:
docker build -t sse_server .

docker run --name "SSE_SERVER-1" -p 8881:8888 sse_server
docker run --name "SSE_SERVER-2" -p 8882:8888 sse_server
docker run --name "SSE_SERVER-3" -p 8883:8888 sse_server
docker run --name "SSE_SERVER-4" -p 8884:8888 sse_server

Запуск клиента:
http://localhost:8888

Материалы:

https://learn.javascript.ru/server-sent-events#format-otveta-servera

Сервер посылает сообщения, разделённые двойным переносом строки \n\n.
Сообщение состоит из следующих полей:
data: – тело сообщения, несколько data подряд интерпретируются как одно сообщение, разделённое переносами строк \n.
id: – обновляет свойство lastEventId, отправляемое в Last-Event-ID при переподключении.
retry: – рекомендованная задержка перед переподключением в миллисекундах. Не может быть установлена с помощью JavaScript.
event: – имя пользовательского события, должно быть указано перед data:.
Сообщение может включать одно или несколько этих полей в любом порядке, но id обычно ставят в конце.
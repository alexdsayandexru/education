cd '/Users/aasudakov/kafka/kafka_2.12-3.6.1'

bin/zookeeper-server-start.sh config/zookeeper.properties
bin/kafka-server-start.sh config/server.properties

bin/kafka-topics.sh --create --topic test1 --bootstrap-server localhost:9092
bin/kafka-topics.sh --describe --topic test1 --bootstrap-server localhost:9092
bin/kafka-topics.sh --bootstrap-server localhost:9092 --topic test1 --alter --partitions 3

bin/kafka-console-producer.sh --topic test1 --bootstrap-server localhost:9092
bin/kafka-console-consumer.sh --topic test1 --from-beginning --bootstrap-server localhost:9092
bin/kafka-console-consumer.sh --topic test1 --from-beginning --bootstrap-server localhost:9092 --group group1


bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --group group1 --describe

//Сброс оффсета на начало
bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --group group1 --reset-offsets --to-earliest --topic test1 --execute

touch /tmp/data && tail -f -n0 /tmp/data | /bin/kafka-console-producer.sh --topic test1 --bootstrap-server=localhost:9092 --sync
for i in $(seq 1 3600); do echo "test${i}" >> /tmp/data; sleep 1; done




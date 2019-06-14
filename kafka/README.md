bin/kafka-topics --create --zookeeper localhost:2181 --replication-factor 1 --partitions 2 --topic topochico


bin/kafka-topics --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic topointermedio --config cleanup.policy=compact

bin/kafka-console-consumer --bootstrap-server localhost:9092 \
    --topic topochico \
    --from-beginning \
    --formatter kafka.tools.DefaultMessageFormatter \
    --property print.key=true \
    --property print.value=true \
    --property key.deserializer=org.apache.kafka.common.serialization.StringDeserializer \
    --property value.deserializer=org.apache.kafka.common.serialization.StringDeserializer

bin/kafka-topics --list --zookeeper localhost:2181
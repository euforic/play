#!/usr/bin/env bash

docker create network eosdev

docker run --name nodeos -d -p 8888:8888 --network eosdev \
-v /tmp/eosio/work:/work -v /tmp/eosio/data:/mnt/dev/data \
-v /tmp/eosio/config:/mnt/dev/config eosio/eos-dev  \
/bin/bash -c "nodeos -e -p eosio --plugin eosio::producer_plugin \
--plugin eosio::history_plugin --plugin eosio::chain_api_plugin \
--plugin eosio::history_api_plugin \
 --plugin eosio::http_plugin -d /mnt/dev/data \
--config-dir /mnt/dev/config \
--http-server-address=0.0.0.0:8888 \
--access-control-allow-origin=* --contracts-console --http-validate-host=false"

docker run -d --name keosd --network=eosdev \
-i eosio/eos-dev /bin/bash -c "keosd --http-server-address=0.0.0.0:9876"

# Useful alias for easy cleos calls
#alias cleos='docker exec -it nodeos /opt/eosio/bin/cleos --url http://127.0.0.1:8888 --wallet-url http://[keosd DOCKER_IP]:9876'


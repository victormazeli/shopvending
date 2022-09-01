#!/bin/bash

sh ./prepare-all.sh

# setup docker machine
docker swarm leave --force
docker swarm init --advertise-addr 192.168.65.1
docker network create --driver overlay --internal shopvending_network

# remove already running services (if any)
docker service rm shopVendingMicroservices_userservice
docker service rm shopVendingMicroservices_gatewayservice
#docker service rm mobileVendingMicroservices_orderservice
#docker service rm mobileVendingMicroservices_paymentservice
#docker service rm mobileVendingMicroservices_ratingservice
#docker service rm mobileVendingMicroservices_sellerservice
#docker service rm mobileVendingMicroservices_userservice
#docker service rm mobileVendingMicroservices_riderservice
#docker service rm mobileVendingMicroservices_locationservice
#docker service rm mobileVendingMicroservices_messagingservice
#docker service rm mobileVendingMicroservices_recommendationengine
#docker service rm mobileVendingMicroservices_deliveryservice
#docker service rm mobileVendingMicroservices_gelftail
#docker service rm mobileVendingMicroservices_foodservice
sh ./deploy-viz.sh

# deploy the services stack
docker stack deploy --compose-file=docker-compose-stage.yml shopVendingMicroservices

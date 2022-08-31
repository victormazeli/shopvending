#!/bin/bash

export CGO_ENABLED=0

echo "Preparing user service"
cd userservice;go get;env GOOS=linux GOARCH=amd64 go build -o user-service-linux-amd64;echo built `pwd`;cd ..
docker image rm cipher2022/userservice
docker build -t cipher2022/userservice userservice/
docker push cipher2022/userservice


echo "Preparing gateway service"
cd gatewayservice;go get;env GOOS=linux GOARCH=amd64 go build -o gateway-service-linux-amd64;echo built `pwd`;cd ..
docker image rm cipher2022/gatewayservice
docker build -t cipher2022/gatewayservice gatewayservice/
docker push cipher2022/gatewayservice


#echo "Preparing rating service"
#cd ratingservice;go get;env GOOS=linux GOARCH=amd64 go build -o ratingservice-linux-amd64;echo built `pwd`;cd ..
#docker image rm mobile-vending-microservices/ratingservice
#docker build -t mobile-vending-microservices/ratingservice ratingservice/
#
#echo "Preparing user service"
#cd userservice;go get;env GOOS=linux GOARCH=amd64 go build -o userservice-linux-amd64;echo built `pwd`;cd ..
#docker image rm mobile-vending-microservices/userservice
#docker build -t mobile-vending-microservices/userservice userservice/
#
#echo "Preparing payment service"
#cd paymentservice;go get;env GOOS=linux GOARCH=amd64 go build -o paymentservice-linux-amd64;echo built `pwd`;cd ..
#docker image rm mobile-vending-microservices/paymentservice
#docker build -t mobile-vending-microservices/paymentservice paymentservice/
#
#echo "Preparing food service"
#cd foodservice;go get;env GOOS=linux GOARCH=amd64 go build -o foodservice-linux-amd64;echo built `pwd`;cd ..
#docker image rm mobile-vending-microservices/foodservice
#docker build -t mobile-vending-microservices/foodservice foodservice/
#
#echo "Preparing order service"
#cd orderservice;go get;env GOOS=linux GOARCH=amd64 go build -o orderservice-linux-amd64;echo built `pwd`;cd ..
#docker image rm mobile-vending-microservices/orderservice
#docker build -t mobile-vending-microservices/orderservice orderservice/
#
#echo "Preparing gateway service"
#cd gatewayservice;go get;env GOOS=linux GOARCH=amd64 go build -o gatewayservice-linux-amd64;echo built `pwd`;cd ..
#docker image rm mobile-vending-microservices/gatewayservice
#docker build -t mobile-vending-microservices/gatewayservice gatewayservice/
#
#echo "Preparing notification service"
#cd notificationservice;go get;env GOOS=linux GOARCH=amd64 go build -o notificationservice-linux-amd64;echo built `pwd`;cd ..
#docker image rm mobile-vending-microservices/notificationservice
#docker build -t mobile-vending-microservices/notificationservice notificationservice/
#
#echo "Preparing rider service"
#cd riderservice;go get;env GOOS=linux GOARCH=amd64 go build -o riderservice-linux-amd64;echo built `pwd`;cd ..
#docker image rm mobile-vending-microservices/riderservice
#docker build -t mobile-vending-microservices/riderservice riderservice/
#
#echo "Preparing location service"
#cd locationservice;go get;env GOOS=linux GOARCH=amd64 go build -o locationservice-linux-amd64;echo built `pwd`;cd ..
#docker image rm mobile-vending-microservices/locationservice
#docker build -t mobile-vending-microservices/locationservice locationservice/
#
#echo "Preparing messaging service"
#cd messagingservice;go get;env GOOS=linux GOARCH=amd64 go build -o messagingservice-linux-amd64;echo built `pwd`;cd ..
#docker image rm mobile-vending-microservices/messagingservice
#docker build -t mobile-vending-microservices/messagingservice messagingservice/
#
#echo "Preparing delivery service"
#cd deliveryservice;go get;env GOOS=linux GOARCH=amd64 go build -o deliveryservice-linux-amd64;echo built `pwd`;cd ..
#docker image rm mobile-vending-microservices/deliveryservice
#docker build -t mobile-vending-microservices/deliveryservice deliveryservice/
#
#echo "Preparing gelftail service"
#cd gelftail;go get;env GOOS=linux GOARCH=amd64 go build -o gelftail-linux-amd64;echo built `pwd`;cd ..
#docker image rm mobile-vending-microservices/gelftail
#docker build -t mobile-vending-microservices/gelftail gelftail/
#
#echo "Preparing recommendation engine"
#docker image rm mobile-vending-microservices/recommendationengine
#cd recommendationengine;chmod +x buildContainer.sh;bash buildContainer.sh;echo built `pwd`;cd ..

echo "Finished preparations"
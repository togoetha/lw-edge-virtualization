#! /bin/bash

sudo docker stop gorestapi
sudo docker rm gorestapi
echo `date +%s%N`
echo `date +%s%N`
sudo docker create --name gorestapi --network host togoetha/go-exe
echo `date +%s%N`
sudo docker start gorestapi

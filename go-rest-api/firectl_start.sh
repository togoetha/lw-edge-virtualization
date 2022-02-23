#! /bin/bash

echo `date +%s%N`
echo `date +%s%N`
sudo docker create --name gorestpie --network host togoetha/go-exe
echo `date +%s%N`
sudo docker start gorestpie
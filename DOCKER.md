# Running on Docker

## Install Docker

From the instructions at https://docs.docker.com/engine/install/debian/:

    sudo apt-get update
    sudo apt-get install ca-certificates curl gnupg lsb-release
    curl -fsSL https://download.docker.com/linux/debian/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
    sudo apt-get update
    sudo apt-get install docker-ce docker-ce-cli containerd.io
    
## Install Golang

From the instructions at https://go.dev/doc/install (choose correct version if not ARM): 

    wget https://go.dev/dl/go1.17.7.linux-arm64.tar.gz
    sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.7.linux-arm64.tar.gz

## Build the container

Depending on the platform, run either build-docker.sh or build-docker-arm.sh in the folder go-rest-api.

## Run!

The commands used to launch this container for the benchmarks are:

    sudo docker create --name gorestapi --network host --cpus=1 togoetha/go-exe-arm
    sudo docker start gorestapi
    
Alternatively, you can use docker_start.sh in the folder go-rest-api, which includes timing statements.

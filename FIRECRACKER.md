# Running on Firecracker

## Install Firecracker

Make sure KVM support is enabled!
Basic instructions at https://github.com/firecracker-microvm/firecracker/blob/main/docs/getting-started.md.

Get the 0.25.2 release specifically:

    wget https://github.com/firecracker-microvm/firecracker/releases/download/v0.25.2/firecracker-v0.25.2-aarch64.tgz
    tar xvzf firecracker-v0.25.2-aarch64.tgz
    cd firecracker-v0.25.2-aarch64
    sudo cp firecracker-v0.25.2-aarch64 /usr/local/bin/firecracker
    
At this point, you should be fine. However, if firecracker doesn't feel like cooperating you may have to set up a socket manually when using firectl in the last step:

    rm -f /tmp/firecracker.socket
    sudo firecracker --api-sock /tmp/firecracker.socket
    
## Get base kernel image

From the go-rest-api directory (chmod may be required):

    ./get_fc_images.sh
    cp hello-* /tmp/
    
## Get firectl

Install Golang from the instructions at https://go.dev/doc/install (choose correct version if not ARM): 

    wget https://go.dev/dl/go1.17.7.linux-arm64.tar.gz
    sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.7.linux-arm64.tar.gz

Get the repo and build using Golang:

    git clone https://github.com/firecracker-microvm/firectl.git
    cd firectl
    make
    
Optionally, copy the resulting firectl executable to /usr/local/bin so it can hang out with the firecracker binary.
The build may get stuck on fetching go libraries, so you may have to retry for no reason.

## Get the Golang filesystem compilation template

Clone the repo and build:

    git clone https://github.com/s8sg/firecracker-go-template.git
    cd firecracker-go-template
    make
    
You may get an error along the lines of <docker image does not exist locally>. If so, pull it manually and retry. Then find something useful to do, as the build takes a while.

## Create filesystem
  
Once the build is done, copy the go-rest-api folder into the firecracker-go-template directory and use the build script:
  
    sudo ./build go-rest-api resttestfs 172.20.0.2/24 172.20.0.1
  
## Run!
  
The easiest way to run the filesystem using the default Firecracker kernel is by using firectl_start.sh in the go-rest-api directory. 
This script automatically sets up a tap device, takes care of iptables rules and launches the microVM using firectl. 
Note: change executable paths and file names if necessary.

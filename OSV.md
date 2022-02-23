# Running on OSv

## Install OSv & Golang

Golang support on ARM is experimental. Get OSv from this specific fork/branch: https://github.com/wkozaczuk/osv/tree/aarch64_syscalls_no_stack

    git clone https://github.com/wkozaczuk/osv.git
    cd osv
    git checkout aarch64_syscalls_no_stack
    
Follow the OSv documentation as normal: https://github.com/cloudius-systems/osv

    git submodule update --init --recursive
    sudo ./scripts/setup.py
    
Make sure KVM support is enabled (see Running OSv section in documentation)!
From the instructions at https://go.dev/doc/install (choose correct version if not ARM): 

    wget https://go.dev/dl/go1.17.7.linux-arm64.tar.gz
    sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.7.linux-arm64.tar.gz

## Copy the service

Copy the folder go-rest-api-osv to <osv dir>/apps/

## Build kernel and service

Go to the OSv base directory and execute the build script:

    sudo ./scripts/build image=go-rest-api
    
Find something else to do in the meantime, because this takes a while.

## Routing rules

Some iptables rules are required to access the REST service beyond the host machine:

    sudo iptables -I FORWARD -o virbr0 -d 192.168.122.76 -p tcp --dport 8080 -j ACCEPT
    sudo iptables -t nat -I PREROUTING -p tcp --dport 8080 -j DNAT --to 192.168.122.76:8080

## Run!

The run command automatically launches the last-built unikernel.

    sudo ./scripts/run.py -nv -c 1

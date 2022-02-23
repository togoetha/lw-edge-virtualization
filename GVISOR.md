# Running on gVisor

Note: at the time of writing, gVisor does not work with Ubuntu >= 21.x, as it requires cgroups v1 which is no longer supported in Ubuntu 21.

## Install Docker

Follow the instructions in [DOCKER.md](DOCKER.md)

## Get gVisor 

From the instructions at https://gvisor.dev/docs/user_guide/install/:

    set -e
    ARCH=$(uname -m)
    URL=https://storage.googleapis.com/gvisor/releases/release/latest/${ARCH}
    wget ${URL}/runsc ${URL}/runsc.sha512 ${URL}/containerd-shim-runsc-v1 ${URL}/containerd-shim-runsc-v1.sha512
    sha512sum -c runsc.sha512 -c containerd-shim-runsc-v1.sha512
    rm -f *.sha512
    chmod a+rx runsc containerd-shim-runsc-v1
    sudo mv runsc containerd-shim-runsc-v1 /usr/local/bin
    sudo runsc install
    sudo systemctl reload docker
    
## Run!

Note: gVisor does not support using the host network namespace at the moment.
Use this command to run the service under gVisor using Docker:

    docker run --runtime=runsc --name gorestapi -p 8080:8080 --cpus=1 togoetha/go-exe-arm

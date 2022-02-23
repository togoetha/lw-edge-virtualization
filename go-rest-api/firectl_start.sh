sudo ip link del tap0
sudo ip tuntap add tap0 mode tap # user $(id -u) group $(id -g)
sudo ip addr add 172.20.0.1/24 dev tap0
sudo ip link set tap0 up
DEVICE_NAME=enp3s0
sudo sh -c "echo 1 > /proc/sys/net/ipv4/ip_forward"
sudo iptables -t nat -A POSTROUTING -o $DEVICE_NAME -j MASQUERADE
sudo iptables -A FORWARD -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
sudo iptables -A FORWARD -i tap0 -o $DEVICE_NAME -j ACCEPT
MAC="$(cat /sys/class/net/tap0/address)"
ROOTFS="$(readlink -f resttestfs)"
sudo ../firectl --kernel=/tmp/hello-vmlinux.bin --root-drive=$ROOTFS --memory=2048 --ncpus=1 --kernel-opts="console=ttyS0 noapic reboot=k panic=1 pci=off nomodules rw" --tap-device=tap0/$MAC

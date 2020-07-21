# kvm-operator

# Prerequisites

   golang: verison 1.13 above

   operator-sdk:operator-sdk-v0.15.2

   qemu-kvm,libvirt-bin bridge-utils virtinst virt-manager
   
# Install:
 golang: verison 1.13 above

 operator-sdk:operator-sdk-v0.15.2
 
  1. sudo apt install qemu-kvm libvirt-bin bridge-utils virtinst virt-manager

  2. check is status 

     sudo systemctl is-active libvirtd

  3. if you want to connect TCP connection

     1. stop the KVM libvirt

         sudo systemctl stop libvirtd

     2. install firewall

        i.  sudo apt install firewalld

        ii. firewall-cmd --add-port=16509/tcp

     3. sudo /usr/sbin/libvirtd --timeout 120 --listen

  4. clone this kvm-operator repository

  5. import the package

     go get github.com/libvirt/libvirt-go

  6. run in local

     operator-sdk run --local

  6. set the kvm-centos.yaml file
     
      https://github.com/raju010193/kvm-operator/blob/master/deploy/crds/kvm.example.com_v1alpha1_kvm_cr.yaml
      
   7. kubectl apply -f kvm-centos.yaml

   if you want to check the running VMs

      go run cmd/kvm_manager/main.go -host "192.168.100.9" --connection "tcp"
   
   
 


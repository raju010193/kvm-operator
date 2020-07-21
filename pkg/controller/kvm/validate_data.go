package kvm
import (
	"strconv"
	kvmv1alpha1 "github.com/raju140/kvm-operator/pkg/apis/kvm/v1alpha1"
)

func ValidateData(instance kvmv1alpha1.Kvm,Name string)(KvmDetails,int)  {
	kvmDetails := KvmDetails{}
	kvmDetails.Image.Name = Name
	if len(instance.Spec.Host)!=0{
		kvmDetails.Host = instance.Spec.Host
	}else{
		kvmDetails.Host = "127.0.0.1" //localhost
	}
	if len(instance.Spec.Image.Path)!=0{
		kvmDetails.Image.Path = instance.Spec.Image.Path
	}
	// memory, err := strconv.Atoi(instance.Spec.Memory)
	if instance.Spec.Memory.Current >=1048576{
		kvmDetails.Memory.Current = instance.Spec.Memory.Current
	}else{
		kvmDetails.Memory.Current = 1048576 //kib
	}
	if instance.Spec.Memory.Max >= 1048567{
		kvmDetails.Memory.Max = instance.Spec.Memory.Max
		
	}else{
		kvmDetails.Memory.Max = 1048567
	}

	if kvmDetails.Memory.Current > kvmDetails.Memory.Max{
		kvmDetails.Memory.Current += instance.Spec.Memory.Max
	}
	if len(instance.Spec.OStype) !=0{
		kvmDetails.OStype = instance.Spec.OStype
	}else{
		kvmDetails.OStype = "hvm"
	}
	if len(instance.Spec.Image.Format) !=0{
		kvmDetails.Image.Format = instance.Spec.Image.Format
	}else{
		kvmDetails.Image.Format = "qcow2"
	}
	current, err := strconv.Atoi(instance.Spec.VCPU.Current)
	max, err := strconv.Atoi(instance.Spec.VCPU.Max)
	// max, err := strconv.Atoi(instance.Spec.VCPU.Max)
	if err!=nil{
		return kvmDetails,1
	}
	if current >=1{
		kvmDetails.vcpu.Current = instance.Spec.VCPU.Current
	}else{
		kvmDetails.vcpu.Current = "1" //kib
		current = 1
	}
	if max > 1{
		kvmDetails.vcpu.Max = instance.Spec.VCPU.Max
	}else{
		kvmDetails.vcpu.Max = "1"
		max = 1
	}
	if current > max{
		max = max+current
		kvmDetails.vcpu.Max = strconv.Itoa(max)
		
	}
	
	if len(instance.Spec.Connection)!=0{
		kvmDetails.ConnectionType = instance.Spec.Connection
	}else{
		kvmDetails.ConnectionType = "tcp"
	}
	if len(instance.Spec.Image.Type)!=0{
		kvmDetails.Image.Type = instance.Spec.Image.Type
	}else{
		kvmDetails.Image.Type = "file"
	}
	if len(instance.Spec.Image.Device)!=0{
		kvmDetails.Image.Device = instance.Spec.Image.Device
	}else{
		kvmDetails.Image.Type = "file"
	}
	if len(instance.Spec.Drive.Drive)!=0{
		kvmDetails.Drive.Drive = instance.Spec.Drive.Drive
	}else{
		kvmDetails.Drive.Drive = "cdrom"
	}
	if len(instance.Spec.Drive.Type)!=0{
		kvmDetails.Drive.Type = instance.Spec.Drive.Type
	}else{
		kvmDetails.Drive.Type = "file"
	}

	if len(instance.Spec.NetworkInterface.InterfaceType)!=0{
		kvmDetails.NetworkInterface.InterfaceType = instance.Spec.NetworkInterface.InterfaceType
	}else{
		kvmDetails.NetworkInterface.InterfaceType = "network"
	}
	if len(instance.Spec.NetworkInterface.Bridge)!=0{
		kvmDetails.NetworkInterface.Bridge = instance.Spec.NetworkInterface.Bridge
	}else{
		return kvmDetails,1
	}
	if len(instance.Spec.NetworkInterface.Model)!=0{
		kvmDetails.NetworkInterface.Model = instance.Spec.NetworkInterface.Model
	}else{
		kvmDetails.NetworkInterface.Model = "0rtl8139"
	}
	if len(instance.Spec.NetworkInterface.MacAddress)!=0{
		kvmDetails.NetworkInterface.MacAddress = instance.Spec.NetworkInterface.MacAddress
	}else{
		kvmDetails.NetworkInterface.MacAddress = "52:54:00:59:30:69"
	}


	return kvmDetails,0
}
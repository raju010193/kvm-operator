package kvm


type KvmDetails struct {
	Image struct{
		Name string
		Path string
		Format string
		Type string
		Device string
	}
	Memory  struct{
		Current uint64
		Max uint64
	}
	vcpu struct{
		Current string
		Max string
	}
	Drive struct{
		Drive string
		Type string
	}
	NetworkInterface struct{
		InterfaceType string
		Bridge string 
		Model string 
		MacAddress string 
	 }
	OStype string
	Host string
	ConnectionType string

  }
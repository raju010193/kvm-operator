package main

import (
	"flag"
	"fmt"
	"os"
	libvirt "github.com/libvirt/libvirt-go"
)
func getConnection(ipAddress string,connectionType string,user string,password string)(*libvirt.Connect,error){
   // get libvirt connection based on connection type
   // return connetio
	callback := func(creds []*libvirt.ConnectCredential) {
		for _, cred := range creds {
			if cred.Type == libvirt.CRED_AUTHNAME {
				cred.Result = user
				cred.ResultLen = len(cred.Result)
			} else if cred.Type == libvirt.CRED_PASSPHRASE {
				cred.Result = password
				cred.ResultLen = len(cred.Result)
			}
		}
	}
	auth := &libvirt.ConnectAuth{
		CredType: []libvirt.ConnectCredentialType{
			libvirt.CRED_AUTHNAME, libvirt.CRED_PASSPHRASE,
		},
		Callback: callback,
	}
	conn,err := libvirt.NewConnectWithAuth("qemu+"+connectionType+"://"+ipAddress+"/system",auth,0)
	return conn,err
}

// list all running domains
func listRunningDomains(conn libvirt.Connect){
	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n %d running domains\n", len(doms))

	fmt.Printf("\n**********Running Domains**********\n")

	for i, dom := range doms {
		name, err := dom.GetName()

		if err == nil {
			fmt.Printf("%d  %6s\n",i+1, name)
		}

		dom.Free()
	}

}
func main() {
	var host = flag.String("host", "", "Host address")
	var connectionType = flag.String("connection", "", "tcp/ssh")
	var runningDomains = flag.Bool("domains",true,"true/false")
	var username = flag.String("user","","User name")
	var password = flag.String("password","","Password")
	
	flag.Parse()

	if *host == "" {
		fmt.Fprintf(os.Stderr, "Missing -host argument\n")
		os.Exit(1)
	}
	if *connectionType == "" {
		fmt.Fprintf(os.Stderr, "Missing -connection tcp/ssh argument\n")
		os.Exit(1)
	}
	if *connectionType == "ssh"{
		if *username ==""{
			fmt.Fprintf(os.Stderr, "Missing -user argument\n")
		}
		if *password == ""{
			fmt.Fprintf(os.Stderr, "Missing -password argument\n")
		}
	}

	conn,err := getConnection(*host,*connectionType,*username,*password)
	if err!=nil{
		fmt.Println(err)
	}
	// calling the running domains
	if *runningDomains == true{
		listRunningDomains(*conn)
	}
	
}
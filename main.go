package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func chekcDomain(domain string){ 
	var hasMX, hasSPF, hasDMARC bool 
	var SPFrecord, DMARCrecord string 

	mxRecords, err:=net.LookupMX(domain)
	if err!=nil{
		log.Fatal(err)
	} 
	if len(mxRecords)>0{
		hasMX=true
	}

	txtRecords,err:=net.LookupTXT(domain)
	if err!=nil{
		log.Fatal(err)
	} 

	for _,record:=range(txtRecords){
		if strings.HasPrefix(record, "v=spf1"){
			hasSPF=true 
			SPFrecord = record
			break
		}
	}


	dmarcRecords,err:=net.LookupTXT("_dmarc."+domain) 
	if err!=nil{
		log.Fatal(err)
	} 

	for _,record:=range(dmarcRecords){
		if strings.HasPrefix(record, "v=DMARC1"){
			hasDMARC=true 
			DMARCrecord = record
			break
		}
	} 

	fmt.Printf("%v %v %v %v %v %v",domain, hasMX, hasSPF, SPFrecord, hasDMARC, DMARCrecord)

}
func main(){
	scanner:=bufio.NewScanner(os.Stdin) 
	fmt.Println("domain, hasMX hasSPF SPFrecord hasDMARC DMARCrecord") 
	for scanner.Scan(){ 
		chekcDomain(scanner.Text())
	} 

	err:=scanner.Err()
	if err!=nil{
		log.Fatal(err)
	}
}
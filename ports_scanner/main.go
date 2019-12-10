package main;
import (
	"time"
	"fmt"
	"net"
	"strings"
	// "log"
	"os"
	// "sync"
)

func IsPortOpen(host string,port int, timeout time.Duration, retry bool)(ok bool,err error){
	addr := fmt.Sprintf("%s:%d",host,port);
	_, err = net.DialTimeout("tcp",addr,timeout)	
	// defer conn.Close();
	if err != nil{
		if strings.Contains(err.Error(), "too many open files"){
			time.Sleep(timeout)
			return IsPortOpen(host,port,timeout,false)
		}
		if networkErr,Vok := err.(net.Error);Vok{
			if networkErr.Timeout() {
				return IsPortOpen(host,port,timeout,false)
			}
			
			err = networkErr
			return
		}
		return
	}
	ok = true
	return;
}

type ScanRes struct{
	Port int
	Status string
	Err error
}

type PortScanner struct{
	Host string
}

func (p *PortScanner) Scan(lower,upper int, timeout time.Duration)(err error){

	ports := make([]int,upper-lower+1)
	portsCh :=  make(chan int, upper-lower+1)
	scanResCh := make(chan ScanRes)

	for c := 0;c<=1024;c++{
		go func (host string,timeout time.Duration, portsC chan int, scanResCh chan ScanRes){
			for p := range portsC{
				scanRes := ScanRes{Port:p}
				ok,err := IsPortOpen(host,p,timeout,true);
				if err != nil{
					scanRes.Err = err
					scanRes.Status = "closed"
				}
				if ok{
					scanRes.Status="open"
				}
				scanResCh <- scanRes
			}
		}(p.Host,timeout,portsCh,scanResCh)
	}
	for i := range ports{
		ports[i] = i+1
	}
	for i := range ports{
		portsCh <- ports[i]
	}
	for range ports{
		scanRes := <- scanResCh 
		if scanRes.Err == nil{
			fmt.Printf("Port:%d Status:%s\n",scanRes.Port,scanRes.Status);
		}
	}
	return
}

func main(){
	usage :=  "Usage:\nscan <host>"
	if len(os.Args) != 2{
		fmt.Println(usage)
		return
	}
	host :=  os.Args[1];
	portScanner := PortScanner{Host:host};
	timeout := time.Millisecond*500
	portScanner.Scan(1,65535,timeout)	
}
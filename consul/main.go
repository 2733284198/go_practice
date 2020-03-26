package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"
)

type ServiceInfo struct {
	ServiceID string
	IP        string
	Port      int
	Load      int
	Timestamp int //load updated ts
}
type ServiceList []ServiceInfo

type KVData struct {
	Load      int `json:"load"`
	Timestamp int `json:"ts"`
}

var (
	servicsMap    = make(map[string]ServiceList)
	serviceLocker = new(sync.Mutex)
	consulClient  *api.Client
	myServiceID   string
	myServiceName string
	myKey         string
)

func CheckErr(err error) {
	if err != nil {
		log.Printf("error: %v", err)
		os.Exit(1)
	}
}
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("check status.")
	fmt.Fprint(w, "status ok!")
}

func StartService(addr string) {
	http.HandleFunc("/status", StatusHandler)
	fmt.Println("start listen...")
	err := http.ListenAndServe(addr, nil)
	CheckErr(err)
}

func main() {
	var statusMonitorAddr, serviceName, serviceIp, consulAddr, foundService string
	var servicePort int
	flag.StringVar(&consulAddr, "consul_addr", "172.17.22.139:8500", "host:port of the service stuats monitor interface")
	flag.StringVar(&statusMonitorAddr, "monitor_addr", "172.20.99.75:54321", "host:port of the service stuats monitor interface")
	flag.StringVar(&serviceName, "service_name", "worker", "name of the service")
	flag.StringVar(&serviceIp, "ip", "172.20.99.75", "service serve ip")
	flag.StringVar(&foundService, "found_service", "worker", "found the target service")
	flag.IntVar(&servicePort, "port", 4300, "service serve port")

	flag.Parse()

	myServiceName = serviceName

	DoRegistService(consulAddr, statusMonitorAddr, serviceName, serviceIp, servicePort)

	go DoDiscover(consulAddr, foundService)

	go StartService(statusMonitorAddr)

	go WaitToUnRegistService()

	go DoUpdateKeyValue(consulAddr, serviceName, serviceIp, servicePort)

	select {}
}

func DoRegistService(consulAddr string, monitorAddr string, serviceName string, ip string, port int) {
	myServiceID = serviceName + "-" + ip
	var tags []string
	service := &api.AgentServiceRegistration{
		ID:      myServiceID,
		Name:    serviceName,
		Port:    port,
		Address: ip,
		Tags:    tags,
		Check: &api.AgentServiceCheck{
			HTTP:     "http://" + monitorAddr + "/status",
			Interval: "5s",
			Timeout:  "1s",
		},
	}

	config := api.DefaultConfig()
	config.Address = consulAddr

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	consulClient = client
	if err := consulClient.Agent().ServiceRegister(service); err != nil {
		log.Fatal(err)
	}
	log.Printf("Registered service %q in consul with tags %q", serviceName, strings.Join(tags, ","))
}

func WaitToUnRegistService() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	if consulClient == nil {
		return
	}
	if err := consulClient.Agent().ServiceDeregister(myServiceID); err != nil {
		log.Fatal(err)
	}
}

func DoDiscover(consulAddr string, foundService string) {
	t := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-t.C:
			DiscoverServices(consulAddr, true, foundService)
		}
	}
}

func DiscoverServices(addr string, healthyOnly bool, service_name string) {
	consulConf := api.DefaultConfig()
	consulConf.Address = addr
	client, err := api.NewClient(consulConf)
	CheckErr(err)

	services, _, err := client.Catalog().Services(&api.QueryOptions{})
	CheckErr(err)

	fmt.Println("--do discover ---:", addr)

	var sers ServiceList
	for name := range services {
		servicesData, _, err := client.Health().Service(name, "", healthyOnly,
			&api.QueryOptions{})
		CheckErr(err)
		for _, entry := range servicesData {
			if service_name != entry.Service.Service {
				continue
			}
			for _, health := range entry.Checks {
				if health.ServiceName != service_name {
					continue
				}
				fmt.Println("  health nodeid:", health.Node, " service_name:", health.ServiceName, " service_id:", health.ServiceID, " status:", health.Status, " ip:", entry.Service.Address, " port:", entry.Service.Port)

				var node ServiceInfo
				node.IP = entry.Service.Address
				node.Port = entry.Service.Port
				node.ServiceID = health.ServiceID

				//get data from kv store
				s := GetKeyValue(service_name, node.IP, node.Port)
				if len(s) > 0 {
					var data KVData
					err = json.Unmarshal([]byte(s), &data)
					if err == nil {
						node.Load = data.Load
						node.Timestamp = data.Timestamp
					}
				}
				fmt.Println("service node updated ip:", node.IP, " port:", node.Port, " serviceid:", node.ServiceID, " load:", node.Load, " ts:", node.Timestamp)
				sers = append(sers, node)
			}
		}
	}

	serviceLocker.Lock()
	servicsMap[service_name] = sers
	serviceLocker.Unlock()
}

func DoUpdateKeyValue(consulAddr string, serviceName string, ip string, port int) {
	t := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-t.C:
			StoreKeyValue(consulAddr, serviceName, ip, port)
		}
	}
}

func StoreKeyValue(consulAddr string, serviceName string, ip string, port int) {

	myKey = myServiceName + "/" + ip + ":" + strconv.Itoa(port)

	var data KVData
	data.Load = rand.Intn(100)
	data.Timestamp = int(time.Now().Unix())
	bys, _ := json.Marshal(&data)

	kv := &api.KVPair{
		Key:   myKey,
		Flags: 0,
		Value: bys,
	}

	_, err := consulClient.KV().Put(kv, nil)
	CheckErr(err)
	fmt.Println(" store data key:", kv.Key, " value:", string(bys))
}

func GetKeyValue(serviceName string, ip string, port int) string {
	key := serviceName + "/" + ip + ":" + strconv.Itoa(port)

	kv, _, err := consulClient.KV().Get(key, nil)
	if kv == nil {
		return ""
	}
	CheckErr(err)

	return string(kv.Value)
}

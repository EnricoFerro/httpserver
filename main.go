package main

import (
	"io"
	"net/http"
	"os"

	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"httpserver/adapters"
	"httpserver/config"
)

type Ips struct {
	Ip          string `json:"ip"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Content struct {
	Resource string `json:"resource,omitempty"`
	Ips      []Ips  `json:"ips,omitempty"`
	Ip       string `json:"ip,omitempty"`
	HostName string `json:"hostname"`
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	conf := config.ReadConfig()
	//Get Resource
	c := Content{}
	if conf.Resource != (config.Resource{}) {
		c = Content{Resource: conf.Resource.Name}
	}

	//Get Ips
	awins, err := adapters.Adapters()
	if err == nil {
		for _, awin := range awins {
			//t.Logf("name=%s, description=%s\n", awin.Name, awin.Description)
			for _, ipnet := range awin.IPNets {
				if conf.Ips.Read == config.All {
					if !ipnet.IP.IsLoopback() && !ipnet.IP.IsUnspecified() {
						ip := Ips{
							Ip:          ipnet.IP.String(),
							Name:        awin.Name,
							Description: awin.Description,
						}
						c.Ips = append(c.Ips, ip)
					}
				} else if conf.Ips.Read == config.Last {
					if !ipnet.IP.IsLoopback() && !ipnet.IP.IsUnspecified() {
						ip := Ips{
							Ip:          ipnet.IP.String(),
							Name:        awin.Name,
							Description: awin.Description,
						}
						c.Ip = ip.Ip
					}
				} else if conf.Ips.Read == config.First {
					if c.Ip == "" {
						if !ipnet.IP.IsLoopback() && !ipnet.IP.IsUnspecified() {
							ip := Ips{
								Ip:          ipnet.IP.String(),
								Name:        awin.Name,
								Description: awin.Description,
							}
							c.Ip = ip.Ip
						}
					}
				}

				//t.Logf("addr=%s, mask=%s\n", ipnet.IP, ipnet.Mask)
			}
		}
	}

	//Get Hostname
	hostname, err := os.Hostname()
	if err == nil {
		c.HostName = hostname
	}

	content, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	io.WriteString(w, string(content))
}

func main() {

	conf := config.ReadConfig()

	// Launch the server
	http.HandleFunc("/", getRoot)
	port := ":" + strconv.FormatInt(conf.Server.Port, 10)
	fmt.Printf("Starting server at port: %s\n", port)
	err := http.ListenAndServe(port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}

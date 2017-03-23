package main
 
import(
"fmt"
"os"
"net/http"
"net"
)

func indexHandler( w http.ResponseWriter, r *http.Request){
    var hostname string
    var ipaddress string
    hostname = get_hostname()
    addrs, err := net.LookupIP(hostname)
    if err != nil {
        //add error handling
    } else {
        for _, a := range addrs {
            ipaddress = a.String()
        }
    }
    fmt.Fprintf(w, "<html><h2>Brian's Golang Web App</h2><p>Spaces are better than tabs</p><p>Version 2.0.4</p><p>Hostname: %s</p><p>IP address: %s</p></html>", hostname, ipaddress)
}
 
func main(){
    http.HandleFunc("/", indexHandler)
    http.ListenAndServe(":8001",nil)
}

func get_hostname() string {
    var result string
    localhost_name, err := os.Hostname()

    if err != nil {
        result = "ERROR: Cannot find server hostname"
    } else {
        result = localhost_name
    }
    return result
}
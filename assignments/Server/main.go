//Khareen Proverbs

package main

import (
	"encoding/json"
	"fmt"
	"log"

	//"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Assignment 3
//  Convert json data article :=
//  `{"id": "BM-1347", "title": "The underage storm", "Content": "The creatives' careers can easily get uncreative but yet creative...", "Summary": "Seeking freedom"}`
//  ...........into  map and send it to client as GET http request
func convertJson(w http.ResponseWriter,_ *http.Request){
	//json data
	article := `{"id": "BM-1347", "title": "The underage storm", "Content": "The creatives' careers can easily get uncreative but yet creative...", "Summary": "Seeking freedom"}`               

	res:= map[string]string{}
	// Unmarshal or Decode the JSON to the interface.
    json.Unmarshal([]byte(article), &res)
	 // Reading each value by its key
	 fmt.Println("id :", res["id"],
	 "\ntitle :", res["title"],
	 "\ncontent :", res["Content"],
	 "\nsummary :", res["Sumary"])
	 w.Header().Set("Content-Type","application/json")//set data type
	json.NewEncoder(w).Encode(article)//convert map into json data (key: value)
}


//____________________________________________________

// Assignment 2
// Write  a go program to get the client IP and reflect/echo back to client page.
// After this add feature to server so that server could block to 
// a particular client with a message add feature to store all past access like 
// date and time to access to the server and display a message that  
// clientIP:xxxxxxxxxxxx  your last access to server is  
// Date  :  xxxx/xx/xxxx time :xxxxxxx


 var lastVisit = make(map[string]time.Time)

func IpHandler(rw http.ResponseWriter, request *http.Request) {
	log.Println(request.Method, request.RequestURI)
	ip := ReadUserIP(request)
	rw.Header().Set("Content-Type", "text/html")
	visit, ok := lastVisit[ip]
	lastVisit[ip] = time.Now()
	if !ok {
		fmt.Fprintln(rw, `<p>This is your first visit.</p>`)
		fmt.Fprintf(rw, "<p>It is %v at %v</p>", lastVisit[ip].Format("January 02, 2006"), lastVisit[ip].Format(time.Kitchen))
		fmt.Fprintf(rw, `<p>Your ip is <pre>%s</pre></p>`, ip)
		log.Printf("IP %v first entry", ip)
	} else {
		fmt.Fprintf(rw, "<p>You last visted on date %v and time %v</p>", visit.Format("January 02, 2006"), visit.Format(time.Kitchen))
		fmt.Fprintf(rw, `<p>Your ip is <pre>%s</pre></p>`, ip)
		log.Printf("IP %v entering again\n", ip)
	}
}

// ReadUserIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.

func ReadUserIP(r *http.Request) string {
    IPAddress := r.Header.Get("X-Real-Ip")
    if IPAddress == "" {
        IPAddress = r.Header.Get("X-Forwarded-For")
    }
    if IPAddress == "" {
        IPAddress = r.RemoteAddr
    }
    return IPAddress
}
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{
		"ip": ReadUserIP(r),
	})
	w.Write(resp)
}
//End of assignment 2
//______________________________________________________

//parameters() write to client, response from client)
func handler(w http.ResponseWriter,_ *http.Request){
	//print to client "w"
	fmt.Fprint(w, "Hello World. Thanks for connecting.")
}
func handler1(w http.ResponseWriter,_ *http.Request){
	//print to client "w"
	fmt.Fprint(w, "Handler 1 US")
}
func Now(w http.ResponseWriter,_ *http.Request){
	//print to client "w"
	now:=time.Now()
	//map key string, value string
	p:=make(map[string]string)
	p["now"]= now.Format(time.ANSIC)
	w.Header().Set("Content-Type","application/json")//set data type
	json.NewEncoder(w).Encode(p)//convert map into json data (key: value)
}


//_____________________________________________________
//Assignment 1
//Time zone Specific API
//which will return time specific to the client time zone
func NowShanghai(w http.ResponseWriter,_ *http.Request){
	//init the location
	loc, _ := time.LoadLocation("Asia/Shanghai")
	//set timezone,  
	now:= time.Now().In(loc)	
	p:=make(map[string]string)//map key string, value string
	p["now"]= now.Format(time.ANSIC)
	w.Header().Set("Content-Type","application/json")//set data type
	json.NewEncoder(w).Encode(p)//convert map into json data (key: value)
}
//_____________________________________________________



// " /" is default path
func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", handler).Methods("GET")//Get from server
    r.HandleFunc("/US", handler1).Methods("GET")//Get from server
    r.HandleFunc("/time", Now).Methods("GET")//Get from server
    r.HandleFunc("/timeAsia", NowShanghai).Methods("GET")//Get from server
    r.HandleFunc("/ping", ExampleHandler).Methods("GET")//Get from server
    r.HandleFunc("/convert", convertJson).Methods("GET")//Get from server
	http.Handle("/",r)
	http.ListenAndServe(":8000",nil)//server port and error handler
}

package main

import ( 
    "fmt" 
    "io" 
    "net" 
    "time"
    "net/http" 
    "net/rpc" 
)

type Watcher int

func (w *Watcher) GetInfo(arg string, result *int64) error { 
    if arg!="test" {
        *result = 0
    }else{
        *result = time.Now().UnixNano()
    }
    
    return nil 
}

func main() {

    http.HandleFunc("/time", MyTest)

    watcher := new(Watcher) 
    rpc.Register(watcher) 
    rpc.HandleHTTP()

    l, err := net.Listen("tcp", ":8888") 
    if err != nil { 
        fmt.Println("Failed,the port may be occupied") 
    } 
    fmt.Println("looking 8888 port") 
    http.Serve(l, nil) 
}

func MyTest(w http.ResponseWriter, r *http.Request) { 
    io.WriteString(w, "<html><body>have getport</body></html>") 
}
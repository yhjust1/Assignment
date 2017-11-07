package main

import ( 
    "fmt" 
    "time"
    "net/rpc" 
)

func main() { 
    client, err := rpc.DialHTTP("tcp", "114.212.86.128:8888") 
    if err != nil { 
        fmt.Println("connect rpc server failed!", err) 
    } 
    var reply int64 
    err = client.Call("Watcher.GetInfo", "test", &reply) 
    sTime:=time.Now().UnixNano()
    if err != nil { 
        fmt.Println("promote program server call failed", err) 
    } 
    if reply == 0{
        fmt.Println("The request was rejected !") 
    }else{
        delay:=(time.Now().UnixNano()-sTime)/2
    fmt.Println("Delay is :",delay)
    fmt.Println("Time is :", time.Unix(0,reply+delay)) 
    }
    
}
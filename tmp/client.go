package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)
func main() { 
   client := http.Client{} 
   resp, err := client.Get("http://localhost:5505/add?a=6&b=8")
   if err != nil { 
         fmt.Println(err) 
         return
   } 

   io.Copy(os.Stdout, resp.Body)

    client = http.Client{}
    resp, err = client.Get("http://localhost:5505/sub?a=8&b=6")
    if err != nil {
        fmt.Println(err)
        return
    }

    io.Copy(os.Stdout, resp.Body)

    client = http.Client{}
    resp, err = client.Get("http://localhost:5505/mul?a=2&b=2")
    if err != nil {
        fmt.Println(err)
        return
    }

    io.Copy(os.Stdout, resp.Body)

    client = http.Client{}
    resp, err = client.Get("http://localhost:5505/div?a=6&b=3")
    if err != nil {
        fmt.Println(err)
        return
    }

    io.Copy(os.Stdout, resp.Body)
}

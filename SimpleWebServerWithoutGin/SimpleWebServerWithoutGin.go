package main

import(
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"
  "encoding/json"
)

type Request struct {
		Site []string
		SearchText string
}

type Response struct {
		FoundAtSite string
}

func main(){
  http.HandleFunc("/", checkText)
  http.ListenAndServe(":8080", nil)
}

func checkText(w http.ResponseWriter, r *http.Request){

  var req Request
  var res Response

  body, err := ioutil.ReadAll(r.Body)
  r.Body.Close()
  if err != nil {
  fmt.Println("Don`t have Body")
  }
  json.Unmarshal(body,&req)

  for n:=0; n<(len(req.Site)); n++ {
    source, err := http.Get(req.Site[n])
    if err != nil {
    fmt.Println("Don`t get Body")
    }
    siteBody, err := ioutil.ReadAll(source.Body)
    source.Body.Close()
    if err != nil {
    fmt.Println("Don`t read Body")
    }
    strSiteBody:=string(siteBody)
    if strings.Contains(strSiteBody, req.SearchText){
      if len(res.FoundAtSite) ==0{
        res.FoundAtSite = req.Site[n]
        resp,err:=json.Marshal(res)
          if err != nil {
          fmt.Println("Don`t encoding to JSON")
          }
        w.Write(resp)
      }else{
        res.FoundAtSite += req.Site[n]
        resp,err:=json.Marshal(res)
        if err != nil {
        fmt.Println("Don`t encoding to JSON")
        }
        w.Write(resp)
      }
    }else{w.WriteHeader(http.StatusNoContent) }
  }
}

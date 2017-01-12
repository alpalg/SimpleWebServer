package main

import(
  "fmt"
  "io/ioutil"
  "net/http"
  "gopkg.in/gin-gonic/gin.v1"
  "strings"
)

type Request struct {
		Site []string
		SearchText string
}

type Response struct {
		FoundAtSite string
}

func main(){
  router := gin.New()
  router.POST("/",checkText())
  router.Run(":8080")
}

func checkText() gin.HandlerFunc{
  return func(c *gin.Context) {
    var req Request
    var res Response
    c.BindJSON(&req)
    fmt.Println(req.SearchText)
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
          fmt.Println(res.FoundAtSite)
          c.JSON(200, res)
        }else{
          c.JSON(200, res)
        }
      }else{
        c.Status(204)
      }
    }
  }
}

package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"golang.org/x/net/html"
)

var client = http.Client{
	Timeout: 2 * time.Second,
}

func IsTimeoutError(err error) bool {
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
	  return true
	}
   return false
   }

   var w = tabwriter.NewWriter(os.Stdout , 50, 0 , 2 , ' ' , 0)
func scrape(WebsiteUrl string , url string , hashMap map[string]string)  {
	var newUrl string
	hashMap[url] = "visited"
	if strings.HasPrefix(url , "/"){
		newUrl = fmt.Sprintf("%s%s" , WebsiteUrl, url)
	}else if strings.HasPrefix(url , WebsiteUrl) {
		newUrl = url
	} else {
		response, err := client.Get(url);
		if err != nil {
			 if IsTimeoutError(err){
				fmt.Fprintf(w,"link:%s\tresponse:%s\n" , url , "Timeout")
				w.Flush()
				return
			}
			return
		}
		if response.StatusCode != 200 {
			fmt.Fprintf(w,"link:%s\tresponse:%d\n" , url , response.StatusCode)
			w.Flush()
		    return 
		}
		return
	}
	
	response , err := client.Get(newUrl);
	if err != nil || response.StatusCode != 200{
		fmt.Fprintf(w,"link:%s\tresponse:%d\n" , url , response.StatusCode)
		w.Flush()
		return 
	} 

	defer response.Body.Close()

	tkn := html.NewTokenizer(response.Body)
	for { 
		
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return 
		case tt == html.StartTagToken:
			t := tkn.Token()

			isAnchor := t.Data == "a"

			if isAnchor{
				for _ , a := range t.Attr{
					if a.Key == "href"{
						_ , ok := hashMap[a.Val];
						if !ok {
							scrape(WebsiteUrl , a.Val , hashMap)
						}
						
					}
				}
			}
		}
	}
}




func main(){
	hashMap := make(map[string]string)
	rootUrl := "https://scrape-me.dreamsofcode.io"
	scrape(rootUrl , rootUrl , hashMap)
}
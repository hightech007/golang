// http client example

// possible reference:
// http://comments.gmane.org/gmane.comp.lang.go.general/61007

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"code.google.com/p/go.net/html"
	"strings"
)


func main() {
    
    //download html
	res, err := http.Get("http://uavsar.jpl.nasa.gov/cgi-bin/product.pl?jobName=Guatml_11201_11017-131_13034-000_0682d_s01_L090_01")
	if err != nil {
		log.Fatal(err)
	}
	
	robots, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	
	//fmt.Printf("%s", robots)
	
	//parse html
	//reference: http://code.google.com/p/go/source/browse/html/example_test.go?repo=net
	doc, err := html.Parse(strings.NewReader(string(robots)))
	if err != nil {
	    log.Fatal(err)
	}
	
	var f func (*html.Node)
	f = func (n *html.Node) {
	    if n.Type == html.ElementNode && n.Data == "a" {
	        for _, a := range n.Attr {
	            if a.Key == "href" {
	                if strings.Contains(a.Val,".kml") {
	                    fmt.Println(a.Val)
	                    break
	                }
	            }
	        }
	    }
	    for c := n.FirstChild; c!=nil; c= c.NextSibling {
	        f(c)
	    }
	}
	f(doc)
}

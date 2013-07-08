// http client example

// possible reference:
// http://comments.gmane.org/gmane.comp.lang.go.general/61007

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)


func main() {
	res, err := http.Get("http://uavsar.jpl.nasa.gov/cgi-bin/product.pl?jobName=Guatml_11201_11017-131_13034-000_0682d_s01_L090_01")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
	
	//parse html
	//extract url for ann and kml
}

/*
* A simple hello World like application for first try with GO
* Compile with go build or run with go run
 */
package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	// Use public function "Println" of package fmt
	fmt.Println("fEEd rEAder ...")
	fmt.Println("==================================")
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")
	for i := 0; i < 5; i++ {
		fmt.Println(feed.Items[i].Title + " --> " + feed.Items[i].Published)
	}
	fmt.Println("==================================")
}

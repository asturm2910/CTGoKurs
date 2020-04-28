package cmd

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func list() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")
	for i := 1; i < 5; i++ {
		fmt.Println(feed.Items[i].GUID + ":    " + feed.Items[i].Title + " (" + feed.Items[i].Published + ")")
	}
}

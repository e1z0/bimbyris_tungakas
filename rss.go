package main

import (
	"fmt"
	"strings"
	"github.com/mmcdole/gofeed"
        "time"
        "context"
)

func news() string {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURLWithContext(settings.News_source, ctx)
	var news_data []string
	news_data = append(news_data,fmt.Sprintf("<b>%s</b>",feed.Title))
	for i,_ := range feed.Items {
		fmt.Printf("%s -> %s\n",feed.Items[i].Title,feed.Items[i].Link)
		news_data = append(news_data,fmt.Sprintf("%d. <a href=\"%s\">%s</a>",i+1,feed.Items[i].Link,feed.Items[i].Title))
  	}
	return strings.Join(news_data, "\n")
}

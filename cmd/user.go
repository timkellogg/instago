package cmd

import (
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"

	"github.com/spf13/cobra"
)

const (
	identifier           = "window._sharedData"
	userDataRoot         = "entry_data.ProfilePage.0.user"
	userMediaRoot        = userDataRoot + ".media.nodes"
	userMediaItemRoot    = userMediaRoot + ".#."
	userItemCaption      = userMediaItemRoot + "caption"
	userItemCommentCount = userMediaItemRoot + "comments.count"
	userItemDate         = userMediaItemRoot + "date"
	userItemLikes        = userMediaItemRoot + "likes.count"
	userItemLinks        = userMediaItemRoot + "thumbnail_src"
)

type item struct {
	caption  string
	comments int64
	date     time.Time
	likes    int64
	link     string
}

type list []item

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User's feed which to scrape",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			data := scrapeInstagram(arg)
			items := sortInstagramItems(data)

			fmt.Printf("Top Posts by %s", arg)

			for i, item := range items {
				fmt.Printf("%d.\tLikes: %d\n \tLink: %s\n\n", i+1, item.likes, item.link)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(userCmd)
}

func scrapeInstagram(profile string) list {
	doc, err := goquery.NewDocument("https://www.instagram.com/" + profile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var data string
	var posts []item

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.Contains(text, identifier) {
			data = strings.Replace(text, identifier, "", 1)
		}
	})

	captions := gjson.Get(data, userItemCaption).Array()
	comments := gjson.Get(data, userItemCommentCount).Array()
	dates := gjson.Get(data, userItemDate).Array()
	likes := gjson.Get(data, userItemLikes).Array()
	links := gjson.Get(data, userItemLinks).Array()

	for i := 0; i < len(captions); i++ {
		caption := captions[i].Str
		comment := comments[i].Int()
		date := dates[i].Time()
		like := likes[i].Int()

		link, err := url.Parse(links[i].Str)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		item := item{
			caption:  caption,
			comments: comment,
			date:     date,
			likes:    like,
			link:     link.String(),
		}

		posts = append(posts, item)
	}

	return posts
}

func sortInstagramItems(items list) list {
	sort.Slice(items[:], func(i, j int) bool {
		return items[i].likes > items[j].likes
	})

	return items
}

package cmd

import (
	"fmt"
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
	userItemLink         = userMediaItemRoot + "thumbnail_src"
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

			for i, item := range items {
				fmt.Printf("%d. %d\n", i+1, item.likes)
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

	results := gjson.GetMany(data,
		// userItemCaption,
		// userItemCommentCount,
		// userItemDate,
		userItemLikes)
	// userItemLink)

	for i, result := range results {
		for j, res := range result.Array() {
			fmt.Printf("result: %s", res)
			fmt.Println("")
			fmt.Printf("inner %d", j)
			fmt.Println("")

			item := item{
				likes: res.Int(),
			}

			posts = append(posts, item)

		}
		// post := result.Array()

		fmt.Println(i)
		fmt.Println(result)

		// caption := post[0].Str
		// comments := post[1].Int()
		// date := post[2].Time()
		// likes := post[3].Int()
		// link := post[4].Str

		// item := item{
		// 	caption:  caption,
		// 	comments: comments,
		// 	date:     date,
		// 	likes:    likes,
		// 	link:     link,
		// }

		// posts = append(posts, item)
	}
	// for i, result := range collection {
	// 	fmt.Println(i)
	// 	for _, res := range result.Array() {
	// 		// if
	// 		post := res.Array()

	// 		fmt.Println(post[0].Str)
	// 		// caption := post[0].Str
	// 		// comments := post[1].Int()
	// 		// date := post[2].Time()
	// 		// likes := post[3].Int()
	// 		// link := post[4].Str

	// 		// item := item{
	// 		// 	caption:  caption,
	// 		// 	comments: comments,
	// 		// 	date:     date,
	// 		// 	likes:    likes,
	// 		// 	link:     link,
	// 		// }

	// 		// posts = append(posts, item)
	// 	}
	// }

	return posts
}

func sortInstagramItems(items list) list {
	sort.Slice(items[:], func(i, j int) bool {
		return items[i].likes < items[j].likes
	})

	return items
}

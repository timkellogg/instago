package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/spf13/cobra"
)

type item struct {
	link  string
	likes int
	date  string
}

type list []item

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User's feed which to scrape",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			// get data
			data := scrapeInstagram(arg)
			items := sortInstagramItems(data)

			for i, item := range items {
				fmt.Printf("%d. %s - %s", i, item.link, item.link)
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

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.Contains(text, "window._sharedData") {
			data = text
		}
	})

	return []item{}
}

func sortInstagramItems(items list) list {
	return []item{}
}

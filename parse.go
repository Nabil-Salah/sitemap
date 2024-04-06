package link

import (
	"golang.org/x/net/html"
	"unicode"
)

// Link represents a link (<a href="...">) in an HTML
// document.
type Link struct {
	Href string
	Text string
}

func Parse(htmlParse *html.Tokenizer) ([]Link, error) {
	var err error
	var links []Link
	for {
		nextnode := htmlParse.Next()
		if nextnode == html.ErrorToken {
			err = html.ErrBufferExceeded
			break
		}
		nodeData := htmlParse.Token()
		if nodeData.Type.String() == "StartTag" && nodeData.Data == "a" {
			link := Link{}
			for _, attribute := range nodeData.Attr {
				if attribute.Key == "href" {
					link.Href = attribute.Val
					break
				}
			}
			link.Text = linkText(htmlParse)
			links = append(links, link)
		}
	}
	return links, err
}
func checkIfStart(char rune, st int) bool {
	return unicode.IsLetter(rune(char)) && st == -1
}

func checkIfEnd(char rune, st int, i int, n int) bool {
	return (!unicode.IsLetter(rune(char)) && st != -1) || (i == n-1 && st != -1)
}

func linkText(htmlParse *html.Tokenizer) string {
	var text string = ""
	for {
		nextnode := htmlParse.Next()
		if nextnode == html.ErrorToken {
			break
		}
		nodeData := htmlParse.Token()
		if nodeData.Type.String() == "EndTag" && nodeData.Data == "a" {
			break
		} else if nodeData.Type.String() == "Comment" {
			continue
		} else {
			var newText = nodeData.Data
			st, n := -1, len(newText)
			for i := 0; i < n; i++ {
				if checkIfStart(rune(newText[i]), st) {
					st = i
				} else if checkIfEnd(rune(newText[i]), st, i, n) {
					if len(text) != 0 {
						text += " "
					}
					text += newText[st:i]
					st = -1
				}
			}

		}

	}
	return text
}

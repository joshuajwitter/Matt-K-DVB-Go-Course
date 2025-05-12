package lesson_11

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"strings"
)

// the raw HTML we will be parsing
var _ = `
<!DOCTYPE html>
<html>
	<body>
		<h1> My First Heading</h1>
		<p>My first paragraph</p
		<p>HTML images are defined with the img tag:</p>
		<img src="xxx.jpg" width="10" height="142">
	</body>
</html>`

func Lesson11() {
	//fmt.Println(raw)

	page, pageErr := fetchURL("https://www.amazon.com")
	if pageErr != nil {
		_, err := fmt.Fprintf(os.Stderr, "page download failed %s", pageErr)
		if err != nil {
			os.Exit(-1)
		}
		os.Exit(-1)
	}

	// this is going to give us a reference to the top node of the html tree structure
	reader := bytes.NewReader([]byte(page))
	doc, err := html.Parse(reader)

	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "parse failed %s", err)
		if err != nil {
			os.Exit(-1)
		}
		os.Exit(-1)
	}

	words, pics := countWordsAndImages(doc)

	fmt.Printf("%d words and %d images", words, pics)
}

// we can make this a recursive function to make it obvious what it does
func countWordsAndImages(doc *html.Node) (int, int) {

	// could this be a closure?
	var words, pics int
	visit(doc, &words, &pics)
	return words, pics
}

func visit(node *html.Node, wordsPtr, picsPtr *int) {

	// if this is a <script> or <style>, drop the whole subtree
	if node.Type == html.ElementNode && (node.Data == "script" || node.Data == "style") {
		return
	}

	if node.Type == html.TextNode {
		for _, w := range strings.Fields(node.Data) {
			fmt.Println(w)
		}
		*wordsPtr += len(strings.Fields(node.Data))
	} else if node.Type == html.ElementNode && node.Data == "img" {
		*picsPtr++
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		visit(c, wordsPtr, picsPtr)
	}
}

func fetchURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

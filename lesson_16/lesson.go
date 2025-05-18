package lesson_16

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const cacheFilename = "lesson_16/posts.txt"

type XKCDPost struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safeTitle"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func DoLesson() {
	var searchTerms []string

	// # go run ./cmd the barrel
	// Fetching posts...
	// Getting posts from the cache...
	// Looking to see if any match the search terms: '[the barrel]'...
	// Comic 1 contains 'the'
	// Comic 1 contains 'barrel'
	// Comic 2 contains 'the'
	// Comic 4 contains 'the'
	// Comic 5 contains 'the'

	for _, term := range os.Args[1:] {
		searchTerms = append(searchTerms, strings.ToLower(term))
	}
	var found bool

	fmt.Println("Fetching posts...")
	posts, err := fetchPosts(500)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Looking to see if any match the search terms: '%v'...\n", searchTerms)
	for _, post := range posts {
		title := strings.ToLower(post.Title)
		transcript := strings.ToLower(post.Transcript)

		for _, searchTerm := range searchTerms {
			if strings.Contains(title, searchTerm) || strings.Contains(transcript, searchTerm) {
				found = true
				fmt.Printf("Comic %d contains '%s'\n", post.Num, searchTerm)
			}
		}
	}

	if !found {
		fmt.Printf("No comic matched the search criteria")
	}
}

func cacheExists() bool {
	_, err := os.Stat(cacheFilename)
	return !os.IsNotExist(err)
}

func fetchPosts(maxPosts int) ([]XKCDPost, error) {
	if !cacheExists() {
		fmt.Println("Getting posts from the network...")
		return fetchPostsFromNetwork(maxPosts)
	} else {
		fmt.Println("Getting posts from the cache...")
		return fetchPostsFromCache()
	}
}

func fetchPostsFromCache() ([]XKCDPost, error) {

	posts := []XKCDPost{}

	// create the cache
	file, err := os.OpenFile(cacheFilename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Could not open cache file")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var post XKCDPost
		line := scanner.Text()
		err := json.Unmarshal([]byte(line), &post)
		if err != nil {
			return nil, fmt.Errorf("Could not unmarshal line %d: %v\n", len(posts)+1, err)
		}
		posts = append(posts, post)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading cache file: %w", err)
	}

	return posts, nil
}

func fetchPostsFromNetwork(maxPosts int) ([]XKCDPost, error) {

	var posts []XKCDPost

	for i := 1; i <= maxPosts; i++ {
		post, err := fetchPostFromNetwork(i)
		if err != nil {
			return nil, fmt.Errorf("Could not fetch post from network: %v\n", err)
		}
		posts = append(posts, post)
		cachePost(post)
	}

	return posts, nil
}

func fetchPostFromNetwork(id int) (XKCDPost, error) {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", id)
	resp, err := http.Get(url)

	if err != nil {
		return XKCDPost{}, fmt.Errorf("Could not fetch post from network: %v\n", err)
	}

	// Defer closing the body, capturing any potential Close() error
	var closeErr error
	defer func() {
		if err := resp.Body.Close(); err != nil {
			closeErr = fmt.Errorf("failed to close response body: %w", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return XKCDPost{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return XKCDPost{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var post XKCDPost
	if err := json.Unmarshal(body, &post); err != nil {
		return XKCDPost{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Check if there was a close error before returning
	if closeErr != nil {
		return XKCDPost{}, closeErr
	}

	return post, nil
}

func cachePost(post XKCDPost) error {
	// create the cache
	file, err := os.OpenFile(cacheFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("Could not open file: %w", err)
	}

	defer file.Close()

	b, err := json.Marshal(post)

	if err != nil {
		return fmt.Errorf("Could not marshal JSON: %w", err)
	}

	_, err = file.Write(append(b, '\n'))

	if err != nil {
		return fmt.Errorf("Could not write to cache: %w", err)
	}

	return nil
}

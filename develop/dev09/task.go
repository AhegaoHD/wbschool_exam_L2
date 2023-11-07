package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// downloadResource скачивает ресурс по URL и сохраняет его в директории outputPath.
func downloadResource(resourceUrl, outputPath string) error {
	resp, err := http.Get(resourceUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned %v status", resp.StatusCode)
	}

	// Создание директории, если она не существует
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

// parseLinks извлекает все ссылки из HTML.
func parseLinks(body io.Reader, base *url.URL) ([]string, error) {
	links := []string{}
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return links, nil
		case html.StartTagToken, html.SelfClosingTagToken:
			token := z.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						// Разрешение относительных ссылок
						link, err := base.Parse(attr.Val)
						if err != nil {
							continue
						}
						links = append(links, link.String())
					}
				}
			}
		}
	}
}

func main() {
	// Флаги командной строки
	baseUrl := flag.String("url", "", "The base URL of the site to download")
	outputDir := flag.String("output", "site_download", "The directory to save the downloaded site")
	flag.Parse()

	if *baseUrl == "" {
		fmt.Println("The base URL is required.")
		os.Exit(1)
	}

	base, err := url.Parse(*baseUrl)
	if err != nil {
		fmt.Println("Invalid base URL.")
		os.Exit(1)
	}

	// Скачиваем базовую страницу
	resp, err := http.Get(*baseUrl)
	if err != nil {
		fmt.Printf("Error downloading the base page: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Server returned %v status\n", resp.StatusCode)
		os.Exit(1)
	}

	links, err := parseLinks(resp.Body, base)
	if err != nil {
		fmt.Printf("Error parsing links: %s\n", err)
		os.Exit(1)
	}

	for _, link := range links {
		fmt.Printf("Downloading link: %s\n", link)

		parsedUrl, err := url.Parse(link)
		if err != nil {
			fmt.Printf("Error parsing link '%s': %s\n", link, err)
			continue
		}

		if strings.HasPrefix(parsedUrl.Path, "/") {
			filePath := filepath.Join(*outputDir, parsedUrl.Path)
			if err := downloadResource(link, filePath); err != nil {
				fmt.Printf("Error downloading resource '%s': %s\n", link, err)
			}
		}
	}
}

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://s1.hdslb.com/bfs/seed/laputa-header/style.css?v=1684231477991", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"`)
	req.Header.Set("Referer", "https://www.bilibili.com/video/BV1NR4y167Hn/?spm_id_from=333.1007.top_right_bar_window_default_collection.content.click&vd_source=5e61197d1a7388a859a7ea2e60555802")
	req.Header.Set("If-None-Match", "91e15fdd774b3408202c24c99ad71e6b")
	req.Header.Set("If-Modified-Since", "Tue, 16 May 2023 10:04:38 GMT")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

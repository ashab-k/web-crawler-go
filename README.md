# Web Crawler to Detect Dead Links

This project is a **web crawler** built in Go, designed to recursively traverse a website, extract hyperlinks (`<a>` tags), and identify **dead links**. Dead links are URLs that return HTTP responses in the range **400-599** (e.g., 404 Not Found, 500 Internal Server Error).

## Features

- **Dead Link Detection**: Identifies and logs links that respond with HTTP status codes in the range 400-599.
- **HTTP Timeout Handling**: Detects links that fail to respond within a set timeout period.
- **Recursive Crawling**: Scans all hyperlinks on a webpage, including relative and absolute links, and avoids revisiting the same URL.
- **Tabular Logging**: Outputs link responses in a clean tabular format for easy identification.
- **URL Normalization**: Handles relative URLs by converting them to absolute URLs.

## How It Works

1. **Initialization**:
   - Starts with a root URL.
   - Maintains a `hashMap` to track visited URLs and avoid revisiting them.

2. **Crawling and Link Validation**:
   - Extracts all hyperlinks from the webpage's HTML.
   - Normalizes relative links to absolute ones.
   - Sends HTTP GET requests to each link.
   - Checks the response status code:
     - **200-299**: Valid link (ignored in logs).
     - **400-599**: Dead link (logged).
     - **Timeout**: Treated as a dead link and logged.

3. **Recursive Exploration**:
   - Continues crawling through new, unvisited links found on each page.




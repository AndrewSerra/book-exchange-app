#!/bin/sh

docker build -t book_scraper:1.0 

docker run -id --mount type=bind,src=/Users/andrewserra/GithubProjects/book-exchange-app/scrapers/out,target=/usr/local/data/scraper --name book-scraper book_scraper:1.0

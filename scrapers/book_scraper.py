from pathlib import Path
import scrapy
import scrapy.spiders
from datetime import datetime

class BestSellerBookScraper(scrapy.Spider):
    name = 'bestsellerbookspider'

    def __init__(self):
        self.name = 'bestsellerbookspider'

        if not Path("/usr/local/data/scraper").exists():
            Path("/usr/local/data/scraper").mkdir(exist_ok=True)

        self.f = open(f"/usr/local/data/scraper/{self.name}_{datetime.now()}.csv", "w")
        
        col_headers = ",".join(["isbn", "title", "author", "pubDate", "publisher", "genre"])
        self.f.write(f"{col_headers}\n")
    
    def __del__(self):
        if not self.f.closed:
            self.f.close()

    def start_requests(self) -> scrapy.spiders.Iterable[scrapy.Request]:
        yield scrapy.Request(url="https://www.barnesandnoble.com/b/books/_/N-1fZ29Z8q8", callback=self.parse)
        yield scrapy.Request(url="https://www.barnesandnoble.com/b/books/_/N-1fZ1z141jfZ29Z8q8", callback=self.parse)
        yield scrapy.Request(url="https://www.barnesandnoble.com/b/paperback/_/N-1fZ2cnh", callback=self.parse)
        yield scrapy.Request(url="https://www.barnesandnoble.com/b/fiction/books/_/N-1fZ2usxZ29Z8q8", callback=self.parse)
        yield scrapy.Request(url="https://www.barnesandnoble.com/b/nonfiction/books/_/N-1fZ2urcZ29Z8q8", callback=self.parse)
        yield scrapy.Request(url="https://www.barnesandnoble.com/b/bn-stores-bestsellers-hardcover-fiction-nonfiction/_/N-1p4n", callback=self.parse)
        yield scrapy.Request(url="https://www.barnesandnoble.com/b/books/teens-ya/_/N-1fZ29Z8q8Z19r4", callback=self.parse)
    
    def parse(self, response: scrapy.spiders.Response) -> scrapy.spiders.Any:
        for item in response.css(".product-info-view"):
            title = item.css(".product-info-title a::text").extract()
            pubDate = item.css(".product-info-title .publ-date::text").extract()
            authors = item.css(".product-shelf-author a::text").extract()
            path = item.css("a::attr(href)").extract()[0]

            request = scrapy.Request(f"https://barnesandnoble.com/{path}", callback=self.parse_book_page)

            request.meta["title"] = title[0].strip()
            request.meta["pubDate"] = pubDate[0].strip()[1:-1]
            request.meta["author"] = ",".join(authors)

            yield request
        

    def parse_book_page(self, response: scrapy.spiders.Response):
        tab = response.css("#ProductDetailsTab")

        isbn = tab.css("tbody tr:nth-child(1) td::text").extract()
        publisher = tab.css("tbody tr:nth-child(2) td span::text").extract()
        genre = response.css(".related-subject-container .related-sub-text a::text").extract()

        data = {
            "isbn": isbn[0],
            "title": response.meta.get("title"),
            "author": response.meta.get("author"),
            "pubDate": response.meta.get("pubDate"),
            "publisher": publisher[0],
            "genre": ",".join(genre)
        }
        
        self.f.write(f"\"{data.get("isbn")}\",\"{data.get("title")}\",\"{data.get("author")}\",\"{data.get("pubDate")}\",\"{data.get("publisher")}\",\"{data.get("genre")}\"\n")
        

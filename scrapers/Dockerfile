FROM python:3.12-bookworm

WORKDIR /app

RUN apt-get update && apt-get install -y cron

COPY crontab /etc/cron.d/crontab
COPY book_scraper.py .

RUN chmod 0644 /etc/cron.d/crontab

RUN touch /var/log/cron.log

# convert this to requirements.txt
RUN pip3 install scrapy

CMD cron && tail -f /var/log/cron.log


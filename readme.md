# Introduction

This project is explained by these videos:

1. Demonstration: TBD
2. Explanation: TBD
3. SPF/DKIM/DMARC Tutorial: TBD

# Installation

Assume the `~/` reflects your project directory, meaning the same directory as the file `docker-compose.yml`.  So if your `docker-compose.yml` is in `/home/stephanie/dmarc/docker-compose.yml`, then all mentions of `~/` in the instructions below will mean `/home/stephanie/dmarc/`.

1. Make a copy of `~/env.sample` to `~/.env`.  Optional - customize settings for better security.
2. Make a copy of `~/parser/env.sample` to `~/parser/.env`.  Optional - customize settings for better security.
3. Put all your zipped DMARC aggregation reports into the `~/parser/logs/zipped/`.  Sample DMARC aggregration reports can be copied from `~/parser/logs/zipped-sample/`.  The zipped DMARC reports should end with any of the following filename extensions: `*.gz` or `*.zip`.
4. Type `cd ~/` to return to project directory.
5. Type `docker-compose up --build -d` to start up the ELK project.
6. Open your web browser and go to `https://<ip address or hostname of kibana>:5601`.
7. Accept any security warnings about untrusted SSL certificates.
8. Login with `elastic` and the password found in the `~/.env` file to ensure the entire ELK stack is up and running.
9. Type `docker exec -it dmarc-parser-1 ./start.sh` to extract, transform and load DMARC aggregation data into your ELK stack.
10. Go to your web browser in Kibana and go to Dashboards to see your DMARC Dashboard.

![Screenshot](screenshot.png "DMARC Dashboard")

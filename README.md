## Woring Flow:
Anita helps to archive cards on Trello and writing the corresponding working reports to the Google sheet regularly .
Anita contains 4 subsystem:
1. Sheet Server 
2. Scheduler
3. Report Server
4. Burndownchart server


The Sheet Server receives the requests from the Schduler and do some processing.
It retreives working reports from the Google sheet if the request is GET.
If the request is POST, it will collect info from Trello and write them to the Google sheet.


The Report Server provides results when querying working reports of a specified account in a time interval.

### To run the service, there are 3 programs to be run.

* Sheet Server
    - Config
        - client_secret.json :  
            To access Google Sheet, Google OAuth is required in this service. This file contains OAuth info for this service.
        - config.ini:
            There is a section [Google Sheets API] in this file.

            ```sh
           [Google Sheets API]
            SheetID : ID of the Google sheet where the woring reports write.
                      For example : https://docs.google.com/spreadsheets/d/1dDbiww9VU68Alr-lo8p4tHghLgyz1feEQo8TH96Etps/edit#gid=1663811628
                      ID is 1dDbiww9VU68Alr-lo8p4tHghLgyz1feEQo8TH96Etps.
            TabName : The tab name of the specified sheet.
            Range   : The range of column the woring reports write to.
    - Program:
        - Sheet Server:
            Sheet Server serves as web server and listen to port 2001.
            It handles GET and POST request.
            For GET request, it will retrieve the working report from the specifed Google sheet and arrange the output and then return the results to users.
            The user name and a time interval is required.
            For POST request, it will do some processing according to the data in POST body and write the data into the GOOGLE sheet.
            
- Report Server
    Report Server serves as web server and listen to port 2001.
    - Config
        - config.ini : 
            These is a section [Sheet Server] in this file.
            ```sh
            [Sheet Server]
                URL: URL of the Sheet server.
    - Program
        - api:
             It can deal with GET request. When it receives GET request, it will send GET request to the Sheet Server whose URL is listed in the config.ini.
             The returned result is the working report of the sepcifed account in a time interval. The account and the time interval are required in the URL.

             === Usage ===
             http://[ip_addr]:5000/wilson.wang/201609-201611

- Scheduler
    - Config
        - config.ini: There are 3 sections in this config.
            ```sh
            [Trello]
            - UserName: The account who has the permission to read and write the Google sheet.
            - BoardName: Board name on Trello. The cards on the board will be precessed by the Working_report_automation.
            - WaitToReportListName
            - ReportedListName
            - Key: Key for accessing Trello API.
            - Token: Token for the key.
            - ArchiveReportedCardsWeekDay: specified day to do archiving cards on Trello.
            - AutoArchiveTime: time to do archiving cards on Trello.

            [Email Setting]

            [Project Setting]
            - ProjectName


    - Program
        - scheduler:
            This programs will set two cron jobs:
            1. Collecting info from trello and writing the results to the Google sheet; i,e. send POST request to the Sheet Server.
            2. Archive the cards on Trello according to the settings in config.ini.

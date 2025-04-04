# DegreeProgressTracker
Automatically updates and displays my degree progress on Discord based off of a Google Sheet

## Summary
My personal project. I used the Google Sheets API and the Discord API to make a monolith Go app that connects to a Google Sheet that I use to keep track of my credits so far. I then take this info to a discord bot that I made to update my server nickname to reflect this change. Unfortunately, I was not able to configure the permissions correctly to get the discord bot to change my nickname on my test server, but the logs show that it's logically sound. Feel free to configure this locally on your machine and let me know on GH issues if there are any problems! Below, I'll walkthrough how to set up the Sheets and Discord APIs.

## Installation and Configuration
Assuming that you have Go installed (see https://go.dev/dl/ for installation), first clone this project on your computer. Then, make a .env file at the root of the file which will be used to load API keys for the config struct in the program to link to the APIs. I will break it down to configuring the .env file and linking the 2 APIs in 3 subsections below:

### .Env File
1. Under root of file run `touch .env file`
2. Configure .env file like so:
```
GOOGLE_SHEETS_JSON=""
SPREADSHEET_ID=""
TOTAL_CREDITS_RANGE=""
DISCORD_BOT_TOKEN=""
ALLOWED_USER_ID=""
```
3. Throughout the rest of the configuration, I'll tell you what to add for each variable

### Google Sheets API
1. Go to Google Cloud Console (https://console.cloud.google.com/) and create a new project.
2. Search for Google Sheets API, select the first link, and then click the "Enable" button
3. Go to nav menu ("APIs & Services" > "Credentials")
4. Click on the "Create Credentials" button on the top of the screen and select "Service Account"
5. Name your service account, configure role (I selected Viewer but Owner may be necessary for add'l features down the road), and create service account
6. Go back to credentials page, go to service you created, then generate a JSON key
7. Save that file under either the project root (maybe rename to `credentials.json`) or internal/config and add the file path to your .env file GOOGLE_SHEETS_JSON variable.
8. 

### Discord API


## Roadmap
- [x] Make a spreadsheet to keep track of CUs completed at WGU
- [x] Make a Go CLI/REPL project
- [x] Implement and test logic for CLI/REPL
- [x] Have Go CLI fetch CUs completed from Google Sheet w/ Sheets API
- [x] Make a Discord Bot (In Go obviously)
- [ ] Have Discord Bot be able to change my server nickname
- [ ] Be Able to Update Server Nickname By Updating Google Sheet

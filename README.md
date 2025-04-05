# DegreeProgressTracker
Automatically updates and displays my degree progress on Discord based off of a Google Sheet

## Summary
My personal project. I used the Google Sheets API and the Discord API to make a monolith Go app that connects to a Google Sheet that I use to keep track of my credits so far. I then take this info to a discord bot that I made to update my server nickname to reflect this change. Feel free to configure this locally on your machine and let me know on GH issues if there are any problems! Below, I'll walkthrough how to set up the Sheets and Discord APIs.

## Install and Configure
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
7. Save that file under either the project root (maybe rename to `credentials.json`) or internal/config and add the file path to your .env file GOOGLE_SHEETS_JSON variable (the .gitignore ignores .json and .env files)
8. Go to your Google Sheet and find your id in the url. The format is
```
https://docs.google.com/spreadsheets/d/SPREADSHEET_ID/edit?gid=SHEET_ID#gid=SHEET_ID
```
9. Place the Google Sheets Id in the `SPREADSHEET_ID` .env variable.
10. Find the cell that contains your total credits and place it into the TOTAL_CREDITS_RANGE .env variable

### Discord API
1. Go to Discord Developer Portal (https://discord.com/developers/applications) and click "New Application" button
2. Navigate to "Bot" tab on left sidebar and click on it
3. Click "View Token" button and add this in your .env as the `DISCORD_BOT_TOKEN` variable
4. Go to "OAuth2" tab -> "URL Generator" and then click scope to be bot
5. Configure permissions to view channels, send messages, read message history. and manage nicknames
6. Go down to the URL Generator and use this URL to invite bot to desired server
7. If able, configure server settings to ensure that bot is high enough role to change nicknames
8. Finally, add your Discord User ID by going into dev mode, right clicking your profile and copy User ID
9. Add this into the `ALLOWED_USER_ID` .env variable

## Usage
This is a Go monorepo with 2 executable apps (REPL CLI and Discord bot). I'll talk about the REPL CLI first.
### REPL
First, to just run this command you can use the shell command `go run ./cmd/fetchcredits` from the project root, which will compile and run the code without outputting an executable. After the prompt starts up `DegreeProgressTracker > `, run the command `help` to see all options. Run `credits` to fetch the number of credits from the spreadsheet. Finally run `exit` to exit the shell.

If you want to build to use as an executable, from the project root run the shell command `go build ./cmd/fetchcredits`. This will save in the root directory. To run on MacOS/Linux, use `./fetchcredits`, or `./fetchcredits.exe` on Windows.

### Discord Bot
Running the program and building the executable are nearly identical as the REPL, but replaced by `./cmd/creditbot` instead. To use the bot on the desired server, send a `!updatecredits` command on an authorized channel. Keep an eye on the terminal to see the logs to see if execution is successful, and if not how to troubleshoot. Finally, to kill the bot execution, just use Ctrl^C and the bot will gracefully shut down.

## Roadmap
- [x] Make a spreadsheet to keep track of CUs completed at WGU
- [x] Make a Go CLI/REPL project
- [x] Implement and test logic for CLI/REPL
- [x] Have Go CLI fetch CUs completed from Google Sheet w/ Sheets API
- [x] Make a Discord Bot (In Go obviously)
- [ ] Have Discord Bot be able to change my server nickname
- [ ] Be Able to Update Server Nickname By Updating Google Sheet

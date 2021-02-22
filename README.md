# TrelloTribbles
A tool to sort trello cards by members who have joined them as a replacement for actual voting.  This is meant to 
enhance a simple retro experience while still allowing the giphy power-up. Cards sorted from most to fewest members.

# Setup
The following environmental variables will need to be set before execution.
```dotenv
TRELLO_API_KEY=
TRELLO_API_TOKEN=
TRELLO_BOARD_ID=
```

### Getting the Trello API Key
Grab your developer API key from https://trello.com/app-key

### Getting the Trello API Token
Create a token for yourself using your API key and the following 
URL: https://trello.com/1/authorize?expiration=1day&name=MyPersonalToken&scope=read&response_type=token&key={YourAPIKey}  
After visiting this page and clicking the green Allow button, you'll be redirected to a page with your token. 

### Getting the Trello Board ID
Trello Board ID is the idea found in your board URL. For example: `https://trello.com/b/M0j3f9k0/MyBoard` the board id is **M0j3f9k0**


# Running
### Linux
```bash
TrelloTribbles
```
### MacOSX
```zsh
TrelloTribbles
```
### Windows
```
TrelloTribbles.exe
```
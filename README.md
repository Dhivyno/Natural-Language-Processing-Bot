# Natural-Language-Processing-Bot 🤖 (Inspired by ChatGPT)
This is a Natural Language Processing Bot made using Wit AI by Facebook and the Wolfram API implemented in Slack using GoLang. (Instructions for usage below)

---

### Instructions for use 📝

If you want to use this bot, you need to make a Slack account and get the Slack API tokens (Slack App Token and Slack OAuth Token). You also need a Wolfram Alpha developer account and create a new app to get the API key. You also need a Facebook account in order to access the Wit AI Server Access Token. Now, create a .env file in your text editor and put those in variables called SLACK_APP_TOKEN, SLACK_OAUTH_TOKEN, WOLFRAM_API_ID and WIT_AI_TOKEN respectively. Then run $ go mod tidy in your command prompt or terminal to import the necessary packages and run $ go run main.go to run the program and use your slack channel to ask the bot questions with the syntax "@[slack_app_name] question - [your_question]" and it will reply with "Received!" and the answer it received from the Wolfram API.

---

![image](https://user-images.githubusercontent.com/63943490/208598141-aa1cbf61-c921-4e00-884d-75f2236c7e5e.png)

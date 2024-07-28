## 🦌 **Odeer**
> **No harm or foul to [ollama guys](https://github.com/ollama/ollama), they did a fantastic job!**

### 🦌 **Download**
**You can download it here:** [Download link](https://github.com/kashifulhaque/odeer/releases/download/v0.1.3/odeer) and run `odeer talk`

### 🏃‍♂️ **Steps to run it**
> Well, I couldn't get the GitHub Actions thingy to work, which should've compiled it and built an executable
1. Register for a [Cloudflare](https://www.cloudflare.com/) account, if not already
2. Goto [Cloudflare Dashboard](https://dash.cloudflare.com/)
3. On the left menu, go to **AI** > **Workers AI** > **Use REST API** button
4. Click on the **Create a Workers AI API Token** button, and save that API key as your environment variable named `CLOUDFLARE_WORKERS_AI_API_KEY`
5. Copy your account ID from that page too, and save it as an environment variable named `CLOUDFLARE_ACCOUNT_ID`
6. Clone this repo, and either run the `run.sh` or `build.sh` file
```sh
git clone https://github.com/kashifulhaque/odeer;
cd odeer;
sh build.sh;
odeer talk;
```

### 🌐 **[WIP] An API server**
> You should be able to run your own API server soon™. It is work in progress as of yet!

`odeer start` starts a server on PORT 8080
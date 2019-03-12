# Line Chatbot Webhook Written in Go Language

## How to use

1. Install Golang
```sh
$ brew install golang
```

2. Install dependencies

* If go version <= 1.10
```sh
$ glide install
```
* If go version >= 1.11
```sh
$ go mod vendor
```

3. Build to binary
```sh
$ make build
```

4. Copy `env.sample` to `.env`, and fill environment variables
```sh
$ cp env.sample .env
```

5. Run binary
```sh
$ ./bin
```

## Click this link for add my line bot channel
<a href="https://line.me/R/ti/p/%40ylf0312k"><img height="36" border="0" alt="Tambah Teman" src="https://scdn.line-apps.com/n/line_add_friends/btn/en.png"></a>

See https://github.com/agungdwiprasetyo/chatbot-ai (repository text mining for processing input chat message)
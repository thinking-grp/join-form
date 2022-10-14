# join-form

When someone apply to join a group, you will know it via Discord WebHook

## 🚀 Usage

### 💨 Run this application

1, Download binary of your environment from [release](https://github.com/thinking-grp/join-form/releases)
2, Rename binary to joinform
3, Running this

```bash
./joinform
```

### 🦮 Main service

Post application/json to http://127.0.0.1:3040/join

- Example

```json
{
  "name": "Hidemaru",
  "twitter_id": "Hidemaru_OwO",
  "github_id": "HidemaruOwO",
  "introduction": "I am thinking member",
  "site": "https://v-sli.me"
}
```

## 🔨 How to build

Runnnig this

```bash
# Target binary of your environment
go build src/main.group

# Target all
bash build.sh
```

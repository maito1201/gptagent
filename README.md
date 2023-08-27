# GPTAgent

This is a simple program to execute linux command by asking chatGPT.


# Usage

## Set env variables

```sh
export OPENAI_API_KEY="your_key" # required
export GPT_MODEL="gpt-4" # default gpt-3.5-turbo
export OPENAI_API_ENDPOINT="https*//domain_name" # default https://api.openai.com/v1/chat/completions

```

## execute

```sh
go run . "display current time and describe in Japanese."
running... date '+%T %Z'
Command output: 22:59:27 JST

現在の時刻は22時59分27秒です。
```
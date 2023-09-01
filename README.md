# GPTAgent

This is a simple program to execute linux command by asking chatGPT.


# Usage

## Set env variables

```sh
export OPENAI_API_KEY="your_key" # required
export GPT_MODEL="gpt-4" # default gpt-3.5-turbo
export OPENAI_API_ENDPOINT="https://domain_name" # default https://api.openai.com/v1/chat/completions
```

## usage examples

```sh
gptagent "please check file of current directory with hidden file"
suggested command is "ls -a". are you sure to execute? (y/n/any opinion): y
Command output: .
..
.git
.gitignore
LICENSE
README.md
command
go.mod
go.sum
gpt
main.go

The current directory contains the following files, including hidden files:

- ".": Represents the current directory.
- "..": Represents the parent directory.
- ".git": Represents the Git repository directory.
- ".gitignore": Represents the Git ignore file.
- "LICENSE": Represents the license file.
- "README.md": Represents the readme file.
- "command": Represents a file named "command".
- "go.mod": Represents the Go module file.
- "go.sum": Represents the Go dependency checksum file.
- "gpt": Represents a file named "gpt".
- "main.go": Represents a Go source code file.
(input your answer): what kind of license is it
suggested command is "cat LICENSE". are you sure to execute? (y/n/any opinion): y
Command output: MIT License

Copyright (c) 2023 maito1201

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

The license for this project is the MIT License.
(input your answer): thanks
You're welcome! If you have any more questions, feel free to ask.
```

```
./gptagent "read ./main.go and describe it in brief"
suggested command is "cat main.go". are you sure to execute? (y/n/any opinion): y
Command output: package main...

The `main.go` file is part of a Go program that uses the GPT (Generative Pre-trained Transformer) model to create a conversational agent that can execute Linux commands. Here is a brief description of the code:

- The program imports necessary packages such as `bufio`, `encoding/json`, `fmt`, `os`, `gptagent/command`, and `gptagent/gpt`.
- There is an initialization function `init()` that sets up the GPT model by retrieving the API endpoint and API key from environment variables.
- The `prompt` variable stores the initial message that is displayed to the user.
- The main function reads the command-line arguments and retrieves the user input.
- The `messages` variable stores a sequence of messages exchanged between the user and the assistant. It includes the system prompt, user's questions, and assistant's responses, including commands to run.
- The `functions` variable defines the available functions that the assistant can execute, in this case, only one function `execute_linux_command`.
- The main function creates a `RequestData` object containing the model, messages, and functions. It also specifies the GPT model to use.
- The `execute` function is called with the `RequestData` object to make API calls to the GPT server and get the assistant's responses.
- The assistant's response is then processed to check if there is additional input or commands to execute.
- If there is more input or commands, the `execute` function is called recursively with updated `RequestData`.

Overall, the code implements a conversational agent that interacts with the user, executes Linux commands as requested, and provides responses based on the GPT model's output.
(input your answer): thanks
You're welcome! If you have any more questions, feel free to ask.
```

```sh
gptagent
Hi there! How can I assist you today?
(input your answer): please tell me the weather of tokyo
suggested command is "curl wttr.in/Tokyo". are you sure to execute? (y/n/any opinion): y
Command output: Weather report: Tokyo

      \   /     Clear
       .-.      +28(32) °C
    ― (   ) ―   ↑ 13 km/h
       `-’      10 km
      /   \     0.0 mm
                                                       ┌─────────────┐
┌──────────────────────────────┬───────────────────────┤  Tue 29 Aug ├───────────────────────┬──────────────────────────────┐
│            Morning           │             Noon      └──────┬──────┘     Evening           │             Night            │
├──────────────────────────────┼──────────────────────────────┼──────────────────────────────┼──────────────────────────────┤
│    \  /       Partly cloudy  │     \   /     Sunny          │     \   /     Sunny          │     \   /     Clear          │
│  _ /"".-.     +28(31) °C     │      .-.      +28(32) °C     │      .-.      +28(32) °C     │      .-.      +28(32) °C     │
│    \_(   ).   ↗ 16-22 km/h   │   ― (   ) ―   ↗ 14-20 km/h   │   ― (   ) ―   ↑ 15-22 km/h   │   ― (   ) ―   ↑ 13-18 km/h   │
│    /(___(__)  10 km          │      `-’      10 km          │      `-’      10 km          │      `-’      10 km          │
│               0.0 mm | 0%    │     /   \     0.0 mm | 0%    │     /   \     0.0 mm | 0%    │     /   \     0.0 mm | 0%    │
└──────────────────────────────┴──────────────────────────────┴──────────────────────────────┴──────────────────────────────┘
                                                       ┌─────────────┐
┌──────────────────────────────┬───────────────────────┤  Wed 30 Aug ├───────────────────────┬──────────────────────────────┐
│            Morning           │             Noon      └──────┬──────┘     Evening           │             Night            │
├──────────────────────────────┼──────────────────────────────┼──────────────────────────────┼──────────────────────────────┤
│  _`/"".-.     Patchy rain po…│  _`/"".-.     Patchy rain po…│     \   /     Sunny          │     \   /     Clear          │
│   ,\_(   ).   +28(32) °C     │   ,\_(   ).   +28(33) °C     │      .-.      +28(33) °C     │      .-.      +28(32) °C     │
│    /(___(__)  ↑ 14-21 km/h   │    /(___(__)  ↑ 14-20 km/h   │   ― (   ) ―   ↑ 15-21 km/h   │   ― (   ) ―   ↑ 17-24 km/h   │
│      ‘ ‘ ‘ ‘  9 km           │      ‘ ‘ ‘ ‘  9 km           │      `-’      10 km          │      `-’      10 km          │
│     ‘ ‘ ‘ ‘   1.6 mm | 85%   │     ‘ ‘ ‘ ‘   0.9 mm | 83%   │     /   \     0.0 mm | 0%    │     /   \     0.0 mm | 0%    │
└──────────────────────────────┴──────────────────────────────┴──────────────────────────────┴──────────────────────────────┘
                                                       ┌─────────────┐
┌──────────────────────────────┬───────────────────────┤  Thu 31 Aug ├───────────────────────┬──────────────────────────────┐
│            Morning           │             Noon      └──────┬──────┘     Evening           │             Night            │
├──────────────────────────────┼──────────────────────────────┼──────────────────────────────┼──────────────────────────────┤
│    \  /       Partly cloudy  │    \  /       Partly cloudy  │     \   /     Sunny          │     \   /     Clear          │
│  _ /"".-.     +27(31) °C     │  _ /"".-.     +28(31) °C     │      .-.      +28(31) °C     │      .-.      +27(31) °C     │
│    \_(   ).   ↖ 12-17 km/h   │    \_(   ).   ↖ 15-20 km/h   │   ― (   ) ―   ↖ 14-20 km/h   │   ― (   ) ―   ↖ 14-19 km/h   │
│    /(___(__)  10 km          │    /(___(__)  10 km          │      `-’      10 km          │      `-’      10 km          │
│               0.0 mm | 0%    │               0.0 mm | 0%    │     /   \     0.0 mm | 0%    │     /   \     0.0 mm | 0%    │
└──────────────────────────────┴──────────────────────────────┴──────────────────────────────┴──────────────────────────────┘
Location: 東京都, 日本 [34.2255804,139.294774527387]

Follow @igor_chubin for wttr.in updates

The current weather in Tokyo is as follows:

- Clear sky
- Temperature: 28°C (feels like 32°C)
- Wind: 13 km/h
- Visibility: 10 km
- No precipitation

Please note that weather conditions may change, so it's always good to check the latest updates.
```

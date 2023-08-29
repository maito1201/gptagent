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
go run . "please check file of current directory with hidden file"
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
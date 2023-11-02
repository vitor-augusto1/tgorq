<h1 align="center">
  Terminal GO ReQuest ( TGORQ)
</h1>

<p align="center">
<b>A Vim-like lightweight CLI tool for performing HTTP requests</b>
</p>

# Features
![image](https://github.com/vitor-augusto1/tgorq/assets/121441594/c290a816-6a2d-4f6d-9066-bcbfdfe9f8de)
![tgorq](https://github.com/vitor-augusto1/tgorq/assets/121441594/876512df-03bc-46fc-87f7-6e3a77a36a39)

* Move with keyboard short-cuts
* Lightweight
* Fast
* Save State option
* Save response to a file

# About TGORQ
Tgorq is a vim-like lightweight TUI (text-based user interface) CLI that allows you to perform fast HTTP requests.
It was built with the [go programming language](https://go.dev) and the [bubble tea framework](https://github.com/charmbracelet/bubbletea).

# Installation
Tgorq requires [go](https://go.dev) to run successfully. To install, you can run the command below or download the binary from the [release page](https://github.com/vitor-augusto1/tgorq/releases/).
```
go install github.com/vitor-augusto1/tgorq@latest
```

# Usage
```
tgorq -h
```
This will display the help manual for the tool. Here are all the flags it supports.
```


__/\\\\\\\\\\\\\\\________________________________________________________
 _\///////\\\/////_________________________________________________________
  _______\/\\\_________/\\\\\\\\________________________________/\\\\\\\\___
   _______\/\\\________/\\\////\\\_____/\\\\\_____/\\/\\\\\\\___/\\\////\\\__
    _______\/\\\_______\//\\\\\\\\\___/\\\///\\\__\/\\\/////\\\_\//\\\\\\\\\__
     _______\/\\\________\///////\\\__/\\\__\//\\\_\/\\\___\///___\///////\\\__
      _______\/\\\________/\\_____\\\_\//\\\__/\\\__\/\\\________________\/\\\__
       _______\/\\\_______\//\\\\\\\\___\///\\\\\/___\/\\\________________\/\\\\_
        _______\///_________\////////______\/////_____\///_________________\////__

    A vim-like TUI (Text User Interface) that allows you to make http requests.
    Example: ./tgorq [ -o | --enable-output ] [ -s | --save-state ]

Usage:
  tgorq [flags]

Flags:
  -o, --enable-output   Stores the response body and headers in the response directory.
  -h, --help            help for tgorq
  -s, --save-state      Save the current application state.
```

# License

*tgorq* is licensed under the MIT license. Take a look at the [MIT License](https://github.com/vitor-augusto1/tgorq/blob/main/LICENSE).

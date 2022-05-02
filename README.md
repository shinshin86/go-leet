# go-leet
[![CI](https://github.com/shinshin86/go-leet/actions/workflows/ci.yml/badge.svg)](https://github.com/shinshin86/go-leet/actions/workflows/ci.yml)

Implements leetspeak converter at Go.  
This is created to learn Go.

## Usage

```sh
# Install
go install github.com/shinshin86/go-leet@latest

go-leet hello
# => ]~[E|_|_<>
```

## About string for leet conversion
As for the string for leet conversion, I using the string described in this page.

[Wikipedia(Ja) - Leet](https://ja.wikipedia.org/wiki/Leet)

## Licence
[MIT](https://github.com/shinshin86/go-leet/blob/main/LICENSE) & [CC-BY-SA 3.0](https://creativecommons.org/licenses/by-sa/3.0/legalcode)

In the getLeetList function, character combinations taken from Wikipedia are used. For this reason, only this function is CC-BY-SA 3.0 Licence.

## Author
[Yuki Shindo](https://shinshin86.com/en)
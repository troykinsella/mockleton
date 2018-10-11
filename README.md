# mockleton

[![Version](https://badge.fury.io/gh/troykinsella%2Fmockleton.svg)](https://badge.fury.io/gh/troykinsella%2Fmockleton)
[![License](https://img.shields.io/github/license/troykinsella/mockleton.svg)](https://github.com/troykinsella/mockleton/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/troykinsella/mockleton.svg?branch=master)](https://travis-ci.org/troykinsella/mockleton)

An executable program mocking tool

## Installation

Checkout [releases](https://github.com/troykinsella/mockleton/releases) and download the appropriate binary for your system.
Put the binary in a convenient place, such as `/usr/local/bin/mockleton`.

Or, run these commands to download and install:
```bash
VERSION=0.1.0
OS=darwin # or linux
curl -SL -o /usr/local/bin/mockleton https://github.com/troykinsella/mockleton/releases/download/v${VERSION}/mockleton_${OS}_amd64
chmod +x /usr/local/bin/mockleton
```

## Usage

### Mocking an Executable

Firstly, you have to make `mockleton` look like the executable you wish to mock:

```bash
$ cd my-project
$ mkdir -p mocks
$ export PATH=$PWD/mocks:$PATH
$ ln -s /usr/local/bin/mockleton $PWD/mocks/program
```

Now, when you run `program` (which points to `mockleton`), it will capture the 
execution details and produce a report file:

```bash
$ echo foo | program bar
$ cat mockleton.out
{
  "mockleton-version": "...",
  "sequence": [
    {
      "exec-spec": {
        "timestamp": "2018-10-10T14:09:31.477900029-07:00",
        "stdin": {
          "content": "foo\n",
          "encoding": "utf-8"
        },
        "args": [
          "program",
          "bar"
        ],
        "env": {
          "EDITOR": "emacs",
          ...
        }
      }
    }
  ]
}
```

If you run `program` again, it will append the new execution details to the `sequence` 
list in the existing `mockleton.out` file.

## Road Map

* stubbed outputs
* basic assertions using gomega

## License

MIT © Troy Kinsella

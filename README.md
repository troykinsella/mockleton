# mockleton

[![Version](https://badge.fury.io/gh/troykinsella%2Fmockleton.svg)](https://badge.fury.io/gh/troykinsella%2Fmockleton)
[![License](https://img.shields.io/github/license/troykinsella/mockleton.svg)](https://github.com/troykinsella/mockleton/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/troykinsella/mockleton.svg?branch=master)](https://travis-ci.org/troykinsella/mockleton)

An executable program mocking tool

## Installation

Checkout [releases](https://github.com/troykinsella/mockleton/releases) and download the appropriate binary for your system.
Put the binary in a convenient place, such as `/usr/local/bin/mockleton`.

Or, run the handy dandy install script:
(Note: go read the script and understand what you're running before trusting it)
```bash
export PREFIX=~ # install into ~/bin
wget -q -O - https://raw.githubusercontent.com/troykinsella/mockleton/master/install.sh | bash
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

This is the basis of how `mockleton` mocks executables. Your scripts or client programs
can continue to call `program` as they normally would without modifying them to understand
a testing context.


## Road Map

* stubbed outputs
* basic assertions using gomega

## License

MIT © Troy Kinsella

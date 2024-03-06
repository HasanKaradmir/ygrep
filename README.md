# `ygrep`: Advanced YAML Search CLI Tool

`ygrep` is a robust command-line interface (CLI) tool, crafted in Go, specifically designed to perform sophisticated searches within YAML files. It stands out in parsing `kubectl` outputs as well as general YAML data, offering functionality akin to the traditional grep but enhanced to suit the intricacies of YAML structures.

In this repository, you'll find comprehensive details about `ygrep`, encompassing everything from the installation process to various usage scenarios. Whether you're dealing with complex `kubectl` outputs or navigating through diverse YAML files, `ygrep` provides the tools you need to search effectively and efficiently.

## Table of Contents
- [What is `ygrep`?](#what-is-ygrep)
- [Examples](#examples)
- [Functionality](#functionality)
- [Enhanced Usage](#enhanced-usage)
- [Installation](#installation)
- [License](#license)

## What is `ygrep`?

**ygrep** is useful for searching within a YAML file.

Here we have a sample YAML file:

<img src="img/ygrep-sample-file.gif" width="640px">

<br>

Here's a directly **key-based search** **`ygrep`** demo:
<br>

<img src="img/ygrep-key-based.gif" width="640px">

<br>

Here's a directly **value-based search** **`ygrep`** demo:
<br>

<img src="img/ygrep-value-based.gif" width="640px">

<br>

... and here's a **`ygrep`** demo received from the pipeline:

<img src="img/ygrep-kubectl.gif" width="640px">

### Examples
```sh
SYNOPSIS

    ygrep PATTERNS [FILE...] 
    STDIN  |  ygrep PATTERNS

# you can directly key based search `command` word in commands.yaml file by default
$ ygrep "search-pattern" sample.yaml

# you can directly value based search `command` word in commands.yaml file
$ ygrep -v "search-pattern" sample.yaml 

# you can also get it from a command that outputs stdout
$ cat commands.yaml | ygrep "search-pattern"

# this output may be cat or kubectl
$ kubectl get pods sample_pod -o yaml | ygrep "search-pattern"

```

-----

## Functionality

`ygrep` operates by searching within YAML files. It initiates a search based on a YAML key you provide, similar to the grep command. It looks for the specified word within the key (without case sensitivity; mere inclusion of the word suffices) and returns the key and its value. This process is conducted throughout the entire file.

-----

## Enhanced Usage

You have the option to input data through stdin (standard input) and combine it with a pipe (`|`). Initially designed to parse kubectl commands, `ygrep` has since evolved with more features and capabilities.

-----

## Installation

Currently, ygrep can be installed using the go install command. Future updates will introduce package-based installation methods. For now, having Golang installed is a prerequisite.

#### Quick Start

Install ygrep on your system and start using it immediately with the following command:

```bash
$ go install github.com/HasanKaradmir/ygrep@latest
```
-----

## License
`ygrep` is released under the [Apache-2.0 License](LICENSE).
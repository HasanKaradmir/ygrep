# `kgrep`: Advanced YAML Search CLI Tool

`kgrep` is a robust command-line interface (CLI) tool, crafted in Go, specifically designed to perform sophisticated searches within YAML files. It stands out in parsing `kubectl` outputs as well as general YAML data, offering functionality akin to the traditional grep but enhanced to suit the intricacies of YAML structures.

In this repository, you'll find comprehensive details about `kgrep`, encompassing everything from the installation process to various usage scenarios. Whether you're dealing with complex `kubectl` outputs or navigating through diverse YAML files, `kgrep` provides the tools you need to search effectively and efficiently.

## Table of Contents
- [What is `kgrep`?](#what-is-kgrep)
- [Examples](#examples)
- [Functionality](#functionality)
- [Enhanced Usage](#enhanced-usage)
- [Installation](#installation)
- [License](#license)

## What is `kgrep`?

**kgrep** is useful for searching within a YAML file.

Here we have a sample YAML file:
![kgrep demo GIF sample yaml file](img/kgrep-sample-yaml.svg)

Here's a **directly** **`kgrep`** demo:
![kgrep demo GIF directly](img/kgrep-directly.gif)

... and here's a **`kgrep`** demo received from the pipeline:
![kgrep demo GIF stdin](img/kgrep-stdin.gif)

### Examples
```sh
SYNOPSIS

    kgrep [FILE...] PATTERNS
    STDIN  |  kgrep PATTERNS

# you can directly search `command` word in commands.yaml file
$ kgrep sample.yaml "search-pattern"

# you can also get it from a command that outputs stdout
$ cat commands.yaml | kgrep "search-pattern"

# this output may be cat or kubectl
$ kubectl get pods sample_pod -o yaml | kgrep "search-pattern"

```

-----

## Functionality

`kgrep` operates by searching within YAML files. It initiates a search based on a YAML key you provide, similar to the grep command. It looks for the specified word within the key (without case sensitivity; mere inclusion of the word suffices) and returns the key and its value. This process is conducted throughout the entire file.

-----

## Enhanced Usage

You have the option to input data through stdin (standard input) and combine it with a pipe (`|`). Initially designed to parse kubectl commands, `kgrep` has since evolved with more features and capabilities.

-----

## Installation

Currently, kgrep can be installed using the go install command. Future updates will introduce package-based installation methods. For now, having Golang installed is a prerequisite.

#### Quick Start

Install kgrep on your system and start using it immediately with the following command:

```bash
$ go install github.com/HasanKaradmir/kgrep@latest
```
-----

## License
`kgrep` is released under the [Apache-2.0 License](LICENSE).
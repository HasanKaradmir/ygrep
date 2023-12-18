# `kgrep`: Grep for YAML

![GitHub stars](https://img.shields.io/github/stars/HasanKaradmir/kgrep.svg?label=github%20stars)

This repository provides kgrep tool.

## What is `kgrep`?

**kgrep** is useful for searching within a YAML file.

Here's a **directly** **`kgrep`** demo:
![kgrep demo GIF directly](img/kgrep-directly.gif)

... and here's a **`kgrep`** demo received from the pipeline:
![kgrep demo GIF stdin](img/kgrep-stdin.gif)

### Examples
```sh
SYNOPSIS

    kgrep [FILE...] PATTERNS
    STDIN | kgrep PATTERNS



$ kgrep commands.yaml command


$ cat commands.yaml | kgrep command

```


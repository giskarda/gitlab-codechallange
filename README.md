## topN

topn.{py,go} is a tool that allows you to sort big files.

At the moment supports only files (unlimited size) composed by numbers one by line.


### Go Version:
To build the go application: ``go build topN.go``, it will create a topN binary in the same folder you run that command.

### Usage:

Both version have a ``-h / --help`` helper.

The python version looks like:

```
[~/src/github.com/giskarda/gitlab-codechallange] ➔  python topN.py  --help
usage: topN.py [-h] --filename FILENAME --top-numbers TOP_NUMBERS

topN: find the top N largest number from a file

optional arguments:
  -h, --help            show this help message and exit
  --filename FILENAME   file to search for top numbers
  --top-numbers TOP_NUMBERS
                        number of top numbers to be retrieved
[~/src/github.com/giskarda/gitlab-codechallange] ➔
```

If you want to understand pro / cons of this approach please read answer.txt
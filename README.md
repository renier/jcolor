# jcolor

Quick and dirty CLI command to colorize structured json logs coming in from standard input.
It will not break if the the line is not parseable JSON, and simply sends it to standard output
as-is.

#### To install

```
$ go install github.com/renier/jcolor
```

#### To use

```
$ some_command 2>&1 | jcolor
```


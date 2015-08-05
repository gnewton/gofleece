# gofleece
Command line json validator

```
Usage of ./gofleece:
  -e=false: Stop if error encountered
  -s=false: Do not produce any output
  -v=false: Verbosity of logging
  -z=false: Acccept a file of zero length as being valid JSON
  ```

## Features
* One or more JSON files on command line can be processed
* If no file arguments, reads from stdin
* Raw json, gzipped and bzip2ed files transparently processed
* Default: only complains about failures
* Default: does not stop when a json file fails
* Handles very large json files with minumal overhead
* Tested with 2GB JSON file, 9MB virtual, 3.6MB resident.
* Properly sets shell error code

## Limitations/dependencies
`gofleece` is basically a wrapper around `encoding/json`, so its interpretation of what is json is what [`encoding/json`](https://golang.org/pkg/encoding/json/) believes it to be.

## TODO
* Presently single threaded: goroutine per file to be validated (limit to max cores; make # routines configuratble)

## Example
```
$ /usr/bin/time -f "%E %M" ./gofleece -v *.json
2015/07/09 23:05:22 util.go:30: Opening: repository_metadata_2014-06-06_150.json
2015/07/09 23:05:22 util.go:49:  File size: 2080284113
2015/07/09 23:05:22 main.go:92: Map
2015/07/09 23:06:47 main.go:55: Valid JSON
2015/07/09 23:06:47 main.go:59: ------------------
2015/07/09 23:06:47 util.go:30: Opening: repository_metadata_2014-06-09_237.json
2015/07/09 23:06:47 util.go:49:  File size: 204
2015/07/09 23:06:47 main.go:92: Map
2015/07/09 23:06:47 main.go:55: Valid JSON
2015/07/09 23:06:47 main.go:59: ------------------
2015/07/09 23:06:47 util.go:30: Opening: repository_metadata_2014-06-10_357.json
2015/07/09 23:06:47 util.go:49:  File size: 506957
2015/07/09 23:06:47 main.go:92: Map
2015/07/09 23:06:47 main.go:55: Valid JSON
2015/07/09 23:06:47 main.go:59: ------------------
2015/07/09 23:06:47 util.go:30: Opening: smallest_array.json
2015/07/09 23:06:47 util.go:49:  File size: 3
2015/07/09 23:06:47 main.go:87: Array
2015/07/09 23:06:47 main.go:55: Valid JSON
2015/07/09 23:06:47 main.go:59: ------------------
2015/07/09 23:06:47 util.go:30: Opening: smallest_map.json
2015/07/09 23:06:47 util.go:49:  File size: 3
2015/07/09 23:06:47 main.go:92: Map
2015/07/09 23:06:47 main.go:55: Valid JSON
2015/07/09 23:06:47 main.go:59: ------------------
1:24.24 3524
$ 
```

## Example 2
```
$ ./gofleece -s *.json 
$ echo $?
0
$ ./gofleece -s main.go
$ echo $?
1
$
```


# License
GPV v2

Copright (C) 2015 Glen Newton

See LICENSE
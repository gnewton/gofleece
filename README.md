# gofleece
Command line json validator

```
Usage of ./gofleece:
  -s=false: Stop if error encountered
  -v=false: Verbosity of logging
  ```

## Features
* One or more JSON files on command line can be processed
* If no file arguments, reads from stdin
* Raw json, gzipped and bzip2ed files transparently processed
* Default: only complains about failures
* Default: does not stop when a json file fails
* Handles very large json files with minumal overhead
* Tested with 2GB JSON file, 11GB resident, 13GB virtual

## TODO
* Presently single threaded: goroutine per file to be validated (limit to max cores; make # routines configuratble)

## Example
```
$ ./gofleece -v *json*
2015/07/09 17:10:11 Opening: repository_metadata_2014-06-06_150.json.bz2
2015/07/09 17:10:11  File size: 454545326
2015/07/09 17:10:11 Map
2015/07/09 17:23:22 Valid JSON
2015/07/09 17:23:22 ------------------
2015/07/09 17:23:22 Opening: repository_metadata_2014-06-09_237.json
2015/07/09 17:23:22  File size: 204
2015/07/09 17:23:22 Map
2015/07/09 17:23:22 Valid JSON
2015/07/09 17:23:22 ------------------
2015/07/09 17:23:22 Opening: repository_metadata_2014-06-10_357.json
2015/07/09 17:23:22  File size: 506957
2015/07/09 17:23:22 Map
2015/07/09 17:23:22 Valid JSON
2015/07/09 17:23:22 ------------------
```
.

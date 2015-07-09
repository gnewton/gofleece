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
* Default only complains about failures
* Default does not stop when a json file fails

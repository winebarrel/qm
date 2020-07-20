# qm

[pt-fingerprint](https://www.percona.com/doc/percona-toolkit/LATEST/pt-fingerprint.html) for JSON Lines.

## Usage

```
Usage of qm:
  -f string
    	fingerprint key (default "fingerprint")
  -q string
    	query key (default "query")
  -version
    	print version and exit
```

```
$ cat data.jsonl
{"query":"select now()","time":6893585}
{"query":"select 1","time":121333}
{"query":"select 2","time":91779}

$ qm data.jsonl # or `cat data.jsonl | qm`
{"query":"select now()","time":6893585,"fingerprint":"select now()"}
{"query":"select 1","time":121333,"fingerprint":"select ?"}
{"query":"select 2","time":91779,"fingerprint":"select ?"}
```

# qm

[pt-fingerprint](https://www.percona.com/doc/percona-toolkit/LATEST/pt-fingerprint.html) for JSON Lines.

## Usage

```
Usage of ./qm:
  -f string
    	fingerprint key (default "fingerprint")
  -q string
    	query key (default "query")
  -sha1
    	append SHA1
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

$ qm -sha1 data.jsonl
{"fingerprint_sha1":"e226cc45856c33b443aeca4c17e37a61d761a3e1","query":"select now()","time":6893585,"fingerprint":"select now()","query_sha1":"e226cc45856c33b443aeca4c17e37a61d761a3e1"}
{"query":"select 1","time":121333,"fingerprint":"select ?","query_sha1":"3232003928f9fe86a9cb634f450d5a53a4025819","fingerprint_sha1":"7ae509fc5e11f3bdd89c7e1a5829d6e86fbd8943"}
{"query":"select 2","time":91779,"fingerprint":"select ?","query_sha1":"63a2df45794572251144e77d007ff287b190732e","fingerprint_sha1":"7ae509fc5e11f3bdd89c7e1a5829d6e86fbd8943"}
```

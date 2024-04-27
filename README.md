# tcl

tcl(tcp check list) is checking tcp health check tool

```bash
Usage of tcl
  -f string
        target json file
  -v    show version

```

## how to

```
$ tcl -f="test.json"

+--------+------------------------------+--------+
|  NAME  |             URL              | STATUS |
+--------+------------------------------+--------+
| yahoo  | https://www.yahoo.co.jp/     |    200 |
| google | https://www.google.co.jp/    |    200 |
+--------+------------------------------+--------+
```

## License

MIT

# Time diff

Simple tool get time diffs. Can only diff against current time for now.

Examples:

```shell
td "2021-08-05 22:00:01"
# 6 days 15 hours 19 minutes 52 seconds

td 2019-12-31
# 1 year 7 months 12 days

td -h "2021-08-05 22:00:01"
# 159.34 hours
```

Usage:

```
Usage: td [flags] [argument]
  -d	diff in days
  -h	diff in hours
  -m	diff in minutes
  -s	diff in seconds
```

## Install / Update

### MacOS

	curl -L https://github.com/fredrik01/td/releases/latest/download/td_darwin_amd64.tar.gz -o td.tar.gz
	mkdir /tmp/td
	tar -xvf td.tar.gz -C /tmp/td
	mv /tmp/td/td /usr/local/bin/td
	rm td.tar.gz
	rm -r /tmp/td

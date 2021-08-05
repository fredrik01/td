# Time diff

Simple tool get time diffs. Can only diff against current time for now.

```shell
td "2021-08-05 22:00:01"
# 1h34m6s

td -h "2021-08-05 22:00:01
# 1.57h
```

## Installation

### MacOS

	curl -L https://github.com/fredrik01/td/releases/latest/download/td_darwin_amd64.tar.gz -o td.tar.gz
	mkdir /tmp/td
	tar -xvf td.tar.gz -C /tmp/td
	mv /tmp/td/td /usr/local/bin/td
	rm td.tar.gz
	rm -r /tmp/td

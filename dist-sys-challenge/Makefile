.PHONY: download-maelstrom
download-maelstrom:
	wget https://github.com/jepsen-io/maelstrom/releases/download/v0.2.3/maelstrom.tar.bz2
	tar -xzf maelstrom.tar.bz2
	rm maelstrom.tar.bz2

.PHONY: 01-echo
01-echo:
	go build -o bin/01-echo 01-echo/main.go
	maelstrom/maelstrom test -w echo --bin bin/01-echo --node-count 1 --time-limit 10

.PHONY: 02-unique-ids
02-unique-ids:
	go build -o bin/02-unique-ids 02-unique-ids/main.go
	maelstrom/maelstrom test -w unique-ids --bin bin/02-unique-ids --node-count 1 --time-limit 10

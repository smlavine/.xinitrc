# Copyright (c) 2022 Sebastian LaVine <mail@smlavine.com>
# SPDX-License-Identifier: CC0-1.0

.POSIX:

PREFIX = $(HOME)

all: .xinitrc

.xinitrc: main.go
	go build

check:
	reuse lint
	go test -v ./...

clean:
	rm .xinitrc

coverage:
	go test -v -coverprofile=cover.out ./...
	go tool cover -html=cover.out


install: .xinitrc
	cp -i .xinitrc $(PREFIX)/.xinitrc
	@# cp -i doesn't preserve x bit?
	chmod +x $(PREFIX)/.xinitrc

.PHONY: all check clean install

// Copyright (c) 2022 Sebastian LaVine <mail@smlavine.com>
// SPDX-License-Identifier: MIT

package main

import "log"

func main() {
	notifications := notifDaemon()
	background := backgroundSetter()
	wm := windowManager()

	if err := notifications.Start(); err != nil {
		log.Println(err)
	}

	if err := background.Run(); err != nil {
		log.Println(err)
	}

	if err := wm.Run(); err != nil {
		log.Fatalln(err)
	}
}

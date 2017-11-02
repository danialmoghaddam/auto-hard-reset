package main

import (
	"fmt"

	"bitbucket.org/kisom/gopush/pushover"
)

//Notify using Pushover notification
func Notify(r *Rig) {
	configFile := ReadConfig()

	if configFile.Pushover == true {
		identity := pushover.Authenticate(
			configFile.PushoverToken,
			configFile.PushoverUser,
		)

		message := fmt.Sprint("Force rebooting ", r.name)
		sent := pushover.Notify(identity, message)
		if !sent {
			log.Error("[!] notification failed.")
		} else {
			log.Notice("Pusherover notification sent")
		}
	}
}

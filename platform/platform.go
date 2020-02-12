package platform

import (
	"log"
)

type Platform int

const (
	Platform_STEAM Platform = iota
	Platform_AGENT
)

func (p Platform) String() string {
	return [...]string{"STEAM", "AGENT"}[p]
}

func PlatformFromString(p string) Platform {

	//
	switch p {
	case Platform_STEAM.String():
		return Platform_STEAM
	case Platform_AGENT.String():
		return Platform_AGENT
	}

	log.Printf("PlatformFromString(): Could not parse string: %s", p)
	return Platform_STEAM
}

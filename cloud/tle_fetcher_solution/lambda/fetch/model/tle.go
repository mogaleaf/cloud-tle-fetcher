package model

import "git.darknebu.la/Satellite/tle"

type Tle struct {
	Text string
	*tle.TLE
}

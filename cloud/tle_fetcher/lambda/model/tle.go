package model

import "git.darknebu.la/Satellite/tle"

type Tle struct {
	Lines []string
	*tle.TLE
}

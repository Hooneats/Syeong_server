package flag

import "flag"

var Flags map[string]*string

func FlagsLoad(fs []*FlagCategory) {
	Flags = make(map[string]*string)
	for _, ca := range fs {
		Flags[ca.Name] = ca.Load()
	}
	flag.Parse()
}

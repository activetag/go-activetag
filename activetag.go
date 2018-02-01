package activetag

import "regexp"

var ActiveTagRegexp = regexp.MustCompile(`^([a-z0-9]+)-([a-z0-9]+(?:-[a-z0-9]+)*)$`)

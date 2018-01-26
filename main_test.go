package main

import (
	"testing"
	"time"
)

func TestGetTimeInLocation(t *testing.T) {
	loadLoc := func(locstr string) *time.Location {
		loc, err := time.LoadLocation(locstr)
		if err != nil {
			t.Fatal(err.Error())
		}
		return loc
	}

	tt := []struct {
		desc string
		in   time.Time
		loc  *time.Location
	}{
		{
			desc: "now",
			in:   time.Now(),
			loc:  loadLoc("America/New_York"),
		},
		{
			desc: "UNIX epoch",
			in:   time.Date(1970, time.January, 1, 0, 0, 0, 0, loadLoc("UTC")),
			loc:  loadLoc("Europe/Paris"),
		},
		{
			desc: "IN THE YEAR TWO THOUUUUU-SAND",
			in:   time.Date(1999, time.December, 31, 23, 59, 59, 999999999, loadLoc("America/Los_Angeles")),
			loc:  loadLoc("Asia/Tokyo"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			out, _ := getTimeInLocation(tc.in, tc.loc.String())
			if !out.Equal(tc.in) {
				t.Errorf("times should be equal, but are not: in: %v, out: %v", tc.in, out)
			}
			if out.Location().String() != tc.loc.String() {
				t.Errorf("out time location and requested location should be equal, but are not: requested: %v, out: %v",
					tc.loc, out.Location())
			}
		})
	}
}

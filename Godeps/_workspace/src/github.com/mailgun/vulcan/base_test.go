/*
Declares gocheck's test suites
*/
package vulcan

import (
	. "github.com/mailgun/vulcand/Godeps/_workspace/src/launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

//This is a simple suite to use if tests dont' need anything
//special
type MainSuite struct {
}

func (s *MainSuite) SetUpTest(c *C) {
}

var _ = Suite(&MainSuite{})

package dateutils

import (
	. "launchpad.net/gocheck"
	"testing"
	"time"
)

type DateUtilSuite struct{}

func Test(t *testing.T)                    { TestingT(t) }
func (t *DateUtilSuite) TearDownTest(c *C) {}
func (t *DateUtilSuite) SetupTest(c *C)    {}

var _ = Suite(&DateUtilSuite{})

func (t *DateUtilSuite) TestDateUtilBasic(c *C) {

	// Setup
	startDate, _ := time.Parse(dateParser, "01/01/1901")
	endDate, _ := time.Parse(dateParser, "31/12/2000")

	// Act
	count, err := countDaysOnFirst(startDate, endDate, time.Sunday)

	// Verify
	c.Check(count, Equals, 171)
	c.Check(err, IsNil)
}

func (t *DateUtilSuite) TestDateUtilRangeErr(c *C) {

	// Setup
	endDate, _ := time.Parse(dateParser, "01/01/1901")
	startDate, _ := time.Parse(dateParser, "31/12/2000")

	// Act
	_, err := countDaysOnFirst(startDate, endDate, time.Sunday)

	// Verify
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "Provided endDate is prior to the startDate")
}

func (t *DateUtilSuite) TestDateUtilNotTheFirstErr(c *C) {

	// Setup
	startDate, _ := time.Parse(dateParser, "02/01/1901")
	endDate, _ := time.Parse(dateParser, "31/12/2000")

	// Act
	_, err := countDaysOnFirst(startDate, endDate, time.Sunday)

	// Verify
	c.Check(err, NotNil)
}

func (t *DateUtilSuite) TestDaysInCentury(c *C) {

	// Act
	count, err := CountDaysOnFirstInCentury(20, time.Sunday)

	// Verify
	c.Check(err, IsNil)
	c.Check(count, Equals, 171)
}

func (t *DateUtilSuite) TestSundaysIn20thCentury(c *C) {

	// Act
	count, err := CountSundaysOnFirstIn20thCentury()

	// Verify
	c.Check(err, IsNil)
	c.Check(count, Equals, 171)
}

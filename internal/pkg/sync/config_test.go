/*
 *
 */

package sync

import (
	"testing"

	"github.com/xelalexv/dregsy/test"
)

//
func TestSyncConfig(t *testing.T) {

	th := test.NewTestHelper(t)

	c, e := LoadConfig(th.GetFixture("skopeo.yaml"))
	th.AssertNoError(e)
	th.AssertNotNil(c)
	th.AssertEqual("skopeo", c.Relay)

	c, e = LoadConfig(th.GetFixture("docker.yaml"))
	th.AssertNoError(e)
	th.AssertNotNil(c)
	th.AssertEqual("docker", c.Relay)
}

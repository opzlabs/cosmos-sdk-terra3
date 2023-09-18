//go:build norace
// +build norace

package network_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/testutil/network"
)

type IntegrationTestSuite struct {
	suite.Suite

	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	var err error
	s.network, err = network.New(s.T(), s.T().TempDir(), network.DefaultConfig())
	s.Require().NoError(err)

	h, err := s.network.WaitForHeight(1)
	s.Require().NoError(err, "stalled at height %d", h)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestNetwork_Liveness() {
	h, err := s.network.WaitForHeightWithTimeout(10, time.Minute)
	s.Require().NoError(err, "expected to reach 10 blocks; got %d", h)
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

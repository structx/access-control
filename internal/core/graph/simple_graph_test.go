package graph_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/structx/access-control/internal/core/graph"
)

type SimpleGraphSuite struct {
	suite.Suite
	graph *graph.SimpleGraph
}

func (suite *SimpleGraphSuite) SetupTest() {
	suite.graph = graph.New()
}

func (suite *SimpleGraphSuite) TestAddVertex() {}

func (suite *SimpleGraphSuite) TestAddEdge() {}

func (suite *SimpleGraphSuite) TestGetVertex() {}

func TestSimpleGraphSuite(t *testing.T) {
	suite.Run(t, new(SimpleGraphSuite))
}

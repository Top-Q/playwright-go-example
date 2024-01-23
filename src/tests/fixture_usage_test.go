package tests

import (
	"fmt"

	"github.com/stretchr/testify/require"
)

func (s *Suite) TestSetupAndTeardownWithFailure() {
	require := require.New(s.T())
	fmt.Println("In the test that suppose to fail")
	require.Nil("String that is not nil")
	fmt.Println("*** After assertion. Should not be printed ***")
}

func (s *Suite) TestSetupAndTeardownWithSuccess() {
	require := require.New(s.T())
	fmt.Println("In the test that suppose to fail")
	require.Nil(nil)
}

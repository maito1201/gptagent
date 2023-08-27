package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCommand(t *testing.T) {
	testCases := []struct {
		command     string
		expectedOut string
		expectedErr error
	}{
		{
			command:     "echo Hello, World!",
			expectedOut: "Hello, World!\n",
		},
		{
			command:     "ls /nonexistent-directory",
			expectedOut: "ls: cannot access '/nonexistent-directory': No such file or directory\n",
		},
		// Add more test cases as needed
	}

	for _, testCase := range testCases {
		c := Command{Command: testCase.command}
		out := c.RunCommand()

		assert.Equal(t, testCase.expectedOut, out, "Output does not match expected value")
	}
}

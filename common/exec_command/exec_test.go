package exec_command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cmd = &LocalCommand{
	Command: "echo 123",
}

func TestLocalCommand_Exec(t *testing.T) {
	result, err := cmd.Exec()
	assert.NoError(t, err)
	assert.Equal(t, "123\n", result)
}

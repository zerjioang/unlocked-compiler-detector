package solidity_version_check

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestDetector(t *testing.T) {
	t.Run("first-app-example", func(t *testing.T) {
		fmt.Println("Checking solidity version")
		file := "testdata/first-app.sol"
		data, err := ioutil.ReadFile(file)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		result, err2 := CheckVersion(file, data)
		assert.NoError(t, err2)
		assert.NotNil(t, result)
		assert.True(t, result.Errored)
		result.Print()
	})
}

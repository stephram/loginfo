package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestProcessLogfile(t *testing.T) {
	t.Run("fail with invalid filename", func(t *testing.T) {
		c, err := processLogFile("", &map[string]int{}, &map[string]int{})
		assert.NotNil(t, err)
		assert.Equal(t, 0, c)
	})

	t.Run("fail with invalid file content", func(t *testing.T) {
		c, _ := processLogFile(os.Getenv("PWD"), nil, nil)
		assert.Equal(t, 0, c)
	})

	t.Run("success row count", func(t *testing.T) {
		iMap := map[string]int{}
		uMap := map[string]int{}

		c, err := processLogFile("../programming-task-example-data.log", &iMap, &uMap)
		assert.Nil(t, err)
		assert.Equal(t, 23, c)
	})

	t.Run("success processing", func(t *testing.T) {
		iMap := map[string]int{}
		uMap := map[string]int{}

		c, err := processLogFile("../testdata/data1.log", &iMap, &uMap)
		assert.Nil(t, err)
		assert.Equal(t, 6, c)

		assert.Equal(t, 3, iMap["1.1.1.1"])
		assert.Equal(t, 2, iMap["2.2.2.2"])
		assert.Equal(t, 1, iMap["3.3.3.3"])

		assert.Equal(t, 3, uMap["/intranet-analytics/1"])
		assert.Equal(t, 2, uMap["/intranet-analytics/2"])
		assert.Equal(t, 1, uMap["/intranet-analytics/3"])
	})
}

func TestAccumulate(t *testing.T) {
	t.Run("fail with invalid args", func(t *testing.T) {
		err := accumulate(nil, "")
		assert.NotNil(t, err)
	})

	t.Run("success", func(t *testing.T) {
		var err error
		rMap := map[string]int{}

		err = accumulate(&rMap, "one")
		assert.Nil(t, err)
		err = accumulate(&rMap, "one")
		assert.Nil(t, err)
		assert.Equal(t, rMap["one"], 2)
		assert.Equal(t, len(rMap), 1)
	})
}

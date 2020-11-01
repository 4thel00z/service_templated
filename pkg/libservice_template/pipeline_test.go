package libservice_template

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadPipeline(t *testing.T) {
	env, err := GetGoEnv()
	assert.Nil(t, err)
	modPath, found := env["GOMOD"]
	assert.True(t, found, "GOMOD not set")
	currentDir := filepath.Dir(modPath)
	fmt.Println(currentDir)
	path := strings.Join([]string{currentDir, "examples", "pipeline.so"}, string(os.PathSeparator))
	pipeline, err := LoadPipeline(path)
	if err != nil {
		t.Fatal(err)
	}

	_, err = pipeline.Init()

	if err != nil {
		t.Fatal(err)
	}

}

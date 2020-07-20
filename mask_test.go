package qm

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEachJsonLine(t *testing.T) {
	assert := assert.New(t)

	stdin := bytes.NewBufferString(`{"query":"select 1","time":100}
	{"query":"select now()","time":200}`)
	lines := []map[string]interface{}{}

	err := EachJsonLine(stdin, "query", "fingerprint", false, func(jl map[string]interface{}) {
		lines = append(lines, jl)
	})

	assert.Equal(nil, err)

	assert.Equal([]map[string]interface{}{
		{"fingerprint": "select ?", "query": "select 1", "time": float64(100)},
		{"fingerprint": "select now()", "query": "select now()", "time": float64(200)},
	}, lines)
}

func TestEachJsonLineWithSHA1(t *testing.T) {
	assert := assert.New(t)

	stdin := bytes.NewBufferString(`{"query":"select 1","time":100}
	{"query":"select now()","time":200}`)
	lines := []map[string]interface{}{}

	err := EachJsonLine(stdin, "query", "fingerprint", true, func(jl map[string]interface{}) {
		lines = append(lines, jl)
	})

	assert.Equal(nil, err)

	assert.Equal([]map[string]interface{}{
		{"fingerprint": "select ?", "fingerprint_sha1": "7ae509fc5e11f3bdd89c7e1a5829d6e86fbd8943", "query": "select 1", "query_sha1": "3232003928f9fe86a9cb634f450d5a53a4025819", "time": float64(100)},
		{"fingerprint": "select now()", "fingerprint_sha1": "e226cc45856c33b443aeca4c17e37a61d761a3e1", "query": "select now()", "query_sha1": "e226cc45856c33b443aeca4c17e37a61d761a3e1", "time": float64(200)},
	}, lines)
}

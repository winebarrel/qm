package qm

import (
	"bufio"
	"fmt"
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/percona/go-mysql/query"
)

func EachJsonLine(file io.Reader, queryKey string, fingerprintKey string, cb func(map[string]interface{})) error {
	reader := bufio.NewReader(file)

	for {
		line, err := readLine(reader)

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		jl := map[string]interface{}{}

		jsoniter.Unmarshal(line, &jl)

		q, ok := jl[queryKey]

		if !ok {
			return fmt.Errorf("query not found: key=%s json=%s", queryKey, string(line))
		}

		fingerprint := query.Fingerprint(q.(string))
		jl[fingerprintKey] = fingerprint

		cb(jl)
	}

	return nil
}

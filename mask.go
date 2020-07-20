package qm

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/percona/go-mysql/query"
)

func EachJsonLine(file io.Reader, queryKey string, fingerprintKey string, appendSHA1 bool, cb func(map[string]interface{})) error {
	reader := bufio.NewReader(file)

	for {
		line, err := readLine(reader)

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		jl := map[string]interface{}{}
		err = jsoniter.Unmarshal(line, &jl)

		if err != nil {
			return err
		}

		rawq, ok := jl[queryKey]

		if !ok {
			return fmt.Errorf("query not found: key=%s json=%s", queryKey, string(line))
		}

		q := rawq.(string)
		fingerprint := query.Fingerprint(q)
		jl[fingerprintKey] = fingerprint

		if appendSHA1 {
			querySHA1 := sha1.Sum([]byte(q))
			jl[queryKey+"_sha1"] = string(hex.EncodeToString(querySHA1[:]))

			fingerprintSHA1 := sha1.Sum([]byte(fingerprint))
			jl[fingerprintKey+"_sha1"] = string(hex.EncodeToString(fingerprintSHA1[:]))
		}

		cb(jl)
	}

	return nil
}

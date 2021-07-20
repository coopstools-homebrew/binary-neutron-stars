package kube

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestKubectl_parseNamespaceResponse(t *testing.T) {
	pastTime := time.Now().UTC().Add((- 60 * 60 * (24 * 3 + 6)) * time.Second)
	resp := fmt.Sprintf(
		"[default,%s][kube-node-lease,2021-07-17T20:08:03Z][kube-public,2021-07-17T20:08:03Z][kube-system,2021-07-17T20:08:03Z]",
		pastTime.Format("2006-01-02T15:04:05.000Z"),
		)
	submatches, _ := parseNamespaceResponse(resp)
	assert.Equal(t, 4, len(submatches))
	assert.Equal(t, "default", submatches[0].Name)
	assert.Equal(t, Age{Days:3, Hours:6}, submatches[0].Age)
}

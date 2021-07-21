package kube

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testResponse = "default\t%s\ningress-nginx\t2021-07-21T01:59:59Z\nkube-node-lease\t2021-07-21T01:55:47Z\nkube-public\t2021-07-21T01:55:47Z\nkube-system\t2021-07-21T01:55:46Z\nprod\t2021-07-21T02:08:15Z\n"

func TestKubectl_parseNamespaceResponse(t *testing.T) {
	pastTime := time.Now().UTC().Add((- 60 * 60 * (24 * 3 + 6)) * time.Second)
	resp := fmt.Sprintf(testResponse, pastTime.Format("2006-01-02T15:04:05.000Z"))
	submatches, _ := parseNamespaceResponse(resp)
	assert.Equal(t, 6, len(submatches))
	assert.Equal(t, "default", submatches[0].Name)
	assert.Equal(t, Age{Days:3, Hours:6}, submatches[0].Age)
}

package kube

import (
	"github.com/pkg/errors"
	"os/exec"
	"strings"
	"time"
)

type Kubectl struct {
	Kubeconfig string
}

type Namespace struct {
	Name 	string	`json:"name"`
	Created string	`json:"created"`
	Age 	Age		`json:"age"`
}

type Age struct {
	Days 	int `json:"days"`
	Hours	int `json:"hours"`
}

var jsonFormat = "jsonpath={range .items[*]}{.metadata.name}{\"\\t\"}{.metadata.creationTimestamp}{\"\\n\"}{end}"

func (ctl Kubectl) ListNamespaces() (parsed []Namespace, err error) {
	var args []string
	if ctl.Kubeconfig != "" {
		args = append(args, "--kubeconfig", ctl.Kubeconfig)
	}
	args = append(args, "-o", jsonFormat, "get", "namespaces")
	cmd := exec.Command("kubectl", args...)
	stdout, err := cmd.Output()
	if err != nil { return parsed, errors.Wrap(err, "Kubectl command failed")}
	return parseNamespaceResponse(string(stdout))
}

func parseNamespaceResponse(resp string) (parsed []Namespace, err error) {
	submatches := strings.Split(strings.TrimSpace(resp), "\n")
	parsed = make([]Namespace, len(submatches), len(submatches))
    for i, submatch := range submatches {
    	parts := strings.Split(submatch,"\t")
    	age, err := parseAge(parts[1])
    	if err != nil { age = Age{} }
    	parsed[i] = Namespace{Name: parts[0], Created: parts[1], Age: age}
	}
	return
}

func parseAge(timeString string) (age Age, err error) {
	layout := "2006-01-02T15:04:05Z"
	t, err := time.Parse(layout, timeString)
	if err != nil { return }
	since := int(time.Now().UTC().Sub(t).Hours())
	age = Age{Days: since / 24, Hours: since % 24}
	return
}

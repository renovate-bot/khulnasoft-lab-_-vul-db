package k8svulndb

import (
	"path/filepath"

	"github.com/khulnasoft-lab/vul-db/pkg/types"
	"github.com/khulnasoft-lab/vul-db/pkg/vulnsrc/osv"
	"github.com/khulnasoft-lab/vul-db/pkg/vulnsrc/vulnerability"
)

var (
	k8sDir = filepath.Join("vuln-list-k8s", "upstream")
)

func NewVulnSrc() osv.OSV {
	sources := map[types.Ecosystem]types.DataSource{
		vulnerability.Kubernetes: {
			ID:   vulnerability.K8sVulnDB,
			Name: "Official Kubernetes CVE Feed",
			URL:  "https://kubernetes.io/docs/reference/issues-security/official-cve-feed/index.json",
		},
	}
	return osv.New(k8sDir, vulnerability.K8sVulnDB, sources, nil)
}

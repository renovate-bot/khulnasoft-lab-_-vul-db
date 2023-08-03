package vulnsrc

import (
	"github.com/khulnasoft-labs/vul-db/pkg/types"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/alma"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/alpine"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/amazon"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/bundler"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/chainguard"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/composer"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/debian"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/ghsa"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/glad"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/mariner"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/node"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/nvd"
	oracleoval "github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/oracle-oval"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/osv"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/photon"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/redhat"
	redhatoval "github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/redhat-oval"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/rocky"
	susecvrf "github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/suse-cvrf"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/ubuntu"
	"github.com/khulnasoft-labs/vul-db/pkg/vulnsrc/wolfi"
)

type VulnSrc interface {
	Name() types.SourceID
	Update(dir string) (err error)
}

var (
	// All holds all data sources
	All = []VulnSrc{
		// NVD
		nvd.NewVulnSrc(),

		// OS packages
		alma.NewVulnSrc(),
		alpine.NewVulnSrc(),
		redhat.NewVulnSrc(),
		redhatoval.NewVulnSrc(),
		debian.NewVulnSrc(),
		ubuntu.NewVulnSrc(),
		amazon.NewVulnSrc(),
		oracleoval.NewVulnSrc(),
		rocky.NewVulnSrc(),
		susecvrf.NewVulnSrc(susecvrf.SUSEEnterpriseLinux),
		susecvrf.NewVulnSrc(susecvrf.OpenSUSE),
		photon.NewVulnSrc(),
		mariner.NewVulnSrc(),
		wolfi.NewVulnSrc(),
		chainguard.NewVulnSrc(),

		// Language-specific packages
		bundler.NewVulnSrc(),
		composer.NewVulnSrc(),
		node.NewVulnSrc(),
		ghsa.NewVulnSrc(),
		glad.NewVulnSrc(),
		osv.NewVulnSrc(),
	}
)

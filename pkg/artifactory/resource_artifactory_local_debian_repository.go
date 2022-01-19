package artifactory

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var debianLocalSchema = mergeSchema(baseLocalRepoSchema, map[string]*schema.Schema{
	"primary_keypair_ref": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Used to sign index files in Debian artifacts. ",
	},
	"secondary_keypair_ref": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Used to sign index files in Debian artifacts. ",
	},
	"trivial_layout": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "When set, the repository will use the deprecated trivial layout.",
		Deprecated:  "You shouldn't be using this",
	},
}, compressionFormats)

func resourceArtifactoryLocalDebianRepository() *schema.Resource {

	return mkResourceSchema(debianLocalSchema, defaultPacker, unPackLocalDebianRepository, func() interface{} {
		return &DebianLocalRepositoryParams{
			LocalRepositoryBaseParams: LocalRepositoryBaseParams{
				PackageType: "debian",
				Rclass:      "local",
			},
		}
	})
}

type DebianLocalRepositoryParams struct {
	LocalRepositoryBaseParams
	TrivialLayout           bool     `hcl:"trivial_layout" json:"debianTrivialLayout,omitempty"`
	IndexCompressionFormats []string `hcl:"index_compression_formats" json:"optionalIndexCompressionFormats,omitempty"`
	PrimaryKeyPairRef       string   `hcl:"primary_keypair_ref" json:"primaryKeyPairRef,omitempty"`
	SecondaryKeyPairRef     string   `hcl:"secondary_keypair_ref" json:"secondaryKeyPairRef,omitempty"`
}

func unPackLocalDebianRepository(data *schema.ResourceData) (interface{}, string, error) {
	d := &ResourceData{ResourceData: data}
	repo := DebianLocalRepositoryParams{
		LocalRepositoryBaseParams: unpackBaseRepo("local", data, "debian"),
		PrimaryKeyPairRef:         d.getString("primary_keypair_ref", false),
		SecondaryKeyPairRef:       d.getString("secondary_keypair_ref", false),
		TrivialLayout:             d.getBool("trivial_layout", false),
		IndexCompressionFormats:   d.getSet("index_compression_formats"),
	}
	return repo, repo.Id(), nil
}

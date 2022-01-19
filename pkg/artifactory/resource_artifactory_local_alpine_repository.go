package artifactory

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var alpineLocalSchema = mergeSchema(baseLocalRepoSchema, map[string]*schema.Schema{
	"primary_keypair_ref": {
		Type:     schema.TypeString,
		Optional: true,
		Description: "Used to sign index files in Alpine Linux repositories. " +
			"See: https://www.jfrog.com/confluence/display/JFROG/Alpine+Linux+Repositories#AlpineLinuxRepositories-SigningAlpineLinuxIndex",
	},
}, compressionFormats)

func resourceArtifactoryLocalAlpineRepository() *schema.Resource {
	return mkResourceSchema(alpineLocalSchema, defaultPacker, unPackLocalAlpineRepository, func() interface{} {
		return &AlpineLocalRepo{
			LocalRepositoryBaseParams: LocalRepositoryBaseParams{
				PackageType: "alpine",
				Rclass:      "local",
			},
		}
	})
}

type AlpineLocalRepo struct {
	LocalRepositoryBaseParams
	PrimaryKeyPairRef string `hcl:"primary_keypair_ref" json:"primaryKeyPairRef"`
}

func unPackLocalAlpineRepository(data *schema.ResourceData) (interface{}, string, error) {
	d := &ResourceData{ResourceData: data}
	repo := AlpineLocalRepo{
		LocalRepositoryBaseParams: unpackBaseRepo("local", data, "alpine"),
		PrimaryKeyPairRef:         d.getString("primary_keypair_ref", false),
	}

	return repo, repo.Id(), nil
}

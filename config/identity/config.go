package identity

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opentelekomcloud_identity_role_assignment_v3", func(r *config.Resource) {
		r.References["project_id"] = config.Reference{
			Type: "ProjectV3",
		}
	})
}

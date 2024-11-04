// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package image

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType   *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion   *string           `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug         *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce         *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError       *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	PersonalAccessToken *string           `mapstructure:"linode_token" cty:"linode_token" hcl:"linode_token"`
	APICAPath           *string           `mapstructure:"api_ca_path" cty:"api_ca_path" hcl:"api_ca_path"`
	Label               *string           `mapstructure:"label" cty:"label" hcl:"label"`
	LabelRegex          *string           `mapstructure:"label_regex" cty:"label_regex" hcl:"label_regex"`
	ID                  *string           `mapstructure:"id" cty:"id" hcl:"id"`
	IDRegex             *string           `mapstructure:"id_regex" cty:"id_regex" hcl:"id_regex"`
	Latest              *bool             `mapstructure:"latest" cty:"latest" hcl:"latest"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":        &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"linode_token":               &hcldec.AttrSpec{Name: "linode_token", Type: cty.String, Required: false},
		"api_ca_path":                &hcldec.AttrSpec{Name: "api_ca_path", Type: cty.String, Required: false},
		"label":                      &hcldec.AttrSpec{Name: "label", Type: cty.String, Required: false},
		"label_regex":                &hcldec.AttrSpec{Name: "label_regex", Type: cty.String, Required: false},
		"id":                         &hcldec.AttrSpec{Name: "id", Type: cty.String, Required: false},
		"id_regex":                   &hcldec.AttrSpec{Name: "id_regex", Type: cty.String, Required: false},
		"latest":                     &hcldec.AttrSpec{Name: "latest", Type: cty.Bool, Required: false},
	}
	return s
}

// FlatDatasourceOutput is an auto-generated flat version of DatasourceOutput.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatDatasourceOutput struct {
	ID           *string  `mapstructure:"id" cty:"id" hcl:"id"`
	Capabilities []string `mapstructure:"capabilities" cty:"capabilities" hcl:"capabilities"`
	Created      *string  `mapstructure:"created" cty:"created" hcl:"created"`
	CreatedBy    *string  `mapstructure:"created_by" cty:"created_by" hcl:"created_by"`
	Deprecated   *bool    `mapstructure:"deprecated" cty:"deprecated" hcl:"deprecated"`
	Description  *string  `mapstructure:"description" cty:"description" hcl:"description"`
	EOL          *string  `mapstructure:"eol" cty:"eol" hcl:"eol"`
	Expiry       *string  `mapstructure:"expiry" cty:"expiry" hcl:"expiry"`
	IsPublic     *bool    `mapstructure:"is_public" cty:"is_public" hcl:"is_public"`
	Label        *string  `mapstructure:"label" cty:"label" hcl:"label"`
	Size         *int     `mapstructure:"size" cty:"size" hcl:"size"`
	Type         *string  `mapstructure:"type" cty:"type" hcl:"type"`
	Updated      *string  `mapstructure:"updated" cty:"updated" hcl:"updated"`
	Vendor       *string  `mapstructure:"vendor" cty:"vendor" hcl:"vendor"`
}

// FlatMapstructure returns a new FlatDatasourceOutput.
// FlatDatasourceOutput is an auto-generated flat version of DatasourceOutput.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*DatasourceOutput) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatDatasourceOutput)
}

// HCL2Spec returns the hcl spec of a DatasourceOutput.
// This spec is used by HCL to read the fields of DatasourceOutput.
// The decoded values from this spec will then be applied to a FlatDatasourceOutput.
func (*FlatDatasourceOutput) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"id":           &hcldec.AttrSpec{Name: "id", Type: cty.String, Required: false},
		"capabilities": &hcldec.AttrSpec{Name: "capabilities", Type: cty.List(cty.String), Required: false},
		"created":      &hcldec.AttrSpec{Name: "created", Type: cty.String, Required: false},
		"created_by":   &hcldec.AttrSpec{Name: "created_by", Type: cty.String, Required: false},
		"deprecated":   &hcldec.AttrSpec{Name: "deprecated", Type: cty.Bool, Required: false},
		"description":  &hcldec.AttrSpec{Name: "description", Type: cty.String, Required: false},
		"eol":          &hcldec.AttrSpec{Name: "eol", Type: cty.String, Required: false},
		"expiry":       &hcldec.AttrSpec{Name: "expiry", Type: cty.String, Required: false},
		"is_public":    &hcldec.AttrSpec{Name: "is_public", Type: cty.Bool, Required: false},
		"label":        &hcldec.AttrSpec{Name: "label", Type: cty.String, Required: false},
		"size":         &hcldec.AttrSpec{Name: "size", Type: cty.Number, Required: false},
		"type":         &hcldec.AttrSpec{Name: "type", Type: cty.String, Required: false},
		"updated":      &hcldec.AttrSpec{Name: "updated", Type: cty.String, Required: false},
		"vendor":       &hcldec.AttrSpec{Name: "vendor", Type: cty.String, Required: false},
	}
	return s
}

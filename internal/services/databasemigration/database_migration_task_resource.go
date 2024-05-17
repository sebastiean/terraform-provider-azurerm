// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasemigration

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tags"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

type TaskResource struct{}

var _ sdk.Resource = TaskResource{}

func (TaskResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"name": {
			Type:     pluginsdk.TypeString,
			Required: true,
		},

		"project_name": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
			// ValidateFunc: validate.ProjectName,
		},

		"service_name": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
			// ValidateFunc: validate.ServiceName,
		},

		"resource_group_name": commonschema.ResourceGroupName(),

		"enable_data_integrity_validation": {
			Type: pluginsdk.TypeBool,
		},

		"enable_query_analysis_validation": {
			Type: pluginsdk.TypeBool,
		},

		"enable_schema_validation": {
			Type: pluginsdk.TypeBool,
		},

		"tags": tags.Schema(),
	}
}

func (TaskResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (TaskResource) ModelObject() interface{} {
	return nil
}

func (TaskResource) ResourceType() string {
	return "azurerm_database_migration_task"
}

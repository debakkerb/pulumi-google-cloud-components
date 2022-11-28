package project

import (
	"fmt"
	"github.com/debakkerb/pulumi-google-cloud-components/util/project"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strings"
)

/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

type GoogleProjectArgs struct {
	ProjectName      string
	GenerateRandomID bool
	ProjectAPIS      []string
	BillingAccountID string
	FolderID         string
	OrganizationID   string
}

type GoogleProject struct {
	pulumi.ResourceState

	Project *organizations.Project `pulumi:"project"`
}

func NewGoogleProject(ctx *pulumi.Context, resourceName string, args *GoogleProjectArgs, opts ...pulumi.ResourceOption) (*GoogleProject, error) {
	if args == nil {
		args = &GoogleProjectArgs{}
	}

	component := &GoogleProject{}
	err := ctx.RegisterComponentResource("google-pulumi-components:resource-manager:Project", resourceName, component, opts...)
	if err != nil {
		return nil, err
	}

	projectSettings := &organizations.ProjectArgs{
		BillingAccount:    pulumi.String(args.BillingAccountID),
		FolderId:          pulumi.String(args.FolderID),
		OrgId:             pulumi.String(args.OrganizationID),
		AutoCreateNetwork: pulumi.Bool(false),
		ProjectId:         pulumi.String(project.GenerateProjectID(args.ProjectName, args.GenerateRandomID)),
	}

	newProject, err := organizations.NewProject(ctx, resourceName, projectSettings, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	for _, api := range args.ProjectAPIS {
		resourceName := fmt.Sprintf("%s-%s-api", newProject.ProjectId, api[0:strings.Index(".", api)-1])
		_, err := projects.NewService(ctx, resourceName, &projects.ServiceArgs{
			DisableDependentServices: pulumi.Bool(true),
			Project:                  newProject.ProjectId,
			Service:                  pulumi.String(api),
		})

		if err != nil {
			return nil, err
		}
	}

	component.Project = newProject

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"projectID":   newProject.ProjectId,
		"projectName": newProject.Name,
	}); err != nil {
		return nil, err
	}
	return component, nil
}

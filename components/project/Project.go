package project

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
		FolderId:          nil,
		OrgId:             nil,
		AutoCreateNetwork: pulumi.Bool(false),
		ProjectId:         pulumi.String(args.ProjectName),
	}

}

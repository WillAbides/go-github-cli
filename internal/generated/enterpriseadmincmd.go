// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import internal "github.com/octo-cli/octo-cli/internal"

type EnterpriseAdminCmd struct {
	AddOrgAccessToSelfHostedRunnerGroupInEnterprise         EnterpriseAdminAddOrgAccessToSelfHostedRunnerGroupInEnterpriseCmd         `cmd:""`
	AddSelfHostedRunnerToGroupForEnterprise                 EnterpriseAdminAddSelfHostedRunnerToGroupForEnterpriseCmd                 `cmd:""`
	CreateRegistrationTokenForEnterprise                    EnterpriseAdminCreateRegistrationTokenForEnterpriseCmd                    `cmd:""`
	CreateRemoveTokenForEnterprise                          EnterpriseAdminCreateRemoveTokenForEnterpriseCmd                          `cmd:""`
	CreateSelfHostedRunnerGroupForEnterprise                EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseCmd                `cmd:""`
	DeleteScimGroupFromEnterprise                           EnterpriseAdminDeleteScimGroupFromEnterpriseCmd                           `cmd:""`
	DeleteSelfHostedRunnerFromEnterprise                    EnterpriseAdminDeleteSelfHostedRunnerFromEnterpriseCmd                    `cmd:""`
	DeleteSelfHostedRunnerGroupFromEnterprise               EnterpriseAdminDeleteSelfHostedRunnerGroupFromEnterpriseCmd               `cmd:""`
	DeleteUserFromEnterprise                                EnterpriseAdminDeleteUserFromEnterpriseCmd                                `cmd:""`
	DisableSelectedOrganizationGithubActionsEnterprise      EnterpriseAdminDisableSelectedOrganizationGithubActionsEnterpriseCmd      `cmd:""`
	EnableSelectedOrganizationGithubActionsEnterprise       EnterpriseAdminEnableSelectedOrganizationGithubActionsEnterpriseCmd       `cmd:""`
	GetAllowedActionsEnterprise                             EnterpriseAdminGetAllowedActionsEnterpriseCmd                             `cmd:""`
	GetGithubActionsPermissionsEnterprise                   EnterpriseAdminGetGithubActionsPermissionsEnterpriseCmd                   `cmd:""`
	GetProvisioningInformationForEnterpriseGroup            EnterpriseAdminGetProvisioningInformationForEnterpriseGroupCmd            `cmd:""`
	GetProvisioningInformationForEnterpriseUser             EnterpriseAdminGetProvisioningInformationForEnterpriseUserCmd             `cmd:""`
	GetSelfHostedRunnerForEnterprise                        EnterpriseAdminGetSelfHostedRunnerForEnterpriseCmd                        `cmd:""`
	GetSelfHostedRunnerGroupForEnterprise                   EnterpriseAdminGetSelfHostedRunnerGroupForEnterpriseCmd                   `cmd:""`
	ListOrgAccessToSelfHostedRunnerGroupInEnterprise        EnterpriseAdminListOrgAccessToSelfHostedRunnerGroupInEnterpriseCmd        `cmd:""`
	ListProvisionedGroupsEnterprise                         EnterpriseAdminListProvisionedGroupsEnterpriseCmd                         `cmd:""`
	ListProvisionedIdentitiesEnterprise                     EnterpriseAdminListProvisionedIdentitiesEnterpriseCmd                     `cmd:""`
	ListRunnerApplicationsForEnterprise                     EnterpriseAdminListRunnerApplicationsForEnterpriseCmd                     `cmd:""`
	ListSelectedOrganizationsEnabledGithubActionsEnterprise EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterpriseCmd `cmd:""`
	ListSelfHostedRunnerGroupsForEnterprise                 EnterpriseAdminListSelfHostedRunnerGroupsForEnterpriseCmd                 `cmd:""`
	ListSelfHostedRunnersForEnterprise                      EnterpriseAdminListSelfHostedRunnersForEnterpriseCmd                      `cmd:""`
	ListSelfHostedRunnersInGroupForEnterprise               EnterpriseAdminListSelfHostedRunnersInGroupForEnterpriseCmd               `cmd:""`
	ProvisionAndInviteEnterpriseGroup                       EnterpriseAdminProvisionAndInviteEnterpriseGroupCmd                       `cmd:""`
	ProvisionAndInviteEnterpriseUser                        EnterpriseAdminProvisionAndInviteEnterpriseUserCmd                        `cmd:""`
	RemoveOrgAccessToSelfHostedRunnerGroupInEnterprise      EnterpriseAdminRemoveOrgAccessToSelfHostedRunnerGroupInEnterpriseCmd      `cmd:""`
	RemoveSelfHostedRunnerFromGroupForEnterprise            EnterpriseAdminRemoveSelfHostedRunnerFromGroupForEnterpriseCmd            `cmd:""`
	SetAllowedActionsEnterprise                             EnterpriseAdminSetAllowedActionsEnterpriseCmd                             `cmd:""`
	SetGithubActionsPermissionsEnterprise                   EnterpriseAdminSetGithubActionsPermissionsEnterpriseCmd                   `cmd:""`
	SetInformationForProvisionedEnterpriseGroup             EnterpriseAdminSetInformationForProvisionedEnterpriseGroupCmd             `cmd:""`
	SetInformationForProvisionedEnterpriseUser              EnterpriseAdminSetInformationForProvisionedEnterpriseUserCmd              `cmd:""`
	SetOrgAccessToSelfHostedRunnerGroupInEnterprise         EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseCmd         `cmd:""`
	SetSelectedOrganizationsEnabledGithubActionsEnterprise  EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseCmd  `cmd:""`
	SetSelfHostedRunnersInGroupForEnterprise                EnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseCmd                `cmd:""`
	UpdateAttributeForEnterpriseGroup                       EnterpriseAdminUpdateAttributeForEnterpriseGroupCmd                       `cmd:""`
	UpdateAttributeForEnterpriseUser                        EnterpriseAdminUpdateAttributeForEnterpriseUserCmd                        `cmd:""`
	UpdateSelfHostedRunnerGroupForEnterprise                EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseCmd                `cmd:""`
}

type EnterpriseAdminAddOrgAccessToSelfHostedRunnerGroupInEnterpriseCmd struct {
	Enterprise    string `name:"enterprise" required:"true"`
	RunnerGroupId int64  `name:"runner_group_id" required:"true"`
	OrgId         int64  `name:"org_id" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminAddOrgAccessToSelfHostedRunnerGroupInEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations/{org_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("runner_group_id", c.RunnerGroupId)
	c.UpdateURLPath("org_id", c.OrgId)
	return c.DoRequest("PUT")
}

type EnterpriseAdminAddSelfHostedRunnerToGroupForEnterpriseCmd struct {
	Enterprise    string `name:"enterprise" required:"true"`
	RunnerGroupId int64  `name:"runner_group_id" required:"true"`
	RunnerId      int64  `name:"runner_id" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminAddSelfHostedRunnerToGroupForEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners/{runner_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("runner_group_id", c.RunnerGroupId)
	c.UpdateURLPath("runner_id", c.RunnerId)
	return c.DoRequest("PUT")
}

type EnterpriseAdminCreateRegistrationTokenForEnterpriseCmd struct {
	Enterprise string `name:"enterprise" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminCreateRegistrationTokenForEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runners/registration-token")
	c.UpdateURLPath("enterprise", c.Enterprise)
	return c.DoRequest("POST")
}

type EnterpriseAdminCreateRemoveTokenForEnterpriseCmd struct {
	Enterprise string `name:"enterprise" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminCreateRemoveTokenForEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runners/remove-token")
	c.UpdateURLPath("enterprise", c.Enterprise)
	return c.DoRequest("POST")
}

type EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseCmd struct {
	Enterprise              string  `name:"enterprise" required:"true"`
	Runners                 []int64 `name:"runners"`
	SelectedOrganizationIds []int64 `name:"selected_organization_ids"`
	Visibility              string  `name:"visibility"`
	Name                    string  `name:"name" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runner-groups")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateBody("name", c.Name)
	c.UpdateBody("runners", c.Runners)
	c.UpdateBody("selected_organization_ids", c.SelectedOrganizationIds)
	c.UpdateBody("visibility", c.Visibility)
	return c.DoRequest("POST")
}

type EnterpriseAdminDeleteScimGroupFromEnterpriseCmd struct {
	Enterprise  string `name:"enterprise" required:"true"`
	ScimGroupId string `name:"scim_group_id" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminDeleteScimGroupFromEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/enterprises/{enterprise}/Groups/{scim_group_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("scim_group_id", c.ScimGroupId)
	return c.DoRequest("DELETE")
}

type EnterpriseAdminDeleteSelfHostedRunnerFromEnterpriseCmd struct {
	Enterprise string `name:"enterprise" required:"true"`
	RunnerId   int64  `name:"runner_id" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminDeleteSelfHostedRunnerFromEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runners/{runner_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("runner_id", c.RunnerId)
	return c.DoRequest("DELETE")
}

type EnterpriseAdminDeleteSelfHostedRunnerGroupFromEnterpriseCmd struct {
	Enterprise    string `name:"enterprise" required:"true"`
	RunnerGroupId int64  `name:"runner_group_id" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminDeleteSelfHostedRunnerGroupFromEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("runner_group_id", c.RunnerGroupId)
	return c.DoRequest("DELETE")
}

type EnterpriseAdminDeleteUserFromEnterpriseCmd struct {
	Enterprise string `name:"enterprise" required:"true"`
	ScimUserId string `name:"scim_user_id" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminDeleteUserFromEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/enterprises/{enterprise}/Users/{scim_user_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("scim_user_id", c.ScimUserId)
	return c.DoRequest("DELETE")
}

type EnterpriseAdminDisableSelectedOrganizationGithubActionsEnterpriseCmd struct {
	Enterprise string `name:"enterprise" required:"true"`
	OrgId      int64  `name:"org_id" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminDisableSelectedOrganizationGithubActionsEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/permissions/organizations/{org_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("org_id", c.OrgId)
	return c.DoRequest("DELETE")
}

type EnterpriseAdminEnableSelectedOrganizationGithubActionsEnterpriseCmd struct {
	Enterprise string `name:"enterprise" required:"true"`
	OrgId      int64  `name:"org_id" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminEnableSelectedOrganizationGithubActionsEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/permissions/organizations/{org_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("org_id", c.OrgId)
	return c.DoRequest("PUT")
}

type EnterpriseAdminGetAllowedActionsEnterpriseCmd struct {
	Enterprise string `name:"enterprise" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminGetAllowedActionsEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/permissions/selected-actions")
	c.UpdateURLPath("enterprise", c.Enterprise)
	return c.DoRequest("GET")
}

type EnterpriseAdminGetGithubActionsPermissionsEnterpriseCmd struct {
	Enterprise string `name:"enterprise" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminGetGithubActionsPermissionsEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/permissions")
	c.UpdateURLPath("enterprise", c.Enterprise)
	return c.DoRequest("GET")
}

type EnterpriseAdminGetProvisioningInformationForEnterpriseGroupCmd struct {
	Enterprise  string `name:"enterprise" required:"true"`
	ScimGroupId string `name:"scim_group_id" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminGetProvisioningInformationForEnterpriseGroupCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/enterprises/{enterprise}/Groups/{scim_group_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("scim_group_id", c.ScimGroupId)
	return c.DoRequest("GET")
}

type EnterpriseAdminGetProvisioningInformationForEnterpriseUserCmd struct {
	Enterprise string `name:"enterprise" required:"true"`
	ScimUserId string `name:"scim_user_id" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminGetProvisioningInformationForEnterpriseUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/enterprises/{enterprise}/Users/{scim_user_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("scim_user_id", c.ScimUserId)
	return c.DoRequest("GET")
}

type EnterpriseAdminGetSelfHostedRunnerForEnterpriseCmd struct {
	Enterprise string `name:"enterprise" required:"true"`
	RunnerId   int64  `name:"runner_id" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminGetSelfHostedRunnerForEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runners/{runner_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("runner_id", c.RunnerId)
	return c.DoRequest("GET")
}

type EnterpriseAdminGetSelfHostedRunnerGroupForEnterpriseCmd struct {
	Enterprise    string `name:"enterprise" required:"true"`
	RunnerGroupId int64  `name:"runner_group_id" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminGetSelfHostedRunnerGroupForEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("runner_group_id", c.RunnerGroupId)
	return c.DoRequest("GET")
}

type EnterpriseAdminListOrgAccessToSelfHostedRunnerGroupInEnterpriseCmd struct {
	Enterprise    string `name:"enterprise" required:"true"`
	RunnerGroupId int64  `name:"runner_group_id" required:"true"`
	Page          int64  `name:"page"`
	PerPage       int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *EnterpriseAdminListOrgAccessToSelfHostedRunnerGroupInEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("runner_group_id", c.RunnerGroupId)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type EnterpriseAdminListProvisionedGroupsEnterpriseCmd struct {
	Enterprise string `name:"enterprise" required:"true"`
	Count      int64  `name:"count"`
	StartIndex int64  `name:"startIndex"`
	internal.BaseCmd
}

func (c *EnterpriseAdminListProvisionedGroupsEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/enterprises/{enterprise}/Groups")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLQuery("startIndex", c.StartIndex)
	c.UpdateURLQuery("count", c.Count)
	return c.DoRequest("GET")
}

type EnterpriseAdminListProvisionedIdentitiesEnterpriseCmd struct {
	Enterprise string `name:"enterprise" required:"true"`
	Count      int64  `name:"count"`
	StartIndex int64  `name:"startIndex"`
	internal.BaseCmd
}

func (c *EnterpriseAdminListProvisionedIdentitiesEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/enterprises/{enterprise}/Users")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLQuery("startIndex", c.StartIndex)
	c.UpdateURLQuery("count", c.Count)
	return c.DoRequest("GET")
}

type EnterpriseAdminListRunnerApplicationsForEnterpriseCmd struct {
	Enterprise string `name:"enterprise" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminListRunnerApplicationsForEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runners/downloads")
	c.UpdateURLPath("enterprise", c.Enterprise)
	return c.DoRequest("GET")
}

type EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterpriseCmd struct {
	Enterprise string `name:"enterprise" required:"true"`
	Page       int64  `name:"page"`
	PerPage    int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/permissions/organizations")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type EnterpriseAdminListSelfHostedRunnerGroupsForEnterpriseCmd struct {
	Enterprise string `name:"enterprise" required:"true"`
	Page       int64  `name:"page"`
	PerPage    int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *EnterpriseAdminListSelfHostedRunnerGroupsForEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runner-groups")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type EnterpriseAdminListSelfHostedRunnersForEnterpriseCmd struct {
	Enterprise string `name:"enterprise" required:"true"`
	Page       int64  `name:"page"`
	PerPage    int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *EnterpriseAdminListSelfHostedRunnersForEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runners")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type EnterpriseAdminListSelfHostedRunnersInGroupForEnterpriseCmd struct {
	Enterprise    string `name:"enterprise" required:"true"`
	RunnerGroupId int64  `name:"runner_group_id" required:"true"`
	Page          int64  `name:"page"`
	PerPage       int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *EnterpriseAdminListSelfHostedRunnersInGroupForEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("runner_group_id", c.RunnerGroupId)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type EnterpriseAdminProvisionAndInviteEnterpriseGroupCmd struct {
	Enterprise  string                `name:"enterprise" required:"true"`
	Members     []internal.JSONObject `name:"members"`
	DisplayName string                `name:"displayName" required:"true"`
	Schemas     []string              `name:"schemas" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminProvisionAndInviteEnterpriseGroupCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/enterprises/{enterprise}/Groups")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateBody("displayName", c.DisplayName)
	c.UpdateBody("members", c.Members)
	c.UpdateBody("schemas", c.Schemas)
	return c.DoRequest("POST")
}

type EnterpriseAdminProvisionAndInviteEnterpriseUserCmd struct {
	Enterprise     string                `name:"enterprise" required:"true"`
	Groups         []internal.JSONObject `name:"groups"`
	Emails         []internal.JSONObject `name:"emails" required:"true"`
	NameFamilyName string                `name:"name.familyName" required:"true"`
	NameGivenName  string                `name:"name.givenName" required:"true"`
	Schemas        []string              `name:"schemas" required:"true"`
	UserName       string                `name:"userName" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminProvisionAndInviteEnterpriseUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/enterprises/{enterprise}/Users")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateBody("emails", c.Emails)
	c.UpdateBody("groups", c.Groups)
	c.UpdateBody("name.familyName", c.NameFamilyName)
	c.UpdateBody("name.givenName", c.NameGivenName)
	c.UpdateBody("schemas", c.Schemas)
	c.UpdateBody("userName", c.UserName)
	return c.DoRequest("POST")
}

type EnterpriseAdminRemoveOrgAccessToSelfHostedRunnerGroupInEnterpriseCmd struct {
	Enterprise    string `name:"enterprise" required:"true"`
	RunnerGroupId int64  `name:"runner_group_id" required:"true"`
	OrgId         int64  `name:"org_id" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminRemoveOrgAccessToSelfHostedRunnerGroupInEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations/{org_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("runner_group_id", c.RunnerGroupId)
	c.UpdateURLPath("org_id", c.OrgId)
	return c.DoRequest("DELETE")
}

type EnterpriseAdminRemoveSelfHostedRunnerFromGroupForEnterpriseCmd struct {
	Enterprise    string `name:"enterprise" required:"true"`
	RunnerGroupId int64  `name:"runner_group_id" required:"true"`
	RunnerId      int64  `name:"runner_id" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminRemoveSelfHostedRunnerFromGroupForEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners/{runner_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("runner_group_id", c.RunnerGroupId)
	c.UpdateURLPath("runner_id", c.RunnerId)
	return c.DoRequest("DELETE")
}

type EnterpriseAdminSetAllowedActionsEnterpriseCmd struct {
	Enterprise         string   `name:"enterprise" required:"true"`
	GithubOwnedAllowed bool     `name:"github_owned_allowed"`
	PatternsAllowed    []string `name:"patterns_allowed"`
	VerifiedAllowed    bool     `name:"verified_allowed"`
	internal.BaseCmd
}

func (c *EnterpriseAdminSetAllowedActionsEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/permissions/selected-actions")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateBody("github_owned_allowed", c.GithubOwnedAllowed)
	c.UpdateBody("patterns_allowed", c.PatternsAllowed)
	c.UpdateBody("verified_allowed", c.VerifiedAllowed)
	return c.DoRequest("PUT")
}

type EnterpriseAdminSetGithubActionsPermissionsEnterpriseCmd struct {
	Enterprise           string `name:"enterprise" required:"true"`
	AllowedActions       string `name:"allowed_actions"`
	EnabledOrganizations string `name:"enabled_organizations" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminSetGithubActionsPermissionsEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/permissions")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateBody("allowed_actions", c.AllowedActions)
	c.UpdateBody("enabled_organizations", c.EnabledOrganizations)
	return c.DoRequest("PUT")
}

type EnterpriseAdminSetInformationForProvisionedEnterpriseGroupCmd struct {
	Enterprise  string                `name:"enterprise" required:"true"`
	ScimGroupId string                `name:"scim_group_id" required:"true"`
	Members     []internal.JSONObject `name:"members"`
	DisplayName string                `name:"displayName" required:"true"`
	Schemas     []string              `name:"schemas" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminSetInformationForProvisionedEnterpriseGroupCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/enterprises/{enterprise}/Groups/{scim_group_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("scim_group_id", c.ScimGroupId)
	c.UpdateBody("displayName", c.DisplayName)
	c.UpdateBody("members", c.Members)
	c.UpdateBody("schemas", c.Schemas)
	return c.DoRequest("PUT")
}

type EnterpriseAdminSetInformationForProvisionedEnterpriseUserCmd struct {
	Enterprise     string                `name:"enterprise" required:"true"`
	ScimUserId     string                `name:"scim_user_id" required:"true"`
	Groups         []internal.JSONObject `name:"groups"`
	Emails         []internal.JSONObject `name:"emails" required:"true"`
	NameFamilyName string                `name:"name.familyName" required:"true"`
	NameGivenName  string                `name:"name.givenName" required:"true"`
	Schemas        []string              `name:"schemas" required:"true"`
	UserName       string                `name:"userName" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminSetInformationForProvisionedEnterpriseUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/enterprises/{enterprise}/Users/{scim_user_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("scim_user_id", c.ScimUserId)
	c.UpdateBody("emails", c.Emails)
	c.UpdateBody("groups", c.Groups)
	c.UpdateBody("name.familyName", c.NameFamilyName)
	c.UpdateBody("name.givenName", c.NameGivenName)
	c.UpdateBody("schemas", c.Schemas)
	c.UpdateBody("userName", c.UserName)
	return c.DoRequest("PUT")
}

type EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseCmd struct {
	Enterprise              string  `name:"enterprise" required:"true"`
	RunnerGroupId           int64   `name:"runner_group_id" required:"true"`
	SelectedOrganizationIds []int64 `name:"selected_organization_ids" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("runner_group_id", c.RunnerGroupId)
	c.UpdateBody("selected_organization_ids", c.SelectedOrganizationIds)
	return c.DoRequest("PUT")
}

type EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseCmd struct {
	Enterprise              string  `name:"enterprise" required:"true"`
	SelectedOrganizationIds []int64 `name:"selected_organization_ids" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/permissions/organizations")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateBody("selected_organization_ids", c.SelectedOrganizationIds)
	return c.DoRequest("PUT")
}

type EnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseCmd struct {
	Enterprise    string  `name:"enterprise" required:"true"`
	RunnerGroupId int64   `name:"runner_group_id" required:"true"`
	Runners       []int64 `name:"runners" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("runner_group_id", c.RunnerGroupId)
	c.UpdateBody("runners", c.Runners)
	return c.DoRequest("PUT")
}

type EnterpriseAdminUpdateAttributeForEnterpriseGroupCmd struct {
	Enterprise  string                `name:"enterprise" required:"true"`
	ScimGroupId string                `name:"scim_group_id" required:"true"`
	Operations  []internal.JSONObject `name:"Operations" required:"true"`
	Schemas     []string              `name:"schemas" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminUpdateAttributeForEnterpriseGroupCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/enterprises/{enterprise}/Groups/{scim_group_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("scim_group_id", c.ScimGroupId)
	c.UpdateBody("Operations", c.Operations)
	c.UpdateBody("schemas", c.Schemas)
	return c.DoRequest("PATCH")
}

type EnterpriseAdminUpdateAttributeForEnterpriseUserCmd struct {
	Enterprise string                `name:"enterprise" required:"true"`
	ScimUserId string                `name:"scim_user_id" required:"true"`
	Operations []internal.JSONObject `name:"Operations" required:"true"`
	Schemas    []string              `name:"schemas" required:"true"`
	internal.BaseCmd
}

func (c *EnterpriseAdminUpdateAttributeForEnterpriseUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/enterprises/{enterprise}/Users/{scim_user_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("scim_user_id", c.ScimUserId)
	c.UpdateBody("Operations", c.Operations)
	c.UpdateBody("schemas", c.Schemas)
	return c.DoRequest("PATCH")
}

type EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseCmd struct {
	Enterprise    string `name:"enterprise" required:"true"`
	RunnerGroupId int64  `name:"runner_group_id" required:"true"`
	Name          string `name:"name"`
	Visibility    string `name:"visibility"`
	internal.BaseCmd
}

func (c *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}")
	c.UpdateURLPath("enterprise", c.Enterprise)
	c.UpdateURLPath("runner_group_id", c.RunnerGroupId)
	c.UpdateBody("name", c.Name)
	c.UpdateBody("visibility", c.Visibility)
	return c.DoRequest("PATCH")
}

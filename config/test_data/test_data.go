package config_test_data
	
	import (
		"github.com/pivotalservices/cf-mgmt/config"
		"github.com/pivotalservices/cf-mgmt/ldap"
		mock "github.com/pivotalservices/cf-mgmt/utils/mocks"
	)
	
	func PopulateWithTestData(utilsMgrMock *mock.MockUtilsManager) error {
		utilsMgrMock.MockFileData["./fixtures/config/ldap.yml"] = ldap.Config{Enabled:true, LdapHost:"127.0.0.1", LdapPort:10389, TLS:false, BindDN:"uid=admin,ou=system", BindPassword:"secret", UserSearchBase:"ou=users,dc=example,dc=com", UserNameAttribute:"uid", UserMailAttribute:"mail", UserObjectClass:"", GroupSearchBase:"ou=groups,dc=example,dc=com", GroupAttribute:"member", Origin:""}
		utilsMgrMock.MockFileData["./fixtures/config/orgs.yml"] = config.Orgs{Orgs:[]string{"test", "test2"}, EnableDeleteOrgs:false, ProtectedOrgs:[]string(nil)}
		utilsMgrMock.MockFileData["./fixtures/config/test/orgConfig.yml"] = config.OrgConfig{Org:"test", BillingManagerGroup:"test_billing_managers", ManagerGroup:"", AuditorGroup:"", BillingManager:config.UserMgmt{LDAPUsers:[]string(nil), Users:[]string(nil), SamlUsers:[]string(nil), LDAPGroup:"", LDAPGroups:[]string(nil)}, Manager:config.UserMgmt{LDAPUsers:[]string(nil), Users:[]string(nil), SamlUsers:[]string(nil), LDAPGroup:"", LDAPGroups:[]string(nil)}, Auditor:config.UserMgmt{LDAPUsers:[]string(nil), Users:[]string(nil), SamlUsers:[]string(nil), LDAPGroup:"", LDAPGroups:[]string(nil)}, PrivateDomains:[]string(nil), RemovePrivateDomains:false, EnableOrgQuota:true, MemoryLimit:10240, InstanceMemoryLimit:-1, TotalRoutes:10, TotalServices:-1, PaidServicePlansAllowed:true, RemoveUsers:false, TotalPrivateDomains:0, TotalReservedRoutePorts:0, TotalServiceKeys:0, AppInstanceLimit:0, DefaultIsoSegment:""}
		utilsMgrMock.MockFileData["./fixtures/config/test/space1/security-group.json"] = "[\n    {\n        \"protocol\": \"all\",\n        \"destination\": \"10.68.192.1-10.68.192.49\"\n    }\n]\n"
		utilsMgrMock.MockFileData["./fixtures/config/test/space1/spaceConfig.yml"] = config.SpaceConfig{Org:"test", Space:"space1", Developer:config.UserMgmt{LDAPUsers:[]string(nil), Users:[]string(nil), SamlUsers:[]string(nil), LDAPGroup:"", LDAPGroups:[]string(nil)}, Manager:config.UserMgmt{LDAPUsers:[]string(nil), Users:[]string(nil), SamlUsers:[]string(nil), LDAPGroup:"", LDAPGroups:[]string(nil)}, Auditor:config.UserMgmt{LDAPUsers:[]string(nil), Users:[]string(nil), SamlUsers:[]string(nil), LDAPGroup:"", LDAPGroups:[]string(nil)}, DeveloperGroup:"", ManagerGroup:"", AuditorGroup:"", AllowSSH:true, EnableSpaceQuota:true, MemoryLimit:10240, InstanceMemoryLimit:-1, TotalRoutes:10, TotalServices:-1, PaidServicePlansAllowed:true, EnableSecurityGroup:true, SecurityGroupContents:"", RemoveUsers:false, TotalPrivateDomains:0, TotalReservedRoutePorts:0, TotalServiceKeys:0, AppInstanceLimit:0, IsoSegment:""}
		utilsMgrMock.MockFileData["./fixtures/config/test/space2/spaceConfig.yml"] = config.SpaceConfig{Org:"test", Space:"space2", Developer:config.UserMgmt{LDAPUsers:[]string(nil), Users:[]string(nil), SamlUsers:[]string(nil), LDAPGroup:"", LDAPGroups:[]string(nil)}, Manager:config.UserMgmt{LDAPUsers:[]string(nil), Users:[]string(nil), SamlUsers:[]string(nil), LDAPGroup:"", LDAPGroups:[]string(nil)}, Auditor:config.UserMgmt{LDAPUsers:[]string(nil), Users:[]string(nil), SamlUsers:[]string(nil), LDAPGroup:"", LDAPGroups:[]string(nil)}, DeveloperGroup:"", ManagerGroup:"", AuditorGroup:"", AllowSSH:true, EnableSpaceQuota:true, MemoryLimit:10240, InstanceMemoryLimit:-1, TotalRoutes:10, TotalServices:-1, PaidServicePlansAllowed:true, EnableSecurityGroup:false, SecurityGroupContents:"", RemoveUsers:false, TotalPrivateDomains:0, TotalReservedRoutePorts:0, TotalServiceKeys:0, AppInstanceLimit:0, IsoSegment:""}
		utilsMgrMock.MockFileData["./fixtures/config/test/spaces.yml"] = config.Spaces{Org:"test", Spaces:[]string{"space1", "space2", "space3", "space4"}, EnableDeleteSpaces:false}
		utilsMgrMock.MockFileData["./fixtures/config/test2/orgConfig.yml"] = config.OrgConfig{Org:"test2", BillingManagerGroup:"test2_billing_managers", ManagerGroup:"", AuditorGroup:"", BillingManager:config.UserMgmt{LDAPUsers:[]string(nil), Users:[]string(nil), SamlUsers:[]string(nil), LDAPGroup:"", LDAPGroups:[]string(nil)}, Manager:config.UserMgmt{LDAPUsers:[]string(nil), Users:[]string(nil), SamlUsers:[]string(nil), LDAPGroup:"", LDAPGroups:[]string(nil)}, Auditor:config.UserMgmt{LDAPUsers:[]string(nil), Users:[]string(nil), SamlUsers:[]string(nil), LDAPGroup:"", LDAPGroups:[]string(nil)}, PrivateDomains:[]string(nil), RemovePrivateDomains:false, EnableOrgQuota:true, MemoryLimit:10240, InstanceMemoryLimit:-1, TotalRoutes:10, TotalServices:-1, PaidServicePlansAllowed:true, RemoveUsers:false, TotalPrivateDomains:0, TotalReservedRoutePorts:0, TotalServiceKeys:0, AppInstanceLimit:0, DefaultIsoSegment:""}
		utilsMgrMock.MockFileData["./fixtures/config/test2/spaces.yml"] = config.Spaces{Org:"test2", Spaces:[]string{"space1a", "space2a", "space3a", "space4a"}, EnableDeleteSpaces:false}
		utilsMgrMock.MockFileData["./fixtures/space-defaults/ldap.yml"] = ldap.Config{Enabled:false, LdapHost:"", LdapPort:0, TLS:false, BindDN:"", BindPassword:"", UserSearchBase:"", UserNameAttribute:"", UserMailAttribute:"", UserObjectClass:"", GroupSearchBase:"", GroupAttribute:"", Origin:"ldap"}
		utilsMgrMock.MockFileData["./fixtures/space-defaults/orgs.yml"] = config.Orgs{Orgs:[]string{"test"}, EnableDeleteOrgs:false, ProtectedOrgs:[]string(nil)}
		utilsMgrMock.MockFileData["./fixtures/space-defaults/spaceDefaults.yml"] = config.SpaceConfig{Org:"", Space:"", Developer:config.UserMgmt{LDAPUsers:[]string{"default-ldap-user"}, Users:[]string{"default-user@test.com"}, SamlUsers:[]string(nil), LDAPGroup:"default-ldap-group", LDAPGroups:[]string(nil)}, Manager:config.UserMgmt{LDAPUsers:[]string{"default-ldap-user"}, Users:[]string{"default-user@test.com"}, SamlUsers:[]string(nil), LDAPGroup:"default-ldap-group", LDAPGroups:[]string(nil)}, Auditor:config.UserMgmt{LDAPUsers:[]string{"default-ldap-user"}, Users:[]string{"default-user@test.com"}, SamlUsers:[]string(nil), LDAPGroup:"default-ldap-group", LDAPGroups:[]string(nil)}, DeveloperGroup:"", ManagerGroup:"", AuditorGroup:"", AllowSSH:false, EnableSpaceQuota:false, MemoryLimit:0, InstanceMemoryLimit:0, TotalRoutes:0, TotalServices:0, PaidServicePlansAllowed:false, EnableSecurityGroup:false, SecurityGroupContents:"", RemoveUsers:false, TotalPrivateDomains:0, TotalReservedRoutePorts:0, TotalServiceKeys:0, AppInstanceLimit:0, IsoSegment:""}
		utilsMgrMock.MockFileData["./fixtures/space-defaults/test/orgConfig.yml"] = config.OrgConfig{Org:"test", BillingManagerGroup:"", ManagerGroup:"", AuditorGroup:"", BillingManager:config.UserMgmt{LDAPUsers:[]string{}, Users:[]string{}, SamlUsers:[]string(nil), LDAPGroup:"", LDAPGroups:[]string(nil)}, Manager:config.UserMgmt{LDAPUsers:[]string{}, Users:[]string{}, SamlUsers:[]string(nil), LDAPGroup:"", LDAPGroups:[]string(nil)}, Auditor:config.UserMgmt{LDAPUsers:[]string{}, Users:[]string{}, SamlUsers:[]string(nil), LDAPGroup:"", LDAPGroups:[]string(nil)}, PrivateDomains:[]string(nil), RemovePrivateDomains:false, EnableOrgQuota:false, MemoryLimit:10240, InstanceMemoryLimit:-1, TotalRoutes:10, TotalServices:-1, PaidServicePlansAllowed:true, RemoveUsers:true, TotalPrivateDomains:0, TotalReservedRoutePorts:0, TotalServiceKeys:0, AppInstanceLimit:0, DefaultIsoSegment:""}
		utilsMgrMock.MockFileData["./fixtures/space-defaults/test/space1/security-group.json"] = "[]"
		utilsMgrMock.MockFileData["./fixtures/space-defaults/test/space1/spaceConfig.yml"] = config.SpaceConfig{Org:"test", Space:"space1", Developer:config.UserMgmt{LDAPUsers:[]string{"space1-ldap-user"}, Users:[]string{"space-1-user@test.com"}, SamlUsers:[]string(nil), LDAPGroup:"space-1-ldap-group", LDAPGroups:[]string(nil)}, Manager:config.UserMgmt{LDAPUsers:[]string{"space1-ldap-user"}, Users:[]string{"space-1-user@test.com"}, SamlUsers:[]string(nil), LDAPGroup:"space-1-ldap-group", LDAPGroups:[]string(nil)}, Auditor:config.UserMgmt{LDAPUsers:[]string{"space1-ldap-user"}, Users:[]string{"space-1-user@test.com"}, SamlUsers:[]string(nil), LDAPGroup:"space-1-ldap-group", LDAPGroups:[]string(nil)}, DeveloperGroup:"", ManagerGroup:"", AuditorGroup:"", AllowSSH:false, EnableSpaceQuota:false, MemoryLimit:10240, InstanceMemoryLimit:-1, TotalRoutes:10, TotalServices:-1, PaidServicePlansAllowed:true, EnableSecurityGroup:false, SecurityGroupContents:"", RemoveUsers:true, TotalPrivateDomains:0, TotalReservedRoutePorts:0, TotalServiceKeys:0, AppInstanceLimit:0, IsoSegment:""}
		utilsMgrMock.MockFileData["./fixtures/space-defaults/test/spaces.yml"] = config.Spaces{Org:"test", Spaces:[]string{"space1"}, EnableDeleteSpaces:false}
		utilsMgrMock.MockFileData["./fixtures/user_config/ldap.yml"] = ldap.Config{Enabled:true, LdapHost:"127.0.0.1", LdapPort:10389, TLS:false, BindDN:"uid=admin,ou=system", BindPassword:"secret", UserSearchBase:"ou=users,dc=example,dc=com", UserNameAttribute:"uid", UserMailAttribute:"mail", UserObjectClass:"", GroupSearchBase:"ou=groups,dc=example,dc=com", GroupAttribute:"member", Origin:""}
		utilsMgrMock.MockFileData["./fixtures/user_config/orgs.yml"] = config.Orgs{Orgs:[]string{"test"}, EnableDeleteOrgs:false, ProtectedOrgs:[]string(nil)}
		utilsMgrMock.MockFileData["./fixtures/user_config/test/orgConfig.yml"] = config.OrgConfig{Org:"test", BillingManagerGroup:"", ManagerGroup:"", AuditorGroup:"", BillingManager:config.UserMgmt{LDAPUsers:[]string{"cwashburn1", "cwashburn2"}, Users:[]string{"cwashburn@testdomain.com", "cwashburn2@testdomain.com"}, SamlUsers:[]string(nil), LDAPGroup:"test_billing_managers", LDAPGroups:[]string{"test_billing_managers_2"}}, Manager:config.UserMgmt{LDAPUsers:[]string{"cwashburn1", "cwashburn2"}, Users:[]string{"cwashburn@testdomain.com", "cwashburn2@testdomain.com"}, SamlUsers:[]string(nil), LDAPGroup:"test_org_managers", LDAPGroups:[]string(nil)}, Auditor:config.UserMgmt{LDAPUsers:[]string{"cwashburn1", "cwashburn2"}, Users:[]string{"cwashburn@testdomain.com", "cwashburn2@testdomain.com"}, SamlUsers:[]string(nil), LDAPGroup:"", LDAPGroups:[]string{"test_org_auditors"}}, PrivateDomains:[]string(nil), RemovePrivateDomains:false, EnableOrgQuota:true, MemoryLimit:10240, InstanceMemoryLimit:-1, TotalRoutes:10, TotalServices:-1, PaidServicePlansAllowed:true, RemoveUsers:false, TotalPrivateDomains:0, TotalReservedRoutePorts:0, TotalServiceKeys:0, AppInstanceLimit:0, DefaultIsoSegment:""}
		utilsMgrMock.MockFileData["./fixtures/user_config/test/space1/spaceConfig.yml"] = config.SpaceConfig{Org:"test", Space:"space1", Developer:config.UserMgmt{LDAPUsers:[]string{"cwashburn1", "cwashburn2"}, Users:[]string{"cwashburn@testdomain.com", "cwashburn2@testdomain.com"}, SamlUsers:[]string(nil), LDAPGroup:"test_space1_developers", LDAPGroups:[]string(nil)}, Manager:config.UserMgmt{LDAPUsers:[]string{"cwashburn1", "cwashburn2"}, Users:[]string{"cwashburn@testdomain.com", "cwashburn2@testdomain.com"}, SamlUsers:[]string(nil), LDAPGroup:"test_space1_managers", LDAPGroups:[]string(nil)}, Auditor:config.UserMgmt{LDAPUsers:[]string{"cwashburn1", "cwashburn2"}, Users:[]string{"cwashburn@testdomain.com", "cwashburn2@testdomain.com"}, SamlUsers:[]string(nil), LDAPGroup:"test_space1_auditors", LDAPGroups:[]string(nil)}, DeveloperGroup:"", ManagerGroup:"", AuditorGroup:"", AllowSSH:true, EnableSpaceQuota:true, MemoryLimit:10240, InstanceMemoryLimit:-1, TotalRoutes:10, TotalServices:-1, PaidServicePlansAllowed:true, EnableSecurityGroup:false, SecurityGroupContents:"", RemoveUsers:false, TotalPrivateDomains:0, TotalReservedRoutePorts:0, TotalServiceKeys:0, AppInstanceLimit:0, IsoSegment:""}
		utilsMgrMock.MockFileData["./fixtures/user_config/test/spaces.yml"] = config.Spaces{Org:"test", Spaces:[]string{"space1", "space2", "space3", "space4"}, EnableDeleteSpaces:false}
		utilsMgrMock.MockFileData["./fixtures/user_config_multiple_groups/test/space1/spaceConfig.yml"] = config.SpaceConfig{Org:"test", Space:"space1", Developer:config.UserMgmt{LDAPUsers:[]string{"cwashburn1", "cwashburn2"}, Users:[]string{"cwashburn@testdomain.com", "cwashburn2@testdomain.com"}, SamlUsers:[]string(nil), LDAPGroup:"test_space1_developers", LDAPGroups:[]string(nil)}, Manager:config.UserMgmt{LDAPUsers:[]string{"cwashburn1", "cwashburn2"}, Users:[]string{"cwashburn@testdomain.com", "cwashburn2@testdomain.com"}, SamlUsers:[]string(nil), LDAPGroup:"test_space1_managers", LDAPGroups:[]string{"test_space1_managers_2"}}, Auditor:config.UserMgmt{LDAPUsers:[]string{"cwashburn1", "cwashburn2"}, Users:[]string{"cwashburn@testdomain.com", "cwashburn2@testdomain.com"}, SamlUsers:[]string(nil), LDAPGroup:"", LDAPGroups:[]string{"test_space1_auditors"}}, DeveloperGroup:"", ManagerGroup:"", AuditorGroup:"", AllowSSH:true, EnableSpaceQuota:true, MemoryLimit:10240, InstanceMemoryLimit:-1, TotalRoutes:10, TotalServices:-1, PaidServicePlansAllowed:true, EnableSecurityGroup:false, SecurityGroupContents:"", RemoveUsers:false, TotalPrivateDomains:0, TotalReservedRoutePorts:0, TotalServiceKeys:0, AppInstanceLimit:0, IsoSegment:""}
		utilsMgrMock.MockFileData["./fixtures/user_update/ldap.yml"] = ldap.Config{Enabled:false, LdapHost:"", LdapPort:0, TLS:false, BindDN:"", BindPassword:"", UserSearchBase:"", UserNameAttribute:"", UserMailAttribute:"", UserObjectClass:"", GroupSearchBase:"", GroupAttribute:"", Origin:"ldap"}
		utilsMgrMock.MockFileData["./fixtures/user_update/orgs.yml"] = config.Orgs{Orgs:[]string{"test"}, EnableDeleteOrgs:false, ProtectedOrgs:[]string(nil)}
		utilsMgrMock.MockFileData["./fixtures/user_update/spaceDefaults.yml"] = config.SpaceConfig{Org:"", Space:"", Developer:config.UserMgmt{LDAPUsers:[]string{"default-ldap-user"}, Users:[]string{"default-user@test.com"}, SamlUsers:[]string(nil), LDAPGroup:"default-ldap-group", LDAPGroups:[]string(nil)}, Manager:config.UserMgmt{LDAPUsers:[]string{"default-ldap-user"}, Users:[]string{"default-user@test.com"}, SamlUsers:[]string(nil), LDAPGroup:"default-ldap-group", LDAPGroups:[]string(nil)}, Auditor:config.UserMgmt{LDAPUsers:[]string{"default-ldap-user"}, Users:[]string{"default-user@test.com"}, SamlUsers:[]string(nil), LDAPGroup:"default-ldap-group", LDAPGroups:[]string(nil)}, DeveloperGroup:"", ManagerGroup:"", AuditorGroup:"", AllowSSH:false, EnableSpaceQuota:false, MemoryLimit:0, InstanceMemoryLimit:0, TotalRoutes:0, TotalServices:0, PaidServicePlansAllowed:false, EnableSecurityGroup:false, SecurityGroupContents:"", RemoveUsers:false, TotalPrivateDomains:0, TotalReservedRoutePorts:0, TotalServiceKeys:0, AppInstanceLimit:0, IsoSegment:""}
		utilsMgrMock.MockFileData["./fixtures/user_update/test/orgConfig.yml"] = config.OrgConfig{Org:"test", BillingManagerGroup:"", ManagerGroup:"", AuditorGroup:"", BillingManager:config.UserMgmt{LDAPUsers:[]string{}, Users:[]string{"C471E3DC50.1C9BFF9A5A"}, SamlUsers:[]string{}, LDAPGroup:"", LDAPGroups:[]string{}}, Manager:config.UserMgmt{LDAPUsers:[]string{}, Users:[]string{}, SamlUsers:[]string{}, LDAPGroup:"", LDAPGroups:[]string{}}, Auditor:config.UserMgmt{LDAPUsers:[]string{}, Users:[]string{}, SamlUsers:[]string{}, LDAPGroup:"", LDAPGroups:[]string{}}, PrivateDomains:[]string{}, RemovePrivateDomains:false, EnableOrgQuota:false, MemoryLimit:10240, InstanceMemoryLimit:-1, TotalRoutes:10, TotalServices:-1, PaidServicePlansAllowed:true, RemoveUsers:true, TotalPrivateDomains:0, TotalReservedRoutePorts:0, TotalServiceKeys:0, AppInstanceLimit:0, DefaultIsoSegment:""}
		utilsMgrMock.MockFileData["./fixtures/user_update/test/space1/security-group.json"] = "[]"
		utilsMgrMock.MockFileData["./fixtures/user_update/test/space1/spaceConfig.yml"] = config.SpaceConfig{Org:"test", Space:"space1", Developer:config.UserMgmt{LDAPUsers:[]string{"space1-ldap-user"}, Users:[]string{"space-1-user@test.com"}, SamlUsers:[]string(nil), LDAPGroup:"space-1-ldap-group", LDAPGroups:[]string(nil)}, Manager:config.UserMgmt{LDAPUsers:[]string{"space1-ldap-user"}, Users:[]string{"space-1-user@test.com"}, SamlUsers:[]string(nil), LDAPGroup:"space-1-ldap-group", LDAPGroups:[]string(nil)}, Auditor:config.UserMgmt{LDAPUsers:[]string{"space1-ldap-user"}, Users:[]string{"space-1-user@test.com"}, SamlUsers:[]string(nil), LDAPGroup:"space-1-ldap-group", LDAPGroups:[]string(nil)}, DeveloperGroup:"", ManagerGroup:"", AuditorGroup:"", AllowSSH:false, EnableSpaceQuota:false, MemoryLimit:10240, InstanceMemoryLimit:-1, TotalRoutes:10, TotalServices:-1, PaidServicePlansAllowed:true, EnableSecurityGroup:false, SecurityGroupContents:"", RemoveUsers:true, TotalPrivateDomains:0, TotalReservedRoutePorts:0, TotalServiceKeys:0, AppInstanceLimit:0, IsoSegment:""}
		utilsMgrMock.MockFileData["./fixtures/user_update/test/spaces.yml"] = config.Spaces{Org:"test", Spaces:[]string{"space1"}, EnableDeleteSpaces:false}

		return nil
	}
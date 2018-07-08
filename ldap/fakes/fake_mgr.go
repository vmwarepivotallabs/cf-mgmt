// This file was generated by counterfeiter
package fakes

import (
	"sync"

	ldapgo_ldap "github.com/go-ldap/ldap"
	"github.com/pivotalservices/cf-mgmt/config"
	"github.com/pivotalservices/cf-mgmt/ldap"
)

type FakeManager struct {
	GetUserIDsStub        func(groupName string) (users []ldap.User, err error)
	getUserIDsMutex       sync.RWMutex
	getUserIDsArgsForCall []struct {
		groupName string
	}
	getUserIDsReturns struct {
		result1 []ldap.User
		result2 error
	}
	GetUserStub        func(userID string) (*ldap.User, error)
	getUserMutex       sync.RWMutex
	getUserArgsForCall []struct {
		userID string
	}
	getUserReturns struct {
		result1 *ldap.User
		result2 error
	}
	GetLdapUserStub        func(userDN string) (*ldap.User, error)
	getLdapUserMutex       sync.RWMutex
	getLdapUserArgsForCall []struct {
		userDN string
	}
	getLdapUserReturns struct {
		result1 *ldap.User
		result2 error
	}
	LdapConnectionStub        func() (*ldapgo_ldap.Conn, error)
	ldapConnectionMutex       sync.RWMutex
	ldapConnectionArgsForCall []struct{}
	ldapConnectionReturns     struct {
		result1 *ldapgo_ldap.Conn
		result2 error
	}
	GetLdapUsersStub        func(groupNames []string, userList []string) ([]ldap.User, error)
	getLdapUsersMutex       sync.RWMutex
	getLdapUsersArgsForCall []struct {
		groupNames []string
		userList   []string
	}
	getLdapUsersReturns struct {
		result1 []ldap.User
		result2 error
	}
	LdapConfigStub        func() *config.LdapConfig
	ldapConfigMutex       sync.RWMutex
	ldapConfigArgsForCall []struct{}
	ldapConfigReturns     struct {
		result1 *config.LdapConfig
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManager) GetUserIDs(groupName string) (users []ldap.User, err error) {
	fake.getUserIDsMutex.Lock()
	fake.getUserIDsArgsForCall = append(fake.getUserIDsArgsForCall, struct {
		groupName string
	}{groupName})
	fake.recordInvocation("GetUserIDs", []interface{}{groupName})
	fake.getUserIDsMutex.Unlock()
	if fake.GetUserIDsStub != nil {
		return fake.GetUserIDsStub(groupName)
	} else {
		return fake.getUserIDsReturns.result1, fake.getUserIDsReturns.result2
	}
}

func (fake *FakeManager) GetUserIDsCallCount() int {
	fake.getUserIDsMutex.RLock()
	defer fake.getUserIDsMutex.RUnlock()
	return len(fake.getUserIDsArgsForCall)
}

func (fake *FakeManager) GetUserIDsArgsForCall(i int) string {
	fake.getUserIDsMutex.RLock()
	defer fake.getUserIDsMutex.RUnlock()
	return fake.getUserIDsArgsForCall[i].groupName
}

func (fake *FakeManager) GetUserIDsReturns(result1 []ldap.User, result2 error) {
	fake.GetUserIDsStub = nil
	fake.getUserIDsReturns = struct {
		result1 []ldap.User
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) GetUser(userID string) (*ldap.User, error) {
	fake.getUserMutex.Lock()
	fake.getUserArgsForCall = append(fake.getUserArgsForCall, struct {
		userID string
	}{userID})
	fake.recordInvocation("GetUser", []interface{}{userID})
	fake.getUserMutex.Unlock()
	if fake.GetUserStub != nil {
		return fake.GetUserStub(userID)
	} else {
		return fake.getUserReturns.result1, fake.getUserReturns.result2
	}
}

func (fake *FakeManager) GetUserCallCount() int {
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	return len(fake.getUserArgsForCall)
}

func (fake *FakeManager) GetUserArgsForCall(i int) string {
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	return fake.getUserArgsForCall[i].userID
}

func (fake *FakeManager) GetUserReturns(result1 *ldap.User, result2 error) {
	fake.GetUserStub = nil
	fake.getUserReturns = struct {
		result1 *ldap.User
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) GetLdapUser(userDN string) (*ldap.User, error) {
	fake.getLdapUserMutex.Lock()
	fake.getLdapUserArgsForCall = append(fake.getLdapUserArgsForCall, struct {
		userDN string
	}{userDN})
	fake.recordInvocation("GetLdapUser", []interface{}{userDN})
	fake.getLdapUserMutex.Unlock()
	if fake.GetLdapUserStub != nil {
		return fake.GetLdapUserStub(userDN)
	} else {
		return fake.getLdapUserReturns.result1, fake.getLdapUserReturns.result2
	}
}

func (fake *FakeManager) GetLdapUserCallCount() int {
	fake.getLdapUserMutex.RLock()
	defer fake.getLdapUserMutex.RUnlock()
	return len(fake.getLdapUserArgsForCall)
}

func (fake *FakeManager) GetLdapUserArgsForCall(i int) string {
	fake.getLdapUserMutex.RLock()
	defer fake.getLdapUserMutex.RUnlock()
	return fake.getLdapUserArgsForCall[i].userDN
}

func (fake *FakeManager) GetLdapUserReturns(result1 *ldap.User, result2 error) {
	fake.GetLdapUserStub = nil
	fake.getLdapUserReturns = struct {
		result1 *ldap.User
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) LdapConnection() (*ldapgo_ldap.Conn, error) {
	fake.ldapConnectionMutex.Lock()
	fake.ldapConnectionArgsForCall = append(fake.ldapConnectionArgsForCall, struct{}{})
	fake.recordInvocation("LdapConnection", []interface{}{})
	fake.ldapConnectionMutex.Unlock()
	if fake.LdapConnectionStub != nil {
		return fake.LdapConnectionStub()
	} else {
		return fake.ldapConnectionReturns.result1, fake.ldapConnectionReturns.result2
	}
}

func (fake *FakeManager) LdapConnectionCallCount() int {
	fake.ldapConnectionMutex.RLock()
	defer fake.ldapConnectionMutex.RUnlock()
	return len(fake.ldapConnectionArgsForCall)
}

func (fake *FakeManager) LdapConnectionReturns(result1 *ldapgo_ldap.Conn, result2 error) {
	fake.LdapConnectionStub = nil
	fake.ldapConnectionReturns = struct {
		result1 *ldapgo_ldap.Conn
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) GetLdapUsers(groupNames []string, userList []string) ([]ldap.User, error) {
	var groupNamesCopy []string
	if groupNames != nil {
		groupNamesCopy = make([]string, len(groupNames))
		copy(groupNamesCopy, groupNames)
	}
	var userListCopy []string
	if userList != nil {
		userListCopy = make([]string, len(userList))
		copy(userListCopy, userList)
	}
	fake.getLdapUsersMutex.Lock()
	fake.getLdapUsersArgsForCall = append(fake.getLdapUsersArgsForCall, struct {
		groupNames []string
		userList   []string
	}{groupNamesCopy, userListCopy})
	fake.recordInvocation("GetLdapUsers", []interface{}{groupNamesCopy, userListCopy})
	fake.getLdapUsersMutex.Unlock()
	if fake.GetLdapUsersStub != nil {
		return fake.GetLdapUsersStub(groupNames, userList)
	} else {
		return fake.getLdapUsersReturns.result1, fake.getLdapUsersReturns.result2
	}
}

func (fake *FakeManager) GetLdapUsersCallCount() int {
	fake.getLdapUsersMutex.RLock()
	defer fake.getLdapUsersMutex.RUnlock()
	return len(fake.getLdapUsersArgsForCall)
}

func (fake *FakeManager) GetLdapUsersArgsForCall(i int) ([]string, []string) {
	fake.getLdapUsersMutex.RLock()
	defer fake.getLdapUsersMutex.RUnlock()
	return fake.getLdapUsersArgsForCall[i].groupNames, fake.getLdapUsersArgsForCall[i].userList
}

func (fake *FakeManager) GetLdapUsersReturns(result1 []ldap.User, result2 error) {
	fake.GetLdapUsersStub = nil
	fake.getLdapUsersReturns = struct {
		result1 []ldap.User
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) LdapConfig() *config.LdapConfig {
	fake.ldapConfigMutex.Lock()
	fake.ldapConfigArgsForCall = append(fake.ldapConfigArgsForCall, struct{}{})
	fake.recordInvocation("LdapConfig", []interface{}{})
	fake.ldapConfigMutex.Unlock()
	if fake.LdapConfigStub != nil {
		return fake.LdapConfigStub()
	} else {
		return fake.ldapConfigReturns.result1
	}
}

func (fake *FakeManager) LdapConfigCallCount() int {
	fake.ldapConfigMutex.RLock()
	defer fake.ldapConfigMutex.RUnlock()
	return len(fake.ldapConfigArgsForCall)
}

func (fake *FakeManager) LdapConfigReturns(result1 *config.LdapConfig) {
	fake.LdapConfigStub = nil
	fake.ldapConfigReturns = struct {
		result1 *config.LdapConfig
	}{result1}
}

func (fake *FakeManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getUserIDsMutex.RLock()
	defer fake.getUserIDsMutex.RUnlock()
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	fake.getLdapUserMutex.RLock()
	defer fake.getLdapUserMutex.RUnlock()
	fake.ldapConnectionMutex.RLock()
	defer fake.ldapConnectionMutex.RUnlock()
	fake.getLdapUsersMutex.RLock()
	defer fake.getLdapUsersMutex.RUnlock()
	fake.ldapConfigMutex.RLock()
	defer fake.ldapConfigMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeManager) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ ldap.Manager = new(FakeManager)
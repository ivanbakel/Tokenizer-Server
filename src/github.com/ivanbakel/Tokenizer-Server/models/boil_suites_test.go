package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Users", testUsers)
	t.Run("Organisations", testOrganisations)
	t.Run("UserTokens", testUserTokens)
	t.Run("Tokens", testTokens)
}

func TestDelete(t *testing.T) {
	t.Run("Users", testUsersDelete)
	t.Run("Organisations", testOrganisationsDelete)
	t.Run("UserTokens", testUserTokensDelete)
	t.Run("Tokens", testTokensDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Users", testUsersQueryDeleteAll)
	t.Run("Organisations", testOrganisationsQueryDeleteAll)
	t.Run("UserTokens", testUserTokensQueryDeleteAll)
	t.Run("Tokens", testTokensQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Users", testUsersSliceDeleteAll)
	t.Run("Organisations", testOrganisationsSliceDeleteAll)
	t.Run("UserTokens", testUserTokensSliceDeleteAll)
	t.Run("Tokens", testTokensSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Users", testUsersExists)
	t.Run("Organisations", testOrganisationsExists)
	t.Run("UserTokens", testUserTokensExists)
	t.Run("Tokens", testTokensExists)
}

func TestFind(t *testing.T) {
	t.Run("Users", testUsersFind)
	t.Run("Organisations", testOrganisationsFind)
	t.Run("UserTokens", testUserTokensFind)
	t.Run("Tokens", testTokensFind)
}

func TestBind(t *testing.T) {
	t.Run("Users", testUsersBind)
	t.Run("Organisations", testOrganisationsBind)
	t.Run("UserTokens", testUserTokensBind)
	t.Run("Tokens", testTokensBind)
}

func TestOne(t *testing.T) {
	t.Run("Users", testUsersOne)
	t.Run("Organisations", testOrganisationsOne)
	t.Run("UserTokens", testUserTokensOne)
	t.Run("Tokens", testTokensOne)
}

func TestAll(t *testing.T) {
	t.Run("Users", testUsersAll)
	t.Run("Organisations", testOrganisationsAll)
	t.Run("UserTokens", testUserTokensAll)
	t.Run("Tokens", testTokensAll)
}

func TestCount(t *testing.T) {
	t.Run("Users", testUsersCount)
	t.Run("Organisations", testOrganisationsCount)
	t.Run("UserTokens", testUserTokensCount)
	t.Run("Tokens", testTokensCount)
}

func TestHooks(t *testing.T) {
	t.Run("Users", testUsersHooks)
	t.Run("Organisations", testOrganisationsHooks)
	t.Run("UserTokens", testUserTokensHooks)
	t.Run("Tokens", testTokensHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
	t.Run("Organisations", testOrganisationsInsert)
	t.Run("Organisations", testOrganisationsInsertWhitelist)
	t.Run("UserTokens", testUserTokensInsert)
	t.Run("UserTokens", testUserTokensInsertWhitelist)
	t.Run("Tokens", testTokensInsert)
	t.Run("Tokens", testTokensInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("UserToOrganisationUsingOrg", testUserToOneOrganisationUsingOrg)
	t.Run("UserTokenToUserUsingUser", testUserTokenToOneUserUsingUser)
	t.Run("UserTokenToOrganisationUsingOrg", testUserTokenToOneOrganisationUsingOrg)
	t.Run("TokenToOrganisationUsingOrg", testTokenToOneOrganisationUsingOrg)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("UserToUserTokens", testUserToManyUserTokens)
	t.Run("OrganisationToOrgUsers", testOrganisationToManyOrgUsers)
	t.Run("OrganisationToOrgUserTokens", testOrganisationToManyOrgUserTokens)
	t.Run("OrganisationToOrgTokens", testOrganisationToManyOrgTokens)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("UserToOrganisationUsingOrg", testUserToOneSetOpOrganisationUsingOrg)
	t.Run("UserTokenToUserUsingUser", testUserTokenToOneSetOpUserUsingUser)
	t.Run("UserTokenToOrganisationUsingOrg", testUserTokenToOneSetOpOrganisationUsingOrg)
	t.Run("TokenToOrganisationUsingOrg", testTokenToOneSetOpOrganisationUsingOrg)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("UserToOrganisationUsingOrg", testUserToOneRemoveOpOrganisationUsingOrg)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("UserToUserTokens", testUserToManyAddOpUserTokens)
	t.Run("OrganisationToOrgUsers", testOrganisationToManyAddOpOrgUsers)
	t.Run("OrganisationToOrgUserTokens", testOrganisationToManyAddOpOrgUserTokens)
	t.Run("OrganisationToOrgTokens", testOrganisationToManyAddOpOrgTokens)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("OrganisationToOrgUsers", testOrganisationToManySetOpOrgUsers)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("OrganisationToOrgUsers", testOrganisationToManyRemoveOpOrgUsers)
}

func TestReload(t *testing.T) {
	t.Run("Users", testUsersReload)
	t.Run("Organisations", testOrganisationsReload)
	t.Run("UserTokens", testUserTokensReload)
	t.Run("Tokens", testTokensReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Users", testUsersReloadAll)
	t.Run("Organisations", testOrganisationsReloadAll)
	t.Run("UserTokens", testUserTokensReloadAll)
	t.Run("Tokens", testTokensReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Users", testUsersSelect)
	t.Run("Organisations", testOrganisationsSelect)
	t.Run("UserTokens", testUserTokensSelect)
	t.Run("Tokens", testTokensSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Users", testUsersUpdate)
	t.Run("Organisations", testOrganisationsUpdate)
	t.Run("UserTokens", testUserTokensUpdate)
	t.Run("Tokens", testTokensUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Users", testUsersSliceUpdateAll)
	t.Run("Organisations", testOrganisationsSliceUpdateAll)
	t.Run("UserTokens", testUserTokensSliceUpdateAll)
	t.Run("Tokens", testTokensSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("Users", testUsersUpsert)
	t.Run("Organisations", testOrganisationsUpsert)
	t.Run("UserTokens", testUserTokensUpsert)
	t.Run("Tokens", testTokensUpsert)
}

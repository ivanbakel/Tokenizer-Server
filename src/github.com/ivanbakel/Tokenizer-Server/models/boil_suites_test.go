package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Tokens", testTokens)
	t.Run("UserTokens", testUserTokens)
	t.Run("Organisations", testOrganisations)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Tokens", testTokensDelete)
	t.Run("UserTokens", testUserTokensDelete)
	t.Run("Organisations", testOrganisationsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Tokens", testTokensQueryDeleteAll)
	t.Run("UserTokens", testUserTokensQueryDeleteAll)
	t.Run("Organisations", testOrganisationsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Tokens", testTokensSliceDeleteAll)
	t.Run("UserTokens", testUserTokensSliceDeleteAll)
	t.Run("Organisations", testOrganisationsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Tokens", testTokensExists)
	t.Run("UserTokens", testUserTokensExists)
	t.Run("Organisations", testOrganisationsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Tokens", testTokensFind)
	t.Run("UserTokens", testUserTokensFind)
	t.Run("Organisations", testOrganisationsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Tokens", testTokensBind)
	t.Run("UserTokens", testUserTokensBind)
	t.Run("Organisations", testOrganisationsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Tokens", testTokensOne)
	t.Run("UserTokens", testUserTokensOne)
	t.Run("Organisations", testOrganisationsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Tokens", testTokensAll)
	t.Run("UserTokens", testUserTokensAll)
	t.Run("Organisations", testOrganisationsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Tokens", testTokensCount)
	t.Run("UserTokens", testUserTokensCount)
	t.Run("Organisations", testOrganisationsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Tokens", testTokensHooks)
	t.Run("UserTokens", testUserTokensHooks)
	t.Run("Organisations", testOrganisationsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Tokens", testTokensInsert)
	t.Run("Tokens", testTokensInsertWhitelist)
	t.Run("UserTokens", testUserTokensInsert)
	t.Run("UserTokens", testUserTokensInsertWhitelist)
	t.Run("Organisations", testOrganisationsInsert)
	t.Run("Organisations", testOrganisationsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("TokenToOrganisationUsingOrg", testTokenToOneOrganisationUsingOrg)
	t.Run("UserTokenToUserUsingUser", testUserTokenToOneUserUsingUser)
	t.Run("UserTokenToTokenUsingToken", testUserTokenToOneTokenUsingToken)
	t.Run("UserToOrganisationUsingOrg", testUserToOneOrganisationUsingOrg)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("TokenToUserTokens", testTokenToManyUserTokens)
	t.Run("OrganisationToOrgTokens", testOrganisationToManyOrgTokens)
	t.Run("OrganisationToOrgUsers", testOrganisationToManyOrgUsers)
	t.Run("UserToUserTokens", testUserToManyUserTokens)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("TokenToOrganisationUsingOrg", testTokenToOneSetOpOrganisationUsingOrg)
	t.Run("UserTokenToUserUsingUser", testUserTokenToOneSetOpUserUsingUser)
	t.Run("UserTokenToTokenUsingToken", testUserTokenToOneSetOpTokenUsingToken)
	t.Run("UserToOrganisationUsingOrg", testUserToOneSetOpOrganisationUsingOrg)
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
	t.Run("TokenToUserTokens", testTokenToManyAddOpUserTokens)
	t.Run("OrganisationToOrgTokens", testOrganisationToManyAddOpOrgTokens)
	t.Run("OrganisationToOrgUsers", testOrganisationToManyAddOpOrgUsers)
	t.Run("UserToUserTokens", testUserToManyAddOpUserTokens)
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
	t.Run("Tokens", testTokensReload)
	t.Run("UserTokens", testUserTokensReload)
	t.Run("Organisations", testOrganisationsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Tokens", testTokensReloadAll)
	t.Run("UserTokens", testUserTokensReloadAll)
	t.Run("Organisations", testOrganisationsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Tokens", testTokensSelect)
	t.Run("UserTokens", testUserTokensSelect)
	t.Run("Organisations", testOrganisationsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Tokens", testTokensUpdate)
	t.Run("UserTokens", testUserTokensUpdate)
	t.Run("Organisations", testOrganisationsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Tokens", testTokensSliceUpdateAll)
	t.Run("UserTokens", testUserTokensSliceUpdateAll)
	t.Run("Organisations", testOrganisationsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("Tokens", testTokensUpsert)
	t.Run("UserTokens", testUserTokensUpsert)
	t.Run("Organisations", testOrganisationsUpsert)
	t.Run("Users", testUsersUpsert)
}

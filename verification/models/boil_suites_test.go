// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("VerificationConfigs", testVerificationConfigs)
	t.Run("VerificationSessions", testVerificationSessions)
	t.Run("VerifiedUsers", testVerifiedUsers)
}

func TestDelete(t *testing.T) {
	t.Run("VerificationConfigs", testVerificationConfigsDelete)
	t.Run("VerificationSessions", testVerificationSessionsDelete)
	t.Run("VerifiedUsers", testVerifiedUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("VerificationConfigs", testVerificationConfigsQueryDeleteAll)
	t.Run("VerificationSessions", testVerificationSessionsQueryDeleteAll)
	t.Run("VerifiedUsers", testVerifiedUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("VerificationConfigs", testVerificationConfigsSliceDeleteAll)
	t.Run("VerificationSessions", testVerificationSessionsSliceDeleteAll)
	t.Run("VerifiedUsers", testVerifiedUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("VerificationConfigs", testVerificationConfigsExists)
	t.Run("VerificationSessions", testVerificationSessionsExists)
	t.Run("VerifiedUsers", testVerifiedUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("VerificationConfigs", testVerificationConfigsFind)
	t.Run("VerificationSessions", testVerificationSessionsFind)
	t.Run("VerifiedUsers", testVerifiedUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("VerificationConfigs", testVerificationConfigsBind)
	t.Run("VerificationSessions", testVerificationSessionsBind)
	t.Run("VerifiedUsers", testVerifiedUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("VerificationConfigs", testVerificationConfigsOne)
	t.Run("VerificationSessions", testVerificationSessionsOne)
	t.Run("VerifiedUsers", testVerifiedUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("VerificationConfigs", testVerificationConfigsAll)
	t.Run("VerificationSessions", testVerificationSessionsAll)
	t.Run("VerifiedUsers", testVerifiedUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("VerificationConfigs", testVerificationConfigsCount)
	t.Run("VerificationSessions", testVerificationSessionsCount)
	t.Run("VerifiedUsers", testVerifiedUsersCount)
}

func TestInsert(t *testing.T) {
	t.Run("VerificationConfigs", testVerificationConfigsInsert)
	t.Run("VerificationConfigs", testVerificationConfigsInsertWhitelist)
	t.Run("VerificationSessions", testVerificationSessionsInsert)
	t.Run("VerificationSessions", testVerificationSessionsInsertWhitelist)
	t.Run("VerifiedUsers", testVerifiedUsersInsert)
	t.Run("VerifiedUsers", testVerifiedUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("VerificationConfigs", testVerificationConfigsReload)
	t.Run("VerificationSessions", testVerificationSessionsReload)
	t.Run("VerifiedUsers", testVerifiedUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("VerificationConfigs", testVerificationConfigsReloadAll)
	t.Run("VerificationSessions", testVerificationSessionsReloadAll)
	t.Run("VerifiedUsers", testVerifiedUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("VerificationConfigs", testVerificationConfigsSelect)
	t.Run("VerificationSessions", testVerificationSessionsSelect)
	t.Run("VerifiedUsers", testVerifiedUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("VerificationConfigs", testVerificationConfigsUpdate)
	t.Run("VerificationSessions", testVerificationSessionsUpdate)
	t.Run("VerifiedUsers", testVerifiedUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("VerificationConfigs", testVerificationConfigsSliceUpdateAll)
	t.Run("VerificationSessions", testVerificationSessionsSliceUpdateAll)
	t.Run("VerifiedUsers", testVerifiedUsersSliceUpdateAll)
}

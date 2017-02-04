package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testUserTokens(t *testing.T) {
	t.Parallel()

	query := UserTokens(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testUserTokensDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userToken := &UserToken{}
	if err = randomize.Struct(seed, userToken, userTokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userToken.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = userToken.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := UserTokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserTokensQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userToken := &UserToken{}
	if err = randomize.Struct(seed, userToken, userTokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userToken.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = UserTokens(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := UserTokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserTokensSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userToken := &UserToken{}
	if err = randomize.Struct(seed, userToken, userTokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userToken.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UserTokenSlice{userToken}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := UserTokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testUserTokensExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userToken := &UserToken{}
	if err = randomize.Struct(seed, userToken, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userToken.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := UserTokenExists(tx, userToken.UserID, userToken.OrgID)
	if err != nil {
		t.Errorf("Unable to check if UserToken exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserTokenExistsG to return true, but got false.")
	}
}
func testUserTokensFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userToken := &UserToken{}
	if err = randomize.Struct(seed, userToken, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userToken.Insert(tx); err != nil {
		t.Error(err)
	}

	userTokenFound, err := FindUserToken(tx, userToken.UserID, userToken.OrgID)
	if err != nil {
		t.Error(err)
	}

	if userTokenFound == nil {
		t.Error("want a record, got nil")
	}
}
func testUserTokensBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userToken := &UserToken{}
	if err = randomize.Struct(seed, userToken, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userToken.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = UserTokens(tx).Bind(userToken); err != nil {
		t.Error(err)
	}
}

func testUserTokensOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userToken := &UserToken{}
	if err = randomize.Struct(seed, userToken, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userToken.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := UserTokens(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserTokensAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userTokenOne := &UserToken{}
	userTokenTwo := &UserToken{}
	if err = randomize.Struct(seed, userTokenOne, userTokenDBTypes, false, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}
	if err = randomize.Struct(seed, userTokenTwo, userTokenDBTypes, false, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userTokenOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = userTokenTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := UserTokens(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserTokensCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userTokenOne := &UserToken{}
	userTokenTwo := &UserToken{}
	if err = randomize.Struct(seed, userTokenOne, userTokenDBTypes, false, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}
	if err = randomize.Struct(seed, userTokenTwo, userTokenDBTypes, false, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userTokenOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = userTokenTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UserTokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func userTokenBeforeInsertHook(e boil.Executor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func userTokenAfterInsertHook(e boil.Executor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func userTokenAfterSelectHook(e boil.Executor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func userTokenBeforeUpdateHook(e boil.Executor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func userTokenAfterUpdateHook(e boil.Executor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func userTokenBeforeDeleteHook(e boil.Executor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func userTokenAfterDeleteHook(e boil.Executor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func userTokenBeforeUpsertHook(e boil.Executor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func userTokenAfterUpsertHook(e boil.Executor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func testUserTokensHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &UserToken{}
	o := &UserToken{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userTokenDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserToken object: %s", err)
	}

	AddUserTokenHook(boil.BeforeInsertHook, userTokenBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userTokenBeforeInsertHooks = []UserTokenHook{}

	AddUserTokenHook(boil.AfterInsertHook, userTokenAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userTokenAfterInsertHooks = []UserTokenHook{}

	AddUserTokenHook(boil.AfterSelectHook, userTokenAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userTokenAfterSelectHooks = []UserTokenHook{}

	AddUserTokenHook(boil.BeforeUpdateHook, userTokenBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userTokenBeforeUpdateHooks = []UserTokenHook{}

	AddUserTokenHook(boil.AfterUpdateHook, userTokenAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userTokenAfterUpdateHooks = []UserTokenHook{}

	AddUserTokenHook(boil.BeforeDeleteHook, userTokenBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userTokenBeforeDeleteHooks = []UserTokenHook{}

	AddUserTokenHook(boil.AfterDeleteHook, userTokenAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userTokenAfterDeleteHooks = []UserTokenHook{}

	AddUserTokenHook(boil.BeforeUpsertHook, userTokenBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userTokenBeforeUpsertHooks = []UserTokenHook{}

	AddUserTokenHook(boil.AfterUpsertHook, userTokenAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userTokenAfterUpsertHooks = []UserTokenHook{}
}
func testUserTokensInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userToken := &UserToken{}
	if err = randomize.Struct(seed, userToken, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userToken.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UserTokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserTokensInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userToken := &UserToken{}
	if err = randomize.Struct(seed, userToken, userTokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userToken.Insert(tx, userTokenColumns...); err != nil {
		t.Error(err)
	}

	count, err := UserTokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserTokenToOneUserUsingUser(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local UserToken
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.User(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UserTokenSlice{&local}
	if err = local.L.LoadUser(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUserTokenToOneOrganisationUsingOrg(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local UserToken
	var foreign Organisation

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, organisationDBTypes, true, organisationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.OrgID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Org(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UserTokenSlice{&local}
	if err = local.L.LoadOrg(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Org == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Org = nil
	if err = local.L.LoadOrg(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Org == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUserTokenToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a UserToken
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userTokenDBTypes, false, strmangle.SetComplement(userTokenPrimaryKeyColumns, userTokenColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserTokens[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		if exists, err := UserTokenExists(tx, a.UserID, a.OrgID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testUserTokenToOneSetOpOrganisationUsingOrg(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a UserToken
	var b, c Organisation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userTokenDBTypes, false, strmangle.SetComplement(userTokenPrimaryKeyColumns, userTokenColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, organisationDBTypes, false, strmangle.SetComplement(organisationPrimaryKeyColumns, organisationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, organisationDBTypes, false, strmangle.SetComplement(organisationPrimaryKeyColumns, organisationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Organisation{&b, &c} {
		err = a.SetOrg(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Org != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OrgUserTokens[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OrgID != x.ID {
			t.Error("foreign key was wrong value", a.OrgID)
		}

		if exists, err := UserTokenExists(tx, a.UserID, a.OrgID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testUserTokensReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userToken := &UserToken{}
	if err = randomize.Struct(seed, userToken, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userToken.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = userToken.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testUserTokensReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userToken := &UserToken{}
	if err = randomize.Struct(seed, userToken, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userToken.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UserTokenSlice{userToken}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testUserTokensSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userToken := &UserToken{}
	if err = randomize.Struct(seed, userToken, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userToken.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := UserTokens(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userTokenDBTypes = map[string]string{`Number`: `smallint`, `OrgID`: `uuid`, `UserID`: `uuid`}
	_                = bytes.MinRead
)

func testUserTokensUpdate(t *testing.T) {
	t.Parallel()

	if len(userTokenColumns) == len(userTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	userToken := &UserToken{}
	if err = randomize.Struct(seed, userToken, userTokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userToken.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UserTokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, userToken, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	if err = userToken.Update(tx); err != nil {
		t.Error(err)
	}
}

func testUserTokensSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userTokenColumns) == len(userTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	userToken := &UserToken{}
	if err = randomize.Struct(seed, userToken, userTokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userToken.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UserTokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, userToken, userTokenDBTypes, true, userTokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userTokenColumns, userTokenPrimaryKeyColumns) {
		fields = userTokenColumns
	} else {
		fields = strmangle.SetComplement(
			userTokenColumns,
			userTokenPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(userToken))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := UserTokenSlice{userToken}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testUserTokensUpsert(t *testing.T) {
	t.Parallel()

	if len(userTokenColumns) == len(userTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	userToken := UserToken{}
	if err = randomize.Struct(seed, &userToken, userTokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userToken.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert UserToken: %s", err)
	}

	count, err := UserTokens(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &userToken, userTokenDBTypes, false, userTokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	if err = userToken.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert UserToken: %s", err)
	}

	count, err = UserTokens(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

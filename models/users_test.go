package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testUsers(t *testing.T) {
	t.Parallel()

	query := Users(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testUsersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = user.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUsersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Users(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUsersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UserSlice{user}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testUsersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := UserExists(tx, user.ID)
	if err != nil {
		t.Errorf("Unable to check if User exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserExistsG to return true, but got false.")
	}
}
func testUsersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	userFound, err := FindUser(tx, user.ID)
	if err != nil {
		t.Error(err)
	}

	if userFound == nil {
		t.Error("want a record, got nil")
	}
}
func testUsersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Users(tx).Bind(user); err != nil {
		t.Error(err)
	}
}

func testUsersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Users(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUsersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userOne := &User{}
	userTwo := &User{}
	if err = randomize.Struct(seed, userOne, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}
	if err = randomize.Struct(seed, userTwo, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = userTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Users(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUsersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userOne := &User{}
	userTwo := &User{}
	if err = randomize.Struct(seed, userOne, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}
	if err = randomize.Struct(seed, userTwo, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = userTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func userBeforeInsertHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userAfterInsertHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userAfterSelectHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userBeforeUpdateHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userAfterUpdateHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userBeforeDeleteHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userAfterDeleteHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userBeforeUpsertHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userAfterUpsertHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func testUsersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &User{}
	o := &User{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userDBTypes, false); err != nil {
		t.Errorf("Unable to randomize User object: %s", err)
	}

	AddUserHook(boil.BeforeInsertHook, userBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userBeforeInsertHooks = []UserHook{}

	AddUserHook(boil.AfterInsertHook, userAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userAfterInsertHooks = []UserHook{}

	AddUserHook(boil.AfterSelectHook, userAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userAfterSelectHooks = []UserHook{}

	AddUserHook(boil.BeforeUpdateHook, userBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userBeforeUpdateHooks = []UserHook{}

	AddUserHook(boil.AfterUpdateHook, userAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userAfterUpdateHooks = []UserHook{}

	AddUserHook(boil.BeforeDeleteHook, userBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userBeforeDeleteHooks = []UserHook{}

	AddUserHook(boil.AfterDeleteHook, userAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userAfterDeleteHooks = []UserHook{}

	AddUserHook(boil.BeforeUpsertHook, userBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userBeforeUpsertHooks = []UserHook{}

	AddUserHook(boil.AfterUpsertHook, userAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userAfterUpsertHooks = []UserHook{}
}
func testUsersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUsersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx, userColumns...); err != nil {
		t.Error(err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserToManyUserTokens(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a User
	var b, c UserToken

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, userTokenDBTypes, false, userTokenColumnsWithDefault...)
	randomize.Struct(seed, &c, userTokenDBTypes, false, userTokenColumnsWithDefault...)

	b.UserID = a.ID
	c.UserID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	userToken, err := a.UserTokens(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range userToken {
		if v.UserID == b.UserID {
			bFound = true
		}
		if v.UserID == c.UserID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := UserSlice{&a}
	if err = a.L.LoadUserTokens(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserTokens); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserTokens = nil
	if err = a.L.LoadUserTokens(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserTokens); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", userToken)
	}
}

func testUserToManyAddOpUserTokens(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a User
	var b, c, d, e UserToken

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserToken{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userTokenDBTypes, false, strmangle.SetComplement(userTokenPrimaryKeyColumns, userTokenColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*UserToken{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserTokens(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.UserID {
			t.Error("foreign key was wrong value", a.ID, first.UserID)
		}
		if a.ID != second.UserID {
			t.Error("foreign key was wrong value", a.ID, second.UserID)
		}

		if first.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserTokens[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserTokens[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserTokens(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testUserToOneOrganisationUsingOrg(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local User
	var foreign Organisation

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, organisationDBTypes, true, organisationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	local.OrgID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.OrgID.String = foreign.ID
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

	slice := UserSlice{&local}
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

func testUserToOneSetOpOrganisationUsingOrg(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a User
	var b, c Organisation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
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

		if x.R.OrgUsers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OrgID.String != x.ID {
			t.Error("foreign key was wrong value", a.OrgID.String)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OrgID.String))
		reflect.Indirect(reflect.ValueOf(&a.OrgID.String)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OrgID.String != x.ID {
			t.Error("foreign key was wrong value", a.OrgID.String, x.ID)
		}
	}
}

func testUserToOneRemoveOpOrganisationUsingOrg(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a User
	var b Organisation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, organisationDBTypes, false, strmangle.SetComplement(organisationPrimaryKeyColumns, organisationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetOrg(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveOrg(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Org(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Org != nil {
		t.Error("R struct entry should be nil")
	}

	if a.OrgID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.OrgUsers) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testUsersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = user.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testUsersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UserSlice{user}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testUsersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Users(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userDBTypes = map[string]string{`FacebookID`: `character varying`, `ID`: `uuid`, `OrgID`: `uuid`}
	_           = bytes.MinRead
)

func testUsersUpdate(t *testing.T) {
	t.Parallel()

	if len(userColumns) == len(userPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err = user.Update(tx); err != nil {
		t.Error(err)
	}
}

func testUsersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userColumns) == len(userPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, user, userDBTypes, true, userPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userColumns, userPrimaryKeyColumns) {
		fields = userColumns
	} else {
		fields = strmangle.SetComplement(
			userColumns,
			userPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(user))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := UserSlice{user}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testUsersUpsert(t *testing.T) {
	t.Parallel()

	if len(userColumns) == len(userPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	user := User{}
	if err = randomize.Struct(seed, &user, userDBTypes, true); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert User: %s", err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &user, userDBTypes, false, userPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err = user.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert User: %s", err)
	}

	count, err = Users(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

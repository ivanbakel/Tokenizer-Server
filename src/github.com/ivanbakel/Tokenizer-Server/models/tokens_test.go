package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testTokens(t *testing.T) {
	t.Parallel()

	query := Tokens(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testTokensDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = token.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTokensQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Tokens(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTokensSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TokenSlice{token}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testTokensExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true, tokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := TokenExists(tx, token.ID)
	if err != nil {
		t.Errorf("Unable to check if Token exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TokenExistsG to return true, but got false.")
	}
}
func testTokensFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true, tokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	tokenFound, err := FindToken(tx, token.ID)
	if err != nil {
		t.Error(err)
	}

	if tokenFound == nil {
		t.Error("want a record, got nil")
	}
}
func testTokensBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true, tokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Tokens(tx).Bind(token); err != nil {
		t.Error(err)
	}
}

func testTokensOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true, tokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Tokens(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTokensAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tokenOne := &Token{}
	tokenTwo := &Token{}
	if err = randomize.Struct(seed, tokenOne, tokenDBTypes, false, tokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}
	if err = randomize.Struct(seed, tokenTwo, tokenDBTypes, false, tokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tokenOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = tokenTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Tokens(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTokensCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tokenOne := &Token{}
	tokenTwo := &Token{}
	if err = randomize.Struct(seed, tokenOne, tokenDBTypes, false, tokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}
	if err = randomize.Struct(seed, tokenTwo, tokenDBTypes, false, tokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tokenOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = tokenTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func tokenBeforeInsertHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func tokenAfterInsertHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func tokenAfterSelectHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func tokenBeforeUpdateHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func tokenAfterUpdateHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func tokenBeforeDeleteHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func tokenAfterDeleteHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func tokenBeforeUpsertHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func tokenAfterUpsertHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func testTokensHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Token{}
	o := &Token{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, tokenDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Token object: %s", err)
	}

	AddTokenHook(boil.BeforeInsertHook, tokenBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	tokenBeforeInsertHooks = []TokenHook{}

	AddTokenHook(boil.AfterInsertHook, tokenAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	tokenAfterInsertHooks = []TokenHook{}

	AddTokenHook(boil.AfterSelectHook, tokenAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	tokenAfterSelectHooks = []TokenHook{}

	AddTokenHook(boil.BeforeUpdateHook, tokenBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	tokenBeforeUpdateHooks = []TokenHook{}

	AddTokenHook(boil.AfterUpdateHook, tokenAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	tokenAfterUpdateHooks = []TokenHook{}

	AddTokenHook(boil.BeforeDeleteHook, tokenBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	tokenBeforeDeleteHooks = []TokenHook{}

	AddTokenHook(boil.AfterDeleteHook, tokenAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	tokenAfterDeleteHooks = []TokenHook{}

	AddTokenHook(boil.BeforeUpsertHook, tokenBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	tokenBeforeUpsertHooks = []TokenHook{}

	AddTokenHook(boil.AfterUpsertHook, tokenAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	tokenAfterUpsertHooks = []TokenHook{}
}
func testTokensInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true, tokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTokensInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx, tokenColumns...); err != nil {
		t.Error(err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTokenToOneOrganisationUsingOrg(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Token
	var foreign Organisation

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, tokenDBTypes, true, tokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
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

	slice := TokenSlice{&local}
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

func testTokenToOneSetOpOrganisationUsingOrg(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Token
	var b, c Organisation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, tokenDBTypes, false, strmangle.SetComplement(tokenPrimaryKeyColumns, tokenColumnsWithoutDefault)...); err != nil {
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

		if x.R.OrgTokens[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OrgID != x.ID {
			t.Error("foreign key was wrong value", a.OrgID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OrgID))
		reflect.Indirect(reflect.ValueOf(&a.OrgID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OrgID != x.ID {
			t.Error("foreign key was wrong value", a.OrgID, x.ID)
		}
	}
}
func testTokensReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true, tokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = token.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testTokensReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true, tokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TokenSlice{token}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testTokensSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true, tokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Tokens(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	tokenDBTypes = map[string]string{`Expires`: `timestamp without time zone`, `ID`: `uuid`, `Name`: `character varying`, `OrgID`: `uuid`}
	_            = bytes.MinRead
)

func testTokensUpdate(t *testing.T) {
	t.Parallel()

	if len(tokenColumns) == len(tokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, token, tokenDBTypes, true, tokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	if err = token.Update(tx); err != nil {
		t.Error(err)
	}
}

func testTokensSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(tokenColumns) == len(tokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, token, tokenDBTypes, true, tokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(tokenColumns, tokenPrimaryKeyColumns) {
		fields = tokenColumns
	} else {
		fields = strmangle.SetComplement(
			tokenColumns,
			tokenPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(token))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := TokenSlice{token}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testTokensUpsert(t *testing.T) {
	t.Parallel()

	if len(tokenColumns) == len(tokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	token := Token{}
	if err = randomize.Struct(seed, &token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Token: %s", err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &token, tokenDBTypes, false, tokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	if err = token.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Token: %s", err)
	}

	count, err = Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

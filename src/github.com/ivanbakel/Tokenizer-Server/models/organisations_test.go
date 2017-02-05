package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testOrganisations(t *testing.T) {
	t.Parallel()

	query := Organisations(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testOrganisationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organisation := &Organisation{}
	if err = randomize.Struct(seed, organisation, organisationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organisation.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = organisation.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Organisations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganisationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organisation := &Organisation{}
	if err = randomize.Struct(seed, organisation, organisationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organisation.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Organisations(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Organisations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganisationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organisation := &Organisation{}
	if err = randomize.Struct(seed, organisation, organisationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organisation.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OrganisationSlice{organisation}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Organisations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testOrganisationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organisation := &Organisation{}
	if err = randomize.Struct(seed, organisation, organisationDBTypes, true, organisationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organisation.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := OrganisationExists(tx, organisation.ID)
	if err != nil {
		t.Errorf("Unable to check if Organisation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OrganisationExistsG to return true, but got false.")
	}
}
func testOrganisationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organisation := &Organisation{}
	if err = randomize.Struct(seed, organisation, organisationDBTypes, true, organisationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organisation.Insert(tx); err != nil {
		t.Error(err)
	}

	organisationFound, err := FindOrganisation(tx, organisation.ID)
	if err != nil {
		t.Error(err)
	}

	if organisationFound == nil {
		t.Error("want a record, got nil")
	}
}
func testOrganisationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organisation := &Organisation{}
	if err = randomize.Struct(seed, organisation, organisationDBTypes, true, organisationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organisation.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Organisations(tx).Bind(organisation); err != nil {
		t.Error(err)
	}
}

func testOrganisationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organisation := &Organisation{}
	if err = randomize.Struct(seed, organisation, organisationDBTypes, true, organisationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organisation.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Organisations(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOrganisationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organisationOne := &Organisation{}
	organisationTwo := &Organisation{}
	if err = randomize.Struct(seed, organisationOne, organisationDBTypes, false, organisationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}
	if err = randomize.Struct(seed, organisationTwo, organisationDBTypes, false, organisationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organisationOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = organisationTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Organisations(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOrganisationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	organisationOne := &Organisation{}
	organisationTwo := &Organisation{}
	if err = randomize.Struct(seed, organisationOne, organisationDBTypes, false, organisationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}
	if err = randomize.Struct(seed, organisationTwo, organisationDBTypes, false, organisationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organisationOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = organisationTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Organisations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func organisationBeforeInsertHook(e boil.Executor, o *Organisation) error {
	*o = Organisation{}
	return nil
}

func organisationAfterInsertHook(e boil.Executor, o *Organisation) error {
	*o = Organisation{}
	return nil
}

func organisationAfterSelectHook(e boil.Executor, o *Organisation) error {
	*o = Organisation{}
	return nil
}

func organisationBeforeUpdateHook(e boil.Executor, o *Organisation) error {
	*o = Organisation{}
	return nil
}

func organisationAfterUpdateHook(e boil.Executor, o *Organisation) error {
	*o = Organisation{}
	return nil
}

func organisationBeforeDeleteHook(e boil.Executor, o *Organisation) error {
	*o = Organisation{}
	return nil
}

func organisationAfterDeleteHook(e boil.Executor, o *Organisation) error {
	*o = Organisation{}
	return nil
}

func organisationBeforeUpsertHook(e boil.Executor, o *Organisation) error {
	*o = Organisation{}
	return nil
}

func organisationAfterUpsertHook(e boil.Executor, o *Organisation) error {
	*o = Organisation{}
	return nil
}

func testOrganisationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Organisation{}
	o := &Organisation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, organisationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Organisation object: %s", err)
	}

	AddOrganisationHook(boil.BeforeInsertHook, organisationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	organisationBeforeInsertHooks = []OrganisationHook{}

	AddOrganisationHook(boil.AfterInsertHook, organisationAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	organisationAfterInsertHooks = []OrganisationHook{}

	AddOrganisationHook(boil.AfterSelectHook, organisationAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	organisationAfterSelectHooks = []OrganisationHook{}

	AddOrganisationHook(boil.BeforeUpdateHook, organisationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	organisationBeforeUpdateHooks = []OrganisationHook{}

	AddOrganisationHook(boil.AfterUpdateHook, organisationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	organisationAfterUpdateHooks = []OrganisationHook{}

	AddOrganisationHook(boil.BeforeDeleteHook, organisationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	organisationBeforeDeleteHooks = []OrganisationHook{}

	AddOrganisationHook(boil.AfterDeleteHook, organisationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	organisationAfterDeleteHooks = []OrganisationHook{}

	AddOrganisationHook(boil.BeforeUpsertHook, organisationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	organisationBeforeUpsertHooks = []OrganisationHook{}

	AddOrganisationHook(boil.AfterUpsertHook, organisationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	organisationAfterUpsertHooks = []OrganisationHook{}
}
func testOrganisationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organisation := &Organisation{}
	if err = randomize.Struct(seed, organisation, organisationDBTypes, true, organisationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organisation.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Organisations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganisationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organisation := &Organisation{}
	if err = randomize.Struct(seed, organisation, organisationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organisation.Insert(tx, organisationColumns...); err != nil {
		t.Error(err)
	}

	count, err := Organisations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganisationToManyOrgTokens(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Organisation
	var b, c Token

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organisationDBTypes, true, organisationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, tokenDBTypes, false, tokenColumnsWithDefault...)
	randomize.Struct(seed, &c, tokenDBTypes, false, tokenColumnsWithDefault...)

	b.OrgID = a.ID
	c.OrgID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	token, err := a.OrgTokens(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range token {
		if v.OrgID == b.OrgID {
			bFound = true
		}
		if v.OrgID == c.OrgID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OrganisationSlice{&a}
	if err = a.L.LoadOrgTokens(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OrgTokens); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.OrgTokens = nil
	if err = a.L.LoadOrgTokens(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OrgTokens); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", token)
	}
}

func testOrganisationToManyOrgUsers(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Organisation
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organisationDBTypes, true, organisationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, userDBTypes, false, userColumnsWithDefault...)
	randomize.Struct(seed, &c, userDBTypes, false, userColumnsWithDefault...)

	b.OrgID.Valid = true
	c.OrgID.Valid = true
	b.OrgID.String = a.ID
	c.OrgID.String = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	user, err := a.OrgUsers(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range user {
		if v.OrgID.String == b.OrgID.String {
			bFound = true
		}
		if v.OrgID.String == c.OrgID.String {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OrganisationSlice{&a}
	if err = a.L.LoadOrgUsers(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OrgUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.OrgUsers = nil
	if err = a.L.LoadOrgUsers(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OrgUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", user)
	}
}

func testOrganisationToManyAddOpOrgTokens(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Organisation
	var b, c, d, e Token

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organisationDBTypes, false, strmangle.SetComplement(organisationPrimaryKeyColumns, organisationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Token{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, tokenDBTypes, false, strmangle.SetComplement(tokenPrimaryKeyColumns, tokenColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Token{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddOrgTokens(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.OrgID {
			t.Error("foreign key was wrong value", a.ID, first.OrgID)
		}
		if a.ID != second.OrgID {
			t.Error("foreign key was wrong value", a.ID, second.OrgID)
		}

		if first.R.Org != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Org != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.OrgTokens[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.OrgTokens[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.OrgTokens(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testOrganisationToManyAddOpOrgUsers(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Organisation
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organisationDBTypes, false, strmangle.SetComplement(organisationPrimaryKeyColumns, organisationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*User{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddOrgUsers(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.OrgID.String {
			t.Error("foreign key was wrong value", a.ID, first.OrgID.String)
		}
		if a.ID != second.OrgID.String {
			t.Error("foreign key was wrong value", a.ID, second.OrgID.String)
		}

		if first.R.Org != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Org != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.OrgUsers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.OrgUsers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.OrgUsers(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testOrganisationToManySetOpOrgUsers(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Organisation
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organisationDBTypes, false, strmangle.SetComplement(organisationPrimaryKeyColumns, organisationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetOrgUsers(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.OrgUsers(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetOrgUsers(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.OrgUsers(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.OrgID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.OrgID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.ID != d.OrgID.String {
		t.Error("foreign key was wrong value", a.ID, d.OrgID.String)
	}
	if a.ID != e.OrgID.String {
		t.Error("foreign key was wrong value", a.ID, e.OrgID.String)
	}

	if b.R.Org != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Org != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Org != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Org != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.OrgUsers[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.OrgUsers[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testOrganisationToManyRemoveOpOrgUsers(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Organisation
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organisationDBTypes, false, strmangle.SetComplement(organisationPrimaryKeyColumns, organisationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddOrgUsers(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.OrgUsers(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveOrgUsers(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.OrgUsers(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.OrgID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.OrgID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Org != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Org != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Org != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Org != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.OrgUsers) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.OrgUsers[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.OrgUsers[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testOrganisationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organisation := &Organisation{}
	if err = randomize.Struct(seed, organisation, organisationDBTypes, true, organisationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organisation.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = organisation.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testOrganisationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organisation := &Organisation{}
	if err = randomize.Struct(seed, organisation, organisationDBTypes, true, organisationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organisation.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OrganisationSlice{organisation}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testOrganisationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organisation := &Organisation{}
	if err = randomize.Struct(seed, organisation, organisationDBTypes, true, organisationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organisation.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Organisations(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	organisationDBTypes = map[string]string{`ID`: `uuid`, `Name`: `character varying`}
	_                   = bytes.MinRead
)

func testOrganisationsUpdate(t *testing.T) {
	t.Parallel()

	if len(organisationColumns) == len(organisationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	organisation := &Organisation{}
	if err = randomize.Struct(seed, organisation, organisationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organisation.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Organisations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, organisation, organisationDBTypes, true, organisationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	if err = organisation.Update(tx); err != nil {
		t.Error(err)
	}
}

func testOrganisationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(organisationColumns) == len(organisationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	organisation := &Organisation{}
	if err = randomize.Struct(seed, organisation, organisationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organisation.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Organisations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, organisation, organisationDBTypes, true, organisationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(organisationColumns, organisationPrimaryKeyColumns) {
		fields = organisationColumns
	} else {
		fields = strmangle.SetComplement(
			organisationColumns,
			organisationPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(organisation))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := OrganisationSlice{organisation}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testOrganisationsUpsert(t *testing.T) {
	t.Parallel()

	if len(organisationColumns) == len(organisationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	organisation := Organisation{}
	if err = randomize.Struct(seed, &organisation, organisationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organisation.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Organisation: %s", err)
	}

	count, err := Organisations(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &organisation, organisationDBTypes, false, organisationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Organisation struct: %s", err)
	}

	if err = organisation.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Organisation: %s", err)
	}

	count, err = Organisations(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

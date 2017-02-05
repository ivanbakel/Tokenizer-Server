package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
)

// Organisation is an object representing the database table.
type Organisation struct {
	ID   string `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *organisationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L organisationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// organisationR is where relationships are stored.
type organisationR struct {
	OrgTokens TokenSlice
	OrgUsers  UserSlice
}

// organisationL is where Load methods for each relationship are stored.
type organisationL struct{}

var (
	organisationColumns               = []string{"id", "name"}
	organisationColumnsWithoutDefault = []string{"name"}
	organisationColumnsWithDefault    = []string{"id"}
	organisationPrimaryKeyColumns     = []string{"id"}
)

type (
	// OrganisationSlice is an alias for a slice of pointers to Organisation.
	// This should generally be used opposed to []Organisation.
	OrganisationSlice []*Organisation
	// OrganisationHook is the signature for custom Organisation hook methods
	OrganisationHook func(boil.Executor, *Organisation) error

	organisationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	organisationType                 = reflect.TypeOf(&Organisation{})
	organisationMapping              = queries.MakeStructMapping(organisationType)
	organisationPrimaryKeyMapping, _ = queries.BindMapping(organisationType, organisationMapping, organisationPrimaryKeyColumns)
	organisationInsertCacheMut       sync.RWMutex
	organisationInsertCache          = make(map[string]insertCache)
	organisationUpdateCacheMut       sync.RWMutex
	organisationUpdateCache          = make(map[string]updateCache)
	organisationUpsertCacheMut       sync.RWMutex
	organisationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var organisationBeforeInsertHooks []OrganisationHook
var organisationBeforeUpdateHooks []OrganisationHook
var organisationBeforeDeleteHooks []OrganisationHook
var organisationBeforeUpsertHooks []OrganisationHook

var organisationAfterInsertHooks []OrganisationHook
var organisationAfterSelectHooks []OrganisationHook
var organisationAfterUpdateHooks []OrganisationHook
var organisationAfterDeleteHooks []OrganisationHook
var organisationAfterUpsertHooks []OrganisationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Organisation) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range organisationBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Organisation) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range organisationBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Organisation) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range organisationBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Organisation) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range organisationBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Organisation) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range organisationAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Organisation) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range organisationAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Organisation) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range organisationAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Organisation) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range organisationAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Organisation) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range organisationAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrganisationHook registers your hook function for all future operations.
func AddOrganisationHook(hookPoint boil.HookPoint, organisationHook OrganisationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		organisationBeforeInsertHooks = append(organisationBeforeInsertHooks, organisationHook)
	case boil.BeforeUpdateHook:
		organisationBeforeUpdateHooks = append(organisationBeforeUpdateHooks, organisationHook)
	case boil.BeforeDeleteHook:
		organisationBeforeDeleteHooks = append(organisationBeforeDeleteHooks, organisationHook)
	case boil.BeforeUpsertHook:
		organisationBeforeUpsertHooks = append(organisationBeforeUpsertHooks, organisationHook)
	case boil.AfterInsertHook:
		organisationAfterInsertHooks = append(organisationAfterInsertHooks, organisationHook)
	case boil.AfterSelectHook:
		organisationAfterSelectHooks = append(organisationAfterSelectHooks, organisationHook)
	case boil.AfterUpdateHook:
		organisationAfterUpdateHooks = append(organisationAfterUpdateHooks, organisationHook)
	case boil.AfterDeleteHook:
		organisationAfterDeleteHooks = append(organisationAfterDeleteHooks, organisationHook)
	case boil.AfterUpsertHook:
		organisationAfterUpsertHooks = append(organisationAfterUpsertHooks, organisationHook)
	}
}

// OneP returns a single organisation record from the query, and panics on error.
func (q organisationQuery) OneP() *Organisation {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single organisation record from the query.
func (q organisationQuery) One() (*Organisation, error) {
	o := &Organisation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for organisations")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Organisation records from the query, and panics on error.
func (q organisationQuery) AllP() OrganisationSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Organisation records from the query.
func (q organisationQuery) All() (OrganisationSlice, error) {
	var o OrganisationSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Organisation slice")
	}

	if len(organisationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Organisation records in the query, and panics on error.
func (q organisationQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Organisation records in the query.
func (q organisationQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count organisations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q organisationQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q organisationQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if organisations exists")
	}

	return count > 0, nil
}

// OrgTokensG retrieves all the token's tokens via org_id column.
func (o *Organisation) OrgTokensG(mods ...qm.QueryMod) tokenQuery {
	return o.OrgTokens(boil.GetDB(), mods...)
}

// OrgTokens retrieves all the token's tokens with an executor via org_id column.
func (o *Organisation) OrgTokens(exec boil.Executor, mods ...qm.QueryMod) tokenQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"org_id\"=?", o.ID),
	)

	query := Tokens(exec, queryMods...)
	queries.SetFrom(query.Query, "\"tokens\" as \"a\"")
	return query
}

// OrgUsersG retrieves all the user's users via org_id column.
func (o *Organisation) OrgUsersG(mods ...qm.QueryMod) userQuery {
	return o.OrgUsers(boil.GetDB(), mods...)
}

// OrgUsers retrieves all the user's users with an executor via org_id column.
func (o *Organisation) OrgUsers(exec boil.Executor, mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"org_id\"=?", o.ID),
	)

	query := Users(exec, queryMods...)
	queries.SetFrom(query.Query, "\"users\" as \"a\"")
	return query
}

// LoadOrgTokens allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (organisationL) LoadOrgTokens(e boil.Executor, singular bool, maybeOrganisation interface{}) error {
	var slice []*Organisation
	var object *Organisation

	count := 1
	if singular {
		object = maybeOrganisation.(*Organisation)
	} else {
		slice = *maybeOrganisation.(*OrganisationSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &organisationR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &organisationR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"tokens\" where \"org_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load tokens")
	}
	defer results.Close()

	var resultSlice []*Token
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice tokens")
	}

	if len(tokenAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.OrgTokens = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.OrgID {
				local.R.OrgTokens = append(local.R.OrgTokens, foreign)
				break
			}
		}
	}

	return nil
}

// LoadOrgUsers allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (organisationL) LoadOrgUsers(e boil.Executor, singular bool, maybeOrganisation interface{}) error {
	var slice []*Organisation
	var object *Organisation

	count := 1
	if singular {
		object = maybeOrganisation.(*Organisation)
	} else {
		slice = *maybeOrganisation.(*OrganisationSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &organisationR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &organisationR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"users\" where \"org_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load users")
	}
	defer results.Close()

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice users")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.OrgUsers = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.OrgID.String {
				local.R.OrgUsers = append(local.R.OrgUsers, foreign)
				break
			}
		}
	}

	return nil
}

// AddOrgTokensG adds the given related objects to the existing relationships
// of the organisation, optionally inserting them as new records.
// Appends related to o.R.OrgTokens.
// Sets related.R.Org appropriately.
// Uses the global database handle.
func (o *Organisation) AddOrgTokensG(insert bool, related ...*Token) error {
	return o.AddOrgTokens(boil.GetDB(), insert, related...)
}

// AddOrgTokensP adds the given related objects to the existing relationships
// of the organisation, optionally inserting them as new records.
// Appends related to o.R.OrgTokens.
// Sets related.R.Org appropriately.
// Panics on error.
func (o *Organisation) AddOrgTokensP(exec boil.Executor, insert bool, related ...*Token) {
	if err := o.AddOrgTokens(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddOrgTokensGP adds the given related objects to the existing relationships
// of the organisation, optionally inserting them as new records.
// Appends related to o.R.OrgTokens.
// Sets related.R.Org appropriately.
// Uses the global database handle and panics on error.
func (o *Organisation) AddOrgTokensGP(insert bool, related ...*Token) {
	if err := o.AddOrgTokens(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddOrgTokens adds the given related objects to the existing relationships
// of the organisation, optionally inserting them as new records.
// Appends related to o.R.OrgTokens.
// Sets related.R.Org appropriately.
func (o *Organisation) AddOrgTokens(exec boil.Executor, insert bool, related ...*Token) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.OrgID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"tokens\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"org_id"}),
				strmangle.WhereClause("\"", "\"", 2, tokenPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.OrgID = o.ID
		}
	}

	if o.R == nil {
		o.R = &organisationR{
			OrgTokens: related,
		}
	} else {
		o.R.OrgTokens = append(o.R.OrgTokens, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &tokenR{
				Org: o,
			}
		} else {
			rel.R.Org = o
		}
	}
	return nil
}

// AddOrgUsersG adds the given related objects to the existing relationships
// of the organisation, optionally inserting them as new records.
// Appends related to o.R.OrgUsers.
// Sets related.R.Org appropriately.
// Uses the global database handle.
func (o *Organisation) AddOrgUsersG(insert bool, related ...*User) error {
	return o.AddOrgUsers(boil.GetDB(), insert, related...)
}

// AddOrgUsersP adds the given related objects to the existing relationships
// of the organisation, optionally inserting them as new records.
// Appends related to o.R.OrgUsers.
// Sets related.R.Org appropriately.
// Panics on error.
func (o *Organisation) AddOrgUsersP(exec boil.Executor, insert bool, related ...*User) {
	if err := o.AddOrgUsers(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddOrgUsersGP adds the given related objects to the existing relationships
// of the organisation, optionally inserting them as new records.
// Appends related to o.R.OrgUsers.
// Sets related.R.Org appropriately.
// Uses the global database handle and panics on error.
func (o *Organisation) AddOrgUsersGP(insert bool, related ...*User) {
	if err := o.AddOrgUsers(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddOrgUsers adds the given related objects to the existing relationships
// of the organisation, optionally inserting them as new records.
// Appends related to o.R.OrgUsers.
// Sets related.R.Org appropriately.
func (o *Organisation) AddOrgUsers(exec boil.Executor, insert bool, related ...*User) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.OrgID.String = o.ID
			rel.OrgID.Valid = true
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"users\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"org_id"}),
				strmangle.WhereClause("\"", "\"", 2, userPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.OrgID.String = o.ID
			rel.OrgID.Valid = true
		}
	}

	if o.R == nil {
		o.R = &organisationR{
			OrgUsers: related,
		}
	} else {
		o.R.OrgUsers = append(o.R.OrgUsers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userR{
				Org: o,
			}
		} else {
			rel.R.Org = o
		}
	}
	return nil
}

// SetOrgUsersG removes all previously related items of the
// organisation replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Org's OrgUsers accordingly.
// Replaces o.R.OrgUsers with related.
// Sets related.R.Org's OrgUsers accordingly.
// Uses the global database handle.
func (o *Organisation) SetOrgUsersG(insert bool, related ...*User) error {
	return o.SetOrgUsers(boil.GetDB(), insert, related...)
}

// SetOrgUsersP removes all previously related items of the
// organisation replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Org's OrgUsers accordingly.
// Replaces o.R.OrgUsers with related.
// Sets related.R.Org's OrgUsers accordingly.
// Panics on error.
func (o *Organisation) SetOrgUsersP(exec boil.Executor, insert bool, related ...*User) {
	if err := o.SetOrgUsers(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetOrgUsersGP removes all previously related items of the
// organisation replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Org's OrgUsers accordingly.
// Replaces o.R.OrgUsers with related.
// Sets related.R.Org's OrgUsers accordingly.
// Uses the global database handle and panics on error.
func (o *Organisation) SetOrgUsersGP(insert bool, related ...*User) {
	if err := o.SetOrgUsers(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetOrgUsers removes all previously related items of the
// organisation replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Org's OrgUsers accordingly.
// Replaces o.R.OrgUsers with related.
// Sets related.R.Org's OrgUsers accordingly.
func (o *Organisation) SetOrgUsers(exec boil.Executor, insert bool, related ...*User) error {
	query := "update \"users\" set \"org_id\" = null where \"org_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.OrgUsers {
			rel.OrgID.Valid = false
			if rel.R == nil {
				continue
			}

			rel.R.Org = nil
		}

		o.R.OrgUsers = nil
	}
	return o.AddOrgUsers(exec, insert, related...)
}

// RemoveOrgUsersG relationships from objects passed in.
// Removes related items from R.OrgUsers (uses pointer comparison, removal does not keep order)
// Sets related.R.Org.
// Uses the global database handle.
func (o *Organisation) RemoveOrgUsersG(related ...*User) error {
	return o.RemoveOrgUsers(boil.GetDB(), related...)
}

// RemoveOrgUsersP relationships from objects passed in.
// Removes related items from R.OrgUsers (uses pointer comparison, removal does not keep order)
// Sets related.R.Org.
// Panics on error.
func (o *Organisation) RemoveOrgUsersP(exec boil.Executor, related ...*User) {
	if err := o.RemoveOrgUsers(exec, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// RemoveOrgUsersGP relationships from objects passed in.
// Removes related items from R.OrgUsers (uses pointer comparison, removal does not keep order)
// Sets related.R.Org.
// Uses the global database handle and panics on error.
func (o *Organisation) RemoveOrgUsersGP(related ...*User) {
	if err := o.RemoveOrgUsers(boil.GetDB(), related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// RemoveOrgUsers relationships from objects passed in.
// Removes related items from R.OrgUsers (uses pointer comparison, removal does not keep order)
// Sets related.R.Org.
func (o *Organisation) RemoveOrgUsers(exec boil.Executor, related ...*User) error {
	var err error
	for _, rel := range related {
		rel.OrgID.Valid = false
		if rel.R != nil {
			rel.R.Org = nil
		}
		if err = rel.Update(exec, "org_id"); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.OrgUsers {
			if rel != ri {
				continue
			}

			ln := len(o.R.OrgUsers)
			if ln > 1 && i < ln-1 {
				o.R.OrgUsers[i] = o.R.OrgUsers[ln-1]
			}
			o.R.OrgUsers = o.R.OrgUsers[:ln-1]
			break
		}
	}

	return nil
}

// OrganisationsG retrieves all records.
func OrganisationsG(mods ...qm.QueryMod) organisationQuery {
	return Organisations(boil.GetDB(), mods...)
}

// Organisations retrieves all the records using an executor.
func Organisations(exec boil.Executor, mods ...qm.QueryMod) organisationQuery {
	mods = append(mods, qm.From("\"organisations\""))
	return organisationQuery{NewQuery(exec, mods...)}
}

// FindOrganisationG retrieves a single record by ID.
func FindOrganisationG(id string, selectCols ...string) (*Organisation, error) {
	return FindOrganisation(boil.GetDB(), id, selectCols...)
}

// FindOrganisationGP retrieves a single record by ID, and panics on error.
func FindOrganisationGP(id string, selectCols ...string) *Organisation {
	retobj, err := FindOrganisation(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindOrganisation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrganisation(exec boil.Executor, id string, selectCols ...string) (*Organisation, error) {
	organisationObj := &Organisation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"organisations\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(organisationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from organisations")
	}

	return organisationObj, nil
}

// FindOrganisationP retrieves a single record by ID with an executor, and panics on error.
func FindOrganisationP(exec boil.Executor, id string, selectCols ...string) *Organisation {
	retobj, err := FindOrganisation(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Organisation) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Organisation) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Organisation) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Organisation) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no organisations provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(organisationColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	organisationInsertCacheMut.RLock()
	cache, cached := organisationInsertCache[key]
	organisationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			organisationColumns,
			organisationColumnsWithDefault,
			organisationColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(organisationType, organisationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(organisationType, organisationMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"organisations\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.query += fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into organisations")
	}

	if !cached {
		organisationInsertCacheMut.Lock()
		organisationInsertCache[key] = cache
		organisationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Organisation record. See Update for
// whitelist behavior description.
func (o *Organisation) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Organisation record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Organisation) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Organisation, and panics on error.
// See Update for whitelist behavior description.
func (o *Organisation) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Organisation.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Organisation) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	organisationUpdateCacheMut.RLock()
	cache, cached := organisationUpdateCache[key]
	organisationUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(organisationColumns, organisationPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update organisations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"organisations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, organisationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(organisationType, organisationMapping, append(wl, organisationPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update organisations row")
	}

	if !cached {
		organisationUpdateCacheMut.Lock()
		organisationUpdateCache[key] = cache
		organisationUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q organisationQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q organisationQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for organisations")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OrganisationSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o OrganisationSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o OrganisationSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrganisationSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organisationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"organisations\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(organisationPrimaryKeyColumns), len(colNames)+1, len(organisationPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in organisation slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Organisation) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Organisation) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Organisation) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Organisation) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no organisations provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(organisationColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	organisationUpsertCacheMut.RLock()
	cache, cached := organisationUpsertCache[key]
	organisationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			organisationColumns,
			organisationColumnsWithDefault,
			organisationColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			organisationColumns,
			organisationPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert organisations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(organisationPrimaryKeyColumns))
			copy(conflict, organisationPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"organisations\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(organisationType, organisationMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(organisationType, organisationMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert organisations")
	}

	if !cached {
		organisationUpsertCacheMut.Lock()
		organisationUpsertCache[key] = cache
		organisationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Organisation record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Organisation) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Organisation record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Organisation) DeleteG() error {
	if o == nil {
		return errors.New("models: no Organisation provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Organisation record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Organisation) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Organisation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Organisation) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Organisation provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), organisationPrimaryKeyMapping)
	sql := "DELETE FROM \"organisations\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from organisations")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q organisationQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q organisationQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no organisationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from organisations")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o OrganisationSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o OrganisationSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Organisation slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o OrganisationSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrganisationSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Organisation slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(organisationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organisationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"organisations\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, organisationPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(organisationPrimaryKeyColumns), 1, len(organisationPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from organisation slice")
	}

	if len(organisationAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Organisation) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Organisation) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Organisation) ReloadG() error {
	if o == nil {
		return errors.New("models: no Organisation provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Organisation) Reload(exec boil.Executor) error {
	ret, err := FindOrganisation(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OrganisationSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OrganisationSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrganisationSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty OrganisationSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrganisationSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	organisations := OrganisationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organisationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"organisations\".* FROM \"organisations\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, organisationPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(organisationPrimaryKeyColumns), 1, len(organisationPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&organisations)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrganisationSlice")
	}

	*o = organisations

	return nil
}

// OrganisationExists checks if the Organisation row exists.
func OrganisationExists(exec boil.Executor, id string) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"organisations\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if organisations exists")
	}

	return exists, nil
}

// OrganisationExistsG checks if the Organisation row exists.
func OrganisationExistsG(id string) (bool, error) {
	return OrganisationExists(boil.GetDB(), id)
}

// OrganisationExistsGP checks if the Organisation row exists. Panics on error.
func OrganisationExistsGP(id string) bool {
	e, err := OrganisationExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// OrganisationExistsP checks if the Organisation row exists. Panics on error.
func OrganisationExistsP(exec boil.Executor, id string) bool {
	e, err := OrganisationExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

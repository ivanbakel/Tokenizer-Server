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

// Token is an object representing the database table.
type Token struct {
	ID      string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name    string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Expires time.Time `boil:"expires" json:"expires" toml:"expires" yaml:"expires"`
	OrgID   string    `boil:"org_id" json:"org_id" toml:"org_id" yaml:"org_id"`

	R *tokenR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tokenL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// tokenR is where relationships are stored.
type tokenR struct {
	Org *Organisation
}

// tokenL is where Load methods for each relationship are stored.
type tokenL struct{}

var (
	tokenColumns               = []string{"id", "name", "expires", "org_id"}
	tokenColumnsWithoutDefault = []string{"name", "expires", "org_id"}
	tokenColumnsWithDefault    = []string{"id"}
	tokenPrimaryKeyColumns     = []string{"id"}
)

type (
	// TokenSlice is an alias for a slice of pointers to Token.
	// This should generally be used opposed to []Token.
	TokenSlice []*Token
	// TokenHook is the signature for custom Token hook methods
	TokenHook func(boil.Executor, *Token) error

	tokenQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tokenType                 = reflect.TypeOf(&Token{})
	tokenMapping              = queries.MakeStructMapping(tokenType)
	tokenPrimaryKeyMapping, _ = queries.BindMapping(tokenType, tokenMapping, tokenPrimaryKeyColumns)
	tokenInsertCacheMut       sync.RWMutex
	tokenInsertCache          = make(map[string]insertCache)
	tokenUpdateCacheMut       sync.RWMutex
	tokenUpdateCache          = make(map[string]updateCache)
	tokenUpsertCacheMut       sync.RWMutex
	tokenUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var tokenBeforeInsertHooks []TokenHook
var tokenBeforeUpdateHooks []TokenHook
var tokenBeforeDeleteHooks []TokenHook
var tokenBeforeUpsertHooks []TokenHook

var tokenAfterInsertHooks []TokenHook
var tokenAfterSelectHooks []TokenHook
var tokenAfterUpdateHooks []TokenHook
var tokenAfterDeleteHooks []TokenHook
var tokenAfterUpsertHooks []TokenHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Token) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range tokenBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Token) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range tokenBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Token) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range tokenBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Token) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range tokenBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Token) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range tokenAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Token) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range tokenAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Token) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range tokenAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Token) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range tokenAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Token) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range tokenAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTokenHook registers your hook function for all future operations.
func AddTokenHook(hookPoint boil.HookPoint, tokenHook TokenHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		tokenBeforeInsertHooks = append(tokenBeforeInsertHooks, tokenHook)
	case boil.BeforeUpdateHook:
		tokenBeforeUpdateHooks = append(tokenBeforeUpdateHooks, tokenHook)
	case boil.BeforeDeleteHook:
		tokenBeforeDeleteHooks = append(tokenBeforeDeleteHooks, tokenHook)
	case boil.BeforeUpsertHook:
		tokenBeforeUpsertHooks = append(tokenBeforeUpsertHooks, tokenHook)
	case boil.AfterInsertHook:
		tokenAfterInsertHooks = append(tokenAfterInsertHooks, tokenHook)
	case boil.AfterSelectHook:
		tokenAfterSelectHooks = append(tokenAfterSelectHooks, tokenHook)
	case boil.AfterUpdateHook:
		tokenAfterUpdateHooks = append(tokenAfterUpdateHooks, tokenHook)
	case boil.AfterDeleteHook:
		tokenAfterDeleteHooks = append(tokenAfterDeleteHooks, tokenHook)
	case boil.AfterUpsertHook:
		tokenAfterUpsertHooks = append(tokenAfterUpsertHooks, tokenHook)
	}
}

// OneP returns a single token record from the query, and panics on error.
func (q tokenQuery) OneP() *Token {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single token record from the query.
func (q tokenQuery) One() (*Token, error) {
	o := &Token{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for tokens")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Token records from the query, and panics on error.
func (q tokenQuery) AllP() TokenSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Token records from the query.
func (q tokenQuery) All() (TokenSlice, error) {
	var o TokenSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Token slice")
	}

	if len(tokenAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Token records in the query, and panics on error.
func (q tokenQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Token records in the query.
func (q tokenQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count tokens rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q tokenQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q tokenQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if tokens exists")
	}

	return count > 0, nil
}

// OrgG pointed to by the foreign key.
func (o *Token) OrgG(mods ...qm.QueryMod) organisationQuery {
	return o.Org(boil.GetDB(), mods...)
}

// Org pointed to by the foreign key.
func (o *Token) Org(exec boil.Executor, mods ...qm.QueryMod) organisationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.OrgID),
	}

	queryMods = append(queryMods, mods...)

	query := Organisations(exec, queryMods...)
	queries.SetFrom(query.Query, "\"organisations\"")

	return query
}

// LoadOrg allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (tokenL) LoadOrg(e boil.Executor, singular bool, maybeToken interface{}) error {
	var slice []*Token
	var object *Token

	count := 1
	if singular {
		object = maybeToken.(*Token)
	} else {
		slice = *maybeToken.(*TokenSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &tokenR{}
		}
		args[0] = object.OrgID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &tokenR{}
			}
			args[i] = obj.OrgID
		}
	}

	query := fmt.Sprintf(
		"select * from \"organisations\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Organisation")
	}
	defer results.Close()

	var resultSlice []*Organisation
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Organisation")
	}

	if len(tokenAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Org = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.OrgID == foreign.ID {
				local.R.Org = foreign
				break
			}
		}
	}

	return nil
}

// SetOrgG of the token to the related item.
// Sets o.R.Org to related.
// Adds o to related.R.OrgTokens.
// Uses the global database handle.
func (o *Token) SetOrgG(insert bool, related *Organisation) error {
	return o.SetOrg(boil.GetDB(), insert, related)
}

// SetOrgP of the token to the related item.
// Sets o.R.Org to related.
// Adds o to related.R.OrgTokens.
// Panics on error.
func (o *Token) SetOrgP(exec boil.Executor, insert bool, related *Organisation) {
	if err := o.SetOrg(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetOrgGP of the token to the related item.
// Sets o.R.Org to related.
// Adds o to related.R.OrgTokens.
// Uses the global database handle and panics on error.
func (o *Token) SetOrgGP(insert bool, related *Organisation) {
	if err := o.SetOrg(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetOrg of the token to the related item.
// Sets o.R.Org to related.
// Adds o to related.R.OrgTokens.
func (o *Token) SetOrg(exec boil.Executor, insert bool, related *Organisation) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"tokens\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"org_id"}),
		strmangle.WhereClause("\"", "\"", 2, tokenPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrgID = related.ID

	if o.R == nil {
		o.R = &tokenR{
			Org: related,
		}
	} else {
		o.R.Org = related
	}

	if related.R == nil {
		related.R = &organisationR{
			OrgTokens: TokenSlice{o},
		}
	} else {
		related.R.OrgTokens = append(related.R.OrgTokens, o)
	}

	return nil
}

// TokensG retrieves all records.
func TokensG(mods ...qm.QueryMod) tokenQuery {
	return Tokens(boil.GetDB(), mods...)
}

// Tokens retrieves all the records using an executor.
func Tokens(exec boil.Executor, mods ...qm.QueryMod) tokenQuery {
	mods = append(mods, qm.From("\"tokens\""))
	return tokenQuery{NewQuery(exec, mods...)}
}

// FindTokenG retrieves a single record by ID.
func FindTokenG(id string, selectCols ...string) (*Token, error) {
	return FindToken(boil.GetDB(), id, selectCols...)
}

// FindTokenGP retrieves a single record by ID, and panics on error.
func FindTokenGP(id string, selectCols ...string) *Token {
	retobj, err := FindToken(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindToken retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindToken(exec boil.Executor, id string, selectCols ...string) (*Token, error) {
	tokenObj := &Token{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"tokens\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(tokenObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from tokens")
	}

	return tokenObj, nil
}

// FindTokenP retrieves a single record by ID with an executor, and panics on error.
func FindTokenP(exec boil.Executor, id string, selectCols ...string) *Token {
	retobj, err := FindToken(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Token) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Token) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Token) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Token) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no tokens provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tokenColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	tokenInsertCacheMut.RLock()
	cache, cached := tokenInsertCache[key]
	tokenInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			tokenColumns,
			tokenColumnsWithDefault,
			tokenColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(tokenType, tokenMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tokenType, tokenMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"tokens\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into tokens")
	}

	if !cached {
		tokenInsertCacheMut.Lock()
		tokenInsertCache[key] = cache
		tokenInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Token record. See Update for
// whitelist behavior description.
func (o *Token) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Token record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Token) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Token, and panics on error.
// See Update for whitelist behavior description.
func (o *Token) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Token.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Token) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	tokenUpdateCacheMut.RLock()
	cache, cached := tokenUpdateCache[key]
	tokenUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(tokenColumns, tokenPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update tokens, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"tokens\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, tokenPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tokenType, tokenMapping, append(wl, tokenPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update tokens row")
	}

	if !cached {
		tokenUpdateCacheMut.Lock()
		tokenUpdateCache[key] = cache
		tokenUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q tokenQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q tokenQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for tokens")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TokenSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o TokenSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o TokenSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TokenSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"tokens\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(tokenPrimaryKeyColumns), len(colNames)+1, len(tokenPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in token slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Token) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Token) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Token) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Token) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no tokens provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tokenColumnsWithDefault, o)

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

	tokenUpsertCacheMut.RLock()
	cache, cached := tokenUpsertCache[key]
	tokenUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			tokenColumns,
			tokenColumnsWithDefault,
			tokenColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			tokenColumns,
			tokenPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert tokens, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(tokenPrimaryKeyColumns))
			copy(conflict, tokenPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"tokens\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(tokenType, tokenMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tokenType, tokenMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert tokens")
	}

	if !cached {
		tokenUpsertCacheMut.Lock()
		tokenUpsertCache[key] = cache
		tokenUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Token record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Token) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Token record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Token) DeleteG() error {
	if o == nil {
		return errors.New("models: no Token provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Token record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Token) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Token record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Token) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Token provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tokenPrimaryKeyMapping)
	sql := "DELETE FROM \"tokens\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from tokens")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q tokenQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q tokenQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no tokenQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from tokens")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o TokenSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o TokenSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Token slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o TokenSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TokenSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Token slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(tokenBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"tokens\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, tokenPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(tokenPrimaryKeyColumns), 1, len(tokenPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from token slice")
	}

	if len(tokenAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Token) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Token) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Token) ReloadG() error {
	if o == nil {
		return errors.New("models: no Token provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Token) Reload(exec boil.Executor) error {
	ret, err := FindToken(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *TokenSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *TokenSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TokenSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty TokenSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TokenSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	tokens := TokenSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"tokens\".* FROM \"tokens\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, tokenPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(tokenPrimaryKeyColumns), 1, len(tokenPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&tokens)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TokenSlice")
	}

	*o = tokens

	return nil
}

// TokenExists checks if the Token row exists.
func TokenExists(exec boil.Executor, id string) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"tokens\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if tokens exists")
	}

	return exists, nil
}

// TokenExistsG checks if the Token row exists.
func TokenExistsG(id string) (bool, error) {
	return TokenExists(boil.GetDB(), id)
}

// TokenExistsGP checks if the Token row exists. Panics on error.
func TokenExistsGP(id string) bool {
	e, err := TokenExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// TokenExistsP checks if the Token row exists. Panics on error.
func TokenExistsP(exec boil.Executor, id string) bool {
	e, err := TokenExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

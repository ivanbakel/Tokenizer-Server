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
	"gopkg.in/nullbio/null.v6"
)

// UserToken is an object representing the database table.
type UserToken struct {
	UserID  string     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	TokenID string     `boil:"token_id" json:"token_id" toml:"token_id" yaml:"token_id"`
	Number  null.Int16 `boil:"number" json:"number,omitempty" toml:"number" yaml:"number,omitempty"`

	R *userTokenR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userTokenL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// userTokenR is where relationships are stored.
type userTokenR struct {
	User  *User
	Token *Token
}

// userTokenL is where Load methods for each relationship are stored.
type userTokenL struct{}

var (
	userTokenColumns               = []string{"user_id", "token_id", "number"}
	userTokenColumnsWithoutDefault = []string{"user_id", "token_id"}
	userTokenColumnsWithDefault    = []string{"number"}
	userTokenPrimaryKeyColumns     = []string{"user_id", "token_id"}
)

type (
	// UserTokenSlice is an alias for a slice of pointers to UserToken.
	// This should generally be used opposed to []UserToken.
	UserTokenSlice []*UserToken
	// UserTokenHook is the signature for custom UserToken hook methods
	UserTokenHook func(boil.Executor, *UserToken) error

	userTokenQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userTokenType                 = reflect.TypeOf(&UserToken{})
	userTokenMapping              = queries.MakeStructMapping(userTokenType)
	userTokenPrimaryKeyMapping, _ = queries.BindMapping(userTokenType, userTokenMapping, userTokenPrimaryKeyColumns)
	userTokenInsertCacheMut       sync.RWMutex
	userTokenInsertCache          = make(map[string]insertCache)
	userTokenUpdateCacheMut       sync.RWMutex
	userTokenUpdateCache          = make(map[string]updateCache)
	userTokenUpsertCacheMut       sync.RWMutex
	userTokenUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var userTokenBeforeInsertHooks []UserTokenHook
var userTokenBeforeUpdateHooks []UserTokenHook
var userTokenBeforeDeleteHooks []UserTokenHook
var userTokenBeforeUpsertHooks []UserTokenHook

var userTokenAfterInsertHooks []UserTokenHook
var userTokenAfterSelectHooks []UserTokenHook
var userTokenAfterUpdateHooks []UserTokenHook
var userTokenAfterDeleteHooks []UserTokenHook
var userTokenAfterUpsertHooks []UserTokenHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserToken) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userTokenBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserToken) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range userTokenBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserToken) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range userTokenBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserToken) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userTokenBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserToken) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userTokenAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserToken) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range userTokenAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserToken) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range userTokenAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserToken) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range userTokenAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserToken) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userTokenAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserTokenHook registers your hook function for all future operations.
func AddUserTokenHook(hookPoint boil.HookPoint, userTokenHook UserTokenHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		userTokenBeforeInsertHooks = append(userTokenBeforeInsertHooks, userTokenHook)
	case boil.BeforeUpdateHook:
		userTokenBeforeUpdateHooks = append(userTokenBeforeUpdateHooks, userTokenHook)
	case boil.BeforeDeleteHook:
		userTokenBeforeDeleteHooks = append(userTokenBeforeDeleteHooks, userTokenHook)
	case boil.BeforeUpsertHook:
		userTokenBeforeUpsertHooks = append(userTokenBeforeUpsertHooks, userTokenHook)
	case boil.AfterInsertHook:
		userTokenAfterInsertHooks = append(userTokenAfterInsertHooks, userTokenHook)
	case boil.AfterSelectHook:
		userTokenAfterSelectHooks = append(userTokenAfterSelectHooks, userTokenHook)
	case boil.AfterUpdateHook:
		userTokenAfterUpdateHooks = append(userTokenAfterUpdateHooks, userTokenHook)
	case boil.AfterDeleteHook:
		userTokenAfterDeleteHooks = append(userTokenAfterDeleteHooks, userTokenHook)
	case boil.AfterUpsertHook:
		userTokenAfterUpsertHooks = append(userTokenAfterUpsertHooks, userTokenHook)
	}
}

// OneP returns a single userToken record from the query, and panics on error.
func (q userTokenQuery) OneP() *UserToken {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single userToken record from the query.
func (q userTokenQuery) One() (*UserToken, error) {
	o := &UserToken{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_tokens")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all UserToken records from the query, and panics on error.
func (q userTokenQuery) AllP() UserTokenSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all UserToken records from the query.
func (q userTokenQuery) All() (UserTokenSlice, error) {
	var o UserTokenSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserToken slice")
	}

	if len(userTokenAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all UserToken records in the query, and panics on error.
func (q userTokenQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all UserToken records in the query.
func (q userTokenQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_tokens rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q userTokenQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q userTokenQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_tokens exists")
	}

	return count > 0, nil
}

// UserG pointed to by the foreign key.
func (o *UserToken) UserG(mods ...qm.QueryMod) userQuery {
	return o.User(boil.GetDB(), mods...)
}

// User pointed to by the foreign key.
func (o *UserToken) User(exec boil.Executor, mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(exec, queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// TokenG pointed to by the foreign key.
func (o *UserToken) TokenG(mods ...qm.QueryMod) tokenQuery {
	return o.Token(boil.GetDB(), mods...)
}

// Token pointed to by the foreign key.
func (o *UserToken) Token(exec boil.Executor, mods ...qm.QueryMod) tokenQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.TokenID),
	}

	queryMods = append(queryMods, mods...)

	query := Tokens(exec, queryMods...)
	queries.SetFrom(query.Query, "\"tokens\"")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (userTokenL) LoadUser(e boil.Executor, singular bool, maybeUserToken interface{}) error {
	var slice []*UserToken
	var object *UserToken

	count := 1
	if singular {
		object = maybeUserToken.(*UserToken)
	} else {
		slice = *maybeUserToken.(*UserTokenSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &userTokenR{}
		}
		args[0] = object.UserID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &userTokenR{}
			}
			args[i] = obj.UserID
		}
	}

	query := fmt.Sprintf(
		"select * from \"users\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}
	defer results.Close()

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if len(userTokenAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.User = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				break
			}
		}
	}

	return nil
}

// LoadToken allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (userTokenL) LoadToken(e boil.Executor, singular bool, maybeUserToken interface{}) error {
	var slice []*UserToken
	var object *UserToken

	count := 1
	if singular {
		object = maybeUserToken.(*UserToken)
	} else {
		slice = *maybeUserToken.(*UserTokenSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &userTokenR{}
		}
		args[0] = object.TokenID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &userTokenR{}
			}
			args[i] = obj.TokenID
		}
	}

	query := fmt.Sprintf(
		"select * from \"tokens\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Token")
	}
	defer results.Close()

	var resultSlice []*Token
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Token")
	}

	if len(userTokenAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Token = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.TokenID == foreign.ID {
				local.R.Token = foreign
				break
			}
		}
	}

	return nil
}

// SetUserG of the user_token to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserTokens.
// Uses the global database handle.
func (o *UserToken) SetUserG(insert bool, related *User) error {
	return o.SetUser(boil.GetDB(), insert, related)
}

// SetUserP of the user_token to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserTokens.
// Panics on error.
func (o *UserToken) SetUserP(exec boil.Executor, insert bool, related *User) {
	if err := o.SetUser(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetUserGP of the user_token to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserTokens.
// Uses the global database handle and panics on error.
func (o *UserToken) SetUserGP(insert bool, related *User) {
	if err := o.SetUser(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetUser of the user_token to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserTokens.
func (o *UserToken) SetUser(exec boil.Executor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_tokens\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, userTokenPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID, o.TokenID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID

	if o.R == nil {
		o.R = &userTokenR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserTokens: UserTokenSlice{o},
		}
	} else {
		related.R.UserTokens = append(related.R.UserTokens, o)
	}

	return nil
}

// SetTokenG of the user_token to the related item.
// Sets o.R.Token to related.
// Adds o to related.R.UserTokens.
// Uses the global database handle.
func (o *UserToken) SetTokenG(insert bool, related *Token) error {
	return o.SetToken(boil.GetDB(), insert, related)
}

// SetTokenP of the user_token to the related item.
// Sets o.R.Token to related.
// Adds o to related.R.UserTokens.
// Panics on error.
func (o *UserToken) SetTokenP(exec boil.Executor, insert bool, related *Token) {
	if err := o.SetToken(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetTokenGP of the user_token to the related item.
// Sets o.R.Token to related.
// Adds o to related.R.UserTokens.
// Uses the global database handle and panics on error.
func (o *UserToken) SetTokenGP(insert bool, related *Token) {
	if err := o.SetToken(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetToken of the user_token to the related item.
// Sets o.R.Token to related.
// Adds o to related.R.UserTokens.
func (o *UserToken) SetToken(exec boil.Executor, insert bool, related *Token) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_tokens\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"token_id"}),
		strmangle.WhereClause("\"", "\"", 2, userTokenPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID, o.TokenID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TokenID = related.ID

	if o.R == nil {
		o.R = &userTokenR{
			Token: related,
		}
	} else {
		o.R.Token = related
	}

	if related.R == nil {
		related.R = &tokenR{
			UserTokens: UserTokenSlice{o},
		}
	} else {
		related.R.UserTokens = append(related.R.UserTokens, o)
	}

	return nil
}

// UserTokensG retrieves all records.
func UserTokensG(mods ...qm.QueryMod) userTokenQuery {
	return UserTokens(boil.GetDB(), mods...)
}

// UserTokens retrieves all the records using an executor.
func UserTokens(exec boil.Executor, mods ...qm.QueryMod) userTokenQuery {
	mods = append(mods, qm.From("\"user_tokens\""))
	return userTokenQuery{NewQuery(exec, mods...)}
}

// FindUserTokenG retrieves a single record by ID.
func FindUserTokenG(userID string, tokenID string, selectCols ...string) (*UserToken, error) {
	return FindUserToken(boil.GetDB(), userID, tokenID, selectCols...)
}

// FindUserTokenGP retrieves a single record by ID, and panics on error.
func FindUserTokenGP(userID string, tokenID string, selectCols ...string) *UserToken {
	retobj, err := FindUserToken(boil.GetDB(), userID, tokenID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindUserToken retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserToken(exec boil.Executor, userID string, tokenID string, selectCols ...string) (*UserToken, error) {
	userTokenObj := &UserToken{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_tokens\" where \"user_id\"=$1 AND \"token_id\"=$2", sel,
	)

	q := queries.Raw(exec, query, userID, tokenID)

	err := q.Bind(userTokenObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_tokens")
	}

	return userTokenObj, nil
}

// FindUserTokenP retrieves a single record by ID with an executor, and panics on error.
func FindUserTokenP(exec boil.Executor, userID string, tokenID string, selectCols ...string) *UserToken {
	retobj, err := FindUserToken(exec, userID, tokenID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *UserToken) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *UserToken) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *UserToken) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *UserToken) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no user_tokens provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userTokenColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	userTokenInsertCacheMut.RLock()
	cache, cached := userTokenInsertCache[key]
	userTokenInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			userTokenColumns,
			userTokenColumnsWithDefault,
			userTokenColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(userTokenType, userTokenMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userTokenType, userTokenMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"user_tokens\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into user_tokens")
	}

	if !cached {
		userTokenInsertCacheMut.Lock()
		userTokenInsertCache[key] = cache
		userTokenInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single UserToken record. See Update for
// whitelist behavior description.
func (o *UserToken) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single UserToken record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *UserToken) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the UserToken, and panics on error.
// See Update for whitelist behavior description.
func (o *UserToken) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the UserToken.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *UserToken) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	userTokenUpdateCacheMut.RLock()
	cache, cached := userTokenUpdateCache[key]
	userTokenUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(userTokenColumns, userTokenPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update user_tokens, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_tokens\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userTokenPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userTokenType, userTokenMapping, append(wl, userTokenPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update user_tokens row")
	}

	if !cached {
		userTokenUpdateCacheMut.Lock()
		userTokenUpdateCache[key] = cache
		userTokenUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q userTokenQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q userTokenQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for user_tokens")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o UserTokenSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o UserTokenSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o UserTokenSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserTokenSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"user_tokens\" SET %s WHERE (\"user_id\",\"token_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(userTokenPrimaryKeyColumns), len(colNames)+1, len(userTokenPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in userToken slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *UserToken) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *UserToken) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *UserToken) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *UserToken) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no user_tokens provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userTokenColumnsWithDefault, o)

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

	userTokenUpsertCacheMut.RLock()
	cache, cached := userTokenUpsertCache[key]
	userTokenUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			userTokenColumns,
			userTokenColumnsWithDefault,
			userTokenColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			userTokenColumns,
			userTokenPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert user_tokens, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userTokenPrimaryKeyColumns))
			copy(conflict, userTokenPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"user_tokens\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(userTokenType, userTokenMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userTokenType, userTokenMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert user_tokens")
	}

	if !cached {
		userTokenUpsertCacheMut.Lock()
		userTokenUpsertCache[key] = cache
		userTokenUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single UserToken record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *UserToken) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single UserToken record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *UserToken) DeleteG() error {
	if o == nil {
		return errors.New("models: no UserToken provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single UserToken record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *UserToken) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single UserToken record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserToken) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no UserToken provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userTokenPrimaryKeyMapping)
	sql := "DELETE FROM \"user_tokens\" WHERE \"user_id\"=$1 AND \"token_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from user_tokens")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q userTokenQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q userTokenQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no userTokenQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from user_tokens")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o UserTokenSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o UserTokenSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no UserToken slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o UserTokenSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserTokenSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no UserToken slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(userTokenBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"user_tokens\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, userTokenPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(userTokenPrimaryKeyColumns), 1, len(userTokenPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from userToken slice")
	}

	if len(userTokenAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *UserToken) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *UserToken) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *UserToken) ReloadG() error {
	if o == nil {
		return errors.New("models: no UserToken provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UserToken) Reload(exec boil.Executor) error {
	ret, err := FindUserToken(exec, o.UserID, o.TokenID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *UserTokenSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *UserTokenSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserTokenSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty UserTokenSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserTokenSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	userTokens := UserTokenSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"user_tokens\".* FROM \"user_tokens\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, userTokenPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(userTokenPrimaryKeyColumns), 1, len(userTokenPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&userTokens)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserTokenSlice")
	}

	*o = userTokens

	return nil
}

// UserTokenExists checks if the UserToken row exists.
func UserTokenExists(exec boil.Executor, userID string, tokenID string) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"user_tokens\" where \"user_id\"=$1 AND \"token_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, userID, tokenID)
	}

	row := exec.QueryRow(sql, userID, tokenID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_tokens exists")
	}

	return exists, nil
}

// UserTokenExistsG checks if the UserToken row exists.
func UserTokenExistsG(userID string, tokenID string) (bool, error) {
	return UserTokenExists(boil.GetDB(), userID, tokenID)
}

// UserTokenExistsGP checks if the UserToken row exists. Panics on error.
func UserTokenExistsGP(userID string, tokenID string) bool {
	e, err := UserTokenExists(boil.GetDB(), userID, tokenID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// UserTokenExistsP checks if the UserToken row exists. Panics on error.
func UserTokenExistsP(exec boil.Executor, userID string, tokenID string) bool {
	e, err := UserTokenExists(exec, userID, tokenID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

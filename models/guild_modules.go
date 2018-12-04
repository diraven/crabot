// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// GuildModule is an object representing the database table.
type GuildModule struct {
	ID       int `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuildID  int `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	ModuleID int `boil:"module_id" json:"module_id" toml:"module_id" yaml:"module_id"`

	R *guildModuleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L guildModuleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GuildModuleColumns = struct {
	ID       string
	GuildID  string
	ModuleID string
}{
	ID:       "id",
	GuildID:  "guild_id",
	ModuleID: "module_id",
}

// GuildModuleRels is where relationship names are stored.
var GuildModuleRels = struct {
	Guild  string
	Module string
}{
	Guild:  "Guild",
	Module: "Module",
}

// guildModuleR is where relationships are stored.
type guildModuleR struct {
	Guild  *Guild
	Module *Module
}

// NewStruct creates a new relationship struct
func (*guildModuleR) NewStruct() *guildModuleR {
	return &guildModuleR{}
}

// guildModuleL is where Load methods for each relationship are stored.
type guildModuleL struct{}

var (
	guildModuleColumns               = []string{"id", "guild_id", "module_id"}
	guildModuleColumnsWithoutDefault = []string{"guild_id", "module_id"}
	guildModuleColumnsWithDefault    = []string{"id"}
	guildModulePrimaryKeyColumns     = []string{"id"}
)

type (
	// GuildModuleSlice is an alias for a slice of pointers to GuildModule.
	// This should generally be used opposed to []GuildModule.
	GuildModuleSlice []*GuildModule
	// GuildModuleHook is the signature for custom GuildModule hook methods
	GuildModuleHook func(context.Context, boil.ContextExecutor, *GuildModule) error

	guildModuleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	guildModuleType                 = reflect.TypeOf(&GuildModule{})
	guildModuleMapping              = queries.MakeStructMapping(guildModuleType)
	guildModulePrimaryKeyMapping, _ = queries.BindMapping(guildModuleType, guildModuleMapping, guildModulePrimaryKeyColumns)
	guildModuleInsertCacheMut       sync.RWMutex
	guildModuleInsertCache          = make(map[string]insertCache)
	guildModuleUpdateCacheMut       sync.RWMutex
	guildModuleUpdateCache          = make(map[string]updateCache)
	guildModuleUpsertCacheMut       sync.RWMutex
	guildModuleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var guildModuleBeforeInsertHooks []GuildModuleHook
var guildModuleBeforeUpdateHooks []GuildModuleHook
var guildModuleBeforeDeleteHooks []GuildModuleHook
var guildModuleBeforeUpsertHooks []GuildModuleHook

var guildModuleAfterInsertHooks []GuildModuleHook
var guildModuleAfterSelectHooks []GuildModuleHook
var guildModuleAfterUpdateHooks []GuildModuleHook
var guildModuleAfterDeleteHooks []GuildModuleHook
var guildModuleAfterUpsertHooks []GuildModuleHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *GuildModule) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range guildModuleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *GuildModule) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range guildModuleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *GuildModule) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range guildModuleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *GuildModule) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range guildModuleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *GuildModule) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range guildModuleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *GuildModule) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range guildModuleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *GuildModule) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range guildModuleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *GuildModule) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range guildModuleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *GuildModule) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range guildModuleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGuildModuleHook registers your hook function for all future operations.
func AddGuildModuleHook(hookPoint boil.HookPoint, guildModuleHook GuildModuleHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		guildModuleBeforeInsertHooks = append(guildModuleBeforeInsertHooks, guildModuleHook)
	case boil.BeforeUpdateHook:
		guildModuleBeforeUpdateHooks = append(guildModuleBeforeUpdateHooks, guildModuleHook)
	case boil.BeforeDeleteHook:
		guildModuleBeforeDeleteHooks = append(guildModuleBeforeDeleteHooks, guildModuleHook)
	case boil.BeforeUpsertHook:
		guildModuleBeforeUpsertHooks = append(guildModuleBeforeUpsertHooks, guildModuleHook)
	case boil.AfterInsertHook:
		guildModuleAfterInsertHooks = append(guildModuleAfterInsertHooks, guildModuleHook)
	case boil.AfterSelectHook:
		guildModuleAfterSelectHooks = append(guildModuleAfterSelectHooks, guildModuleHook)
	case boil.AfterUpdateHook:
		guildModuleAfterUpdateHooks = append(guildModuleAfterUpdateHooks, guildModuleHook)
	case boil.AfterDeleteHook:
		guildModuleAfterDeleteHooks = append(guildModuleAfterDeleteHooks, guildModuleHook)
	case boil.AfterUpsertHook:
		guildModuleAfterUpsertHooks = append(guildModuleAfterUpsertHooks, guildModuleHook)
	}
}

// One returns a single guildModule record from the query.
func (q guildModuleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*GuildModule, error) {
	o := &GuildModule{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for guild_modules")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all GuildModule records from the query.
func (q guildModuleQuery) All(ctx context.Context, exec boil.ContextExecutor) (GuildModuleSlice, error) {
	var o []*GuildModule

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to GuildModule slice")
	}

	if len(guildModuleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all GuildModule records in the query.
func (q guildModuleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count guild_modules rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q guildModuleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if guild_modules exists")
	}

	return count > 0, nil
}

// Guild pointed to by the foreign key.
func (o *GuildModule) Guild(mods ...qm.QueryMod) guildQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.GuildID),
	}

	queryMods = append(queryMods, mods...)

	query := Guilds(queryMods...)
	queries.SetFrom(query.Query, "\"guild\"")

	return query
}

// Module pointed to by the foreign key.
func (o *GuildModule) Module(mods ...qm.QueryMod) moduleQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.ModuleID),
	}

	queryMods = append(queryMods, mods...)

	query := Modules(queryMods...)
	queries.SetFrom(query.Query, "\"module\"")

	return query
}

// LoadGuild allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (guildModuleL) LoadGuild(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGuildModule interface{}, mods queries.Applicator) error {
	var slice []*GuildModule
	var object *GuildModule

	if singular {
		object = maybeGuildModule.(*GuildModule)
	} else {
		slice = *maybeGuildModule.(*[]*GuildModule)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &guildModuleR{}
		}
		args = append(args, object.GuildID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &guildModuleR{}
			}

			for _, a := range args {
				if a == obj.GuildID {
					continue Outer
				}
			}

			args = append(args, obj.GuildID)

		}
	}

	query := NewQuery(qm.From(`guild`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Guild")
	}

	var resultSlice []*Guild
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Guild")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for guild")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for guild")
	}

	if len(guildModuleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Guild = foreign
		if foreign.R == nil {
			foreign.R = &guildR{}
		}
		foreign.R.GuildModules = append(foreign.R.GuildModules, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.GuildID == foreign.ID {
				local.R.Guild = foreign
				if foreign.R == nil {
					foreign.R = &guildR{}
				}
				foreign.R.GuildModules = append(foreign.R.GuildModules, local)
				break
			}
		}
	}

	return nil
}

// LoadModule allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (guildModuleL) LoadModule(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGuildModule interface{}, mods queries.Applicator) error {
	var slice []*GuildModule
	var object *GuildModule

	if singular {
		object = maybeGuildModule.(*GuildModule)
	} else {
		slice = *maybeGuildModule.(*[]*GuildModule)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &guildModuleR{}
		}
		args = append(args, object.ModuleID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &guildModuleR{}
			}

			for _, a := range args {
				if a == obj.ModuleID {
					continue Outer
				}
			}

			args = append(args, obj.ModuleID)

		}
	}

	query := NewQuery(qm.From(`module`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Module")
	}

	var resultSlice []*Module
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Module")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for module")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for module")
	}

	if len(guildModuleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Module = foreign
		if foreign.R == nil {
			foreign.R = &moduleR{}
		}
		foreign.R.GuildModules = append(foreign.R.GuildModules, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ModuleID == foreign.ID {
				local.R.Module = foreign
				if foreign.R == nil {
					foreign.R = &moduleR{}
				}
				foreign.R.GuildModules = append(foreign.R.GuildModules, local)
				break
			}
		}
	}

	return nil
}

// SetGuild of the guildModule to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.GuildModules.
func (o *GuildModule) SetGuild(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Guild) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"guild_modules\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"guild_id"}),
		strmangle.WhereClause("\"", "\"", 2, guildModulePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GuildID = related.ID
	if o.R == nil {
		o.R = &guildModuleR{
			Guild: related,
		}
	} else {
		o.R.Guild = related
	}

	if related.R == nil {
		related.R = &guildR{
			GuildModules: GuildModuleSlice{o},
		}
	} else {
		related.R.GuildModules = append(related.R.GuildModules, o)
	}

	return nil
}

// SetModule of the guildModule to the related item.
// Sets o.R.Module to related.
// Adds o to related.R.GuildModules.
func (o *GuildModule) SetModule(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Module) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"guild_modules\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"module_id"}),
		strmangle.WhereClause("\"", "\"", 2, guildModulePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ModuleID = related.ID
	if o.R == nil {
		o.R = &guildModuleR{
			Module: related,
		}
	} else {
		o.R.Module = related
	}

	if related.R == nil {
		related.R = &moduleR{
			GuildModules: GuildModuleSlice{o},
		}
	} else {
		related.R.GuildModules = append(related.R.GuildModules, o)
	}

	return nil
}

// GuildModules retrieves all the records using an executor.
func GuildModules(mods ...qm.QueryMod) guildModuleQuery {
	mods = append(mods, qm.From("\"guild_modules\""))
	return guildModuleQuery{NewQuery(mods...)}
}

// FindGuildModule retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGuildModule(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*GuildModule, error) {
	guildModuleObj := &GuildModule{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"guild_modules\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, guildModuleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from guild_modules")
	}

	return guildModuleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GuildModule) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no guild_modules provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(guildModuleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	guildModuleInsertCacheMut.RLock()
	cache, cached := guildModuleInsertCache[key]
	guildModuleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			guildModuleColumns,
			guildModuleColumnsWithDefault,
			guildModuleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(guildModuleType, guildModuleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(guildModuleType, guildModuleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"guild_modules\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"guild_modules\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into guild_modules")
	}

	if !cached {
		guildModuleInsertCacheMut.Lock()
		guildModuleInsertCache[key] = cache
		guildModuleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the GuildModule.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GuildModule) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	guildModuleUpdateCacheMut.RLock()
	cache, cached := guildModuleUpdateCache[key]
	guildModuleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			guildModuleColumns,
			guildModulePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update guild_modules, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"guild_modules\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, guildModulePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(guildModuleType, guildModuleMapping, append(wl, guildModulePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update guild_modules row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for guild_modules")
	}

	if !cached {
		guildModuleUpdateCacheMut.Lock()
		guildModuleUpdateCache[key] = cache
		guildModuleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q guildModuleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for guild_modules")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for guild_modules")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GuildModuleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), guildModulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"guild_modules\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, guildModulePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in guildModule slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all guildModule")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *GuildModule) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no guild_modules provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(guildModuleColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
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
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	guildModuleUpsertCacheMut.RLock()
	cache, cached := guildModuleUpsertCache[key]
	guildModuleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			guildModuleColumns,
			guildModuleColumnsWithDefault,
			guildModuleColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			guildModuleColumns,
			guildModulePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert guild_modules, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(guildModulePrimaryKeyColumns))
			copy(conflict, guildModulePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"guild_modules\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(guildModuleType, guildModuleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(guildModuleType, guildModuleMapping, ret)
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
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert guild_modules")
	}

	if !cached {
		guildModuleUpsertCacheMut.Lock()
		guildModuleUpsertCache[key] = cache
		guildModuleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single GuildModule record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GuildModule) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no GuildModule provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), guildModulePrimaryKeyMapping)
	sql := "DELETE FROM \"guild_modules\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from guild_modules")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for guild_modules")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q guildModuleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no guildModuleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from guild_modules")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for guild_modules")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GuildModuleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no GuildModule slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(guildModuleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), guildModulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"guild_modules\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, guildModulePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from guildModule slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for guild_modules")
	}

	if len(guildModuleAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *GuildModule) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGuildModule(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GuildModuleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GuildModuleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), guildModulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"guild_modules\".* FROM \"guild_modules\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, guildModulePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in GuildModuleSlice")
	}

	*o = slice

	return nil
}

// GuildModuleExists checks if the GuildModule row exists.
func GuildModuleExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"guild_modules\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if guild_modules exists")
	}

	return exists, nil
}

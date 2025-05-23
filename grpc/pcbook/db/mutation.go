// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/JrMarcco/go-learning/grpc/pcbook/db/laptop"
	"github.com/JrMarcco/go-learning/grpc/pcbook/db/predicate"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeLaptop = "Laptop"
)

// LaptopMutation represents an operation that mutates the Laptop nodes in the graph.
type LaptopMutation struct {
	config
	op              Op
	typ             string
	id              *uint64
	uid             *string
	brand           *string
	laptop_name     *string
	weight          *float64
	addweight       *float64
	price_rmb       *uint32
	addprice_rmb    *int32
	release_year    *uint32
	addrelease_year *int32
	created_at      *time.Time
	updated_at      *time.Time
	clearedFields   map[string]struct{}
	done            bool
	oldValue        func(context.Context) (*Laptop, error)
	predicates      []predicate.Laptop
}

var _ ent.Mutation = (*LaptopMutation)(nil)

// laptopOption allows management of the mutation configuration using functional options.
type laptopOption func(*LaptopMutation)

// newLaptopMutation creates new mutation for the Laptop entity.
func newLaptopMutation(c config, op Op, opts ...laptopOption) *LaptopMutation {
	m := &LaptopMutation{
		config:        c,
		op:            op,
		typ:           TypeLaptop,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withLaptopID sets the ID field of the mutation.
func withLaptopID(id uint64) laptopOption {
	return func(m *LaptopMutation) {
		var (
			err   error
			once  sync.Once
			value *Laptop
		)
		m.oldValue = func(ctx context.Context) (*Laptop, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Laptop.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withLaptop sets the old Laptop of the mutation.
func withLaptop(node *Laptop) laptopOption {
	return func(m *LaptopMutation) {
		m.oldValue = func(context.Context) (*Laptop, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m LaptopMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m LaptopMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("db: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *LaptopMutation) ID() (id uint64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *LaptopMutation) IDs(ctx context.Context) ([]uint64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uint64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Laptop.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUID sets the "uid" field.
func (m *LaptopMutation) SetUID(s string) {
	m.uid = &s
}

// UID returns the value of the "uid" field in the mutation.
func (m *LaptopMutation) UID() (r string, exists bool) {
	v := m.uid
	if v == nil {
		return
	}
	return *v, true
}

// OldUID returns the old "uid" field's value of the Laptop entity.
// If the Laptop object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LaptopMutation) OldUID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUID: %w", err)
	}
	return oldValue.UID, nil
}

// ResetUID resets all changes to the "uid" field.
func (m *LaptopMutation) ResetUID() {
	m.uid = nil
}

// SetBrand sets the "brand" field.
func (m *LaptopMutation) SetBrand(s string) {
	m.brand = &s
}

// Brand returns the value of the "brand" field in the mutation.
func (m *LaptopMutation) Brand() (r string, exists bool) {
	v := m.brand
	if v == nil {
		return
	}
	return *v, true
}

// OldBrand returns the old "brand" field's value of the Laptop entity.
// If the Laptop object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LaptopMutation) OldBrand(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBrand is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBrand requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBrand: %w", err)
	}
	return oldValue.Brand, nil
}

// ResetBrand resets all changes to the "brand" field.
func (m *LaptopMutation) ResetBrand() {
	m.brand = nil
}

// SetLaptopName sets the "laptop_name" field.
func (m *LaptopMutation) SetLaptopName(s string) {
	m.laptop_name = &s
}

// LaptopName returns the value of the "laptop_name" field in the mutation.
func (m *LaptopMutation) LaptopName() (r string, exists bool) {
	v := m.laptop_name
	if v == nil {
		return
	}
	return *v, true
}

// OldLaptopName returns the old "laptop_name" field's value of the Laptop entity.
// If the Laptop object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LaptopMutation) OldLaptopName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLaptopName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLaptopName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLaptopName: %w", err)
	}
	return oldValue.LaptopName, nil
}

// ResetLaptopName resets all changes to the "laptop_name" field.
func (m *LaptopMutation) ResetLaptopName() {
	m.laptop_name = nil
}

// SetWeight sets the "weight" field.
func (m *LaptopMutation) SetWeight(f float64) {
	m.weight = &f
	m.addweight = nil
}

// Weight returns the value of the "weight" field in the mutation.
func (m *LaptopMutation) Weight() (r float64, exists bool) {
	v := m.weight
	if v == nil {
		return
	}
	return *v, true
}

// OldWeight returns the old "weight" field's value of the Laptop entity.
// If the Laptop object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LaptopMutation) OldWeight(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldWeight is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldWeight requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldWeight: %w", err)
	}
	return oldValue.Weight, nil
}

// AddWeight adds f to the "weight" field.
func (m *LaptopMutation) AddWeight(f float64) {
	if m.addweight != nil {
		*m.addweight += f
	} else {
		m.addweight = &f
	}
}

// AddedWeight returns the value that was added to the "weight" field in this mutation.
func (m *LaptopMutation) AddedWeight() (r float64, exists bool) {
	v := m.addweight
	if v == nil {
		return
	}
	return *v, true
}

// ResetWeight resets all changes to the "weight" field.
func (m *LaptopMutation) ResetWeight() {
	m.weight = nil
	m.addweight = nil
}

// SetPriceRmb sets the "price_rmb" field.
func (m *LaptopMutation) SetPriceRmb(u uint32) {
	m.price_rmb = &u
	m.addprice_rmb = nil
}

// PriceRmb returns the value of the "price_rmb" field in the mutation.
func (m *LaptopMutation) PriceRmb() (r uint32, exists bool) {
	v := m.price_rmb
	if v == nil {
		return
	}
	return *v, true
}

// OldPriceRmb returns the old "price_rmb" field's value of the Laptop entity.
// If the Laptop object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LaptopMutation) OldPriceRmb(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPriceRmb is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPriceRmb requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPriceRmb: %w", err)
	}
	return oldValue.PriceRmb, nil
}

// AddPriceRmb adds u to the "price_rmb" field.
func (m *LaptopMutation) AddPriceRmb(u int32) {
	if m.addprice_rmb != nil {
		*m.addprice_rmb += u
	} else {
		m.addprice_rmb = &u
	}
}

// AddedPriceRmb returns the value that was added to the "price_rmb" field in this mutation.
func (m *LaptopMutation) AddedPriceRmb() (r int32, exists bool) {
	v := m.addprice_rmb
	if v == nil {
		return
	}
	return *v, true
}

// ResetPriceRmb resets all changes to the "price_rmb" field.
func (m *LaptopMutation) ResetPriceRmb() {
	m.price_rmb = nil
	m.addprice_rmb = nil
}

// SetReleaseYear sets the "release_year" field.
func (m *LaptopMutation) SetReleaseYear(u uint32) {
	m.release_year = &u
	m.addrelease_year = nil
}

// ReleaseYear returns the value of the "release_year" field in the mutation.
func (m *LaptopMutation) ReleaseYear() (r uint32, exists bool) {
	v := m.release_year
	if v == nil {
		return
	}
	return *v, true
}

// OldReleaseYear returns the old "release_year" field's value of the Laptop entity.
// If the Laptop object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LaptopMutation) OldReleaseYear(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldReleaseYear is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldReleaseYear requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldReleaseYear: %w", err)
	}
	return oldValue.ReleaseYear, nil
}

// AddReleaseYear adds u to the "release_year" field.
func (m *LaptopMutation) AddReleaseYear(u int32) {
	if m.addrelease_year != nil {
		*m.addrelease_year += u
	} else {
		m.addrelease_year = &u
	}
}

// AddedReleaseYear returns the value that was added to the "release_year" field in this mutation.
func (m *LaptopMutation) AddedReleaseYear() (r int32, exists bool) {
	v := m.addrelease_year
	if v == nil {
		return
	}
	return *v, true
}

// ResetReleaseYear resets all changes to the "release_year" field.
func (m *LaptopMutation) ResetReleaseYear() {
	m.release_year = nil
	m.addrelease_year = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *LaptopMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *LaptopMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Laptop entity.
// If the Laptop object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LaptopMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *LaptopMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *LaptopMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *LaptopMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Laptop entity.
// If the Laptop object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LaptopMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *LaptopMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the LaptopMutation builder.
func (m *LaptopMutation) Where(ps ...predicate.Laptop) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *LaptopMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Laptop).
func (m *LaptopMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *LaptopMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.uid != nil {
		fields = append(fields, laptop.FieldUID)
	}
	if m.brand != nil {
		fields = append(fields, laptop.FieldBrand)
	}
	if m.laptop_name != nil {
		fields = append(fields, laptop.FieldLaptopName)
	}
	if m.weight != nil {
		fields = append(fields, laptop.FieldWeight)
	}
	if m.price_rmb != nil {
		fields = append(fields, laptop.FieldPriceRmb)
	}
	if m.release_year != nil {
		fields = append(fields, laptop.FieldReleaseYear)
	}
	if m.created_at != nil {
		fields = append(fields, laptop.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, laptop.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *LaptopMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case laptop.FieldUID:
		return m.UID()
	case laptop.FieldBrand:
		return m.Brand()
	case laptop.FieldLaptopName:
		return m.LaptopName()
	case laptop.FieldWeight:
		return m.Weight()
	case laptop.FieldPriceRmb:
		return m.PriceRmb()
	case laptop.FieldReleaseYear:
		return m.ReleaseYear()
	case laptop.FieldCreatedAt:
		return m.CreatedAt()
	case laptop.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *LaptopMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case laptop.FieldUID:
		return m.OldUID(ctx)
	case laptop.FieldBrand:
		return m.OldBrand(ctx)
	case laptop.FieldLaptopName:
		return m.OldLaptopName(ctx)
	case laptop.FieldWeight:
		return m.OldWeight(ctx)
	case laptop.FieldPriceRmb:
		return m.OldPriceRmb(ctx)
	case laptop.FieldReleaseYear:
		return m.OldReleaseYear(ctx)
	case laptop.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case laptop.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Laptop field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LaptopMutation) SetField(name string, value ent.Value) error {
	switch name {
	case laptop.FieldUID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUID(v)
		return nil
	case laptop.FieldBrand:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBrand(v)
		return nil
	case laptop.FieldLaptopName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLaptopName(v)
		return nil
	case laptop.FieldWeight:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetWeight(v)
		return nil
	case laptop.FieldPriceRmb:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPriceRmb(v)
		return nil
	case laptop.FieldReleaseYear:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetReleaseYear(v)
		return nil
	case laptop.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case laptop.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Laptop field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *LaptopMutation) AddedFields() []string {
	var fields []string
	if m.addweight != nil {
		fields = append(fields, laptop.FieldWeight)
	}
	if m.addprice_rmb != nil {
		fields = append(fields, laptop.FieldPriceRmb)
	}
	if m.addrelease_year != nil {
		fields = append(fields, laptop.FieldReleaseYear)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *LaptopMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case laptop.FieldWeight:
		return m.AddedWeight()
	case laptop.FieldPriceRmb:
		return m.AddedPriceRmb()
	case laptop.FieldReleaseYear:
		return m.AddedReleaseYear()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LaptopMutation) AddField(name string, value ent.Value) error {
	switch name {
	case laptop.FieldWeight:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddWeight(v)
		return nil
	case laptop.FieldPriceRmb:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPriceRmb(v)
		return nil
	case laptop.FieldReleaseYear:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddReleaseYear(v)
		return nil
	}
	return fmt.Errorf("unknown Laptop numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *LaptopMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *LaptopMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *LaptopMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Laptop nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *LaptopMutation) ResetField(name string) error {
	switch name {
	case laptop.FieldUID:
		m.ResetUID()
		return nil
	case laptop.FieldBrand:
		m.ResetBrand()
		return nil
	case laptop.FieldLaptopName:
		m.ResetLaptopName()
		return nil
	case laptop.FieldWeight:
		m.ResetWeight()
		return nil
	case laptop.FieldPriceRmb:
		m.ResetPriceRmb()
		return nil
	case laptop.FieldReleaseYear:
		m.ResetReleaseYear()
		return nil
	case laptop.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case laptop.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Laptop field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *LaptopMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *LaptopMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *LaptopMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *LaptopMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *LaptopMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *LaptopMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *LaptopMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Laptop unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *LaptopMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Laptop edge %s", name)
}

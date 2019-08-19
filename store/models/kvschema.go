// Code generated by "kvschema"; DO NOT EDIT.

package models

import (
	"github.com/google/note-maps/kv"
)

// Txn provides entities, components, and indexes backed by a key-value store.
type Txn struct{ kv.Partitioned }

func New(t kv.Txn) Txn { return Txn{kv.Partitioned{t, 0}} }

// SetIIs sets the IIs associated with e to v.
//
// Corresponding indexes are updated.
func (s Txn) SetIIs(e kv.Entity, v IIs) error {
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	IIsPrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	var old IIs
	if err := s.Get(key, old.Decode); err != nil {
		return err
	}
	if err := s.Set(key, v.Encode()); err != nil {
		return err
	}
	lek := len(key)
	kv.Entity(0).EncodeAt(key[10:])
	key = append(key, kv.Component(0).Encode()...)
	var (
		lik = len(key)
		es  kv.EntitySlice
	)

	// Update Literal index
	key = key[:lek].AppendComponent(LiteralPrefix)
	for _, iv := range old.IndexLiteral() {
		key = append(key[:lik], iv.Encode()...)
		if err := s.Get(key, es.Decode); err != nil {
			return err
		}
		if es.Remove(e) {
			if err := s.Set(key, es.Encode()); err != nil {
				return err
			}
		}
	}
	for _, iv := range v.IndexLiteral() {
		key = append(key[:lik], iv.Encode()...)
		if err := s.Get(key, es.Decode); err != nil {
			return err
		}
		if es.Insert(e) {
			if err := s.Set(key, es.Encode()); err != nil {
				return err
			}
		}
	}
	return nil
}

// DeleteIIs deletes the IIs associated with e.
//
// Corresponding indexes are updated.
func (s Txn) DeleteIIs(e kv.Entity) error {
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	IIsPrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	var old IIs
	if err := s.Get(key, old.Decode); err != nil {
		return err
	}
	if err := s.Delete(key); err != nil {
		return err
	}
	lek := len(key)
	kv.Entity(0).EncodeAt(key[10:])
	key = append(key, kv.Component(0).Encode()...)
	var (
		lik = len(key)
		es  kv.EntitySlice
	)

	// Update Literal index
	key = key[:lek].AppendComponent(LiteralPrefix)
	for _, iv := range old.IndexLiteral() {
		key = append(key[:lik], iv.Encode()...)
		if err := s.Get(key, es.Decode); err != nil {
			return err
		}
		if es.Remove(e) {
			if err := s.Set(key, es.Encode()); err != nil {
				return err
			}
		}
	}
	return nil
}

// GetIIs returns the IIs associated with e.
//
// If no IIs has been explicitly set for e, and GetIIs will return
// the result of decoding a IIs from an empty slice of bytes.
func (s Txn) GetIIs(e kv.Entity) (IIs, error) {
	var v IIs
	vs, err := s.GetIIsSlice([]kv.Entity{e})
	if len(vs) >= 1 {
		v = vs[0]
	}
	return v, err
}

// GetIIsSlice returns a IIs for each entity in es.
//
// If no IIs has been explicitly set for an entity, and the result will
// be a IIs that has been decoded from an empty slice of bytes.
func (s Txn) GetIIsSlice(es []kv.Entity) ([]IIs, error) {
	result := make([]IIs, len(es))
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	IIsPrefix.EncodeAt(key[8:])
	for i, e := range es {
		e.EncodeAt(key[10:])
		err := s.Get(key, (&result[i]).Decode)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// AllIIsEntities returns the first n entities that have a IIs, beginning
// with the first entity greater than or equal to *start.
//
// A nil start value will be interpreted as a pointer to zero.
//
// A value of n less than or equal to zero will be interpretted as the largest
// possible value.
func (s Txn) AllIIsEntities(start *kv.Entity, n int) (es []kv.Entity, err error) {
	return s.AllComponentEntities(IIsPrefix, start, n)
}

// EntitiesMatchingIIsLiteral returns entities with IIs values that return a matching kv.String from their IndexLiteral method.
//
// The returned EntitySlice is already sorted.
func (s Txn) EntitiesMatchingIIsLiteral(v kv.String) (kv.EntitySlice, error) {
	key := make(kv.Prefix, 8+2+8+2)
	s.Partition.EncodeAt(key)
	IIsPrefix.EncodeAt(key[8:])
	kv.Entity(0).EncodeAt(key[10:])
	LiteralPrefix.EncodeAt(key[18:])
	key = append(key, v.Encode()...)
	var es kv.EntitySlice
	return es, s.Get(key, es.Decode)
}

// EntitiesByIIsLiteral returns entities with
// IIs values ordered by the kv.String values from their
// IndexLiteral method.
//
// Reading begins at cursor, and ends when the length of the returned Entity
// slice is less than n. When reading is not complete, cursor is updated such
// that using it in a subequent call to ByLiteral would return next n
// entities.
func (s Txn) EntitiesByIIsLiteral(cursor *kv.IndexCursor, n int) (es []kv.Entity, err error) {
	return s.EntitiesByComponentIndex(IIsPrefix, LiteralPrefix, cursor, n)
}

// SetName sets the Name associated with e to v.
//
// Corresponding indexes are updated.
func (s Txn) SetName(e kv.Entity, v *Name) error {
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	NamePrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	var old Name
	if err := s.Get(key, old.Decode); err != nil {
		return err
	}
	if err := s.Set(key, v.Encode()); err != nil {
		return err
	}
	lek := len(key)
	kv.Entity(0).EncodeAt(key[10:])
	key = append(key, kv.Component(0).Encode()...)
	var (
		lik = len(key)
		es  kv.EntitySlice
	)

	// Update Value index
	key = key[:lek].AppendComponent(ValuePrefix)
	for _, iv := range old.IndexValue() {
		key = append(key[:lik], iv.Encode()...)
		if err := s.Get(key, es.Decode); err != nil {
			return err
		}
		if es.Remove(e) {
			if err := s.Set(key, es.Encode()); err != nil {
				return err
			}
		}
	}
	for _, iv := range v.IndexValue() {
		key = append(key[:lik], iv.Encode()...)
		if err := s.Get(key, es.Decode); err != nil {
			return err
		}
		if es.Insert(e) {
			if err := s.Set(key, es.Encode()); err != nil {
				return err
			}
		}
	}
	return nil
}

// DeleteName deletes the Name associated with e.
//
// Corresponding indexes are updated.
func (s Txn) DeleteName(e kv.Entity) error {
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	NamePrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	var old Name
	if err := s.Get(key, old.Decode); err != nil {
		return err
	}
	if err := s.Delete(key); err != nil {
		return err
	}
	lek := len(key)
	kv.Entity(0).EncodeAt(key[10:])
	key = append(key, kv.Component(0).Encode()...)
	var (
		lik = len(key)
		es  kv.EntitySlice
	)

	// Update Value index
	key = key[:lek].AppendComponent(ValuePrefix)
	for _, iv := range old.IndexValue() {
		key = append(key[:lik], iv.Encode()...)
		if err := s.Get(key, es.Decode); err != nil {
			return err
		}
		if es.Remove(e) {
			if err := s.Set(key, es.Encode()); err != nil {
				return err
			}
		}
	}
	return nil
}

// GetName returns the Name associated with e.
//
// If no Name has been explicitly set for e, and GetName will return
// the result of decoding a Name from an empty slice of bytes.
func (s Txn) GetName(e kv.Entity) (Name, error) {
	var v Name
	vs, err := s.GetNameSlice([]kv.Entity{e})
	if len(vs) >= 1 {
		v = vs[0]
	}
	return v, err
}

// GetNameSlice returns a Name for each entity in es.
//
// If no Name has been explicitly set for an entity, and the result will
// be a Name that has been decoded from an empty slice of bytes.
func (s Txn) GetNameSlice(es []kv.Entity) ([]Name, error) {
	result := make([]Name, len(es))
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	NamePrefix.EncodeAt(key[8:])
	for i, e := range es {
		e.EncodeAt(key[10:])
		err := s.Get(key, (&result[i]).Decode)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// AllNameEntities returns the first n entities that have a Name, beginning
// with the first entity greater than or equal to *start.
//
// A nil start value will be interpreted as a pointer to zero.
//
// A value of n less than or equal to zero will be interpretted as the largest
// possible value.
func (s Txn) AllNameEntities(start *kv.Entity, n int) (es []kv.Entity, err error) {
	return s.AllComponentEntities(NamePrefix, start, n)
}

// EntitiesMatchingNameValue returns entities with Name values that return a matching kv.String from their IndexValue method.
//
// The returned EntitySlice is already sorted.
func (s Txn) EntitiesMatchingNameValue(v kv.String) (kv.EntitySlice, error) {
	key := make(kv.Prefix, 8+2+8+2)
	s.Partition.EncodeAt(key)
	NamePrefix.EncodeAt(key[8:])
	kv.Entity(0).EncodeAt(key[10:])
	ValuePrefix.EncodeAt(key[18:])
	key = append(key, v.Encode()...)
	var es kv.EntitySlice
	return es, s.Get(key, es.Decode)
}

// EntitiesByNameValue returns entities with
// Name values ordered by the kv.String values from their
// IndexValue method.
//
// Reading begins at cursor, and ends when the length of the returned Entity
// slice is less than n. When reading is not complete, cursor is updated such
// that using it in a subequent call to ByValue would return next n
// entities.
func (s Txn) EntitiesByNameValue(cursor *kv.IndexCursor, n int) (es []kv.Entity, err error) {
	return s.EntitiesByComponentIndex(NamePrefix, ValuePrefix, cursor, n)
}

// SetOccurrence sets the Occurrence associated with e to v.
//
// Corresponding indexes are updated.
func (s Txn) SetOccurrence(e kv.Entity, v *Occurrence) error {
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	OccurrencePrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	var old Occurrence
	if err := s.Get(key, old.Decode); err != nil {
		return err
	}
	if err := s.Set(key, v.Encode()); err != nil {
		return err
	}
	lek := len(key)
	kv.Entity(0).EncodeAt(key[10:])
	key = append(key, kv.Component(0).Encode()...)
	var (
		lik = len(key)
		es  kv.EntitySlice
	)

	// Update Value index
	key = key[:lek].AppendComponent(ValuePrefix)
	for _, iv := range old.IndexValue() {
		key = append(key[:lik], iv.Encode()...)
		if err := s.Get(key, es.Decode); err != nil {
			return err
		}
		if es.Remove(e) {
			if err := s.Set(key, es.Encode()); err != nil {
				return err
			}
		}
	}
	for _, iv := range v.IndexValue() {
		key = append(key[:lik], iv.Encode()...)
		if err := s.Get(key, es.Decode); err != nil {
			return err
		}
		if es.Insert(e) {
			if err := s.Set(key, es.Encode()); err != nil {
				return err
			}
		}
	}
	return nil
}

// DeleteOccurrence deletes the Occurrence associated with e.
//
// Corresponding indexes are updated.
func (s Txn) DeleteOccurrence(e kv.Entity) error {
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	OccurrencePrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	var old Occurrence
	if err := s.Get(key, old.Decode); err != nil {
		return err
	}
	if err := s.Delete(key); err != nil {
		return err
	}
	lek := len(key)
	kv.Entity(0).EncodeAt(key[10:])
	key = append(key, kv.Component(0).Encode()...)
	var (
		lik = len(key)
		es  kv.EntitySlice
	)

	// Update Value index
	key = key[:lek].AppendComponent(ValuePrefix)
	for _, iv := range old.IndexValue() {
		key = append(key[:lik], iv.Encode()...)
		if err := s.Get(key, es.Decode); err != nil {
			return err
		}
		if es.Remove(e) {
			if err := s.Set(key, es.Encode()); err != nil {
				return err
			}
		}
	}
	return nil
}

// GetOccurrence returns the Occurrence associated with e.
//
// If no Occurrence has been explicitly set for e, and GetOccurrence will return
// the result of decoding a Occurrence from an empty slice of bytes.
func (s Txn) GetOccurrence(e kv.Entity) (Occurrence, error) {
	var v Occurrence
	vs, err := s.GetOccurrenceSlice([]kv.Entity{e})
	if len(vs) >= 1 {
		v = vs[0]
	}
	return v, err
}

// GetOccurrenceSlice returns a Occurrence for each entity in es.
//
// If no Occurrence has been explicitly set for an entity, and the result will
// be a Occurrence that has been decoded from an empty slice of bytes.
func (s Txn) GetOccurrenceSlice(es []kv.Entity) ([]Occurrence, error) {
	result := make([]Occurrence, len(es))
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	OccurrencePrefix.EncodeAt(key[8:])
	for i, e := range es {
		e.EncodeAt(key[10:])
		err := s.Get(key, (&result[i]).Decode)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// AllOccurrenceEntities returns the first n entities that have a Occurrence, beginning
// with the first entity greater than or equal to *start.
//
// A nil start value will be interpreted as a pointer to zero.
//
// A value of n less than or equal to zero will be interpretted as the largest
// possible value.
func (s Txn) AllOccurrenceEntities(start *kv.Entity, n int) (es []kv.Entity, err error) {
	return s.AllComponentEntities(OccurrencePrefix, start, n)
}

// EntitiesMatchingOccurrenceValue returns entities with Occurrence values that return a matching kv.String from their IndexValue method.
//
// The returned EntitySlice is already sorted.
func (s Txn) EntitiesMatchingOccurrenceValue(v kv.String) (kv.EntitySlice, error) {
	key := make(kv.Prefix, 8+2+8+2)
	s.Partition.EncodeAt(key)
	OccurrencePrefix.EncodeAt(key[8:])
	kv.Entity(0).EncodeAt(key[10:])
	ValuePrefix.EncodeAt(key[18:])
	key = append(key, v.Encode()...)
	var es kv.EntitySlice
	return es, s.Get(key, es.Decode)
}

// EntitiesByOccurrenceValue returns entities with
// Occurrence values ordered by the kv.String values from their
// IndexValue method.
//
// Reading begins at cursor, and ends when the length of the returned Entity
// slice is less than n. When reading is not complete, cursor is updated such
// that using it in a subequent call to ByValue would return next n
// entities.
func (s Txn) EntitiesByOccurrenceValue(cursor *kv.IndexCursor, n int) (es []kv.Entity, err error) {
	return s.EntitiesByComponentIndex(OccurrencePrefix, ValuePrefix, cursor, n)
}

// SetSIs sets the SIs associated with e to v.
//
// Corresponding indexes are updated.
func (s Txn) SetSIs(e kv.Entity, v SIs) error {
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	SIsPrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	var old SIs
	if err := s.Get(key, old.Decode); err != nil {
		return err
	}
	if err := s.Set(key, v.Encode()); err != nil {
		return err
	}
	lek := len(key)
	kv.Entity(0).EncodeAt(key[10:])
	key = append(key, kv.Component(0).Encode()...)
	var (
		lik = len(key)
		es  kv.EntitySlice
	)

	// Update Literal index
	key = key[:lek].AppendComponent(LiteralPrefix)
	for _, iv := range old.IndexLiteral() {
		key = append(key[:lik], iv.Encode()...)
		if err := s.Get(key, es.Decode); err != nil {
			return err
		}
		if es.Remove(e) {
			if err := s.Set(key, es.Encode()); err != nil {
				return err
			}
		}
	}
	for _, iv := range v.IndexLiteral() {
		key = append(key[:lik], iv.Encode()...)
		if err := s.Get(key, es.Decode); err != nil {
			return err
		}
		if es.Insert(e) {
			if err := s.Set(key, es.Encode()); err != nil {
				return err
			}
		}
	}
	return nil
}

// DeleteSIs deletes the SIs associated with e.
//
// Corresponding indexes are updated.
func (s Txn) DeleteSIs(e kv.Entity) error {
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	SIsPrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	var old SIs
	if err := s.Get(key, old.Decode); err != nil {
		return err
	}
	if err := s.Delete(key); err != nil {
		return err
	}
	lek := len(key)
	kv.Entity(0).EncodeAt(key[10:])
	key = append(key, kv.Component(0).Encode()...)
	var (
		lik = len(key)
		es  kv.EntitySlice
	)

	// Update Literal index
	key = key[:lek].AppendComponent(LiteralPrefix)
	for _, iv := range old.IndexLiteral() {
		key = append(key[:lik], iv.Encode()...)
		if err := s.Get(key, es.Decode); err != nil {
			return err
		}
		if es.Remove(e) {
			if err := s.Set(key, es.Encode()); err != nil {
				return err
			}
		}
	}
	return nil
}

// GetSIs returns the SIs associated with e.
//
// If no SIs has been explicitly set for e, and GetSIs will return
// the result of decoding a SIs from an empty slice of bytes.
func (s Txn) GetSIs(e kv.Entity) (SIs, error) {
	var v SIs
	vs, err := s.GetSIsSlice([]kv.Entity{e})
	if len(vs) >= 1 {
		v = vs[0]
	}
	return v, err
}

// GetSIsSlice returns a SIs for each entity in es.
//
// If no SIs has been explicitly set for an entity, and the result will
// be a SIs that has been decoded from an empty slice of bytes.
func (s Txn) GetSIsSlice(es []kv.Entity) ([]SIs, error) {
	result := make([]SIs, len(es))
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	SIsPrefix.EncodeAt(key[8:])
	for i, e := range es {
		e.EncodeAt(key[10:])
		err := s.Get(key, (&result[i]).Decode)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// AllSIsEntities returns the first n entities that have a SIs, beginning
// with the first entity greater than or equal to *start.
//
// A nil start value will be interpreted as a pointer to zero.
//
// A value of n less than or equal to zero will be interpretted as the largest
// possible value.
func (s Txn) AllSIsEntities(start *kv.Entity, n int) (es []kv.Entity, err error) {
	return s.AllComponentEntities(SIsPrefix, start, n)
}

// EntitiesMatchingSIsLiteral returns entities with SIs values that return a matching kv.String from their IndexLiteral method.
//
// The returned EntitySlice is already sorted.
func (s Txn) EntitiesMatchingSIsLiteral(v kv.String) (kv.EntitySlice, error) {
	key := make(kv.Prefix, 8+2+8+2)
	s.Partition.EncodeAt(key)
	SIsPrefix.EncodeAt(key[8:])
	kv.Entity(0).EncodeAt(key[10:])
	LiteralPrefix.EncodeAt(key[18:])
	key = append(key, v.Encode()...)
	var es kv.EntitySlice
	return es, s.Get(key, es.Decode)
}

// EntitiesBySIsLiteral returns entities with
// SIs values ordered by the kv.String values from their
// IndexLiteral method.
//
// Reading begins at cursor, and ends when the length of the returned Entity
// slice is less than n. When reading is not complete, cursor is updated such
// that using it in a subequent call to ByLiteral would return next n
// entities.
func (s Txn) EntitiesBySIsLiteral(cursor *kv.IndexCursor, n int) (es []kv.Entity, err error) {
	return s.EntitiesByComponentIndex(SIsPrefix, LiteralPrefix, cursor, n)
}

// SetSLs sets the SLs associated with e to v.
//
// Corresponding indexes are updated.
func (s Txn) SetSLs(e kv.Entity, v SLs) error {
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	SLsPrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	var old SLs
	if err := s.Get(key, old.Decode); err != nil {
		return err
	}
	if err := s.Set(key, v.Encode()); err != nil {
		return err
	}
	lek := len(key)
	kv.Entity(0).EncodeAt(key[10:])
	key = append(key, kv.Component(0).Encode()...)
	var (
		lik = len(key)
		es  kv.EntitySlice
	)

	// Update Literal index
	key = key[:lek].AppendComponent(LiteralPrefix)
	for _, iv := range old.IndexLiteral() {
		key = append(key[:lik], iv.Encode()...)
		if err := s.Get(key, es.Decode); err != nil {
			return err
		}
		if es.Remove(e) {
			if err := s.Set(key, es.Encode()); err != nil {
				return err
			}
		}
	}
	for _, iv := range v.IndexLiteral() {
		key = append(key[:lik], iv.Encode()...)
		if err := s.Get(key, es.Decode); err != nil {
			return err
		}
		if es.Insert(e) {
			if err := s.Set(key, es.Encode()); err != nil {
				return err
			}
		}
	}
	return nil
}

// DeleteSLs deletes the SLs associated with e.
//
// Corresponding indexes are updated.
func (s Txn) DeleteSLs(e kv.Entity) error {
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	SLsPrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	var old SLs
	if err := s.Get(key, old.Decode); err != nil {
		return err
	}
	if err := s.Delete(key); err != nil {
		return err
	}
	lek := len(key)
	kv.Entity(0).EncodeAt(key[10:])
	key = append(key, kv.Component(0).Encode()...)
	var (
		lik = len(key)
		es  kv.EntitySlice
	)

	// Update Literal index
	key = key[:lek].AppendComponent(LiteralPrefix)
	for _, iv := range old.IndexLiteral() {
		key = append(key[:lik], iv.Encode()...)
		if err := s.Get(key, es.Decode); err != nil {
			return err
		}
		if es.Remove(e) {
			if err := s.Set(key, es.Encode()); err != nil {
				return err
			}
		}
	}
	return nil
}

// GetSLs returns the SLs associated with e.
//
// If no SLs has been explicitly set for e, and GetSLs will return
// the result of decoding a SLs from an empty slice of bytes.
func (s Txn) GetSLs(e kv.Entity) (SLs, error) {
	var v SLs
	vs, err := s.GetSLsSlice([]kv.Entity{e})
	if len(vs) >= 1 {
		v = vs[0]
	}
	return v, err
}

// GetSLsSlice returns a SLs for each entity in es.
//
// If no SLs has been explicitly set for an entity, and the result will
// be a SLs that has been decoded from an empty slice of bytes.
func (s Txn) GetSLsSlice(es []kv.Entity) ([]SLs, error) {
	result := make([]SLs, len(es))
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	SLsPrefix.EncodeAt(key[8:])
	for i, e := range es {
		e.EncodeAt(key[10:])
		err := s.Get(key, (&result[i]).Decode)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// AllSLsEntities returns the first n entities that have a SLs, beginning
// with the first entity greater than or equal to *start.
//
// A nil start value will be interpreted as a pointer to zero.
//
// A value of n less than or equal to zero will be interpretted as the largest
// possible value.
func (s Txn) AllSLsEntities(start *kv.Entity, n int) (es []kv.Entity, err error) {
	return s.AllComponentEntities(SLsPrefix, start, n)
}

// EntitiesMatchingSLsLiteral returns entities with SLs values that return a matching kv.String from their IndexLiteral method.
//
// The returned EntitySlice is already sorted.
func (s Txn) EntitiesMatchingSLsLiteral(v kv.String) (kv.EntitySlice, error) {
	key := make(kv.Prefix, 8+2+8+2)
	s.Partition.EncodeAt(key)
	SLsPrefix.EncodeAt(key[8:])
	kv.Entity(0).EncodeAt(key[10:])
	LiteralPrefix.EncodeAt(key[18:])
	key = append(key, v.Encode()...)
	var es kv.EntitySlice
	return es, s.Get(key, es.Decode)
}

// EntitiesBySLsLiteral returns entities with
// SLs values ordered by the kv.String values from their
// IndexLiteral method.
//
// Reading begins at cursor, and ends when the length of the returned Entity
// slice is less than n. When reading is not complete, cursor is updated such
// that using it in a subequent call to ByLiteral would return next n
// entities.
func (s Txn) EntitiesBySLsLiteral(cursor *kv.IndexCursor, n int) (es []kv.Entity, err error) {
	return s.EntitiesByComponentIndex(SLsPrefix, LiteralPrefix, cursor, n)
}

// SetTopicMapInfo sets the TopicMapInfo associated with e to v.
//
// Corresponding indexes are updated.
func (s Txn) SetTopicMapInfo(e kv.Entity, v *TopicMapInfo) error {
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	TopicMapInfoPrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	return s.Set(key, v.Encode())
}

// DeleteTopicMapInfo deletes the TopicMapInfo associated with e.
//
// Corresponding indexes are updated.
func (s Txn) DeleteTopicMapInfo(e kv.Entity) error {
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	TopicMapInfoPrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	return s.Delete(key)
}

// GetTopicMapInfo returns the TopicMapInfo associated with e.
//
// If no TopicMapInfo has been explicitly set for e, and GetTopicMapInfo will return
// the result of decoding a TopicMapInfo from an empty slice of bytes.
func (s Txn) GetTopicMapInfo(e kv.Entity) (TopicMapInfo, error) {
	var v TopicMapInfo
	vs, err := s.GetTopicMapInfoSlice([]kv.Entity{e})
	if len(vs) >= 1 {
		v = vs[0]
	}
	return v, err
}

// GetTopicMapInfoSlice returns a TopicMapInfo for each entity in es.
//
// If no TopicMapInfo has been explicitly set for an entity, and the result will
// be a TopicMapInfo that has been decoded from an empty slice of bytes.
func (s Txn) GetTopicMapInfoSlice(es []kv.Entity) ([]TopicMapInfo, error) {
	result := make([]TopicMapInfo, len(es))
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	TopicMapInfoPrefix.EncodeAt(key[8:])
	for i, e := range es {
		e.EncodeAt(key[10:])
		err := s.Get(key, (&result[i]).Decode)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// AllTopicMapInfoEntities returns the first n entities that have a TopicMapInfo, beginning
// with the first entity greater than or equal to *start.
//
// A nil start value will be interpreted as a pointer to zero.
//
// A value of n less than or equal to zero will be interpretted as the largest
// possible value.
func (s Txn) AllTopicMapInfoEntities(start *kv.Entity, n int) (es []kv.Entity, err error) {
	return s.AllComponentEntities(TopicMapInfoPrefix, start, n)
}

// SetTopicNames sets the TopicNames associated with e to v.
//
// Corresponding indexes are updated.
func (s Txn) SetTopicNames(e kv.Entity, v TopicNames) error {
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	TopicNamesPrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	return s.Set(key, v.Encode())
}

// DeleteTopicNames deletes the TopicNames associated with e.
//
// Corresponding indexes are updated.
func (s Txn) DeleteTopicNames(e kv.Entity) error {
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	TopicNamesPrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	return s.Delete(key)
}

// GetTopicNames returns the TopicNames associated with e.
//
// If no TopicNames has been explicitly set for e, and GetTopicNames will return
// the result of decoding a TopicNames from an empty slice of bytes.
func (s Txn) GetTopicNames(e kv.Entity) (TopicNames, error) {
	var v TopicNames
	vs, err := s.GetTopicNamesSlice([]kv.Entity{e})
	if len(vs) >= 1 {
		v = vs[0]
	}
	return v, err
}

// GetTopicNamesSlice returns a TopicNames for each entity in es.
//
// If no TopicNames has been explicitly set for an entity, and the result will
// be a TopicNames that has been decoded from an empty slice of bytes.
func (s Txn) GetTopicNamesSlice(es []kv.Entity) ([]TopicNames, error) {
	result := make([]TopicNames, len(es))
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	TopicNamesPrefix.EncodeAt(key[8:])
	for i, e := range es {
		e.EncodeAt(key[10:])
		err := s.Get(key, (&result[i]).Decode)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// AllTopicNamesEntities returns the first n entities that have a TopicNames, beginning
// with the first entity greater than or equal to *start.
//
// A nil start value will be interpreted as a pointer to zero.
//
// A value of n less than or equal to zero will be interpretted as the largest
// possible value.
func (s Txn) AllTopicNamesEntities(start *kv.Entity, n int) (es []kv.Entity, err error) {
	return s.AllComponentEntities(TopicNamesPrefix, start, n)
}

// SetTopicOccurrences sets the TopicOccurrences associated with e to v.
//
// Corresponding indexes are updated.
func (s Txn) SetTopicOccurrences(e kv.Entity, v TopicOccurrences) error {
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	TopicOccurrencesPrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	return s.Set(key, v.Encode())
}

// DeleteTopicOccurrences deletes the TopicOccurrences associated with e.
//
// Corresponding indexes are updated.
func (s Txn) DeleteTopicOccurrences(e kv.Entity) error {
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	TopicOccurrencesPrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	return s.Delete(key)
}

// GetTopicOccurrences returns the TopicOccurrences associated with e.
//
// If no TopicOccurrences has been explicitly set for e, and GetTopicOccurrences will return
// the result of decoding a TopicOccurrences from an empty slice of bytes.
func (s Txn) GetTopicOccurrences(e kv.Entity) (TopicOccurrences, error) {
	var v TopicOccurrences
	vs, err := s.GetTopicOccurrencesSlice([]kv.Entity{e})
	if len(vs) >= 1 {
		v = vs[0]
	}
	return v, err
}

// GetTopicOccurrencesSlice returns a TopicOccurrences for each entity in es.
//
// If no TopicOccurrences has been explicitly set for an entity, and the result will
// be a TopicOccurrences that has been decoded from an empty slice of bytes.
func (s Txn) GetTopicOccurrencesSlice(es []kv.Entity) ([]TopicOccurrences, error) {
	result := make([]TopicOccurrences, len(es))
	key := make(kv.Prefix, 8+2+8)
	s.Partition.EncodeAt(key)
	TopicOccurrencesPrefix.EncodeAt(key[8:])
	for i, e := range es {
		e.EncodeAt(key[10:])
		err := s.Get(key, (&result[i]).Decode)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// AllTopicOccurrencesEntities returns the first n entities that have a TopicOccurrences, beginning
// with the first entity greater than or equal to *start.
//
// A nil start value will be interpreted as a pointer to zero.
//
// A value of n less than or equal to zero will be interpretted as the largest
// possible value.
func (s Txn) AllTopicOccurrencesEntities(start *kv.Entity, n int) (es []kv.Entity, err error) {
	return s.AllComponentEntities(TopicOccurrencesPrefix, start, n)
}

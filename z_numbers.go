// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package column

import (
	"sync"

	"github.com/kelindar/bitmap"
)

// --------------------------- float32s ----------------------------

// columnFloat32 represents a generic column
type columnFloat32 struct {
	sync.RWMutex
	fill bitmap.Bitmap // The fill-list
	data []float32     // The actual values
}

// makeFloat32s creates a new vector or float32s
func makeFloat32s() Column {
	return &columnFloat32{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]float32, 0, 64),
	}
}

// Update sets a value at a specified index
func (c *columnFloat32) Update(idx uint32, value interface{}) {
	c.Lock()
	defer c.Unlock()

	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}

	// Set the data at index
	c.fill.Set(idx)
	c.data[idx] = value.(float32)
}

// UpdateMany atomically updates a set of key/value pairs in the columns.
func (c *columnFloat32) UpdateMany(keys *bitmap.Bitmap, values []interface{}) {
	c.Lock()
	defer c.Unlock()

	i := 0
	c.fill.Or(*keys)
	keys.Range(func(idx uint32) bool {
		c.data[idx] = values[i].(float32)
		i++
		return true
	})
}

// Value retrieves a value at a specified index
func (c *columnFloat32) Value(idx uint32) (v interface{}, ok bool) {
	v = float32(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnFloat32) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnFloat32) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnFloat32) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Delete removes a value at a specified index
func (c *columnFloat32) Delete(idx uint32) {
	c.Lock()
	c.fill.Remove(idx)
	c.Unlock()
}

// DeleteMany deletes a set of items from the column.
func (c *columnFloat32) DeleteMany(items *bitmap.Bitmap) {
	c.Lock()
	c.fill.AndNot(*items)
	c.Unlock()
}

// Contains checks whether the column has a value at a specified index.
func (c *columnFloat32) Contains(idx uint32) bool {
	c.RLock()
	defer c.RUnlock()
	return c.fill.Contains(idx)
}

// And performs a logical and operation and updates the destination bitmap.
func (c *columnFloat32) And(dst *bitmap.Bitmap) {
	c.RLock()
	dst.And(c.fill)
	c.RUnlock()
}

// And performs a logical and not operation and updates the destination bitmap.
func (c *columnFloat32) AndNot(dst *bitmap.Bitmap) {
	c.RLock()
	dst.AndNot(c.fill)
	c.RUnlock()
}

// Or performs a logical or operation and updates the destination bitmap.
func (c *columnFloat32) Or(dst *bitmap.Bitmap) {
	c.RLock()
	dst.Or(c.fill)
	c.RUnlock()
}

// --------------------------- float64s ----------------------------

// columnFloat64 represents a generic column
type columnFloat64 struct {
	sync.RWMutex
	fill bitmap.Bitmap // The fill-list
	data []float64     // The actual values
}

// makeFloat64s creates a new vector or float64s
func makeFloat64s() Column {
	return &columnFloat64{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]float64, 0, 64),
	}
}

// Update sets a value at a specified index
func (c *columnFloat64) Update(idx uint32, value interface{}) {
	c.Lock()
	defer c.Unlock()

	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}

	// Set the data at index
	c.fill.Set(idx)
	c.data[idx] = value.(float64)
}

// UpdateMany atomically updates a set of key/value pairs in the columns.
func (c *columnFloat64) UpdateMany(keys *bitmap.Bitmap, values []interface{}) {
	c.Lock()
	defer c.Unlock()

	i := 0
	c.fill.Or(*keys)
	keys.Range(func(idx uint32) bool {
		c.data[idx] = values[i].(float64)
		i++
		return true
	})
}

// Value retrieves a value at a specified index
func (c *columnFloat64) Value(idx uint32) (v interface{}, ok bool) {
	v = float64(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnFloat64) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnFloat64) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnFloat64) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Delete removes a value at a specified index
func (c *columnFloat64) Delete(idx uint32) {
	c.Lock()
	c.fill.Remove(idx)
	c.Unlock()
}

// DeleteMany deletes a set of items from the column.
func (c *columnFloat64) DeleteMany(items *bitmap.Bitmap) {
	c.Lock()
	c.fill.AndNot(*items)
	c.Unlock()
}

// Contains checks whether the column has a value at a specified index.
func (c *columnFloat64) Contains(idx uint32) bool {
	c.RLock()
	defer c.RUnlock()
	return c.fill.Contains(idx)
}

// And performs a logical and operation and updates the destination bitmap.
func (c *columnFloat64) And(dst *bitmap.Bitmap) {
	c.RLock()
	dst.And(c.fill)
	c.RUnlock()
}

// And performs a logical and not operation and updates the destination bitmap.
func (c *columnFloat64) AndNot(dst *bitmap.Bitmap) {
	c.RLock()
	dst.AndNot(c.fill)
	c.RUnlock()
}

// Or performs a logical or operation and updates the destination bitmap.
func (c *columnFloat64) Or(dst *bitmap.Bitmap) {
	c.RLock()
	dst.Or(c.fill)
	c.RUnlock()
}

// --------------------------- ints ----------------------------

// columnInt represents a generic column
type columnInt struct {
	sync.RWMutex
	fill bitmap.Bitmap // The fill-list
	data []int         // The actual values
}

// makeInts creates a new vector or ints
func makeInts() Column {
	return &columnInt{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]int, 0, 64),
	}
}

// Update sets a value at a specified index
func (c *columnInt) Update(idx uint32, value interface{}) {
	c.Lock()
	defer c.Unlock()

	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}

	// Set the data at index
	c.fill.Set(idx)
	c.data[idx] = value.(int)
}

// UpdateMany atomically updates a set of key/value pairs in the columns.
func (c *columnInt) UpdateMany(keys *bitmap.Bitmap, values []interface{}) {
	c.Lock()
	defer c.Unlock()

	i := 0
	c.fill.Or(*keys)
	keys.Range(func(idx uint32) bool {
		c.data[idx] = values[i].(int)
		i++
		return true
	})
}

// Value retrieves a value at a specified index
func (c *columnInt) Value(idx uint32) (v interface{}, ok bool) {
	v = int(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnInt) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnInt) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnInt) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Delete removes a value at a specified index
func (c *columnInt) Delete(idx uint32) {
	c.Lock()
	c.fill.Remove(idx)
	c.Unlock()
}

// DeleteMany deletes a set of items from the column.
func (c *columnInt) DeleteMany(items *bitmap.Bitmap) {
	c.Lock()
	c.fill.AndNot(*items)
	c.Unlock()
}

// Contains checks whether the column has a value at a specified index.
func (c *columnInt) Contains(idx uint32) bool {
	c.RLock()
	defer c.RUnlock()
	return c.fill.Contains(idx)
}

// And performs a logical and operation and updates the destination bitmap.
func (c *columnInt) And(dst *bitmap.Bitmap) {
	c.RLock()
	dst.And(c.fill)
	c.RUnlock()
}

// And performs a logical and not operation and updates the destination bitmap.
func (c *columnInt) AndNot(dst *bitmap.Bitmap) {
	c.RLock()
	dst.AndNot(c.fill)
	c.RUnlock()
}

// Or performs a logical or operation and updates the destination bitmap.
func (c *columnInt) Or(dst *bitmap.Bitmap) {
	c.RLock()
	dst.Or(c.fill)
	c.RUnlock()
}

// --------------------------- int16s ----------------------------

// columnInt16 represents a generic column
type columnInt16 struct {
	sync.RWMutex
	fill bitmap.Bitmap // The fill-list
	data []int16       // The actual values
}

// makeInt16s creates a new vector or int16s
func makeInt16s() Column {
	return &columnInt16{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]int16, 0, 64),
	}
}

// Update sets a value at a specified index
func (c *columnInt16) Update(idx uint32, value interface{}) {
	c.Lock()
	defer c.Unlock()

	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}

	// Set the data at index
	c.fill.Set(idx)
	c.data[idx] = value.(int16)
}

// UpdateMany atomically updates a set of key/value pairs in the columns.
func (c *columnInt16) UpdateMany(keys *bitmap.Bitmap, values []interface{}) {
	c.Lock()
	defer c.Unlock()

	i := 0
	c.fill.Or(*keys)
	keys.Range(func(idx uint32) bool {
		c.data[idx] = values[i].(int16)
		i++
		return true
	})
}

// Value retrieves a value at a specified index
func (c *columnInt16) Value(idx uint32) (v interface{}, ok bool) {
	v = int16(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnInt16) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnInt16) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnInt16) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Delete removes a value at a specified index
func (c *columnInt16) Delete(idx uint32) {
	c.Lock()
	c.fill.Remove(idx)
	c.Unlock()
}

// DeleteMany deletes a set of items from the column.
func (c *columnInt16) DeleteMany(items *bitmap.Bitmap) {
	c.Lock()
	c.fill.AndNot(*items)
	c.Unlock()
}

// Contains checks whether the column has a value at a specified index.
func (c *columnInt16) Contains(idx uint32) bool {
	c.RLock()
	defer c.RUnlock()
	return c.fill.Contains(idx)
}

// And performs a logical and operation and updates the destination bitmap.
func (c *columnInt16) And(dst *bitmap.Bitmap) {
	c.RLock()
	dst.And(c.fill)
	c.RUnlock()
}

// And performs a logical and not operation and updates the destination bitmap.
func (c *columnInt16) AndNot(dst *bitmap.Bitmap) {
	c.RLock()
	dst.AndNot(c.fill)
	c.RUnlock()
}

// Or performs a logical or operation and updates the destination bitmap.
func (c *columnInt16) Or(dst *bitmap.Bitmap) {
	c.RLock()
	dst.Or(c.fill)
	c.RUnlock()
}

// --------------------------- int32s ----------------------------

// columnInt32 represents a generic column
type columnInt32 struct {
	sync.RWMutex
	fill bitmap.Bitmap // The fill-list
	data []int32       // The actual values
}

// makeInt32s creates a new vector or int32s
func makeInt32s() Column {
	return &columnInt32{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]int32, 0, 64),
	}
}

// Update sets a value at a specified index
func (c *columnInt32) Update(idx uint32, value interface{}) {
	c.Lock()
	defer c.Unlock()

	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}

	// Set the data at index
	c.fill.Set(idx)
	c.data[idx] = value.(int32)
}

// UpdateMany atomically updates a set of key/value pairs in the columns.
func (c *columnInt32) UpdateMany(keys *bitmap.Bitmap, values []interface{}) {
	c.Lock()
	defer c.Unlock()

	i := 0
	c.fill.Or(*keys)
	keys.Range(func(idx uint32) bool {
		c.data[idx] = values[i].(int32)
		i++
		return true
	})
}

// Value retrieves a value at a specified index
func (c *columnInt32) Value(idx uint32) (v interface{}, ok bool) {
	v = int32(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnInt32) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnInt32) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnInt32) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Delete removes a value at a specified index
func (c *columnInt32) Delete(idx uint32) {
	c.Lock()
	c.fill.Remove(idx)
	c.Unlock()
}

// DeleteMany deletes a set of items from the column.
func (c *columnInt32) DeleteMany(items *bitmap.Bitmap) {
	c.Lock()
	c.fill.AndNot(*items)
	c.Unlock()
}

// Contains checks whether the column has a value at a specified index.
func (c *columnInt32) Contains(idx uint32) bool {
	c.RLock()
	defer c.RUnlock()
	return c.fill.Contains(idx)
}

// And performs a logical and operation and updates the destination bitmap.
func (c *columnInt32) And(dst *bitmap.Bitmap) {
	c.RLock()
	dst.And(c.fill)
	c.RUnlock()
}

// And performs a logical and not operation and updates the destination bitmap.
func (c *columnInt32) AndNot(dst *bitmap.Bitmap) {
	c.RLock()
	dst.AndNot(c.fill)
	c.RUnlock()
}

// Or performs a logical or operation and updates the destination bitmap.
func (c *columnInt32) Or(dst *bitmap.Bitmap) {
	c.RLock()
	dst.Or(c.fill)
	c.RUnlock()
}

// --------------------------- int64s ----------------------------

// columnInt64 represents a generic column
type columnInt64 struct {
	sync.RWMutex
	fill bitmap.Bitmap // The fill-list
	data []int64       // The actual values
}

// makeInt64s creates a new vector or int64s
func makeInt64s() Column {
	return &columnInt64{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]int64, 0, 64),
	}
}

// Update sets a value at a specified index
func (c *columnInt64) Update(idx uint32, value interface{}) {
	c.Lock()
	defer c.Unlock()

	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}

	// Set the data at index
	c.fill.Set(idx)
	c.data[idx] = value.(int64)
}

// UpdateMany atomically updates a set of key/value pairs in the columns.
func (c *columnInt64) UpdateMany(keys *bitmap.Bitmap, values []interface{}) {
	c.Lock()
	defer c.Unlock()

	i := 0
	c.fill.Or(*keys)
	keys.Range(func(idx uint32) bool {
		c.data[idx] = values[i].(int64)
		i++
		return true
	})
}

// Value retrieves a value at a specified index
func (c *columnInt64) Value(idx uint32) (v interface{}, ok bool) {
	v = int64(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnInt64) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnInt64) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnInt64) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Delete removes a value at a specified index
func (c *columnInt64) Delete(idx uint32) {
	c.Lock()
	c.fill.Remove(idx)
	c.Unlock()
}

// DeleteMany deletes a set of items from the column.
func (c *columnInt64) DeleteMany(items *bitmap.Bitmap) {
	c.Lock()
	c.fill.AndNot(*items)
	c.Unlock()
}

// Contains checks whether the column has a value at a specified index.
func (c *columnInt64) Contains(idx uint32) bool {
	c.RLock()
	defer c.RUnlock()
	return c.fill.Contains(idx)
}

// And performs a logical and operation and updates the destination bitmap.
func (c *columnInt64) And(dst *bitmap.Bitmap) {
	c.RLock()
	dst.And(c.fill)
	c.RUnlock()
}

// And performs a logical and not operation and updates the destination bitmap.
func (c *columnInt64) AndNot(dst *bitmap.Bitmap) {
	c.RLock()
	dst.AndNot(c.fill)
	c.RUnlock()
}

// Or performs a logical or operation and updates the destination bitmap.
func (c *columnInt64) Or(dst *bitmap.Bitmap) {
	c.RLock()
	dst.Or(c.fill)
	c.RUnlock()
}

// --------------------------- uints ----------------------------

// columnUint represents a generic column
type columnUint struct {
	sync.RWMutex
	fill bitmap.Bitmap // The fill-list
	data []uint        // The actual values
}

// makeUints creates a new vector or uints
func makeUints() Column {
	return &columnUint{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]uint, 0, 64),
	}
}

// Update sets a value at a specified index
func (c *columnUint) Update(idx uint32, value interface{}) {
	c.Lock()
	defer c.Unlock()

	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}

	// Set the data at index
	c.fill.Set(idx)
	c.data[idx] = value.(uint)
}

// UpdateMany atomically updates a set of key/value pairs in the columns.
func (c *columnUint) UpdateMany(keys *bitmap.Bitmap, values []interface{}) {
	c.Lock()
	defer c.Unlock()

	i := 0
	c.fill.Or(*keys)
	keys.Range(func(idx uint32) bool {
		c.data[idx] = values[i].(uint)
		i++
		return true
	})
}

// Value retrieves a value at a specified index
func (c *columnUint) Value(idx uint32) (v interface{}, ok bool) {
	v = uint(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnUint) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnUint) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnUint) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Delete removes a value at a specified index
func (c *columnUint) Delete(idx uint32) {
	c.Lock()
	c.fill.Remove(idx)
	c.Unlock()
}

// DeleteMany deletes a set of items from the column.
func (c *columnUint) DeleteMany(items *bitmap.Bitmap) {
	c.Lock()
	c.fill.AndNot(*items)
	c.Unlock()
}

// Contains checks whether the column has a value at a specified index.
func (c *columnUint) Contains(idx uint32) bool {
	c.RLock()
	defer c.RUnlock()
	return c.fill.Contains(idx)
}

// And performs a logical and operation and updates the destination bitmap.
func (c *columnUint) And(dst *bitmap.Bitmap) {
	c.RLock()
	dst.And(c.fill)
	c.RUnlock()
}

// And performs a logical and not operation and updates the destination bitmap.
func (c *columnUint) AndNot(dst *bitmap.Bitmap) {
	c.RLock()
	dst.AndNot(c.fill)
	c.RUnlock()
}

// Or performs a logical or operation and updates the destination bitmap.
func (c *columnUint) Or(dst *bitmap.Bitmap) {
	c.RLock()
	dst.Or(c.fill)
	c.RUnlock()
}

// --------------------------- uint16s ----------------------------

// columnUint16 represents a generic column
type columnUint16 struct {
	sync.RWMutex
	fill bitmap.Bitmap // The fill-list
	data []uint16      // The actual values
}

// makeUint16s creates a new vector or uint16s
func makeUint16s() Column {
	return &columnUint16{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]uint16, 0, 64),
	}
}

// Update sets a value at a specified index
func (c *columnUint16) Update(idx uint32, value interface{}) {
	c.Lock()
	defer c.Unlock()

	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}

	// Set the data at index
	c.fill.Set(idx)
	c.data[idx] = value.(uint16)
}

// UpdateMany atomically updates a set of key/value pairs in the columns.
func (c *columnUint16) UpdateMany(keys *bitmap.Bitmap, values []interface{}) {
	c.Lock()
	defer c.Unlock()

	i := 0
	c.fill.Or(*keys)
	keys.Range(func(idx uint32) bool {
		c.data[idx] = values[i].(uint16)
		i++
		return true
	})
}

// Value retrieves a value at a specified index
func (c *columnUint16) Value(idx uint32) (v interface{}, ok bool) {
	v = uint16(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnUint16) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnUint16) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnUint16) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Delete removes a value at a specified index
func (c *columnUint16) Delete(idx uint32) {
	c.Lock()
	c.fill.Remove(idx)
	c.Unlock()
}

// DeleteMany deletes a set of items from the column.
func (c *columnUint16) DeleteMany(items *bitmap.Bitmap) {
	c.Lock()
	c.fill.AndNot(*items)
	c.Unlock()
}

// Contains checks whether the column has a value at a specified index.
func (c *columnUint16) Contains(idx uint32) bool {
	c.RLock()
	defer c.RUnlock()
	return c.fill.Contains(idx)
}

// And performs a logical and operation and updates the destination bitmap.
func (c *columnUint16) And(dst *bitmap.Bitmap) {
	c.RLock()
	dst.And(c.fill)
	c.RUnlock()
}

// And performs a logical and not operation and updates the destination bitmap.
func (c *columnUint16) AndNot(dst *bitmap.Bitmap) {
	c.RLock()
	dst.AndNot(c.fill)
	c.RUnlock()
}

// Or performs a logical or operation and updates the destination bitmap.
func (c *columnUint16) Or(dst *bitmap.Bitmap) {
	c.RLock()
	dst.Or(c.fill)
	c.RUnlock()
}

// --------------------------- uint32s ----------------------------

// columnUint32 represents a generic column
type columnUint32 struct {
	sync.RWMutex
	fill bitmap.Bitmap // The fill-list
	data []uint32      // The actual values
}

// makeUint32s creates a new vector or uint32s
func makeUint32s() Column {
	return &columnUint32{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]uint32, 0, 64),
	}
}

// Update sets a value at a specified index
func (c *columnUint32) Update(idx uint32, value interface{}) {
	c.Lock()
	defer c.Unlock()

	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}

	// Set the data at index
	c.fill.Set(idx)
	c.data[idx] = value.(uint32)
}

// UpdateMany atomically updates a set of key/value pairs in the columns.
func (c *columnUint32) UpdateMany(keys *bitmap.Bitmap, values []interface{}) {
	c.Lock()
	defer c.Unlock()

	i := 0
	c.fill.Or(*keys)
	keys.Range(func(idx uint32) bool {
		c.data[idx] = values[i].(uint32)
		i++
		return true
	})
}

// Value retrieves a value at a specified index
func (c *columnUint32) Value(idx uint32) (v interface{}, ok bool) {
	v = uint32(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnUint32) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnUint32) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnUint32) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Delete removes a value at a specified index
func (c *columnUint32) Delete(idx uint32) {
	c.Lock()
	c.fill.Remove(idx)
	c.Unlock()
}

// DeleteMany deletes a set of items from the column.
func (c *columnUint32) DeleteMany(items *bitmap.Bitmap) {
	c.Lock()
	c.fill.AndNot(*items)
	c.Unlock()
}

// Contains checks whether the column has a value at a specified index.
func (c *columnUint32) Contains(idx uint32) bool {
	c.RLock()
	defer c.RUnlock()
	return c.fill.Contains(idx)
}

// And performs a logical and operation and updates the destination bitmap.
func (c *columnUint32) And(dst *bitmap.Bitmap) {
	c.RLock()
	dst.And(c.fill)
	c.RUnlock()
}

// And performs a logical and not operation and updates the destination bitmap.
func (c *columnUint32) AndNot(dst *bitmap.Bitmap) {
	c.RLock()
	dst.AndNot(c.fill)
	c.RUnlock()
}

// Or performs a logical or operation and updates the destination bitmap.
func (c *columnUint32) Or(dst *bitmap.Bitmap) {
	c.RLock()
	dst.Or(c.fill)
	c.RUnlock()
}

// --------------------------- uint64s ----------------------------

// columnUint64 represents a generic column
type columnUint64 struct {
	sync.RWMutex
	fill bitmap.Bitmap // The fill-list
	data []uint64      // The actual values
}

// makeUint64s creates a new vector or uint64s
func makeUint64s() Column {
	return &columnUint64{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]uint64, 0, 64),
	}
}

// Update sets a value at a specified index
func (c *columnUint64) Update(idx uint32, value interface{}) {
	c.Lock()
	defer c.Unlock()

	size := uint32(len(c.data))
	for i := size; i <= idx; i++ {
		c.data = append(c.data, 0)
	}

	// Set the data at index
	c.fill.Set(idx)
	c.data[idx] = value.(uint64)
}

// UpdateMany atomically updates a set of key/value pairs in the columns.
func (c *columnUint64) UpdateMany(keys *bitmap.Bitmap, values []interface{}) {
	c.Lock()
	defer c.Unlock()

	i := 0
	c.fill.Or(*keys)
	keys.Range(func(idx uint32) bool {
		c.data[idx] = values[i].(uint64)
		i++
		return true
	})
}

// Value retrieves a value at a specified index
func (c *columnUint64) Value(idx uint32) (v interface{}, ok bool) {
	v = uint64(0)
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = c.data[idx], true
	}
	c.RUnlock()
	return
}

// Float64 retrieves a float64 value at a specified index
func (c *columnUint64) Float64(idx uint32) (v float64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = float64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Int64 retrieves an int64 value at a specified index
func (c *columnUint64) Int64(idx uint32) (v int64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = int64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Uint64 retrieves an uint64 value at a specified index
func (c *columnUint64) Uint64(idx uint32) (v uint64, ok bool) {
	c.RLock()
	if idx < uint32(len(c.data)) && c.fill.Contains(idx) {
		v, ok = uint64(c.data[idx]), true
	}
	c.RUnlock()
	return
}

// Delete removes a value at a specified index
func (c *columnUint64) Delete(idx uint32) {
	c.Lock()
	c.fill.Remove(idx)
	c.Unlock()
}

// DeleteMany deletes a set of items from the column.
func (c *columnUint64) DeleteMany(items *bitmap.Bitmap) {
	c.Lock()
	c.fill.AndNot(*items)
	c.Unlock()
}

// Contains checks whether the column has a value at a specified index.
func (c *columnUint64) Contains(idx uint32) bool {
	c.RLock()
	defer c.RUnlock()
	return c.fill.Contains(idx)
}

// And performs a logical and operation and updates the destination bitmap.
func (c *columnUint64) And(dst *bitmap.Bitmap) {
	c.RLock()
	dst.And(c.fill)
	c.RUnlock()
}

// And performs a logical and not operation and updates the destination bitmap.
func (c *columnUint64) AndNot(dst *bitmap.Bitmap) {
	c.RLock()
	dst.AndNot(c.fill)
	c.RUnlock()
}

// Or performs a logical or operation and updates the destination bitmap.
func (c *columnUint64) Or(dst *bitmap.Bitmap) {
	c.RLock()
	dst.Or(c.fill)
	c.RUnlock()
}

// Code generated by gen, DO NOT EDIT.

package id

import (
	"encoding/json"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestNewLayerID(t *testing.T) {
	id := NewLayerID()
	assert.NotNil(t, id)
	u, err := ulid.Parse(id.String())
	assert.NotNil(t, u)
	assert.Nil(t, err)
}

func TestLayerIDFrom(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected struct {
			result LayerID
			err    error
		}
	}{
		{
			name:  "Fail:Not valid string",
			input: "testMustFail",
			expected: struct {
				result LayerID
				err    error
			}{
				result: LayerID{},
				err:    ErrInvalidID,
			},
		},
		{
			name:  "Fail:Not valid string",
			input: "",
			expected: struct {
				result LayerID
				err    error
			}{
				result: LayerID{},
				err:    ErrInvalidID,
			},
		},
		{
			name:  "success:valid string",
			input: "01f2r7kg1fvvffp0gmexgy5hxy",
			expected: struct {
				result LayerID
				err    error
			}{
				result: LayerID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
				err:    nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, err := LayerIDFrom(tt.input)
			assert.Equal(t, tt.expected.result, result)
			if tt.expected.err != nil {
				assert.Equal(t, tt.expected.err, err)
			}
		})
	}
}

func TestMustLayerID(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		shouldPanic bool
		expected    LayerID
	}{
		{
			name:        "Fail:Not valid string",
			input:       "testMustFail",
			shouldPanic: true,
		},
		{
			name:        "Fail:Not valid string",
			input:       "",
			shouldPanic: true,
		},
		{
			name:        "success:valid string",
			input:       "01f2r7kg1fvvffp0gmexgy5hxy",
			shouldPanic: false,
			expected:    LayerID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.shouldPanic {
				assert.Panics(t, func() { MustBeID(tt.input) })
				return
			}
			result := MustLayerID(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLayerIDFromRef(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *LayerID
	}{
		{
			name:     "Fail:Not valid string",
			input:    "testMustFail",
			expected: nil,
		},
		{
			name:     "Fail:Not valid string",
			input:    "",
			expected: nil,
		},
		{
			name:     "success:valid string",
			input:    "01f2r7kg1fvvffp0gmexgy5hxy",
			expected: &LayerID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := LayerIDFromRef(&tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLayerIDFromRefID(t *testing.T) {
	id := New()
	id2 := LayerIDFromRefID(&id)
	assert.Equal(t, id.id, id2.id)
	assert.Nil(t, LayerIDFromRefID(nil))
	assert.Nil(t, LayerIDFromRefID(&ID{}))
}

func TestLayerID_ID(t *testing.T) {
	id := New()
	id2 := LayerIDFromRefID(&id)
	assert.Equal(t, id, id2.ID())
}

func TestLayerID_String(t *testing.T) {
	id := New()
	id2 := LayerIDFromRefID(&id)
	assert.Equal(t, id.String(), id2.String())
	assert.Equal(t, "", LayerID{}.String())
}

func TestLayerID_RefString(t *testing.T) {
	id := NewLayerID()
	assert.Equal(t, id.String(), *id.RefString())
	assert.Nil(t, LayerID{}.RefString())
}

func TestLayerID_GoString(t *testing.T) {
	id := New()
	id2 := LayerIDFromRefID(&id)
	assert.Equal(t, "LayerID("+id.String()+")", id2.GoString())
	assert.Equal(t, "LayerID()", LayerID{}.GoString())
}

func TestLayerID_Ref(t *testing.T) {
	id := NewLayerID()
	assert.Equal(t, LayerID(id), *id.Ref())
	assert.Nil(t, (&LayerID{}).Ref())
}

func TestLayerID_Contains(t *testing.T) {
	id := NewLayerID()
	id2 := NewLayerID()
	assert.True(t, id.Contains([]LayerID{id, id2}))
	assert.False(t, LayerID{}.Contains([]LayerID{id, id2, {}}))
	assert.False(t, id.Contains([]LayerID{id2}))
}

func TestLayerID_CopyRef(t *testing.T) {
	id := NewLayerID().Ref()
	id2 := id.CopyRef()
	assert.Equal(t, id, id2)
	assert.NotSame(t, id, id2)
	assert.Nil(t, (*LayerID)(nil).CopyRef())
}

func TestLayerID_IDRef(t *testing.T) {
	id := New()
	id2 := LayerIDFromRefID(&id)
	assert.Equal(t, &id, id2.IDRef())
	assert.Nil(t, (&LayerID{}).IDRef())
	assert.Nil(t, (*LayerID)(nil).IDRef())
}

func TestLayerID_StringRef(t *testing.T) {
	id := NewLayerID()
	assert.Equal(t, id.String(), *id.StringRef())
	assert.Nil(t, (&LayerID{}).StringRef())
	assert.Nil(t, (*LayerID)(nil).StringRef())
}

func TestLayerID_MarhsalJSON(t *testing.T) {
	id := NewLayerID()
	res, err := id.MarhsalJSON()
	assert.Nil(t, err)
	exp, _ := json.Marshal(id.String())
	assert.Equal(t, exp, res)

	res, err = (&LayerID{}).MarhsalJSON()
	assert.Nil(t, err)
	assert.Nil(t, res)

	res, err = (*LayerID)(nil).MarhsalJSON()
	assert.Nil(t, err)
	assert.Nil(t, res)
}

func TestLayerID_UnmarhsalJSON(t *testing.T) {
	jsonString := "\"01f3zhkysvcxsnzepyyqtq21fb\""
	id := MustLayerID("01f3zhkysvcxsnzepyyqtq21fb")
	id2 := &LayerID{}
	err := id2.UnmarhsalJSON([]byte(jsonString))
	assert.Nil(t, err)
	assert.Equal(t, id, *id2)
}

func TestLayerID_MarshalText(t *testing.T) {
	id := New()
	res, err := LayerIDFromRefID(&id).MarshalText()
	assert.Nil(t, err)
	assert.Equal(t, []byte(id.String()), res)

	res, err = (&LayerID{}).MarshalText()
	assert.Nil(t, err)
	assert.Nil(t, res)

	res, err = (*LayerID)(nil).MarshalText()
	assert.Nil(t, err)
	assert.Nil(t, res)
}

func TestLayerID_UnmarshalText(t *testing.T) {
	text := []byte("01f3zhcaq35403zdjnd6dcm0t2")
	id2 := &LayerID{}
	err := id2.UnmarshalText(text)
	assert.Nil(t, err)
	assert.Equal(t, "01f3zhcaq35403zdjnd6dcm0t2", id2.String())
}

func TestLayerID_IsNil(t *testing.T) {
	assert.True(t, LayerID{}.IsNil())
	assert.False(t, NewLayerID().IsNil())
}

func TestLayerID_IsNilRef(t *testing.T) {
	assert.True(t, LayerID{}.Ref().IsNilRef())
	assert.True(t, (*LayerID)(nil).IsNilRef())
	assert.False(t, NewLayerID().Ref().IsNilRef())
}

func TestLayerIDsToStrings(t *testing.T) {
	tests := []struct {
		name     string
		input    []LayerID
		expected []string
	}{
		{
			name:     "Empty slice",
			input:    make([]LayerID, 0),
			expected: make([]string, 0),
		},
		{
			name:     "1 element",
			input:    []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
		},
		{
			name: "multiple elements",
			input: []LayerID{
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []string{
				"01f3zhcaq35403zdjnd6dcm0t1",
				"01f3zhcaq35403zdjnd6dcm0t2",
				"01f3zhcaq35403zdjnd6dcm0t3",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, LayerIDsToStrings(tt.input))
		})
	}
}

func TestLayerIDsFrom(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected struct {
			res []LayerID
			err error
		}
	}{
		{
			name:  "Empty slice",
			input: make([]string, 0),
			expected: struct {
				res []LayerID
				err error
			}{
				res: make([]LayerID, 0),
				err: nil,
			},
		},
		{
			name:  "1 element",
			input: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
			expected: struct {
				res []LayerID
				err error
			}{
				res: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t2")},
				err: nil,
			},
		},
		{
			name: "multiple elements",
			input: []string{
				"01f3zhcaq35403zdjnd6dcm0t1",
				"01f3zhcaq35403zdjnd6dcm0t2",
				"01f3zhcaq35403zdjnd6dcm0t3",
			},
			expected: struct {
				res []LayerID
				err error
			}{
				res: []LayerID{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
				err: nil,
			},
		},
		{
			name: "error",
			input: []string{
				"01f3zhcaq35403zdjnd6dcm0t1",
				"x",
				"01f3zhcaq35403zdjnd6dcm0t3",
			},
			expected: struct {
				res []LayerID
				err error
			}{
				res: nil,
				err: ErrInvalidID,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := LayerIDsFrom(tc.input)
			if tc.expected.err != nil {
				assert.Equal(t, tc.expected.err, err)
				assert.Nil(t, res)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expected.res, res)
			}
		})
	}
}

func TestLayerIDsFromID(t *testing.T) {
	tests := []struct {
		name     string
		input    []ID
		expected []LayerID
	}{
		{
			name:     "Empty slice",
			input:    make([]ID, 0),
			expected: make([]LayerID, 0),
		},
		{
			name:     "1 element",
			input:    []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []ID{
				MustBeID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []LayerID{
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := LayerIDsFromID(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestLayerIDsFromIDRef(t *testing.T) {
	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")

	tests := []struct {
		name     string
		input    []*ID
		expected []LayerID
	}{
		{
			name:     "Empty slice",
			input:    make([]*ID, 0),
			expected: make([]LayerID, 0),
		},
		{
			name:     "1 element",
			input:    []*ID{&id1},
			expected: []LayerID{MustLayerID(id1.String())},
		},
		{
			name:  "multiple elements",
			input: []*ID{&id1, &id2, &id3},
			expected: []LayerID{
				MustLayerID(id1.String()),
				MustLayerID(id2.String()),
				MustLayerID(id3.String()),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := LayerIDsFromIDRef(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestLayerIDsToID(t *testing.T) {
	tests := []struct {
		name     string
		input    []LayerID
		expected []ID
	}{
		{
			name:     "Empty slice",
			input:    make([]LayerID, 0),
			expected: make([]ID, 0),
		},
		{
			name:     "1 element",
			input:    []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []LayerID{
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []ID{
				MustBeID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := LayerIDsToID(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestLayerIDsToIDRef(t *testing.T) {
	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	id21 := MustLayerID(id1.String())
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	id22 := MustLayerID(id2.String())
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")
	id23 := MustLayerID(id3.String())

	tests := []struct {
		name     string
		input    []*LayerID
		expected []*ID
	}{
		{
			name:     "Empty slice",
			input:    make([]*LayerID, 0),
			expected: make([]*ID, 0),
		},
		{
			name:     "1 element",
			input:    []*LayerID{&id21},
			expected: []*ID{&id1},
		},
		{
			name:     "multiple elements",
			input:    []*LayerID{&id21, &id22, &id23},
			expected: []*ID{&id1, &id2, &id3},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := LayerIDsToIDRef(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestNewLayerIDSet(t *testing.T) {
	LayerIdSet := NewLayerIDSet()
	assert.NotNil(t, LayerIdSet)
	assert.Empty(t, LayerIdSet.m)
	assert.Empty(t, LayerIdSet.s)
}

func TestLayerIDSet_Add(t *testing.T) {
	tests := []struct {
		name     string
		input    []LayerID
		expected *LayerIDSet
	}{
		{
			name:  "Empty slice",
			input: make([]LayerID, 0),
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{},
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: []LayerID{
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []LayerID{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
		{
			name: "multiple elements with duplication",
			input: []LayerID{
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []LayerID{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			set := NewLayerIDSet()
			set.Add(tc.input...)
			assert.Equal(t, tc.expected, set)
		})
	}
}

func TestLayerIDSet_AddRef(t *testing.T) {
	tests := []struct {
		name     string
		input    *LayerID
		expected *LayerIDSet
	}{
		{
			name:  "Empty slice",
			input: nil,
			expected: &LayerIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: MustLayerID("01f3zhcaq35403zdjnd6dcm0t1").Ref(),
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			set := NewLayerIDSet()
			set.AddRef(tc.input)
			assert.Equal(t, tc.expected, set)
		})
	}
}

func TestLayerIDSet_Has(t *testing.T) {
	tests := []struct {
		name     string
		target   *LayerIDSet
		input    LayerID
		expected bool
	}{
		{
			name:     "Empty Set",
			target:   &LayerIDSet{},
			input:    MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
			expected: false,
		},
		{
			name: "Set Contains the element",
			target: &LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			input:    MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
			expected: true,
		},
		{
			name: "Set does not Contains the element",
			target: &LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			input:    MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
			expected: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, tc.target.Has(tc.input))
		})
	}
}

func TestLayerIDSet_Clear(t *testing.T) {
	tests := []struct {
		name     string
		input    *LayerIDSet
		expected *LayerIDSet
	}{
		{
			name:     "Empty set",
			input:    &LayerIDSet{},
			expected: &LayerIDSet{},
		},
		{
			name:     "Nil set",
			input:    nil,
			expected: nil,
		},
		{
			name: "Contains the element",
			input: &LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: &LayerIDSet{
				m: nil,
				s: nil,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.input.Clear()
			assert.Equal(t, tc.expected, tc.input)
		})
	}
}

func TestLayerIDSet_All(t *testing.T) {
	tests := []struct {
		name     string
		input    *LayerIDSet
		expected []LayerID
	}{
		{
			name: "Empty",
			input: &LayerIDSet{
				m: map[LayerID]struct{}{},
				s: nil,
			},
			expected: make([]LayerID, 0),
		},
		{
			name:     "Nil",
			input:    nil,
			expected: nil,
		},
		{
			name: "1 element",
			input: &LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
		},
		{
			name: "multiple elements",
			input: &LayerIDSet{
				m: map[LayerID]struct{}{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []LayerID{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: []LayerID{
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, tc.input.All())
		})
	}
}

func TestLayerIDSet_Clone(t *testing.T) {
	tests := []struct {
		name     string
		input    *LayerIDSet
		expected *LayerIDSet
	}{
		{
			name:     "nil set",
			input:    nil,
			expected: NewLayerIDSet(),
		},
		{
			name:     "Empty set",
			input:    NewLayerIDSet(),
			expected: NewLayerIDSet(),
		},
		{
			name: "1 element",
			input: &LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: &LayerIDSet{
				m: map[LayerID]struct{}{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []LayerID{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []LayerID{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			clone := tc.input.Clone()
			assert.Equal(t, tc.expected, clone)
			assert.NotSame(t, tc.input, clone)
		})
	}
}

func TestLayerIDSet_Merge(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			a *LayerIDSet
			b *LayerIDSet
		}
		expected *LayerIDSet
	}{
		{
			name: "Nil Set",
			input: struct {
				a *LayerIDSet
				b *LayerIDSet
			}{
				a: &LayerIDSet{
					m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: nil,
			},
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "Empty Set",
			input: struct {
				a *LayerIDSet
				b *LayerIDSet
			}{
				a: &LayerIDSet{},
				b: &LayerIDSet{},
			},
			expected: &LayerIDSet{},
		},
		{
			name: "1 Empty Set",
			input: struct {
				a *LayerIDSet
				b *LayerIDSet
			}{
				a: &LayerIDSet{
					m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &LayerIDSet{},
			},
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "2 non Empty Set",
			input: struct {
				a *LayerIDSet
				b *LayerIDSet
			}{
				a: &LayerIDSet{
					m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &LayerIDSet{
					m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"): {}},
					s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t2")},
				},
			},
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"): {},
				},
				s: []LayerID{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, tc.input.a.Merge(tc.input.b))
		})
	}
}

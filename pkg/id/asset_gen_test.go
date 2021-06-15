// Code generated by gen, DO NOT EDIT.

package id

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestNewAssetID(t *testing.T) {
	id := NewAssetID()
	assert.NotNil(t, id)
	ulID, err := ulid.Parse(id.String())

	assert.NotNil(t, ulID)
	assert.Nil(t, err)
}

func TestAssetIDFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    string
		expected struct {
			result AssetID
			err    error
		}
	}{
		{
			name:  "Fail:Not valid string",
			input: "testMustFail",
			expected: struct {
				result AssetID
				err    error
			}{
				AssetID{},
				ErrInvalidID,
			},
		},
		{
			name:  "Fail:Not valid string",
			input: "",
			expected: struct {
				result AssetID
				err    error
			}{
				AssetID{},
				ErrInvalidID,
			},
		},
		{
			name:  "success:valid string",
			input: "01f2r7kg1fvvffp0gmexgy5hxy",
			expected: struct {
				result AssetID
				err    error
			}{
				AssetID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
				nil,
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result, err := AssetIDFrom(tc.input)
			assert.Equal(tt, tc.expected.result, result)
			if err != nil {
				assert.True(tt, errors.As(tc.expected.err, &err))
			}
		})
	}
}

func TestMustAssetID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		input       string
		shouldPanic bool
		expected    AssetID
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
			expected:    AssetID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			if tc.shouldPanic {
				assert.Panics(tt, func() { MustBeID(tc.input) })
				return
			}
			result := MustAssetID(tc.input)
			assert.Equal(tt, tc.expected, result)
		})
	}
}

func TestAssetIDFromRef(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected *AssetID
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
			expected: &AssetID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result := AssetIDFromRef(&tc.input)
			assert.Equal(tt, tc.expected, result)
			if tc.expected != nil {
				assert.Equal(tt, *tc.expected, *result)
			}
		})
	}
}

func TestAssetIDFromRefID(t *testing.T) {
	id := New()

	subId := AssetIDFromRefID(&id)

	assert.NotNil(t, subId)
	assert.Equal(t, subId.id, id.id)
}

func TestAssetID_ID(t *testing.T) {
	id := New()
	subId := AssetIDFromRefID(&id)

	idOrg := subId.ID()

	assert.Equal(t, id, idOrg)
}

func TestAssetID_String(t *testing.T) {
	id := New()
	subId := AssetIDFromRefID(&id)

	assert.Equal(t, subId.String(), id.String())
}

func TestAssetID_GoString(t *testing.T) {
	id := New()
	subId := AssetIDFromRefID(&id)

	assert.Equal(t, subId.GoString(), "id.AssetID("+id.String()+")")
}

func TestAssetID_RefString(t *testing.T) {
	id := New()
	subId := AssetIDFromRefID(&id)

	refString := subId.StringRef()

	assert.NotNil(t, refString)
	assert.Equal(t, *refString, id.String())
}

func TestAssetID_Ref(t *testing.T) {
	id := New()
	subId := AssetIDFromRefID(&id)

	subIdRef := subId.Ref()

	assert.Equal(t, *subId, *subIdRef)
}

func TestAssetID_CopyRef(t *testing.T) {
	id := New()
	subId := AssetIDFromRefID(&id)

	subIdCopyRef := subId.CopyRef()

	assert.Equal(t, *subId, *subIdCopyRef)
	assert.NotSame(t, subId, subIdCopyRef)
}

func TestAssetID_IDRef(t *testing.T) {
	id := New()
	subId := AssetIDFromRefID(&id)

	assert.Equal(t, id, *subId.IDRef())
}

func TestAssetID_StringRef(t *testing.T) {
	id := New()
	subId := AssetIDFromRefID(&id)

	assert.Equal(t, *subId.StringRef(), id.String())
}

func TestAssetID_MarhsalJSON(t *testing.T) {
	id := New()
	subId := AssetIDFromRefID(&id)

	res, err := subId.MarhsalJSON()
	exp, _ := json.Marshal(subId.String())

	assert.Nil(t, err)
	assert.Equal(t, exp, res)
}

func TestAssetID_UnmarhsalJSON(t *testing.T) {
	jsonString := "\"01f3zhkysvcxsnzepyyqtq21fb\""

	subId := &AssetID{}

	err := subId.UnmarhsalJSON([]byte(jsonString))

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhkysvcxsnzepyyqtq21fb", subId.String())
}

func TestAssetID_MarshalText(t *testing.T) {
	id := New()
	subId := AssetIDFromRefID(&id)

	res, err := subId.MarshalText()

	assert.Nil(t, err)
	assert.Equal(t, []byte(id.String()), res)
}

func TestAssetID_UnmarshalText(t *testing.T) {
	text := []byte("01f3zhcaq35403zdjnd6dcm0t2")

	subId := &AssetID{}

	err := subId.UnmarshalText(text)

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhcaq35403zdjnd6dcm0t2", subId.String())

}

func TestAssetID_IsNil(t *testing.T) {
	subId := AssetID{}

	assert.True(t, subId.IsNil())

	id := New()
	subId = *AssetIDFromRefID(&id)

	assert.False(t, subId.IsNil())
}

func TestAssetIDToKeys(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []AssetID
		expected []string
	}{
		{
			name:     "Empty slice",
			input:    make([]AssetID, 0),
			expected: make([]string, 0),
		},
		{
			name:     "1 element",
			input:    []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
		},
		{
			name: "multiple elements",
			input: []AssetID{
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []string{
				"01f3zhcaq35403zdjnd6dcm0t1",
				"01f3zhcaq35403zdjnd6dcm0t2",
				"01f3zhcaq35403zdjnd6dcm0t3",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.expected, AssetIDToKeys(tc.input))
		})
	}

}

func TestAssetIDsFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []string
		expected struct {
			res []AssetID
			err error
		}
	}{
		{
			name:  "Empty slice",
			input: make([]string, 0),
			expected: struct {
				res []AssetID
				err error
			}{
				res: make([]AssetID, 0),
				err: nil,
			},
		},
		{
			name:  "1 element",
			input: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
			expected: struct {
				res []AssetID
				err error
			}{
				res: []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t2")},
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
				res []AssetID
				err error
			}{
				res: []AssetID{
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
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
				res []AssetID
				err error
			}{
				res: nil,
				err: ErrInvalidID,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			if tc.expected.err != nil {
				_, err := AssetIDsFrom(tc.input)
				assert.True(tt, errors.As(ErrInvalidID, &err))
			} else {
				res, err := AssetIDsFrom(tc.input)
				assert.Equal(tt, tc.expected.res, res)
				assert.Nil(tt, err)
			}

		})
	}
}

func TestAssetIDsFromID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []ID
		expected []AssetID
	}{
		{
			name:     "Empty slice",
			input:    make([]ID, 0),
			expected: make([]AssetID, 0),
		},
		{
			name:     "1 element",
			input:    []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []ID{
				MustBeID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []AssetID{
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := AssetIDsFromID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestAssetIDsFromIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")

	testCases := []struct {
		name     string
		input    []*ID
		expected []AssetID
	}{
		{
			name:     "Empty slice",
			input:    make([]*ID, 0),
			expected: make([]AssetID, 0),
		},
		{
			name:     "1 element",
			input:    []*ID{&id1},
			expected: []AssetID{MustAssetID(id1.String())},
		},
		{
			name:  "multiple elements",
			input: []*ID{&id1, &id2, &id3},
			expected: []AssetID{
				MustAssetID(id1.String()),
				MustAssetID(id2.String()),
				MustAssetID(id3.String()),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := AssetIDsFromIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestAssetIDsToID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []AssetID
		expected []ID
	}{
		{
			name:     "Empty slice",
			input:    make([]AssetID, 0),
			expected: make([]ID, 0),
		},
		{
			name:     "1 element",
			input:    []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []AssetID{
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []ID{
				MustBeID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := AssetIDsToID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestAssetIDsToIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	subId1 := MustAssetID(id1.String())
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	subId2 := MustAssetID(id2.String())
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")
	subId3 := MustAssetID(id3.String())

	testCases := []struct {
		name     string
		input    []*AssetID
		expected []*ID
	}{
		{
			name:     "Empty slice",
			input:    make([]*AssetID, 0),
			expected: make([]*ID, 0),
		},
		{
			name:     "1 element",
			input:    []*AssetID{&subId1},
			expected: []*ID{&id1},
		},
		{
			name:     "multiple elements",
			input:    []*AssetID{&subId1, &subId2, &subId3},
			expected: []*ID{&id1, &id2, &id3},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := AssetIDsToIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestNewAssetIDSet(t *testing.T) {
	AssetIdSet := NewAssetIDSet()

	assert.NotNil(t, AssetIdSet)
	assert.Empty(t, AssetIdSet.m)
	assert.Empty(t, AssetIdSet.s)
}

func TestAssetIDSet_Add(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []AssetID
		expected *AssetIDSet
	}{
		{
			name:  "Empty slice",
			input: make([]AssetID, 0),
			expected: &AssetIDSet{
				m: map[AssetID]struct{}{},
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: &AssetIDSet{
				m: map[AssetID]struct{}{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: []AssetID{
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &AssetIDSet{
				m: map[AssetID]struct{}{
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []AssetID{
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
		{
			name: "multiple elements with duplication",
			input: []AssetID{
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &AssetIDSet{
				m: map[AssetID]struct{}{
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []AssetID{
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewAssetIDSet()
			set.Add(tc.input...)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestAssetIDSet_AddRef(t *testing.T) {
	t.Parallel()

	AssetId := MustAssetID("01f3zhcaq35403zdjnd6dcm0t1")

	testCases := []struct {
		name     string
		input    *AssetID
		expected *AssetIDSet
	}{
		{
			name:  "Empty slice",
			input: nil,
			expected: &AssetIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: &AssetId,
			expected: &AssetIDSet{
				m: map[AssetID]struct{}{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewAssetIDSet()
			set.AddRef(tc.input)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestAssetIDSet_Has(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			AssetIDSet
			AssetID
		}
		expected bool
	}{
		{
			name: "Empty Set",
			input: struct {
				AssetIDSet
				AssetID
			}{AssetIDSet: AssetIDSet{}, AssetID: MustAssetID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: false,
		},
		{
			name: "Set Contains the element",
			input: struct {
				AssetIDSet
				AssetID
			}{AssetIDSet: AssetIDSet{
				m: map[AssetID]struct{}{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, AssetID: MustAssetID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: true,
		},
		{
			name: "Set does not Contains the element",
			input: struct {
				AssetIDSet
				AssetID
			}{AssetIDSet: AssetIDSet{
				m: map[AssetID]struct{}{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, AssetID: MustAssetID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.expected, tc.input.AssetIDSet.Has(tc.input.AssetID))
		})
	}
}

func TestAssetIDSet_Clear(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    AssetIDSet
		expected AssetIDSet
	}{
		{
			name:  "Empty Set",
			input: AssetIDSet{},
			expected: AssetIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name: "Set Contains the element",
			input: AssetIDSet{
				m: map[AssetID]struct{}{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: AssetIDSet{
				m: nil,
				s: nil,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			set := tc.input
			p := &set
			p.Clear()
			assert.Equal(tt, tc.expected, *p)
		})
	}
}

func TestAssetIDSet_All(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *AssetIDSet
		expected []AssetID
	}{
		{
			name: "Empty slice",
			input: &AssetIDSet{
				m: map[AssetID]struct{}{},
				s: nil,
			},
			expected: make([]AssetID, 0),
		},
		{
			name: "1 element",
			input: &AssetIDSet{
				m: map[AssetID]struct{}{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1")},
		},
		{
			name: "multiple elements",
			input: &AssetIDSet{
				m: map[AssetID]struct{}{
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []AssetID{
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: []AssetID{
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustAssetID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			assert.Equal(tt, tc.expected, tc.input.All())
		})
	}
}

func TestAssetIDSet_Clone(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *AssetIDSet
		expected *AssetIDSet
	}{
		{
			name:     "nil set",
			input:    nil,
			expected: NewAssetIDSet(),
		},
		{
			name:     "Empty set",
			input:    NewAssetIDSet(),
			expected: NewAssetIDSet(),
		},
		{
			name: "1 element",
			input: &AssetIDSet{
				m: map[AssetID]struct{}{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: &AssetIDSet{
				m: map[AssetID]struct{}{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: &AssetIDSet{
				m: map[AssetID]struct{}{
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []AssetID{
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: &AssetIDSet{
				m: map[AssetID]struct{}{
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []AssetID{
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			clone := tc.input.Clone()
			assert.Equal(tt, tc.expected, clone)
			assert.False(tt, tc.input == clone)
		})
	}
}

func TestAssetIDSet_Merge(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			a *AssetIDSet
			b *AssetIDSet
		}
		expected *AssetIDSet
	}{
		{
			name: "Empty Set",
			input: struct {
				a *AssetIDSet
				b *AssetIDSet
			}{
				a: &AssetIDSet{},
				b: &AssetIDSet{},
			},
			expected: &AssetIDSet{},
		},
		{
			name: "1 Empty Set",
			input: struct {
				a *AssetIDSet
				b *AssetIDSet
			}{
				a: &AssetIDSet{
					m: map[AssetID]struct{}{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
					s: []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &AssetIDSet{},
			},
			expected: &AssetIDSet{
				m: map[AssetID]struct{}{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "2 non Empty Set",
			input: struct {
				a *AssetIDSet
				b *AssetIDSet
			}{
				a: &AssetIDSet{
					m: map[AssetID]struct{}{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
					s: []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &AssetIDSet{
					m: map[AssetID]struct{}{MustAssetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{}},
					s: []AssetID{MustAssetID("01f3zhcaq35403zdjnd6dcm0t2")},
				},
			},
			expected: &AssetIDSet{
				m: map[AssetID]struct{}{
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
				},
				s: []AssetID{
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustAssetID("01f3zhcaq35403zdjnd6dcm0t2"),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			assert.Equal(tt, tc.expected, tc.input.a.Merge(tc.input.b))
		})
	}
}

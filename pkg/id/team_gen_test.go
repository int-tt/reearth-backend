// Code generated by gen, DO NOT EDIT.

package id

import (
	"encoding/json"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestNewTeamID(t *testing.T) {
	id := NewTeamID()
	assert.NotNil(t, id)
	u, err := ulid.Parse(id.String())
	assert.NotNil(t, u)
	assert.Nil(t, err)
}

func TestTeamIDFrom(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected struct {
			result TeamID
			err    error
		}
	}{
		{
			name:  "Fail:Not valid string",
			input: "testMustFail",
			expected: struct {
				result TeamID
				err    error
			}{
				result: TeamID{},
				err:    ErrInvalidID,
			},
		},
		{
			name:  "Fail:Not valid string",
			input: "",
			expected: struct {
				result TeamID
				err    error
			}{
				result: TeamID{},
				err:    ErrInvalidID,
			},
		},
		{
			name:  "success:valid string",
			input: "01f2r7kg1fvvffp0gmexgy5hxy",
			expected: struct {
				result TeamID
				err    error
			}{
				result: TeamID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
				err:    nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, err := TeamIDFrom(tt.input)
			assert.Equal(t, tt.expected.result, result)
			if tt.expected.err != nil {
				assert.Equal(t, tt.expected.err, err)
			}
		})
	}
}

func TestMustTeamID(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		shouldPanic bool
		expected    TeamID
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
			expected:    TeamID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
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
			result := MustTeamID(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTeamIDFromRef(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *TeamID
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
			expected: &TeamID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := TeamIDFromRef(&tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTeamIDFromRefID(t *testing.T) {
	id := New()
	id2 := TeamIDFromRefID(&id)
	assert.Equal(t, id.id, id2.id)
	assert.Nil(t, TeamIDFromRefID(nil))
	assert.Nil(t, TeamIDFromRefID(&ID{}))
}

func TestTeamID_ID(t *testing.T) {
	id := New()
	id2 := TeamIDFromRefID(&id)
	assert.Equal(t, id, id2.ID())
}

func TestTeamID_String(t *testing.T) {
	id := New()
	id2 := TeamIDFromRefID(&id)
	assert.Equal(t, id.String(), id2.String())
	assert.Equal(t, "", TeamID{}.String())
}

func TestTeamID_RefString(t *testing.T) {
	id := NewTeamID()
	assert.Equal(t, id.String(), *id.RefString())
	assert.Nil(t, TeamID{}.RefString())
}

func TestTeamID_GoString(t *testing.T) {
	id := New()
	id2 := TeamIDFromRefID(&id)
	assert.Equal(t, "TeamID("+id.String()+")", id2.GoString())
	assert.Equal(t, "TeamID()", TeamID{}.GoString())
}

func TestTeamID_Ref(t *testing.T) {
	id := NewTeamID()
	assert.Equal(t, TeamID(id), *id.Ref())
	assert.Nil(t, (&TeamID{}).Ref())
}

func TestTeamID_Contains(t *testing.T) {
	id := NewTeamID()
	id2 := NewTeamID()
	assert.True(t, id.Contains([]TeamID{id, id2}))
	assert.False(t, TeamID{}.Contains([]TeamID{id, id2, {}}))
	assert.False(t, id.Contains([]TeamID{id2}))
}

func TestTeamID_CopyRef(t *testing.T) {
	id := NewTeamID().Ref()
	id2 := id.CopyRef()
	assert.Equal(t, id, id2)
	assert.NotSame(t, id, id2)
	assert.Nil(t, (*TeamID)(nil).CopyRef())
}

func TestTeamID_IDRef(t *testing.T) {
	id := New()
	id2 := TeamIDFromRefID(&id)
	assert.Equal(t, &id, id2.IDRef())
	assert.Nil(t, (&TeamID{}).IDRef())
	assert.Nil(t, (*TeamID)(nil).IDRef())
}

func TestTeamID_StringRef(t *testing.T) {
	id := NewTeamID()
	assert.Equal(t, id.String(), *id.StringRef())
	assert.Nil(t, (&TeamID{}).StringRef())
	assert.Nil(t, (*TeamID)(nil).StringRef())
}

func TestTeamID_MarhsalJSON(t *testing.T) {
	id := NewTeamID()
	res, err := id.MarhsalJSON()
	assert.Nil(t, err)
	exp, _ := json.Marshal(id.String())
	assert.Equal(t, exp, res)

	res, err = (&TeamID{}).MarhsalJSON()
	assert.Nil(t, err)
	assert.Nil(t, res)

	res, err = (*TeamID)(nil).MarhsalJSON()
	assert.Nil(t, err)
	assert.Nil(t, res)
}

func TestTeamID_UnmarhsalJSON(t *testing.T) {
	jsonString := "\"01f3zhkysvcxsnzepyyqtq21fb\""
	id := MustTeamID("01f3zhkysvcxsnzepyyqtq21fb")
	id2 := &TeamID{}
	err := id2.UnmarhsalJSON([]byte(jsonString))
	assert.Nil(t, err)
	assert.Equal(t, id, *id2)
}

func TestTeamID_MarshalText(t *testing.T) {
	id := New()
	res, err := TeamIDFromRefID(&id).MarshalText()
	assert.Nil(t, err)
	assert.Equal(t, []byte(id.String()), res)

	res, err = (&TeamID{}).MarshalText()
	assert.Nil(t, err)
	assert.Nil(t, res)

	res, err = (*TeamID)(nil).MarshalText()
	assert.Nil(t, err)
	assert.Nil(t, res)
}

func TestTeamID_UnmarshalText(t *testing.T) {
	text := []byte("01f3zhcaq35403zdjnd6dcm0t2")
	id2 := &TeamID{}
	err := id2.UnmarshalText(text)
	assert.Nil(t, err)
	assert.Equal(t, "01f3zhcaq35403zdjnd6dcm0t2", id2.String())
}

func TestTeamID_IsNil(t *testing.T) {
	assert.True(t, TeamID{}.IsNil())
	assert.False(t, NewTeamID().IsNil())
}

func TestTeamID_IsNilRef(t *testing.T) {
	assert.True(t, TeamID{}.Ref().IsNilRef())
	assert.True(t, (*TeamID)(nil).IsNilRef())
	assert.False(t, NewTeamID().Ref().IsNilRef())
}

func TestTeamIDsToStrings(t *testing.T) {
	tests := []struct {
		name     string
		input    []TeamID
		expected []string
	}{
		{
			name:     "Empty slice",
			input:    make([]TeamID, 0),
			expected: make([]string, 0),
		},
		{
			name:     "1 element",
			input:    []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
		},
		{
			name: "multiple elements",
			input: []TeamID{
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t3"),
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
			assert.Equal(t, tt.expected, TeamIDsToStrings(tt.input))
		})
	}
}

func TestTeamIDsFrom(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected struct {
			res []TeamID
			err error
		}
	}{
		{
			name:  "Empty slice",
			input: make([]string, 0),
			expected: struct {
				res []TeamID
				err error
			}{
				res: make([]TeamID, 0),
				err: nil,
			},
		},
		{
			name:  "1 element",
			input: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
			expected: struct {
				res []TeamID
				err error
			}{
				res: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t2")},
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
				res []TeamID
				err error
			}{
				res: []TeamID{
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t3"),
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
				res []TeamID
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
			res, err := TeamIDsFrom(tc.input)
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

func TestTeamIDsFromID(t *testing.T) {
	tests := []struct {
		name     string
		input    []ID
		expected []TeamID
	}{
		{
			name:     "Empty slice",
			input:    make([]ID, 0),
			expected: make([]TeamID, 0),
		},
		{
			name:     "1 element",
			input:    []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []ID{
				MustBeID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []TeamID{
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := TeamIDsFromID(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestTeamIDsFromIDRef(t *testing.T) {
	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")

	tests := []struct {
		name     string
		input    []*ID
		expected []TeamID
	}{
		{
			name:     "Empty slice",
			input:    make([]*ID, 0),
			expected: make([]TeamID, 0),
		},
		{
			name:     "1 element",
			input:    []*ID{&id1},
			expected: []TeamID{MustTeamID(id1.String())},
		},
		{
			name:  "multiple elements",
			input: []*ID{&id1, &id2, &id3},
			expected: []TeamID{
				MustTeamID(id1.String()),
				MustTeamID(id2.String()),
				MustTeamID(id3.String()),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := TeamIDsFromIDRef(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestTeamIDsToID(t *testing.T) {
	tests := []struct {
		name     string
		input    []TeamID
		expected []ID
	}{
		{
			name:     "Empty slice",
			input:    make([]TeamID, 0),
			expected: make([]ID, 0),
		},
		{
			name:     "1 element",
			input:    []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []TeamID{
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t3"),
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
			res := TeamIDsToID(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestTeamIDsToIDRef(t *testing.T) {
	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	id21 := MustTeamID(id1.String())
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	id22 := MustTeamID(id2.String())
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")
	id23 := MustTeamID(id3.String())

	tests := []struct {
		name     string
		input    []*TeamID
		expected []*ID
	}{
		{
			name:     "Empty slice",
			input:    make([]*TeamID, 0),
			expected: make([]*ID, 0),
		},
		{
			name:     "1 element",
			input:    []*TeamID{&id21},
			expected: []*ID{&id1},
		},
		{
			name:     "multiple elements",
			input:    []*TeamID{&id21, &id22, &id23},
			expected: []*ID{&id1, &id2, &id3},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := TeamIDsToIDRef(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestNewTeamIDSet(t *testing.T) {
	TeamIdSet := NewTeamIDSet()
	assert.NotNil(t, TeamIdSet)
	assert.Empty(t, TeamIdSet.m)
	assert.Empty(t, TeamIdSet.s)
}

func TestTeamIDSet_Add(t *testing.T) {
	tests := []struct {
		name     string
		input    []TeamID
		expected *TeamIDSet
	}{
		{
			name:  "Empty slice",
			input: make([]TeamID, 0),
			expected: &TeamIDSet{
				m: map[TeamID]struct{}{},
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: &TeamIDSet{
				m: map[TeamID]struct{}{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: []TeamID{
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &TeamIDSet{
				m: map[TeamID]struct{}{
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []TeamID{
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
		{
			name: "multiple elements with duplication",
			input: []TeamID{
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &TeamIDSet{
				m: map[TeamID]struct{}{
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []TeamID{
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			set := NewTeamIDSet()
			set.Add(tc.input...)
			assert.Equal(t, tc.expected, set)
		})
	}
}

func TestTeamIDSet_AddRef(t *testing.T) {
	tests := []struct {
		name     string
		input    *TeamID
		expected *TeamIDSet
	}{
		{
			name:  "Empty slice",
			input: nil,
			expected: &TeamIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: MustTeamID("01f3zhcaq35403zdjnd6dcm0t1").Ref(),
			expected: &TeamIDSet{
				m: map[TeamID]struct{}{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			set := NewTeamIDSet()
			set.AddRef(tc.input)
			assert.Equal(t, tc.expected, set)
		})
	}
}

func TestTeamIDSet_Has(t *testing.T) {
	tests := []struct {
		name     string
		target   *TeamIDSet
		input    TeamID
		expected bool
	}{
		{
			name:     "Empty Set",
			target:   &TeamIDSet{},
			input:    MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"),
			expected: false,
		},
		{
			name: "Set Contains the element",
			target: &TeamIDSet{
				m: map[TeamID]struct{}{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			input:    MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"),
			expected: true,
		},
		{
			name: "Set does not Contains the element",
			target: &TeamIDSet{
				m: map[TeamID]struct{}{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			input:    MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"),
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

func TestTeamIDSet_Clear(t *testing.T) {
	tests := []struct {
		name     string
		input    *TeamIDSet
		expected *TeamIDSet
	}{
		{
			name:     "Empty set",
			input:    &TeamIDSet{},
			expected: &TeamIDSet{},
		},
		{
			name:     "Nil set",
			input:    nil,
			expected: nil,
		},
		{
			name: "Contains the element",
			input: &TeamIDSet{
				m: map[TeamID]struct{}{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: &TeamIDSet{
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

func TestTeamIDSet_All(t *testing.T) {
	tests := []struct {
		name     string
		input    *TeamIDSet
		expected []TeamID
	}{
		{
			name: "Empty",
			input: &TeamIDSet{
				m: map[TeamID]struct{}{},
				s: nil,
			},
			expected: make([]TeamID, 0),
		},
		{
			name:     "Nil",
			input:    nil,
			expected: nil,
		},
		{
			name: "1 element",
			input: &TeamIDSet{
				m: map[TeamID]struct{}{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1")},
		},
		{
			name: "multiple elements",
			input: &TeamIDSet{
				m: map[TeamID]struct{}{
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []TeamID{
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: []TeamID{
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustTeamID("01f3zhcaq35403zdjnd6dcm0t3"),
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

func TestTeamIDSet_Clone(t *testing.T) {
	tests := []struct {
		name     string
		input    *TeamIDSet
		expected *TeamIDSet
	}{
		{
			name:     "nil set",
			input:    nil,
			expected: NewTeamIDSet(),
		},
		{
			name:     "Empty set",
			input:    NewTeamIDSet(),
			expected: NewTeamIDSet(),
		},
		{
			name: "1 element",
			input: &TeamIDSet{
				m: map[TeamID]struct{}{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: &TeamIDSet{
				m: map[TeamID]struct{}{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: &TeamIDSet{
				m: map[TeamID]struct{}{
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []TeamID{
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: &TeamIDSet{
				m: map[TeamID]struct{}{
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []TeamID{
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t3"),
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

func TestTeamIDSet_Merge(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			a *TeamIDSet
			b *TeamIDSet
		}
		expected *TeamIDSet
	}{
		{
			name: "Nil Set",
			input: struct {
				a *TeamIDSet
				b *TeamIDSet
			}{
				a: &TeamIDSet{
					m: map[TeamID]struct{}{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: nil,
			},
			expected: &TeamIDSet{
				m: map[TeamID]struct{}{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "Empty Set",
			input: struct {
				a *TeamIDSet
				b *TeamIDSet
			}{
				a: &TeamIDSet{},
				b: &TeamIDSet{},
			},
			expected: &TeamIDSet{},
		},
		{
			name: "1 Empty Set",
			input: struct {
				a *TeamIDSet
				b *TeamIDSet
			}{
				a: &TeamIDSet{
					m: map[TeamID]struct{}{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &TeamIDSet{},
			},
			expected: &TeamIDSet{
				m: map[TeamID]struct{}{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "2 non Empty Set",
			input: struct {
				a *TeamIDSet
				b *TeamIDSet
			}{
				a: &TeamIDSet{
					m: map[TeamID]struct{}{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &TeamIDSet{
					m: map[TeamID]struct{}{MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"): {}},
					s: []TeamID{MustTeamID("01f3zhcaq35403zdjnd6dcm0t2")},
				},
			},
			expected: &TeamIDSet{
				m: map[TeamID]struct{}{
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"): {},
				},
				s: []TeamID{
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustTeamID("01f3zhcaq35403zdjnd6dcm0t2"),
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

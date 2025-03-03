// Code generated by gen, DO NOT EDIT.

package id

import "encoding/json"

// UserID is an ID for User.
type UserID ID

// NewUserID generates a new UserId.
func NewUserID() UserID {
	return UserID(New())
}

// UserIDFrom generates a new UserID from a string.
func UserIDFrom(i string) (nid UserID, err error) {
	var did ID
	did, err = FromID(i)
	if err != nil {
		return
	}
	nid = UserID(did)
	return
}

// MustUserID generates a new UserID from a string, but panics if the string cannot be parsed.
func MustUserID(i string) UserID {
	did, err := FromID(i)
	if err != nil {
		panic(err)
	}
	return UserID(did)
}

// UserIDFromRef generates a new UserID from a string ref.
func UserIDFromRef(i *string) *UserID {
	did := FromIDRef(i)
	if did == nil {
		return nil
	}
	nid := UserID(*did)
	return &nid
}

// UserIDFromRefID generates a new UserID from a ref of a generic ID.
func UserIDFromRefID(i *ID) *UserID {
	if i == nil || i.IsNil() {
		return nil
	}
	nid := UserID(*i)
	return &nid
}

// ID returns a domain ID.
func (d UserID) ID() ID {
	return ID(d)
}

// String returns a string representation.
func (d UserID) String() string {
	if d.IsNil() {
		return ""
	}
	return ID(d).String()
}

// StringRef returns a reference of the string representation.
func (d UserID) RefString() *string {
	if d.IsNil() {
		return nil
	}
	str := d.String()
	return &str
}

// GoString implements fmt.GoStringer interface.
func (d UserID) GoString() string {
	return "UserID(" + d.String() + ")"
}

// Ref returns a reference.
func (d UserID) Ref() *UserID {
	if d.IsNil() {
		return nil
	}
	d2 := d
	return &d2
}

// Contains returns whether the id is contained in the slice.
func (d UserID) Contains(ids []UserID) bool {
	if d.IsNil() {
		return false
	}
	for _, i := range ids {
		if d.ID().Equal(i.ID()) {
			return true
		}
	}
	return false
}

// CopyRef returns a copy of a reference.
func (d *UserID) CopyRef() *UserID {
	if d.IsNilRef() {
		return nil
	}
	d2 := *d
	return &d2
}

// IDRef returns a reference of a domain id.
func (d *UserID) IDRef() *ID {
	if d.IsNilRef() {
		return nil
	}
	id := ID(*d)
	return &id
}

// StringRef returns a reference of a string representation.
func (d *UserID) StringRef() *string {
	if d.IsNilRef() {
		return nil
	}
	id := ID(*d).String()
	return &id
}

// MarhsalJSON implements json.Marhsaler interface
func (d *UserID) MarhsalJSON() ([]byte, error) {
	if d.IsNilRef() {
		return nil, nil
	}
	return json.Marshal(d.String())
}

// UnmarhsalJSON implements json.Unmarshaler interface
func (d *UserID) UnmarhsalJSON(bs []byte) (err error) {
	var idstr string
	if err = json.Unmarshal(bs, &idstr); err != nil {
		return
	}
	*d, err = UserIDFrom(idstr)
	return
}

// MarshalText implements encoding.TextMarshaler interface
func (d *UserID) MarshalText() ([]byte, error) {
	if d.IsNilRef() {
		return nil, nil
	}
	return []byte(d.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface
func (d *UserID) UnmarshalText(text []byte) (err error) {
	*d, err = UserIDFrom(string(text))
	return
}

// IsNil returns true if a ID is zero-value
func (d UserID) IsNil() bool {
	return ID(d).IsNil()
}

// IsNilRef returns true if a ID is nil or zero-value
func (d *UserID) IsNilRef() bool {
	return d == nil || ID(*d).IsNil()
}

// UserIDsToStrings converts IDs into a string slice.
func UserIDsToStrings(ids []UserID) []string {
	strs := make([]string, 0, len(ids))
	for _, i := range ids {
		strs = append(strs, i.String())
	}
	return strs
}

// UserIDsFrom converts a string slice into a ID slice.
func UserIDsFrom(ids []string) ([]UserID, error) {
	dids := make([]UserID, 0, len(ids))
	for _, i := range ids {
		did, err := UserIDFrom(i)
		if err != nil {
			return nil, err
		}
		dids = append(dids, did)
	}
	return dids, nil
}

// UserIDsFromID converts a generic ID slice into a ID slice.
func UserIDsFromID(ids []ID) []UserID {
	dids := make([]UserID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, UserID(i))
	}
	return dids
}

// UserIDsFromIDRef converts a ref of a generic ID slice into a ID slice.
func UserIDsFromIDRef(ids []*ID) []UserID {
	dids := make([]UserID, 0, len(ids))
	for _, i := range ids {
		if i != nil {
			dids = append(dids, UserID(*i))
		}
	}
	return dids
}

// UserIDsToID converts a ID slice into a generic ID slice.
func UserIDsToID(ids []UserID) []ID {
	dids := make([]ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.ID())
	}
	return dids
}

// UserIDsToIDRef converts a ID ref slice into a generic ID ref slice.
func UserIDsToIDRef(ids []*UserID) []*ID {
	dids := make([]*ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.IDRef())
	}
	return dids
}

// UserIDSet represents a set of UserIDs
type UserIDSet struct {
	m map[UserID]struct{}
	s []UserID
}

// NewUserIDSet creates a new UserIDSet
func NewUserIDSet() *UserIDSet {
	return &UserIDSet{}
}

// Add adds a new ID if it does not exists in the set
func (s *UserIDSet) Add(p ...UserID) {
	if s == nil || p == nil {
		return
	}
	if s.m == nil {
		s.m = map[UserID]struct{}{}
	}
	for _, i := range p {
		if _, ok := s.m[i]; !ok {
			if s.s == nil {
				s.s = []UserID{}
			}
			s.m[i] = struct{}{}
			s.s = append(s.s, i)
		}
	}
}

// AddRef adds a new ID ref if it does not exists in the set
func (s *UserIDSet) AddRef(p *UserID) {
	if s == nil || p == nil {
		return
	}
	s.Add(*p)
}

// Has checks if the ID exists in the set
func (s *UserIDSet) Has(p UserID) bool {
	if s == nil || s.m == nil {
		return false
	}
	_, ok := s.m[p]
	return ok
}

// Clear clears all stored IDs
func (s *UserIDSet) Clear() {
	if s == nil {
		return
	}
	s.m = nil
	s.s = nil
}

// All returns stored all IDs as a slice
func (s *UserIDSet) All() []UserID {
	if s == nil {
		return nil
	}
	return append([]UserID{}, s.s...)
}

// Clone returns a cloned set
func (s *UserIDSet) Clone() *UserIDSet {
	if s == nil {
		return NewUserIDSet()
	}
	s2 := NewUserIDSet()
	s2.Add(s.s...)
	return s2
}

// Merge returns a merged set
func (s *UserIDSet) Merge(s2 *UserIDSet) *UserIDSet {
	s3 := s.Clone()
	if s2 == nil {
		return s3
	}
	s3.Add(s2.s...)
	return s3
}

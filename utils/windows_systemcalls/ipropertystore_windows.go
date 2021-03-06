package windows_systemcalls

import (
	"github.com/go-ole/go-ole"
	"unsafe"
)

var IPropertyStoreGuid = ole.NewGUID("886d8eeb-8cf2-4446-8d02-cdba1dbdcf99")

type IPropertyStore struct {
	ole.IUnknown
}

type IPropertyStoreVtbl struct {
	ole.IUnknownVtbl
	GetCount uintptr
	GetAt    uintptr
	GetValue uintptr
	SetValue uintptr
	Commit   uintptr
}

func (v *IPropertyStore) VTable() *IPropertyStoreVtbl {
	return (*IPropertyStoreVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IPropertyStore) GetCount(count *uint32) (err error) {
	err = psGetCount(v, count)
	return
}

func (v *IPropertyStore) GetAt(iProp uint32, pkey *PROPERTYKEY) error {
	return psGetAt(v, iProp, pkey)
}

func (v *IPropertyStore) GetValue(key *PROPERTYKEY, pv *PROPVARIANT) (err error) {
	err = psGetValue(v, key, pv)
	return
}

func (v *IPropertyStore) SetValue() (err error) {
	err = psSetValue()
	return
}

func (v *IPropertyStore) Commit() (err error) {
	err = psCommit()
	return
}

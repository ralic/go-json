package json_test

import (
	"bytes"
	"testing"

	"github.com/goccy/go-json"
)

func TestCoverBool(t *testing.T) {
	type structBool struct {
		A bool `json:"a"`
	}
	type structBoolOmitEmpty struct {
		A bool `json:"a,omitempty"`
	}
	type structBoolString struct {
		A bool `json:"a,string"`
	}

	type structBoolPtr struct {
		A *bool `json:"a"`
	}
	type structBoolPtrOmitEmpty struct {
		A *bool `json:"a,omitempty"`
	}
	type structBoolPtrString struct {
		A *bool `json:"a,string"`
	}

	tests := []struct {
		name string
		data interface{}
	}{
		{
			name: "Bool",
			data: bool(true),
		},
		{
			name: "BoolPtr",
			data: boolptr(true),
		},
		{
			name: "BoolPtr3",
			data: boolptr3(true),
		},
		{
			name: "BoolPtrNil",
			data: (*bool)(nil),
		},
		{
			name: "BoolPtr3Nil",
			data: (***bool)(nil),
		},

		// HeadBoolZero
		{
			name: "HeadBoolZero",
			data: struct {
				A bool `json:"a"`
			}{},
		},
		{
			name: "HeadBoolZeroOmitEmpty",
			data: struct {
				A bool `json:"a,omitempty"`
			}{},
		},
		{
			name: "HeadBoolZeroString",
			data: struct {
				A bool `json:"a,string"`
			}{},
		},

		// HeadBool
		{
			name: "HeadBool",
			data: struct {
				A bool `json:"a"`
			}{A: true},
		},
		{
			name: "HeadBoolOmitEmpty",
			data: struct {
				A bool `json:"a,omitempty"`
			}{A: true},
		},
		{
			name: "HeadBoolString",
			data: struct {
				A bool `json:"a,string"`
			}{A: true},
		},

		// HeadBoolPtr
		{
			name: "HeadBoolPtr",
			data: struct {
				A *bool `json:"a"`
			}{A: boolptr(true)},
		},
		{
			name: "HeadBoolPtrOmitEmpty",
			data: struct {
				A *bool `json:"a,omitempty"`
			}{A: boolptr(true)},
		},
		{
			name: "HeadBoolPtrString",
			data: struct {
				A *bool `json:"a,string"`
			}{A: boolptr(true)},
		},

		// HeadBoolPtrNil
		{
			name: "HeadBoolPtrNil",
			data: struct {
				A *bool `json:"a"`
			}{A: nil},
		},
		{
			name: "HeadBoolPtrNilOmitEmpty",
			data: struct {
				A *bool `json:"a,omitempty"`
			}{A: nil},
		},
		{
			name: "HeadBoolPtrNilString",
			data: struct {
				A *bool `json:"a,string"`
			}{A: nil},
		},

		// PtrHeadBoolZero
		{
			name: "PtrHeadBoolZero",
			data: &struct {
				A bool `json:"a"`
			}{},
		},
		{
			name: "PtrHeadBoolZeroOmitEmpty",
			data: &struct {
				A bool `json:"a,omitempty"`
			}{},
		},
		{
			name: "PtrHeadBoolZeroString",
			data: &struct {
				A bool `json:"a,string"`
			}{},
		},

		// PtrHeadBool
		{
			name: "PtrHeadBool",
			data: &struct {
				A bool `json:"a"`
			}{A: true},
		},
		{
			name: "PtrHeadBoolOmitEmpty",
			data: &struct {
				A bool `json:"a,omitempty"`
			}{A: true},
		},
		{
			name: "PtrHeadBoolString",
			data: &struct {
				A bool `json:"a,string"`
			}{A: true},
		},

		// PtrHeadBoolPtr
		{
			name: "PtrHeadBoolPtr",
			data: &struct {
				A *bool `json:"a"`
			}{A: boolptr(true)},
		},
		{
			name: "PtrHeadBoolPtrOmitEmpty",
			data: &struct {
				A *bool `json:"a,omitempty"`
			}{A: boolptr(true)},
		},
		{
			name: "PtrHeadBoolPtrString",
			data: &struct {
				A *bool `json:"a,string"`
			}{A: boolptr(true)},
		},

		// PtrHeadBoolPtrNil
		{
			name: "PtrHeadBoolPtrNil",
			data: &struct {
				A *bool `json:"a"`
			}{A: nil},
		},
		{
			name: "PtrHeadBoolPtrNilOmitEmpty",
			data: &struct {
				A *bool `json:"a,omitempty"`
			}{A: nil},
		},
		{
			name: "PtrHeadBoolPtrNilString",
			data: &struct {
				A *bool `json:"a,string"`
			}{A: nil},
		},

		// PtrHeadBoolNil
		{
			name: "PtrHeadBoolNil",
			data: (*struct {
				A *bool `json:"a"`
			})(nil),
		},
		{
			name: "PtrHeadBoolNilOmitEmpty",
			data: (*struct {
				A *bool `json:"a,omitempty"`
			})(nil),
		},
		{
			name: "PtrHeadBoolNilString",
			data: (*struct {
				A *bool `json:"a,string"`
			})(nil),
		},

		// HeadBoolZeroMultiFields
		{
			name: "HeadBoolZeroMultiFields",
			data: struct {
				A bool `json:"a"`
				B bool `json:"b"`
				C bool `json:"c"`
			}{},
		},
		{
			name: "HeadBoolZeroMultiFieldsOmitEmpty",
			data: struct {
				A bool `json:"a,omitempty"`
				B bool `json:"b,omitempty"`
				C bool `json:"c,omitempty"`
			}{},
		},
		{
			name: "HeadBoolZeroMultiFields",
			data: struct {
				A bool `json:"a,string"`
				B bool `json:"b,string"`
				C bool `json:"c,string"`
			}{},
		},

		// HeadBoolMultiFields
		{
			name: "HeadBoolMultiFields",
			data: struct {
				A bool `json:"a"`
				B bool `json:"b"`
				C bool `json:"c"`
			}{A: true, B: false, C: true},
		},
		{
			name: "HeadBoolMultiFieldsOmitEmpty",
			data: struct {
				A bool `json:"a,omitempty"`
				B bool `json:"b,omitempty"`
				C bool `json:"c,omitempty"`
			}{A: true, B: false, C: true},
		},
		{
			name: "HeadBoolMultiFieldsString",
			data: struct {
				A bool `json:"a,string"`
				B bool `json:"b,string"`
				C bool `json:"c,string"`
			}{A: true, B: false, C: true},
		},

		// HeadBoolPtrMultiFields
		{
			name: "HeadBoolPtrMultiFields",
			data: struct {
				A *bool `json:"a"`
				B *bool `json:"b"`
				C *bool `json:"c"`
			}{A: boolptr(true), B: boolptr(false), C: boolptr(true)},
		},
		{
			name: "HeadBoolPtrMultiFieldsOmitEmpty",
			data: struct {
				A *bool `json:"a,omitempty"`
				B *bool `json:"b,omitempty"`
				C *bool `json:"c,omitempty"`
			}{A: boolptr(true), B: boolptr(false), C: boolptr(true)},
		},
		{
			name: "HeadBoolPtrMultiFieldsString",
			data: struct {
				A *bool `json:"a,string"`
				B *bool `json:"b,string"`
				C *bool `json:"c,string"`
			}{A: boolptr(true), B: boolptr(false), C: boolptr(true)},
		},

		// HeadBoolPtrNilMultiFields
		{
			name: "HeadBoolPtrNilMultiFields",
			data: struct {
				A *bool `json:"a"`
				B *bool `json:"b"`
				C *bool `json:"c"`
			}{A: nil, B: nil, C: nil},
		},
		{
			name: "HeadBoolPtrNilMultiFieldsOmitEmpty",
			data: struct {
				A *bool `json:"a,omitempty"`
				B *bool `json:"b,omitempty"`
				C *bool `json:"c,omitempty"`
			}{A: nil, B: nil, C: nil},
		},
		{
			name: "HeadBoolPtrNilMultiFieldsString",
			data: struct {
				A *bool `json:"a,string"`
				B *bool `json:"b,string"`
				C *bool `json:"c,string"`
			}{A: nil, B: nil, C: nil},
		},

		// PtrHeadBoolZeroMultiFields
		{
			name: "PtrHeadBoolZeroMultiFields",
			data: &struct {
				A bool `json:"a"`
				B bool `json:"b"`
			}{},
		},
		{
			name: "PtrHeadBoolZeroMultiFieldsOmitEmpty",
			data: &struct {
				A bool `json:"a,omitempty"`
				B bool `json:"b,omitempty"`
			}{},
		},
		{
			name: "PtrHeadBoolZeroMultiFieldsString",
			data: &struct {
				A bool `json:"a,string"`
				B bool `json:"b,string"`
			}{},
		},

		// PtrHeadBoolMultiFields
		{
			name: "PtrHeadBoolMultiFields",
			data: &struct {
				A bool `json:"a"`
				B bool `json:"b"`
			}{A: true, B: false},
		},
		{
			name: "PtrHeadBoolMultiFieldsOmitEmpty",
			data: &struct {
				A bool `json:"a,omitempty"`
				B bool `json:"b,omitempty"`
			}{A: true, B: false},
		},
		{
			name: "PtrHeadBoolMultiFieldsString",
			data: &struct {
				A bool `json:"a,string"`
				B bool `json:"b,string"`
			}{A: true, B: false},
		},

		// PtrHeadBoolPtrMultiFields
		{
			name: "PtrHeadBoolPtrMultiFields",
			data: &struct {
				A *bool `json:"a"`
				B *bool `json:"b"`
			}{A: boolptr(true), B: boolptr(false)},
		},
		{
			name: "PtrHeadBoolPtrMultiFieldsOmitEmpty",
			data: &struct {
				A *bool `json:"a,omitempty"`
				B *bool `json:"b,omitempty"`
			}{A: boolptr(true), B: boolptr(false)},
		},
		{
			name: "PtrHeadBoolPtrMultiFieldsString",
			data: &struct {
				A *bool `json:"a,string"`
				B *bool `json:"b,string"`
			}{A: boolptr(true), B: boolptr(false)},
		},

		// PtrHeadBoolPtrNilMultiFields
		{
			name: "PtrHeadBoolPtrNilMultiFields",
			data: &struct {
				A *bool `json:"a"`
				B *bool `json:"b"`
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadBoolPtrNilMultiFieldsOmitEmpty",
			data: &struct {
				A *bool `json:"a,omitempty"`
				B *bool `json:"b,omitempty"`
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadBoolPtrNilMultiFieldsString",
			data: &struct {
				A *bool `json:"a,string"`
				B *bool `json:"b,string"`
			}{A: nil, B: nil},
		},

		// PtrHeadBoolNilMultiFields
		{
			name: "PtrHeadBoolNilMultiFields",
			data: (*struct {
				A *bool `json:"a"`
				B *bool `json:"b"`
			})(nil),
		},
		{
			name: "PtrHeadBoolNilMultiFieldsOmitEmpty",
			data: (*struct {
				A *bool `json:"a,omitempty"`
				B *bool `json:"b,omitempty"`
			})(nil),
		},
		{
			name: "PtrHeadBoolNilMultiFieldsString",
			data: (*struct {
				A *bool `json:"a,string"`
				B *bool `json:"b,string"`
			})(nil),
		},

		// HeadBoolZeroNotRoot
		{
			name: "HeadBoolZeroNotRoot",
			data: struct {
				A struct {
					A bool `json:"a"`
				}
			}{},
		},
		{
			name: "HeadBoolZeroNotRootOmitEmpty",
			data: struct {
				A struct {
					A bool `json:"a,omitempty"`
				}
			}{},
		},
		{
			name: "HeadBoolZeroNotRootString",
			data: struct {
				A struct {
					A bool `json:"a,string"`
				}
			}{},
		},

		// HeadBoolNotRoot
		{
			name: "HeadBoolNotRoot",
			data: struct {
				A struct {
					A bool `json:"a"`
				}
			}{A: struct {
				A bool `json:"a"`
			}{A: true}},
		},
		{
			name: "HeadBoolNotRootOmitEmpty",
			data: struct {
				A struct {
					A bool `json:"a,omitempty"`
				}
			}{A: struct {
				A bool `json:"a,omitempty"`
			}{A: true}},
		},
		{
			name: "HeadBoolNotRootString",
			data: struct {
				A struct {
					A bool `json:"a,string"`
				}
			}{A: struct {
				A bool `json:"a,string"`
			}{A: true}},
		},

		// HeadBoolPtrNotRoot
		{
			name: "HeadBoolPtrNotRoot",
			data: struct {
				A struct {
					A *bool `json:"a"`
				}
			}{A: struct {
				A *bool `json:"a"`
			}{boolptr(true)}},
		},
		{
			name: "HeadBoolPtrNotRootOmitEmpty",
			data: struct {
				A struct {
					A *bool `json:"a,omitempty"`
				}
			}{A: struct {
				A *bool `json:"a,omitempty"`
			}{boolptr(true)}},
		},
		{
			name: "HeadBoolPtrNotRootString",
			data: struct {
				A struct {
					A *bool `json:"a,string"`
				}
			}{A: struct {
				A *bool `json:"a,string"`
			}{boolptr(true)}},
		},

		// HeadBoolPtrNilNotRoot
		{
			name: "HeadBoolPtrNilNotRoot",
			data: struct {
				A struct {
					A *bool `json:"a"`
				}
			}{},
		},
		{
			name: "HeadBoolPtrNilNotRootOmitEmpty",
			data: struct {
				A struct {
					A *bool `json:"a,omitempty"`
				}
			}{},
		},
		{
			name: "HeadBoolPtrNilNotRootString",
			data: struct {
				A struct {
					A *bool `json:"a,string"`
				}
			}{},
		},

		// PtrHeadBoolZeroNotRoot
		{
			name: "PtrHeadBoolZeroNotRoot",
			data: struct {
				A *struct {
					A bool `json:"a"`
				}
			}{A: new(struct {
				A bool `json:"a"`
			})},
		},
		{
			name: "PtrHeadBoolZeroNotRootOmitEmpty",
			data: struct {
				A *struct {
					A bool `json:"a,omitempty"`
				}
			}{A: new(struct {
				A bool `json:"a,omitempty"`
			})},
		},
		{
			name: "PtrHeadBoolZeroNotRootString",
			data: struct {
				A *struct {
					A bool `json:"a,string"`
				}
			}{A: new(struct {
				A bool `json:"a,string"`
			})},
		},

		// PtrHeadBoolNotRoot
		{
			name: "PtrHeadBoolNotRoot",
			data: struct {
				A *struct {
					A bool `json:"a"`
				}
			}{A: &(struct {
				A bool `json:"a"`
			}{A: true})},
		},
		{
			name: "PtrHeadBoolNotRootOmitEmpty",
			data: struct {
				A *struct {
					A bool `json:"a,omitempty"`
				}
			}{A: &(struct {
				A bool `json:"a,omitempty"`
			}{A: true})},
		},
		{
			name: "PtrHeadBoolNotRootString",
			data: struct {
				A *struct {
					A bool `json:"a,string"`
				}
			}{A: &(struct {
				A bool `json:"a,string"`
			}{A: true})},
		},

		// PtrHeadBoolPtrNotRoot
		{
			name: "PtrHeadBoolPtrNotRoot",
			data: struct {
				A *struct {
					A *bool `json:"a"`
				}
			}{A: &(struct {
				A *bool `json:"a"`
			}{A: boolptr(true)})},
		},
		{
			name: "PtrHeadBoolPtrNotRootOmitEmpty",
			data: struct {
				A *struct {
					A *bool `json:"a,omitempty"`
				}
			}{A: &(struct {
				A *bool `json:"a,omitempty"`
			}{A: boolptr(true)})},
		},
		{
			name: "PtrHeadBoolPtrNotRootString",
			data: struct {
				A *struct {
					A *bool `json:"a,string"`
				}
			}{A: &(struct {
				A *bool `json:"a,string"`
			}{A: boolptr(true)})},
		},

		// PtrHeadBoolPtrNilNotRoot
		{
			name: "PtrHeadBoolPtrNilNotRoot",
			data: struct {
				A *struct {
					A *bool `json:"a"`
				}
			}{A: &(struct {
				A *bool `json:"a"`
			}{A: nil})},
		},
		{
			name: "PtrHeadBoolPtrNilNotRootOmitEmpty",
			data: struct {
				A *struct {
					A *bool `json:"a,omitempty"`
				}
			}{A: &(struct {
				A *bool `json:"a,omitempty"`
			}{A: nil})},
		},
		{
			name: "PtrHeadBoolPtrNilNotRootString",
			data: struct {
				A *struct {
					A *bool `json:"a,string"`
				}
			}{A: &(struct {
				A *bool `json:"a,string"`
			}{A: nil})},
		},

		// PtrHeadBoolNilNotRoot
		{
			name: "PtrHeadBoolNilNotRoot",
			data: struct {
				A *struct {
					A *bool `json:"a"`
				}
			}{A: nil},
		},
		{
			name: "PtrHeadBoolNilNotRootOmitEmpty",
			data: struct {
				A *struct {
					A *bool `json:"a,omitempty"`
				} `json:",omitempty"`
			}{A: nil},
		},
		{
			name: "PtrHeadBoolNilNotRootString",
			data: struct {
				A *struct {
					A *bool `json:"a,string"`
				} `json:",string"`
			}{A: nil},
		},

		// HeadBoolZeroMultiFieldsNotRoot
		{
			name: "HeadBoolZeroMultiFieldsNotRoot",
			data: struct {
				A struct {
					A bool `json:"a"`
				}
				B struct {
					B bool `json:"b"`
				}
			}{},
		},
		{
			name: "HeadBoolZeroMultiFieldsNotRootOmitEmpty",
			data: struct {
				A struct {
					A bool `json:"a,omitempty"`
				}
				B struct {
					B bool `json:"b,omitempty"`
				}
			}{},
		},
		{
			name: "HeadBoolZeroMultiFieldsNotRootString",
			data: struct {
				A struct {
					A bool `json:"a,string"`
				}
				B struct {
					B bool `json:"b,string"`
				}
			}{},
		},

		// HeadBoolMultiFieldsNotRoot
		{
			name: "HeadBoolMultiFieldsNotRoot",
			data: struct {
				A struct {
					A bool `json:"a"`
				}
				B struct {
					B bool `json:"b"`
				}
			}{A: struct {
				A bool `json:"a"`
			}{A: true}, B: struct {
				B bool `json:"b"`
			}{B: false}},
		},
		{
			name: "HeadBoolMultiFieldsNotRootOmitEmpty",
			data: struct {
				A struct {
					A bool `json:"a,omitempty"`
				}
				B struct {
					B bool `json:"b,omitempty"`
				}
			}{A: struct {
				A bool `json:"a,omitempty"`
			}{A: true}, B: struct {
				B bool `json:"b,omitempty"`
			}{B: false}},
		},
		{
			name: "HeadBoolMultiFieldsNotRootString",
			data: struct {
				A struct {
					A bool `json:"a,string"`
				}
				B struct {
					B bool `json:"b,string"`
				}
			}{A: struct {
				A bool `json:"a,string"`
			}{A: true}, B: struct {
				B bool `json:"b,string"`
			}{B: false}},
		},

		// HeadBoolPtrMultiFieldsNotRoot
		{
			name: "HeadBoolPtrMultiFieldsNotRoot",
			data: struct {
				A struct {
					A *bool `json:"a"`
				}
				B struct {
					B *bool `json:"b"`
				}
			}{A: struct {
				A *bool `json:"a"`
			}{A: boolptr(true)}, B: struct {
				B *bool `json:"b"`
			}{B: boolptr(false)}},
		},
		{
			name: "HeadBoolPtrMultiFieldsNotRootOmitEmpty",
			data: struct {
				A struct {
					A *bool `json:"a,omitempty"`
				}
				B struct {
					B *bool `json:"b,omitempty"`
				}
			}{A: struct {
				A *bool `json:"a,omitempty"`
			}{A: boolptr(true)}, B: struct {
				B *bool `json:"b,omitempty"`
			}{B: boolptr(false)}},
		},
		{
			name: "HeadBoolPtrMultiFieldsNotRootString",
			data: struct {
				A struct {
					A *bool `json:"a,string"`
				}
				B struct {
					B *bool `json:"b,string"`
				}
			}{A: struct {
				A *bool `json:"a,string"`
			}{A: boolptr(true)}, B: struct {
				B *bool `json:"b,string"`
			}{B: boolptr(false)}},
		},

		// HeadBoolPtrNilMultiFieldsNotRoot
		{
			name: "HeadBoolPtrNilMultiFieldsNotRoot",
			data: struct {
				A struct {
					A *bool `json:"a"`
				}
				B struct {
					B *bool `json:"b"`
				}
			}{A: struct {
				A *bool `json:"a"`
			}{A: nil}, B: struct {
				B *bool `json:"b"`
			}{B: nil}},
		},
		{
			name: "HeadBoolPtrNilMultiFieldsNotRootOmitEmpty",
			data: struct {
				A struct {
					A *bool `json:"a,omitempty"`
				}
				B struct {
					B *bool `json:"b,omitempty"`
				}
			}{A: struct {
				A *bool `json:"a,omitempty"`
			}{A: nil}, B: struct {
				B *bool `json:"b,omitempty"`
			}{B: nil}},
		},
		{
			name: "HeadBoolPtrNilMultiFieldsNotRootString",
			data: struct {
				A struct {
					A *bool `json:"a,string"`
				}
				B struct {
					B *bool `json:"b,string"`
				}
			}{A: struct {
				A *bool `json:"a,string"`
			}{A: nil}, B: struct {
				B *bool `json:"b,string"`
			}{B: nil}},
		},

		// PtrHeadBoolZeroMultiFieldsNotRoot
		{
			name: "PtrHeadBoolZeroMultiFieldsNotRoot",
			data: &struct {
				A struct {
					A bool `json:"a"`
				}
				B struct {
					B bool `json:"b"`
				}
			}{},
		},
		{
			name: "PtrHeadBoolZeroMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A struct {
					A bool `json:"a,omitempty"`
				}
				B struct {
					B bool `json:"b,omitempty"`
				}
			}{},
		},
		{
			name: "PtrHeadBoolZeroMultiFieldsNotRootString",
			data: &struct {
				A struct {
					A bool `json:"a,string"`
				}
				B struct {
					B bool `json:"b,string"`
				}
			}{},
		},

		// PtrHeadBoolMultiFieldsNotRoot
		{
			name: "PtrHeadBoolMultiFieldsNotRoot",
			data: &struct {
				A struct {
					A bool `json:"a"`
				}
				B struct {
					B bool `json:"b"`
				}
			}{A: struct {
				A bool `json:"a"`
			}{A: true}, B: struct {
				B bool `json:"b"`
			}{B: false}},
		},
		{
			name: "PtrHeadBoolMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A struct {
					A bool `json:"a,omitempty"`
				}
				B struct {
					B bool `json:"b,omitempty"`
				}
			}{A: struct {
				A bool `json:"a,omitempty"`
			}{A: true}, B: struct {
				B bool `json:"b,omitempty"`
			}{B: false}},
		},
		{
			name: "PtrHeadBoolMultiFieldsNotRootString",
			data: &struct {
				A struct {
					A bool `json:"a,string"`
				}
				B struct {
					B bool `json:"b,string"`
				}
			}{A: struct {
				A bool `json:"a,string"`
			}{A: true}, B: struct {
				B bool `json:"b,string"`
			}{B: false}},
		},

		// PtrHeadBoolPtrMultiFieldsNotRoot
		{
			name: "PtrHeadBoolPtrMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A *bool `json:"a"`
				}
				B *struct {
					B *bool `json:"b"`
				}
			}{A: &(struct {
				A *bool `json:"a"`
			}{A: boolptr(true)}), B: &(struct {
				B *bool `json:"b"`
			}{B: boolptr(false)})},
		},
		{
			name: "PtrHeadBoolPtrMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A *bool `json:"a,omitempty"`
				}
				B *struct {
					B *bool `json:"b,omitempty"`
				}
			}{A: &(struct {
				A *bool `json:"a,omitempty"`
			}{A: boolptr(true)}), B: &(struct {
				B *bool `json:"b,omitempty"`
			}{B: boolptr(false)})},
		},
		{
			name: "PtrHeadBoolPtrMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A *bool `json:"a,string"`
				}
				B *struct {
					B *bool `json:"b,string"`
				}
			}{A: &(struct {
				A *bool `json:"a,string"`
			}{A: boolptr(true)}), B: &(struct {
				B *bool `json:"b,string"`
			}{B: boolptr(false)})},
		},

		// PtrHeadBoolPtrNilMultiFieldsNotRoot
		{
			name: "PtrHeadBoolPtrNilMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A *bool `json:"a"`
				}
				B *struct {
					B *bool `json:"b"`
				}
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadBoolPtrNilMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A *bool `json:"a,omitempty"`
				} `json:",omitempty"`
				B *struct {
					B *bool `json:"b,omitempty"`
				} `json:",omitempty"`
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadBoolPtrNilMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A *bool `json:"a,string"`
				} `json:",string"`
				B *struct {
					B *bool `json:"b,string"`
				} `json:",string"`
			}{A: nil, B: nil},
		},

		// PtrHeadBoolNilMultiFieldsNotRoot
		{
			name: "PtrHeadBoolNilMultiFieldsNotRoot",
			data: (*struct {
				A *struct {
					A *bool `json:"a"`
				}
				B *struct {
					B *bool `json:"b"`
				}
			})(nil),
		},
		{
			name: "PtrHeadBoolNilMultiFieldsNotRootOmitEmpty",
			data: (*struct {
				A *struct {
					A *bool `json:"a,omitempty"`
				}
				B *struct {
					B *bool `json:"b,omitempty"`
				}
			})(nil),
		},
		{
			name: "PtrHeadBoolNilMultiFieldsNotRootString",
			data: (*struct {
				A *struct {
					A *bool `json:"a,string"`
				}
				B *struct {
					B *bool `json:"b,string"`
				}
			})(nil),
		},

		// PtrHeadBoolDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadBoolDoubleMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A bool `json:"a"`
					B bool `json:"b"`
				}
				B *struct {
					A bool `json:"a"`
					B bool `json:"b"`
				}
			}{A: &(struct {
				A bool `json:"a"`
				B bool `json:"b"`
			}{A: true, B: false}), B: &(struct {
				A bool `json:"a"`
				B bool `json:"b"`
			}{A: true, B: false})},
		},
		{
			name: "PtrHeadBoolDoubleMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A bool `json:"a,omitempty"`
					B bool `json:"b,omitempty"`
				}
				B *struct {
					A bool `json:"a,omitempty"`
					B bool `json:"b,omitempty"`
				}
			}{A: &(struct {
				A bool `json:"a,omitempty"`
				B bool `json:"b,omitempty"`
			}{A: true, B: false}), B: &(struct {
				A bool `json:"a,omitempty"`
				B bool `json:"b,omitempty"`
			}{A: true, B: false})},
		},
		{
			name: "PtrHeadBoolDoubleMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A bool `json:"a,string"`
					B bool `json:"b,string"`
				}
				B *struct {
					A bool `json:"a,string"`
					B bool `json:"b,string"`
				}
			}{A: &(struct {
				A bool `json:"a,string"`
				B bool `json:"b,string"`
			}{A: true, B: false}), B: &(struct {
				A bool `json:"a,string"`
				B bool `json:"b,string"`
			}{A: true, B: false})},
		},

		// PtrHeadBoolNilDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadBoolNilDoubleMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A bool `json:"a"`
					B bool `json:"b"`
				}
				B *struct {
					A bool `json:"a"`
					B bool `json:"b"`
				}
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadBoolNilDoubleMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A bool `json:"a,omitempty"`
					B bool `json:"b,omitempty"`
				} `json:",omitempty"`
				B *struct {
					A bool `json:"a,omitempty"`
					B bool `json:"b,omitempty"`
				} `json:",omitempty"`
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadBoolNilDoubleMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A bool `json:"a,string"`
					B bool `json:"b,string"`
				}
				B *struct {
					A bool `json:"a,string"`
					B bool `json:"b,string"`
				}
			}{A: nil, B: nil},
		},

		// PtrHeadBoolNilDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadBoolNilDoubleMultiFieldsNotRoot",
			data: (*struct {
				A *struct {
					A bool `json:"a"`
					B bool `json:"b"`
				}
				B *struct {
					A bool `json:"a"`
					B bool `json:"b"`
				}
			})(nil),
		},
		{
			name: "PtrHeadBoolNilDoubleMultiFieldsNotRootOmitEmpty",
			data: (*struct {
				A *struct {
					A bool `json:"a,omitempty"`
					B bool `json:"b,omitempty"`
				}
				B *struct {
					A bool `json:"a,omitempty"`
					B bool `json:"b,omitempty"`
				}
			})(nil),
		},
		{
			name: "PtrHeadBoolNilDoubleMultiFieldsNotRootString",
			data: (*struct {
				A *struct {
					A bool `json:"a,string"`
					B bool `json:"b,string"`
				}
				B *struct {
					A bool `json:"a,string"`
					B bool `json:"b,string"`
				}
			})(nil),
		},

		// PtrHeadBoolPtrDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadBoolPtrDoubleMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A *bool `json:"a"`
					B *bool `json:"b"`
				}
				B *struct {
					A *bool `json:"a"`
					B *bool `json:"b"`
				}
			}{A: &(struct {
				A *bool `json:"a"`
				B *bool `json:"b"`
			}{A: boolptr(true), B: boolptr(false)}), B: &(struct {
				A *bool `json:"a"`
				B *bool `json:"b"`
			}{A: boolptr(true), B: boolptr(false)})},
		},
		{
			name: "PtrHeadBoolPtrDoubleMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A *bool `json:"a,omitempty"`
					B *bool `json:"b,omitempty"`
				}
				B *struct {
					A *bool `json:"a,omitempty"`
					B *bool `json:"b,omitempty"`
				}
			}{A: &(struct {
				A *bool `json:"a,omitempty"`
				B *bool `json:"b,omitempty"`
			}{A: boolptr(true), B: boolptr(false)}), B: &(struct {
				A *bool `json:"a,omitempty"`
				B *bool `json:"b,omitempty"`
			}{A: boolptr(true), B: boolptr(false)})},
		},
		{
			name: "PtrHeadBoolPtrDoubleMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A *bool `json:"a,string"`
					B *bool `json:"b,string"`
				}
				B *struct {
					A *bool `json:"a,string"`
					B *bool `json:"b,string"`
				}
			}{A: &(struct {
				A *bool `json:"a,string"`
				B *bool `json:"b,string"`
			}{A: boolptr(true), B: boolptr(false)}), B: &(struct {
				A *bool `json:"a,string"`
				B *bool `json:"b,string"`
			}{A: boolptr(true), B: boolptr(false)})},
		},

		// PtrHeadBoolPtrNilDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadBoolPtrNilDoubleMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A *bool `json:"a"`
					B *bool `json:"b"`
				}
				B *struct {
					A *bool `json:"a"`
					B *bool `json:"b"`
				}
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadBoolPtrNilDoubleMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A *bool `json:"a,omitempty"`
					B *bool `json:"b,omitempty"`
				} `json:",omitempty"`
				B *struct {
					A *bool `json:"a,omitempty"`
					B *bool `json:"b,omitempty"`
				} `json:",omitempty"`
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadBoolPtrNilDoubleMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A *bool `json:"a,string"`
					B *bool `json:"b,string"`
				}
				B *struct {
					A *bool `json:"a,string"`
					B *bool `json:"b,string"`
				}
			}{A: nil, B: nil},
		},

		// PtrHeadBoolPtrNilDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadBoolPtrNilDoubleMultiFieldsNotRoot",
			data: (*struct {
				A *struct {
					A *bool `json:"a"`
					B *bool `json:"b"`
				}
				B *struct {
					A *bool `json:"a"`
					B *bool `json:"b"`
				}
			})(nil),
		},
		{
			name: "PtrHeadBoolPtrNilDoubleMultiFieldsNotRootOmitEmpty",
			data: (*struct {
				A *struct {
					A *bool `json:"a,omitempty"`
					B *bool `json:"b,omitempty"`
				}
				B *struct {
					A *bool `json:"a,omitempty"`
					B *bool `json:"b,omitempty"`
				}
			})(nil),
		},
		{
			name: "PtrHeadBoolPtrNilDoubleMultiFieldsNotRootString",
			data: (*struct {
				A *struct {
					A *bool `json:"a,string"`
					B *bool `json:"b,string"`
				}
				B *struct {
					A *bool `json:"a,string"`
					B *bool `json:"b,string"`
				}
			})(nil),
		},

		// AnonymousHeadBool
		{
			name: "AnonymousHeadBool",
			data: struct {
				structBool
				B bool `json:"b"`
			}{
				structBool: structBool{A: true},
				B:          false,
			},
		},
		{
			name: "AnonymousHeadBoolOmitEmpty",
			data: struct {
				structBoolOmitEmpty
				B bool `json:"b,omitempty"`
			}{
				structBoolOmitEmpty: structBoolOmitEmpty{A: true},
				B:                   false,
			},
		},
		{
			name: "AnonymousHeadBoolString",
			data: struct {
				structBoolString
				B bool `json:"b,string"`
			}{
				structBoolString: structBoolString{A: true},
				B:                false,
			},
		},

		// PtrAnonymousHeadBool
		{
			name: "PtrAnonymousHeadBool",
			data: struct {
				*structBool
				B bool `json:"b"`
			}{
				structBool: &structBool{A: true},
				B:          false,
			},
		},
		{
			name: "PtrAnonymousHeadBoolOmitEmpty",
			data: struct {
				*structBoolOmitEmpty
				B bool `json:"b,omitempty"`
			}{
				structBoolOmitEmpty: &structBoolOmitEmpty{A: true},
				B:                   false,
			},
		},
		{
			name: "PtrAnonymousHeadBoolString",
			data: struct {
				*structBoolString
				B bool `json:"b,string"`
			}{
				structBoolString: &structBoolString{A: true},
				B:                false,
			},
		},

		// NilPtrAnonymousHeadBool
		{
			name: "NilPtrAnonymousHeadBool",
			data: struct {
				*structBool
				B bool `json:"b"`
			}{
				structBool: nil,
				B:          true,
			},
		},
		{
			name: "NilPtrAnonymousHeadBoolOmitEmpty",
			data: struct {
				*structBoolOmitEmpty
				B bool `json:"b,omitempty"`
			}{
				structBoolOmitEmpty: nil,
				B:                   true,
			},
		},
		{
			name: "NilPtrAnonymousHeadBoolString",
			data: struct {
				*structBoolString
				B bool `json:"b,string"`
			}{
				structBoolString: nil,
				B:                true,
			},
		},

		// AnonymousHeadBoolPtr
		{
			name: "AnonymousHeadBoolPtr",
			data: struct {
				structBoolPtr
				B *bool `json:"b"`
			}{
				structBoolPtr: structBoolPtr{A: boolptr(true)},
				B:             boolptr(false),
			},
		},
		{
			name: "AnonymousHeadBoolPtrOmitEmpty",
			data: struct {
				structBoolPtrOmitEmpty
				B *bool `json:"b,omitempty"`
			}{
				structBoolPtrOmitEmpty: structBoolPtrOmitEmpty{A: boolptr(true)},
				B:                      boolptr(false),
			},
		},
		{
			name: "AnonymousHeadBoolPtrString",
			data: struct {
				structBoolPtrString
				B *bool `json:"b,string"`
			}{
				structBoolPtrString: structBoolPtrString{A: boolptr(true)},
				B:                   boolptr(false),
			},
		},

		// AnonymousHeadBoolPtrNil
		{
			name: "AnonymousHeadBoolPtrNil",
			data: struct {
				structBoolPtr
				B *bool `json:"b"`
			}{
				structBoolPtr: structBoolPtr{A: nil},
				B:             boolptr(true),
			},
		},
		{
			name: "AnonymousHeadBoolPtrNilOmitEmpty",
			data: struct {
				structBoolPtrOmitEmpty
				B *bool `json:"b,omitempty"`
			}{
				structBoolPtrOmitEmpty: structBoolPtrOmitEmpty{A: nil},
				B:                      boolptr(true),
			},
		},
		{
			name: "AnonymousHeadBoolPtrNilString",
			data: struct {
				structBoolPtrString
				B *bool `json:"b,string"`
			}{
				structBoolPtrString: structBoolPtrString{A: nil},
				B:                   boolptr(true),
			},
		},

		// PtrAnonymousHeadBoolPtr
		{
			name: "PtrAnonymousHeadBoolPtr",
			data: struct {
				*structBoolPtr
				B *bool `json:"b"`
			}{
				structBoolPtr: &structBoolPtr{A: boolptr(true)},
				B:             boolptr(false),
			},
		},
		{
			name: "PtrAnonymousHeadBoolPtrOmitEmpty",
			data: struct {
				*structBoolPtrOmitEmpty
				B *bool `json:"b,omitempty"`
			}{
				structBoolPtrOmitEmpty: &structBoolPtrOmitEmpty{A: boolptr(true)},
				B:                      boolptr(false),
			},
		},
		{
			name: "PtrAnonymousHeadBoolPtrString",
			data: struct {
				*structBoolPtrString
				B *bool `json:"b,string"`
			}{
				structBoolPtrString: &structBoolPtrString{A: boolptr(true)},
				B:                   boolptr(false),
			},
		},

		// NilPtrAnonymousHeadBoolPtr
		{
			name: "NilPtrAnonymousHeadBoolPtr",
			data: struct {
				*structBoolPtr
				B *bool `json:"b"`
			}{
				structBoolPtr: nil,
				B:             boolptr(true),
			},
		},
		{
			name: "NilPtrAnonymousHeadBoolPtrOmitEmpty",
			data: struct {
				*structBoolPtrOmitEmpty
				B *bool `json:"b,omitempty"`
			}{
				structBoolPtrOmitEmpty: nil,
				B:                      boolptr(true),
			},
		},
		{
			name: "NilPtrAnonymousHeadBoolPtrString",
			data: struct {
				*structBoolPtrString
				B *bool `json:"b,string"`
			}{
				structBoolPtrString: nil,
				B:                   boolptr(true),
			},
		},

		// AnonymousHeadBoolOnly
		{
			name: "AnonymousHeadBoolOnly",
			data: struct {
				structBool
			}{
				structBool: structBool{A: true},
			},
		},
		{
			name: "AnonymousHeadBoolOnlyOmitEmpty",
			data: struct {
				structBoolOmitEmpty
			}{
				structBoolOmitEmpty: structBoolOmitEmpty{A: true},
			},
		},
		{
			name: "AnonymousHeadBoolOnlyString",
			data: struct {
				structBoolString
			}{
				structBoolString: structBoolString{A: true},
			},
		},

		// PtrAnonymousHeadBoolOnly
		{
			name: "PtrAnonymousHeadBoolOnly",
			data: struct {
				*structBool
			}{
				structBool: &structBool{A: true},
			},
		},
		{
			name: "PtrAnonymousHeadBoolOnlyOmitEmpty",
			data: struct {
				*structBoolOmitEmpty
			}{
				structBoolOmitEmpty: &structBoolOmitEmpty{A: true},
			},
		},
		{
			name: "PtrAnonymousHeadBoolOnlyString",
			data: struct {
				*structBoolString
			}{
				structBoolString: &structBoolString{A: true},
			},
		},

		// NilPtrAnonymousHeadBoolOnly
		{
			name: "NilPtrAnonymousHeadBoolOnly",
			data: struct {
				*structBool
			}{
				structBool: nil,
			},
		},
		{
			name: "NilPtrAnonymousHeadBoolOnlyOmitEmpty",
			data: struct {
				*structBoolOmitEmpty
			}{
				structBoolOmitEmpty: nil,
			},
		},
		{
			name: "NilPtrAnonymousHeadBoolOnlyString",
			data: struct {
				*structBoolString
			}{
				structBoolString: nil,
			},
		},

		// AnonymousHeadBoolPtrOnly
		{
			name: "AnonymousHeadBoolPtrOnly",
			data: struct {
				structBoolPtr
			}{
				structBoolPtr: structBoolPtr{A: boolptr(true)},
			},
		},
		{
			name: "AnonymousHeadBoolPtrOnlyOmitEmpty",
			data: struct {
				structBoolPtrOmitEmpty
			}{
				structBoolPtrOmitEmpty: structBoolPtrOmitEmpty{A: boolptr(true)},
			},
		},
		{
			name: "AnonymousHeadBoolPtrOnlyString",
			data: struct {
				structBoolPtrString
			}{
				structBoolPtrString: structBoolPtrString{A: boolptr(true)},
			},
		},

		// AnonymousHeadBoolPtrNilOnly
		{
			name: "AnonymousHeadBoolPtrNilOnly",
			data: struct {
				structBoolPtr
			}{
				structBoolPtr: structBoolPtr{A: nil},
			},
		},
		{
			name: "AnonymousHeadBoolPtrNilOnlyOmitEmpty",
			data: struct {
				structBoolPtrOmitEmpty
			}{
				structBoolPtrOmitEmpty: structBoolPtrOmitEmpty{A: nil},
			},
		},
		{
			name: "AnonymousHeadBoolPtrNilOnlyString",
			data: struct {
				structBoolPtrString
			}{
				structBoolPtrString: structBoolPtrString{A: nil},
			},
		},

		// PtrAnonymousHeadBoolPtrOnly
		{
			name: "PtrAnonymousHeadBoolPtrOnly",
			data: struct {
				*structBoolPtr
			}{
				structBoolPtr: &structBoolPtr{A: boolptr(true)},
			},
		},
		{
			name: "PtrAnonymousHeadBoolPtrOnlyOmitEmpty",
			data: struct {
				*structBoolPtrOmitEmpty
			}{
				structBoolPtrOmitEmpty: &structBoolPtrOmitEmpty{A: boolptr(true)},
			},
		},
		{
			name: "PtrAnonymousHeadBoolPtrOnlyString",
			data: struct {
				*structBoolPtrString
			}{
				structBoolPtrString: &structBoolPtrString{A: boolptr(true)},
			},
		},

		// NilPtrAnonymousHeadBoolPtrOnly
		{
			name: "NilPtrAnonymousHeadBoolPtrOnly",
			data: struct {
				*structBoolPtr
			}{
				structBoolPtr: nil,
			},
		},
		{
			name: "NilPtrAnonymousHeadBoolPtrOnlyOmitEmpty",
			data: struct {
				*structBoolPtrOmitEmpty
			}{
				structBoolPtrOmitEmpty: nil,
			},
		},
		{
			name: "NilPtrAnonymousHeadBoolPtrOnlyString",
			data: struct {
				*structBoolPtrString
			}{
				structBoolPtrString: nil,
			},
		},
	}
	for _, test := range tests {
		for _, indent := range []bool{true, false} {
			for _, htmlEscape := range []bool{true, false} {
				var buf bytes.Buffer
				enc := json.NewEncoder(&buf)
				enc.SetEscapeHTML(htmlEscape)
				if indent {
					enc.SetIndent("", "  ")
				}
				if err := enc.Encode(test.data); err != nil {
					t.Fatalf("%s(htmlEscape:%v,indent:%v): %+v: %s", test.name, htmlEscape, indent, test.data, err)
				}
				stdresult := encodeByEncodingJSON(test.data, indent, htmlEscape)
				if buf.String() != stdresult {
					t.Errorf("%s(htmlEscape:%v,indent:%v): doesn't compatible with encoding/json. expected %q but got %q", test.name, htmlEscape, indent, stdresult, buf.String())
				}
			}
		}
	}
}

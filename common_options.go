package gitlab

import (
	"net/url"
	"strconv"
)

// APIOption abstracts all options, so we're easier to construct requests
type APIOption interface {
	// accepts an url.Values so we can combine multiple APIOptions by
	//   opt1.Encode(opt2.Encode(nil))
	// or even
	//   var val url.Values
	//   for _, o := range options {
	//       val = o.Encode(val)
	//   }
	Encode(v url.Values) url.Values
}

// Few helper function might be useful for structs implementing APIOption.
// These functions are named in "optTypeSuffix" format.
// Functions without suffix detects for zero value.

func optIntGT(val, expect int, v url.Values, key string) {
	if val > expect {
		v.Set(key, strconv.Itoa(val))
	}
}
func optInt(val int, v url.Values, key string) {
	optIntGT(val, 0, v, key)
}
func optStr(val string, v url.Values, key string) {
	if val != "" {
		v.Set(key, val)
	}
}

// ListOption represents pagination related options
type ListOption struct {
	Page    int
	PerPage int
}

func (o ListOption) Encode(v url.Values) url.Values {
	ret := v
	if v == nil {
		ret = url.Values{}
	}

	optIntGT(o.Page, 1, ret, "page")

	// gitlab accepts 20-100
	if o.PerPage > 0 {
		pp := o.PerPage
		if pp < 20 {
			pp = 20
		}
		if pp > 100 {
			pp = 100
		}
		ret.Set("per_page", strconv.Itoa(pp))
	}

	return ret
}

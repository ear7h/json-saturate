package main

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
)

func main() {
	var v interface{}
	err := json.NewDecoder(os.Stdin).Decode(&v)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(os.Stdout).Encode(Summarize(v))
}

func Summarize(v interface{}) interface{} {
	switch vv := v.(type) {
	default:
		return v

	case map[string]interface{}:
		ret := make(map[string]interface{}, len(vv))

		for k, v := range vv {
			ret[k] = Summarize(v)
		}

		return ret

	case []interface{}:
		ret := []interface{}{}

	L:
		for _, sv := range vv {
			for ei, ev := range ret {
				u, ok := Unify(ev, sv)
				if ok {
					ret[ei] = u
					continue L
				}
			}

			ret = append(ret, sv)
		}

		return ret
	}

	return v
}

func Unify(a, b interface{}) (interface{}, bool) {
	switch aa := a.(type) {
	default:

		if a == nil {
			return b, true
		}

		if b == nil {
			return a, true
		}

		av := reflect.ValueOf(a)
		bv := reflect.ValueOf(b)

		if av.Type() != bv.Type() {
			return nil, false
		}

		if av.IsZero() {
			return b, true
		}

		return a, true

	case []interface{}:
		bb, ok := b.([]interface{})
		if !ok {
			return nil, false
		}

		return Summarize(append(aa, bb...)), true

	case map[string]interface{}:
		bb, ok := b.(map[string]interface{})
		if !ok {
			return nil, false
		}

		keys := map[string]struct{}{}

		for k := range aa {
			keys[k] = struct{}{}
		}

		for k := range bb {
			keys[k] = struct{}{}
		}

		ret := make(map[string]interface{}, len(keys))

		for k := range keys {
			av, ok := aa[k]
			if !ok {
				ret[k] = bb[k]
				continue
			}

			bv, ok := bb[k]
			if !ok {
				ret[k] = aa[k]
				continue
			}

			u, ok := Unify(av, bv)
			if !ok {
				return nil, false
			}

			ret[k] = u
		}

		return ret, true
	}
}

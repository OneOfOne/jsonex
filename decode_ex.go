package jsonex

import (
	"reflect"
)

func (dec *Decoder) DecodeValue(v any) error {
	tok, err := dec.Token()
	if err != nil {
		return err
	}
	switch tok {
	case Delim('['):
		return dec.Decode(v)
	case Delim('{'):
		dec.scanp--
		dec.tokenState = tokenArrayValue
		return dec.Decode(v)
	case nil:
		return nil
	default:
		reflect.ValueOf(v).Elem().Set(reflect.ValueOf(tok))
		return nil
	}
}

func (dec *Decoder) DecodeKey() (string, error) {
	tok, err := dec.Token()
	if err != nil {
		return "", err
	}
	switch tok {
	case Delim('{'):
		tok, err := dec.Token()
		if err != nil {
			return "", err
		}
		return tok.(string), nil
	case nil:
		return "", nil
	default:
		if key, ok := tok.(string); ok {
			return key, nil
		}
		return "", &SyntaxError{"invalid key: ", dec.InputOffset()}
	}
}

// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package DTO

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson13d317efDecodeBrabusInternalDTO(in *jlexer.Lexer, out *Disk) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "space":
			out.Space = string(in.String())
		case "usage":
			out.Usage = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson13d317efEncodeBrabusInternalDTO(out *jwriter.Writer, in Disk) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"space\":"
		out.RawString(prefix[1:])
		out.String(string(in.Space))
	}
	{
		const prefix string = ",\"usage\":"
		out.RawString(prefix)
		out.String(string(in.Usage))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Disk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson13d317efEncodeBrabusInternalDTO(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Disk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson13d317efEncodeBrabusInternalDTO(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Disk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson13d317efDecodeBrabusInternalDTO(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Disk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson13d317efDecodeBrabusInternalDTO(l, v)
}

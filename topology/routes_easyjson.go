// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package topology

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

func easyjson43db4356DecodeGithubComSkydiveProjectSkydiveTopology(in *jlexer.Lexer, out *RoutingTables) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(RoutingTables, 0, 8)
			} else {
				*out = RoutingTables{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *RoutingTable
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(RoutingTable)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson43db4356EncodeGithubComSkydiveProjectSkydiveTopology(out *jwriter.Writer, in RoutingTables) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				(*v3).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v RoutingTables) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson43db4356EncodeGithubComSkydiveProjectSkydiveTopology(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RoutingTables) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson43db4356EncodeGithubComSkydiveProjectSkydiveTopology(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RoutingTables) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson43db4356DecodeGithubComSkydiveProjectSkydiveTopology(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RoutingTables) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson43db4356DecodeGithubComSkydiveProjectSkydiveTopology(l, v)
}
func easyjson43db4356DecodeGithubComSkydiveProjectSkydiveTopology1(in *jlexer.Lexer, out *RoutingTable) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ID":
			out.ID = int64(in.Int64())
		case "Src":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.Src).UnmarshalText(data))
			}
		case "Routes":
			if in.IsNull() {
				in.Skip()
				out.Routes = nil
			} else {
				in.Delim('[')
				if out.Routes == nil {
					if !in.IsDelim(']') {
						out.Routes = make([]*Route, 0, 8)
					} else {
						out.Routes = []*Route{}
					}
				} else {
					out.Routes = (out.Routes)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *Route
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(Route)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Routes = append(out.Routes, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson43db4356EncodeGithubComSkydiveProjectSkydiveTopology1(out *jwriter.Writer, in RoutingTable) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"Src\":"
		out.RawString(prefix)
		out.RawText((in.Src).MarshalText())
	}
	{
		const prefix string = ",\"Routes\":"
		out.RawString(prefix)
		if in.Routes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Routes {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RoutingTable) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson43db4356EncodeGithubComSkydiveProjectSkydiveTopology1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RoutingTable) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson43db4356EncodeGithubComSkydiveProjectSkydiveTopology1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RoutingTable) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson43db4356DecodeGithubComSkydiveProjectSkydiveTopology1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RoutingTable) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson43db4356DecodeGithubComSkydiveProjectSkydiveTopology1(l, v)
}
func easyjson43db4356DecodeGithubComSkydiveProjectSkydiveTopology2(in *jlexer.Lexer, out *Route) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Protocol":
			out.Protocol = int64(in.Int64())
		case "Prefix":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Prefix).UnmarshalJSON(data))
			}
		case "NextHops":
			if in.IsNull() {
				in.Skip()
				out.NextHops = nil
			} else {
				in.Delim('[')
				if out.NextHops == nil {
					if !in.IsDelim(']') {
						out.NextHops = make([]*NextHop, 0, 8)
					} else {
						out.NextHops = []*NextHop{}
					}
				} else {
					out.NextHops = (out.NextHops)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *NextHop
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(NextHop)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.NextHops = append(out.NextHops, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson43db4356EncodeGithubComSkydiveProjectSkydiveTopology2(out *jwriter.Writer, in Route) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Protocol\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Protocol))
	}
	{
		const prefix string = ",\"Prefix\":"
		out.RawString(prefix)
		out.Raw((in.Prefix).MarshalJSON())
	}
	{
		const prefix string = ",\"NextHops\":"
		out.RawString(prefix)
		if in.NextHops == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.NextHops {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					(*v9).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Route) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson43db4356EncodeGithubComSkydiveProjectSkydiveTopology2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Route) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson43db4356EncodeGithubComSkydiveProjectSkydiveTopology2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Route) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson43db4356DecodeGithubComSkydiveProjectSkydiveTopology2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Route) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson43db4356DecodeGithubComSkydiveProjectSkydiveTopology2(l, v)
}

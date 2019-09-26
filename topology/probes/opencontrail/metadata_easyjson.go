// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package opencontrail

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

func easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail(in *jlexer.Lexer, out *RoutingTable) {
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
		case "InterfacesUUID":
			if in.IsNull() {
				in.Skip()
				out.InterfacesUUID = nil
			} else {
				in.Delim('[')
				if out.InterfacesUUID == nil {
					if !in.IsDelim(']') {
						out.InterfacesUUID = make([]string, 0, 4)
					} else {
						out.InterfacesUUID = []string{}
					}
				} else {
					out.InterfacesUUID = (out.InterfacesUUID)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.InterfacesUUID = append(out.InterfacesUUID, v1)
					in.WantComma()
				}
				in.Delim(']')
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
					var v2 *Route
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(Route)
						}
						(*v2).UnmarshalEasyJSON(in)
					}
					out.Routes = append(out.Routes, v2)
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
func easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail(out *jwriter.Writer, in RoutingTable) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"InterfacesUUID\":"
		out.RawString(prefix[1:])
		if in.InterfacesUUID == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.InterfacesUUID {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
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
	easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RoutingTable) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RoutingTable) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RoutingTable) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail(l, v)
}
func easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail1(in *jlexer.Lexer, out *Route) {
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
		case "Family":
			out.Family = string(in.String())
		case "Prefix":
			out.Prefix = string(in.String())
		case "NhID":
			out.NhID = int64(in.Int64())
		case "Protocol":
			out.Protocol = int64(in.Int64())
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
func easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail1(out *jwriter.Writer, in Route) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Family != "" {
		const prefix string = ",\"Family\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Family))
	}
	if in.Prefix != "" {
		const prefix string = ",\"Prefix\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Prefix))
	}
	if in.NhID != 0 {
		const prefix string = ",\"NhID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.NhID))
	}
	if in.Protocol != 0 {
		const prefix string = ",\"Protocol\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Protocol))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Route) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Route) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Route) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Route) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail1(l, v)
}
func easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail2(in *jlexer.Lexer, out *Metadata) {
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
		case "UUID":
			out.UUID = string(in.String())
		case "MAC":
			out.MAC = string(in.String())
		case "VRF":
			out.VRF = string(in.String())
		case "VRFID":
			out.VRFID = int64(in.Int64())
		case "LocalIP":
			out.LocalIP = string(in.String())
		case "RoutingTable":
			if in.IsNull() {
				in.Skip()
				out.RoutingTable = nil
			} else {
				in.Delim('[')
				if out.RoutingTable == nil {
					if !in.IsDelim(']') {
						out.RoutingTable = make([]*Route, 0, 8)
					} else {
						out.RoutingTable = []*Route{}
					}
				} else {
					out.RoutingTable = (out.RoutingTable)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *Route
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(Route)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.RoutingTable = append(out.RoutingTable, v7)
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
func easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail2(out *jwriter.Writer, in Metadata) {
	out.RawByte('{')
	first := true
	_ = first
	if in.UUID != "" {
		const prefix string = ",\"UUID\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.UUID))
	}
	if in.MAC != "" {
		const prefix string = ",\"MAC\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MAC))
	}
	if in.VRF != "" {
		const prefix string = ",\"VRF\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.VRF))
	}
	if in.VRFID != 0 {
		const prefix string = ",\"VRFID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.VRFID))
	}
	if in.LocalIP != "" {
		const prefix string = ",\"LocalIP\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LocalIP))
	}
	if len(in.RoutingTable) != 0 {
		const prefix string = ",\"RoutingTable\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.RoutingTable {
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
func (v Metadata) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Metadata) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Metadata) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Metadata) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesOpencontrail2(l, v)
}

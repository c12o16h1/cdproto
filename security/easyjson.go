// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package security

import (
	json "encoding/json"
	cdp "github.com/c12o16h1/cdproto/cdp"
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

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity(in *jlexer.Lexer, out *VisibleSecurityState) {
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
		case "securityState":
			(out.SecurityState).UnmarshalEasyJSON(in)
		case "certificateSecurityState":
			if in.IsNull() {
				in.Skip()
				out.CertificateSecurityState = nil
			} else {
				if out.CertificateSecurityState == nil {
					out.CertificateSecurityState = new(CertificateSecurityState)
				}
				(*out.CertificateSecurityState).UnmarshalEasyJSON(in)
			}
		case "securityStateIssueIds":
			if in.IsNull() {
				in.Skip()
				out.SecurityStateIssueIds = nil
			} else {
				in.Delim('[')
				if out.SecurityStateIssueIds == nil {
					if !in.IsDelim(']') {
						out.SecurityStateIssueIds = make([]string, 0, 4)
					} else {
						out.SecurityStateIssueIds = []string{}
					}
				} else {
					out.SecurityStateIssueIds = (out.SecurityStateIssueIds)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.SecurityStateIssueIds = append(out.SecurityStateIssueIds, v1)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity(out *jwriter.Writer, in VisibleSecurityState) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"securityState\":"
		out.RawString(prefix[1:])
		(in.SecurityState).MarshalEasyJSON(out)
	}
	if in.CertificateSecurityState != nil {
		const prefix string = ",\"certificateSecurityState\":"
		out.RawString(prefix)
		(*in.CertificateSecurityState).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"securityStateIssueIds\":"
		out.RawString(prefix)
		if in.SecurityStateIssueIds == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.SecurityStateIssueIds {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VisibleSecurityState) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v VisibleSecurityState) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VisibleSecurityState) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *VisibleSecurityState) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity1(in *jlexer.Lexer, out *StateExplanation) {
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
		case "securityState":
			(out.SecurityState).UnmarshalEasyJSON(in)
		case "title":
			out.Title = string(in.String())
		case "summary":
			out.Summary = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "mixedContentType":
			(out.MixedContentType).UnmarshalEasyJSON(in)
		case "certificate":
			if in.IsNull() {
				in.Skip()
				out.Certificate = nil
			} else {
				in.Delim('[')
				if out.Certificate == nil {
					if !in.IsDelim(']') {
						out.Certificate = make([]string, 0, 4)
					} else {
						out.Certificate = []string{}
					}
				} else {
					out.Certificate = (out.Certificate)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Certificate = append(out.Certificate, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "recommendations":
			if in.IsNull() {
				in.Skip()
				out.Recommendations = nil
			} else {
				in.Delim('[')
				if out.Recommendations == nil {
					if !in.IsDelim(']') {
						out.Recommendations = make([]string, 0, 4)
					} else {
						out.Recommendations = []string{}
					}
				} else {
					out.Recommendations = (out.Recommendations)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.Recommendations = append(out.Recommendations, v5)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity1(out *jwriter.Writer, in StateExplanation) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"securityState\":"
		out.RawString(prefix[1:])
		(in.SecurityState).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"summary\":"
		out.RawString(prefix)
		out.String(string(in.Summary))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"mixedContentType\":"
		out.RawString(prefix)
		(in.MixedContentType).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"certificate\":"
		out.RawString(prefix)
		if in.Certificate == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Certificate {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	if len(in.Recommendations) != 0 {
		const prefix string = ",\"recommendations\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v8, v9 := range in.Recommendations {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StateExplanation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StateExplanation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StateExplanation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StateExplanation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity1(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity2(in *jlexer.Lexer, out *SetIgnoreCertificateErrorsParams) {
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
		case "ignore":
			out.Ignore = bool(in.Bool())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity2(out *jwriter.Writer, in SetIgnoreCertificateErrorsParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ignore\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Ignore))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetIgnoreCertificateErrorsParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetIgnoreCertificateErrorsParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetIgnoreCertificateErrorsParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetIgnoreCertificateErrorsParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity2(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity3(in *jlexer.Lexer, out *EventVisibleSecurityStateChanged) {
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
		case "visibleSecurityState":
			if in.IsNull() {
				in.Skip()
				out.VisibleSecurityState = nil
			} else {
				if out.VisibleSecurityState == nil {
					out.VisibleSecurityState = new(VisibleSecurityState)
				}
				(*out.VisibleSecurityState).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity3(out *jwriter.Writer, in EventVisibleSecurityStateChanged) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"visibleSecurityState\":"
		out.RawString(prefix[1:])
		if in.VisibleSecurityState == nil {
			out.RawString("null")
		} else {
			(*in.VisibleSecurityState).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventVisibleSecurityStateChanged) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventVisibleSecurityStateChanged) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventVisibleSecurityStateChanged) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventVisibleSecurityStateChanged) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity3(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity4(in *jlexer.Lexer, out *EventSecurityStateChanged) {
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
		case "securityState":
			(out.SecurityState).UnmarshalEasyJSON(in)
		case "explanations":
			if in.IsNull() {
				in.Skip()
				out.Explanations = nil
			} else {
				in.Delim('[')
				if out.Explanations == nil {
					if !in.IsDelim(']') {
						out.Explanations = make([]*StateExplanation, 0, 8)
					} else {
						out.Explanations = []*StateExplanation{}
					}
				} else {
					out.Explanations = (out.Explanations)[:0]
				}
				for !in.IsDelim(']') {
					var v10 *StateExplanation
					if in.IsNull() {
						in.Skip()
						v10 = nil
					} else {
						if v10 == nil {
							v10 = new(StateExplanation)
						}
						(*v10).UnmarshalEasyJSON(in)
					}
					out.Explanations = append(out.Explanations, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "summary":
			out.Summary = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity4(out *jwriter.Writer, in EventSecurityStateChanged) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"securityState\":"
		out.RawString(prefix[1:])
		(in.SecurityState).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"explanations\":"
		out.RawString(prefix)
		if in.Explanations == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Explanations {
				if v11 > 0 {
					out.RawByte(',')
				}
				if v12 == nil {
					out.RawString("null")
				} else {
					(*v12).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Summary != "" {
		const prefix string = ",\"summary\":"
		out.RawString(prefix)
		out.String(string(in.Summary))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventSecurityStateChanged) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventSecurityStateChanged) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventSecurityStateChanged) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventSecurityStateChanged) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity4(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity5(in *jlexer.Lexer, out *EnableParams) {
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity5(out *jwriter.Writer, in EnableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EnableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EnableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EnableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EnableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity5(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity6(in *jlexer.Lexer, out *DisableParams) {
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity6(out *jwriter.Writer, in DisableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DisableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DisableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DisableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DisableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity6(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity7(in *jlexer.Lexer, out *CertificateSecurityState) {
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
		case "protocol":
			out.Protocol = string(in.String())
		case "keyExchange":
			out.KeyExchange = string(in.String())
		case "keyExchangeGroup":
			out.KeyExchangeGroup = string(in.String())
		case "cipher":
			out.Cipher = string(in.String())
		case "mac":
			out.Mac = string(in.String())
		case "certificate":
			if in.IsNull() {
				in.Skip()
				out.Certificate = nil
			} else {
				in.Delim('[')
				if out.Certificate == nil {
					if !in.IsDelim(']') {
						out.Certificate = make([]string, 0, 4)
					} else {
						out.Certificate = []string{}
					}
				} else {
					out.Certificate = (out.Certificate)[:0]
				}
				for !in.IsDelim(']') {
					var v13 string
					v13 = string(in.String())
					out.Certificate = append(out.Certificate, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "subjectName":
			out.SubjectName = string(in.String())
		case "issuer":
			out.Issuer = string(in.String())
		case "validFrom":
			if in.IsNull() {
				in.Skip()
				out.ValidFrom = nil
			} else {
				if out.ValidFrom == nil {
					out.ValidFrom = new(cdp.TimeSinceEpoch)
				}
				(*out.ValidFrom).UnmarshalEasyJSON(in)
			}
		case "validTo":
			if in.IsNull() {
				in.Skip()
				out.ValidTo = nil
			} else {
				if out.ValidTo == nil {
					out.ValidTo = new(cdp.TimeSinceEpoch)
				}
				(*out.ValidTo).UnmarshalEasyJSON(in)
			}
		case "certifcateHasWeakSignature":
			out.CertifcateHasWeakSignature = bool(in.Bool())
		case "modernSSL":
			out.ModernSSL = bool(in.Bool())
		case "obsoleteSslProtocol":
			out.ObsoleteSslProtocol = bool(in.Bool())
		case "obsoleteSslKeyExchange":
			out.ObsoleteSslKeyExchange = bool(in.Bool())
		case "obsoleteSslCipher":
			out.ObsoleteSslCipher = bool(in.Bool())
		case "obsoleteSslSignature":
			out.ObsoleteSslSignature = bool(in.Bool())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity7(out *jwriter.Writer, in CertificateSecurityState) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"protocol\":"
		out.RawString(prefix[1:])
		out.String(string(in.Protocol))
	}
	{
		const prefix string = ",\"keyExchange\":"
		out.RawString(prefix)
		out.String(string(in.KeyExchange))
	}
	if in.KeyExchangeGroup != "" {
		const prefix string = ",\"keyExchangeGroup\":"
		out.RawString(prefix)
		out.String(string(in.KeyExchangeGroup))
	}
	{
		const prefix string = ",\"cipher\":"
		out.RawString(prefix)
		out.String(string(in.Cipher))
	}
	if in.Mac != "" {
		const prefix string = ",\"mac\":"
		out.RawString(prefix)
		out.String(string(in.Mac))
	}
	{
		const prefix string = ",\"certificate\":"
		out.RawString(prefix)
		if in.Certificate == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Certificate {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.String(string(v15))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"subjectName\":"
		out.RawString(prefix)
		out.String(string(in.SubjectName))
	}
	{
		const prefix string = ",\"issuer\":"
		out.RawString(prefix)
		out.String(string(in.Issuer))
	}
	{
		const prefix string = ",\"validFrom\":"
		out.RawString(prefix)
		if in.ValidFrom == nil {
			out.RawString("null")
		} else {
			(*in.ValidFrom).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"validTo\":"
		out.RawString(prefix)
		if in.ValidTo == nil {
			out.RawString("null")
		} else {
			(*in.ValidTo).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"certifcateHasWeakSignature\":"
		out.RawString(prefix)
		out.Bool(bool(in.CertifcateHasWeakSignature))
	}
	{
		const prefix string = ",\"modernSSL\":"
		out.RawString(prefix)
		out.Bool(bool(in.ModernSSL))
	}
	{
		const prefix string = ",\"obsoleteSslProtocol\":"
		out.RawString(prefix)
		out.Bool(bool(in.ObsoleteSslProtocol))
	}
	{
		const prefix string = ",\"obsoleteSslKeyExchange\":"
		out.RawString(prefix)
		out.Bool(bool(in.ObsoleteSslKeyExchange))
	}
	{
		const prefix string = ",\"obsoleteSslCipher\":"
		out.RawString(prefix)
		out.Bool(bool(in.ObsoleteSslCipher))
	}
	{
		const prefix string = ",\"obsoleteSslSignature\":"
		out.RawString(prefix)
		out.Bool(bool(in.ObsoleteSslSignature))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CertificateSecurityState) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CertificateSecurityState) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoSecurity7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CertificateSecurityState) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CertificateSecurityState) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoSecurity7(l, v)
}

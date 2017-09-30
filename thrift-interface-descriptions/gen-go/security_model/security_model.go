// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package security_model

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

// Attributes:
//  - AccessToken
//  - ClaimsMap
type AuthzToken struct {
  AccessToken string `thrift:"accessToken,1,required" db:"accessToken" json:"accessToken"`
  ClaimsMap map[string]string `thrift:"claimsMap,2" db:"claimsMap" json:"claimsMap,omitempty"`
}

func NewAuthzToken() *AuthzToken {
  return &AuthzToken{}
}


func (p *AuthzToken) GetAccessToken() string {
  return p.AccessToken
}
var AuthzToken_ClaimsMap_DEFAULT map[string]string

func (p *AuthzToken) GetClaimsMap() map[string]string {
  return p.ClaimsMap
}
func (p *AuthzToken) IsSetClaimsMap() bool {
  return p.ClaimsMap != nil
}

func (p *AuthzToken) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetAccessToken bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
      issetAccessToken = true
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetAccessToken{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field AccessToken is not set"));
  }
  return nil
}

func (p *AuthzToken)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.AccessToken = v
}
  return nil
}

func (p *AuthzToken)  ReadField2(iprot thrift.TProtocol) error {
  _, _, size, err := iprot.ReadMapBegin()
  if err != nil {
    return thrift.PrependError("error reading map begin: ", err)
  }
  tMap := make(map[string]string, size)
  p.ClaimsMap =  tMap
  for i := 0; i < size; i ++ {
var _key0 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _key0 = v
}
var _val1 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _val1 = v
}
    p.ClaimsMap[_key0] = _val1
  }
  if err := iprot.ReadMapEnd(); err != nil {
    return thrift.PrependError("error reading map end: ", err)
  }
  return nil
}

func (p *AuthzToken) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("AuthzToken"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *AuthzToken) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("accessToken", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:accessToken: ", p), err) }
  if err := oprot.WriteString(string(p.AccessToken)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.accessToken (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:accessToken: ", p), err) }
  return err
}

func (p *AuthzToken) writeField2(oprot thrift.TProtocol) (err error) {
  if p.IsSetClaimsMap() {
    if err := oprot.WriteFieldBegin("claimsMap", thrift.MAP, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:claimsMap: ", p), err) }
    if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.ClaimsMap)); err != nil {
      return thrift.PrependError("error writing map begin: ", err)
    }
    for k, v := range p.ClaimsMap {
      if err := oprot.WriteString(string(k)); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
      if err := oprot.WriteString(string(v)); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    }
    if err := oprot.WriteMapEnd(); err != nil {
      return thrift.PrependError("error writing map end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:claimsMap: ", p), err) }
  }
  return err
}

func (p *AuthzToken) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("AuthzToken(%+v)", *p)
}


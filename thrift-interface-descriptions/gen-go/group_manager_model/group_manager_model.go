// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package group_manager_model

import (
	"bytes"
	"reflect"
	"database/sql/driver"
	"errors"
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"airavata_commons"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

var _ = airavata_commons.GoUnusedProtection__
type ResourceType int64
const (
  ResourceType_PROJECT ResourceType = 0
  ResourceType_EXPERIMENT ResourceType = 1
  ResourceType_DATA ResourceType = 2
  ResourceType_OTHER ResourceType = 3
)

func (p ResourceType) String() string {
  switch p {
  case ResourceType_PROJECT: return "PROJECT"
  case ResourceType_EXPERIMENT: return "EXPERIMENT"
  case ResourceType_DATA: return "DATA"
  case ResourceType_OTHER: return "OTHER"
  }
  return "<UNSET>"
}

func ResourceTypeFromString(s string) (ResourceType, error) {
  switch s {
  case "PROJECT": return ResourceType_PROJECT, nil 
  case "EXPERIMENT": return ResourceType_EXPERIMENT, nil 
  case "DATA": return ResourceType_DATA, nil 
  case "OTHER": return ResourceType_OTHER, nil 
  }
  return ResourceType(0), fmt.Errorf("not a valid ResourceType string")
}


func ResourceTypePtr(v ResourceType) *ResourceType { return &v }

func (p ResourceType) MarshalText() ([]byte, error) {
return []byte(p.String()), nil
}

func (p *ResourceType) UnmarshalText(text []byte) error {
q, err := ResourceTypeFromString(string(text))
if (err != nil) {
return err
}
*p = q
return nil
}

func (p *ResourceType) Scan(value interface{}) error {
v, ok := value.(int64)
if !ok {
return errors.New("Scan value is not int64")
}
*p = ResourceType(v)
return nil
}

func (p * ResourceType) Value() (driver.Value, error) {
  if p == nil {
    return nil, nil
  }
return int64(*p), nil
}
type ResourcePermissionType int64
const (
  ResourcePermissionType_WRITE ResourcePermissionType = 0
  ResourcePermissionType_READ ResourcePermissionType = 1
  ResourcePermissionType_OWNER ResourcePermissionType = 2
)

func (p ResourcePermissionType) String() string {
  switch p {
  case ResourcePermissionType_WRITE: return "WRITE"
  case ResourcePermissionType_READ: return "READ"
  case ResourcePermissionType_OWNER: return "OWNER"
  }
  return "<UNSET>"
}

func ResourcePermissionTypeFromString(s string) (ResourcePermissionType, error) {
  switch s {
  case "WRITE": return ResourcePermissionType_WRITE, nil 
  case "READ": return ResourcePermissionType_READ, nil 
  case "OWNER": return ResourcePermissionType_OWNER, nil 
  }
  return ResourcePermissionType(0), fmt.Errorf("not a valid ResourcePermissionType string")
}


func ResourcePermissionTypePtr(v ResourcePermissionType) *ResourcePermissionType { return &v }

func (p ResourcePermissionType) MarshalText() ([]byte, error) {
return []byte(p.String()), nil
}

func (p *ResourcePermissionType) UnmarshalText(text []byte) error {
q, err := ResourcePermissionTypeFromString(string(text))
if (err != nil) {
return err
}
*p = q
return nil
}

func (p *ResourcePermissionType) Scan(value interface{}) error {
v, ok := value.(int64)
if !ok {
return errors.New("Scan value is not int64")
}
*p = ResourcePermissionType(v)
return nil
}

func (p * ResourcePermissionType) Value() (driver.Value, error) {
  if p == nil {
    return nil, nil
  }
return int64(*p), nil
}
// Attributes:
//  - ID
//  - Name
//  - OwnerId
//  - Description
//  - Members
type GroupModel struct {
  ID *string `thrift:"id,1" db:"id" json:"id,omitempty"`
  Name *string `thrift:"name,2" db:"name" json:"name,omitempty"`
  OwnerId *string `thrift:"ownerId,3" db:"ownerId" json:"ownerId,omitempty"`
  Description *string `thrift:"description,4" db:"description" json:"description,omitempty"`
  Members []string `thrift:"members,5" db:"members" json:"members,omitempty"`
}

func NewGroupModel() *GroupModel {
  return &GroupModel{}
}

var GroupModel_ID_DEFAULT string
func (p *GroupModel) GetID() string {
  if !p.IsSetID() {
    return GroupModel_ID_DEFAULT
  }
return *p.ID
}
var GroupModel_Name_DEFAULT string
func (p *GroupModel) GetName() string {
  if !p.IsSetName() {
    return GroupModel_Name_DEFAULT
  }
return *p.Name
}
var GroupModel_OwnerId_DEFAULT string
func (p *GroupModel) GetOwnerId() string {
  if !p.IsSetOwnerId() {
    return GroupModel_OwnerId_DEFAULT
  }
return *p.OwnerId
}
var GroupModel_Description_DEFAULT string
func (p *GroupModel) GetDescription() string {
  if !p.IsSetDescription() {
    return GroupModel_Description_DEFAULT
  }
return *p.Description
}
var GroupModel_Members_DEFAULT []string

func (p *GroupModel) GetMembers() []string {
  return p.Members
}
func (p *GroupModel) IsSetID() bool {
  return p.ID != nil
}

func (p *GroupModel) IsSetName() bool {
  return p.Name != nil
}

func (p *GroupModel) IsSetOwnerId() bool {
  return p.OwnerId != nil
}

func (p *GroupModel) IsSetDescription() bool {
  return p.Description != nil
}

func (p *GroupModel) IsSetMembers() bool {
  return p.Members != nil
}

func (p *GroupModel) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField4(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 5:
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField5(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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
  return nil
}

func (p *GroupModel)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.ID = &v
}
  return nil
}

func (p *GroupModel)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Name = &v
}
  return nil
}

func (p *GroupModel)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.OwnerId = &v
}
  return nil
}

func (p *GroupModel)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Description = &v
}
  return nil
}

func (p *GroupModel)  ReadField5(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]string, 0, size)
  p.Members =  tSlice
  for i := 0; i < size; i ++ {
var _elem0 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem0 = v
}
    p.Members = append(p.Members, _elem0)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *GroupModel) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("GroupModel"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
    if err := p.writeField5(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GroupModel) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetID() {
    if err := oprot.WriteFieldBegin("id", thrift.STRING, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err) }
    if err := oprot.WriteString(string(*p.ID)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err) }
  }
  return err
}

func (p *GroupModel) writeField2(oprot thrift.TProtocol) (err error) {
  if p.IsSetName() {
    if err := oprot.WriteFieldBegin("name", thrift.STRING, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:name: ", p), err) }
    if err := oprot.WriteString(string(*p.Name)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.name (2) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:name: ", p), err) }
  }
  return err
}

func (p *GroupModel) writeField3(oprot thrift.TProtocol) (err error) {
  if p.IsSetOwnerId() {
    if err := oprot.WriteFieldBegin("ownerId", thrift.STRING, 3); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:ownerId: ", p), err) }
    if err := oprot.WriteString(string(*p.OwnerId)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.ownerId (3) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 3:ownerId: ", p), err) }
  }
  return err
}

func (p *GroupModel) writeField4(oprot thrift.TProtocol) (err error) {
  if p.IsSetDescription() {
    if err := oprot.WriteFieldBegin("description", thrift.STRING, 4); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:description: ", p), err) }
    if err := oprot.WriteString(string(*p.Description)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.description (4) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 4:description: ", p), err) }
  }
  return err
}

func (p *GroupModel) writeField5(oprot thrift.TProtocol) (err error) {
  if p.IsSetMembers() {
    if err := oprot.WriteFieldBegin("members", thrift.LIST, 5); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:members: ", p), err) }
    if err := oprot.WriteListBegin(thrift.STRING, len(p.Members)); err != nil {
      return thrift.PrependError("error writing list begin: ", err)
    }
    for _, v := range p.Members {
      if err := oprot.WriteString(string(v)); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    }
    if err := oprot.WriteListEnd(); err != nil {
      return thrift.PrependError("error writing list end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 5:members: ", p), err) }
  }
  return err
}

func (p *GroupModel) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GroupModel(%+v)", *p)
}


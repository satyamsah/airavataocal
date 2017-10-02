// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package scheduling_model

import (
	"bytes"
	"reflect"
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

// ComputationalResourceSchedulingModel:
// 
// 
// 
// Attributes:
//  - ResourceHostId
//  - TotalCPUCount
//  - NodeCount
//  - NumberOfThreads
//  - QueueName
//  - WallTimeLimit
//  - TotalPhysicalMemory
//  - ChessisNumber
//  - StaticWorkingDir
//  - OverrideLoginUserName
//  - OverrideScratchLocation
//  - OverrideAllocationProjectNumber
type ComputationalResourceSchedulingModel struct {
  ResourceHostId *string `thrift:"resourceHostId,1" db:"resourceHostId" json:"resourceHostId,omitempty"`
  TotalCPUCount *int32 `thrift:"totalCPUCount,2" db:"totalCPUCount" json:"totalCPUCount,omitempty"`
  NodeCount *int32 `thrift:"nodeCount,3" db:"nodeCount" json:"nodeCount,omitempty"`
  NumberOfThreads *int32 `thrift:"numberOfThreads,4" db:"numberOfThreads" json:"numberOfThreads,omitempty"`
  QueueName *string `thrift:"queueName,5" db:"queueName" json:"queueName,omitempty"`
  WallTimeLimit *int32 `thrift:"wallTimeLimit,6" db:"wallTimeLimit" json:"wallTimeLimit,omitempty"`
  TotalPhysicalMemory *int32 `thrift:"totalPhysicalMemory,7" db:"totalPhysicalMemory" json:"totalPhysicalMemory,omitempty"`
  ChessisNumber *string `thrift:"chessisNumber,8" db:"chessisNumber" json:"chessisNumber,omitempty"`
  StaticWorkingDir *string `thrift:"staticWorkingDir,9" db:"staticWorkingDir" json:"staticWorkingDir,omitempty"`
  OverrideLoginUserName *string `thrift:"overrideLoginUserName,10" db:"overrideLoginUserName" json:"overrideLoginUserName,omitempty"`
  OverrideScratchLocation *string `thrift:"overrideScratchLocation,11" db:"overrideScratchLocation" json:"overrideScratchLocation,omitempty"`
  OverrideAllocationProjectNumber *string `thrift:"overrideAllocationProjectNumber,12" db:"overrideAllocationProjectNumber" json:"overrideAllocationProjectNumber,omitempty"`
}

func NewComputationalResourceSchedulingModel() *ComputationalResourceSchedulingModel {
  return &ComputationalResourceSchedulingModel{}
}

var ComputationalResourceSchedulingModel_ResourceHostId_DEFAULT string
func (p *ComputationalResourceSchedulingModel) GetResourceHostId() string {
  if !p.IsSetResourceHostId() {
    return ComputationalResourceSchedulingModel_ResourceHostId_DEFAULT
  }
return *p.ResourceHostId
}
var ComputationalResourceSchedulingModel_TotalCPUCount_DEFAULT int32
func (p *ComputationalResourceSchedulingModel) GetTotalCPUCount() int32 {
  if !p.IsSetTotalCPUCount() {
    return ComputationalResourceSchedulingModel_TotalCPUCount_DEFAULT
  }
return *p.TotalCPUCount
}
var ComputationalResourceSchedulingModel_NodeCount_DEFAULT int32
func (p *ComputationalResourceSchedulingModel) GetNodeCount() int32 {
  if !p.IsSetNodeCount() {
    return ComputationalResourceSchedulingModel_NodeCount_DEFAULT
  }
return *p.NodeCount
}
var ComputationalResourceSchedulingModel_NumberOfThreads_DEFAULT int32
func (p *ComputationalResourceSchedulingModel) GetNumberOfThreads() int32 {
  if !p.IsSetNumberOfThreads() {
    return ComputationalResourceSchedulingModel_NumberOfThreads_DEFAULT
  }
return *p.NumberOfThreads
}
var ComputationalResourceSchedulingModel_QueueName_DEFAULT string
func (p *ComputationalResourceSchedulingModel) GetQueueName() string {
  if !p.IsSetQueueName() {
    return ComputationalResourceSchedulingModel_QueueName_DEFAULT
  }
return *p.QueueName
}
var ComputationalResourceSchedulingModel_WallTimeLimit_DEFAULT int32
func (p *ComputationalResourceSchedulingModel) GetWallTimeLimit() int32 {
  if !p.IsSetWallTimeLimit() {
    return ComputationalResourceSchedulingModel_WallTimeLimit_DEFAULT
  }
return *p.WallTimeLimit
}
var ComputationalResourceSchedulingModel_TotalPhysicalMemory_DEFAULT int32
func (p *ComputationalResourceSchedulingModel) GetTotalPhysicalMemory() int32 {
  if !p.IsSetTotalPhysicalMemory() {
    return ComputationalResourceSchedulingModel_TotalPhysicalMemory_DEFAULT
  }
return *p.TotalPhysicalMemory
}
var ComputationalResourceSchedulingModel_ChessisNumber_DEFAULT string
func (p *ComputationalResourceSchedulingModel) GetChessisNumber() string {
  if !p.IsSetChessisNumber() {
    return ComputationalResourceSchedulingModel_ChessisNumber_DEFAULT
  }
return *p.ChessisNumber
}
var ComputationalResourceSchedulingModel_StaticWorkingDir_DEFAULT string
func (p *ComputationalResourceSchedulingModel) GetStaticWorkingDir() string {
  if !p.IsSetStaticWorkingDir() {
    return ComputationalResourceSchedulingModel_StaticWorkingDir_DEFAULT
  }
return *p.StaticWorkingDir
}
var ComputationalResourceSchedulingModel_OverrideLoginUserName_DEFAULT string
func (p *ComputationalResourceSchedulingModel) GetOverrideLoginUserName() string {
  if !p.IsSetOverrideLoginUserName() {
    return ComputationalResourceSchedulingModel_OverrideLoginUserName_DEFAULT
  }
return *p.OverrideLoginUserName
}
var ComputationalResourceSchedulingModel_OverrideScratchLocation_DEFAULT string
func (p *ComputationalResourceSchedulingModel) GetOverrideScratchLocation() string {
  if !p.IsSetOverrideScratchLocation() {
    return ComputationalResourceSchedulingModel_OverrideScratchLocation_DEFAULT
  }
return *p.OverrideScratchLocation
}
var ComputationalResourceSchedulingModel_OverrideAllocationProjectNumber_DEFAULT string
func (p *ComputationalResourceSchedulingModel) GetOverrideAllocationProjectNumber() string {
  if !p.IsSetOverrideAllocationProjectNumber() {
    return ComputationalResourceSchedulingModel_OverrideAllocationProjectNumber_DEFAULT
  }
return *p.OverrideAllocationProjectNumber
}
func (p *ComputationalResourceSchedulingModel) IsSetResourceHostId() bool {
  return p.ResourceHostId != nil
}

func (p *ComputationalResourceSchedulingModel) IsSetTotalCPUCount() bool {
  return p.TotalCPUCount != nil
}

func (p *ComputationalResourceSchedulingModel) IsSetNodeCount() bool {
  return p.NodeCount != nil
}

func (p *ComputationalResourceSchedulingModel) IsSetNumberOfThreads() bool {
  return p.NumberOfThreads != nil
}

func (p *ComputationalResourceSchedulingModel) IsSetQueueName() bool {
  return p.QueueName != nil
}

func (p *ComputationalResourceSchedulingModel) IsSetWallTimeLimit() bool {
  return p.WallTimeLimit != nil
}

func (p *ComputationalResourceSchedulingModel) IsSetTotalPhysicalMemory() bool {
  return p.TotalPhysicalMemory != nil
}

func (p *ComputationalResourceSchedulingModel) IsSetChessisNumber() bool {
  return p.ChessisNumber != nil
}

func (p *ComputationalResourceSchedulingModel) IsSetStaticWorkingDir() bool {
  return p.StaticWorkingDir != nil
}

func (p *ComputationalResourceSchedulingModel) IsSetOverrideLoginUserName() bool {
  return p.OverrideLoginUserName != nil
}

func (p *ComputationalResourceSchedulingModel) IsSetOverrideScratchLocation() bool {
  return p.OverrideScratchLocation != nil
}

func (p *ComputationalResourceSchedulingModel) IsSetOverrideAllocationProjectNumber() bool {
  return p.OverrideAllocationProjectNumber != nil
}

func (p *ComputationalResourceSchedulingModel) Read(iprot thrift.TProtocol) error {
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
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField4(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 5:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField5(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 6:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField6(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 7:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField7(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 8:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField8(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 9:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField9(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 10:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField10(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 11:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField11(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 12:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField12(iprot); err != nil {
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

func (p *ComputationalResourceSchedulingModel)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.ResourceHostId = &v
}
  return nil
}

func (p *ComputationalResourceSchedulingModel)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.TotalCPUCount = &v
}
  return nil
}

func (p *ComputationalResourceSchedulingModel)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.NodeCount = &v
}
  return nil
}

func (p *ComputationalResourceSchedulingModel)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.NumberOfThreads = &v
}
  return nil
}

func (p *ComputationalResourceSchedulingModel)  ReadField5(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 5: ", err)
} else {
  p.QueueName = &v
}
  return nil
}

func (p *ComputationalResourceSchedulingModel)  ReadField6(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 6: ", err)
} else {
  p.WallTimeLimit = &v
}
  return nil
}

func (p *ComputationalResourceSchedulingModel)  ReadField7(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 7: ", err)
} else {
  p.TotalPhysicalMemory = &v
}
  return nil
}

func (p *ComputationalResourceSchedulingModel)  ReadField8(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 8: ", err)
} else {
  p.ChessisNumber = &v
}
  return nil
}

func (p *ComputationalResourceSchedulingModel)  ReadField9(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 9: ", err)
} else {
  p.StaticWorkingDir = &v
}
  return nil
}

func (p *ComputationalResourceSchedulingModel)  ReadField10(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 10: ", err)
} else {
  p.OverrideLoginUserName = &v
}
  return nil
}

func (p *ComputationalResourceSchedulingModel)  ReadField11(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 11: ", err)
} else {
  p.OverrideScratchLocation = &v
}
  return nil
}

func (p *ComputationalResourceSchedulingModel)  ReadField12(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 12: ", err)
} else {
  p.OverrideAllocationProjectNumber = &v
}
  return nil
}

func (p *ComputationalResourceSchedulingModel) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("ComputationalResourceSchedulingModel"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
    if err := p.writeField5(oprot); err != nil { return err }
    if err := p.writeField6(oprot); err != nil { return err }
    if err := p.writeField7(oprot); err != nil { return err }
    if err := p.writeField8(oprot); err != nil { return err }
    if err := p.writeField9(oprot); err != nil { return err }
    if err := p.writeField10(oprot); err != nil { return err }
    if err := p.writeField11(oprot); err != nil { return err }
    if err := p.writeField12(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ComputationalResourceSchedulingModel) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetResourceHostId() {
    if err := oprot.WriteFieldBegin("resourceHostId", thrift.STRING, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:resourceHostId: ", p), err) }
    if err := oprot.WriteString(string(*p.ResourceHostId)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.resourceHostId (1) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:resourceHostId: ", p), err) }
  }
  return err
}

func (p *ComputationalResourceSchedulingModel) writeField2(oprot thrift.TProtocol) (err error) {
  if p.IsSetTotalCPUCount() {
    if err := oprot.WriteFieldBegin("totalCPUCount", thrift.I32, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:totalCPUCount: ", p), err) }
    if err := oprot.WriteI32(int32(*p.TotalCPUCount)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.totalCPUCount (2) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:totalCPUCount: ", p), err) }
  }
  return err
}

func (p *ComputationalResourceSchedulingModel) writeField3(oprot thrift.TProtocol) (err error) {
  if p.IsSetNodeCount() {
    if err := oprot.WriteFieldBegin("nodeCount", thrift.I32, 3); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:nodeCount: ", p), err) }
    if err := oprot.WriteI32(int32(*p.NodeCount)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.nodeCount (3) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 3:nodeCount: ", p), err) }
  }
  return err
}

func (p *ComputationalResourceSchedulingModel) writeField4(oprot thrift.TProtocol) (err error) {
  if p.IsSetNumberOfThreads() {
    if err := oprot.WriteFieldBegin("numberOfThreads", thrift.I32, 4); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:numberOfThreads: ", p), err) }
    if err := oprot.WriteI32(int32(*p.NumberOfThreads)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.numberOfThreads (4) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 4:numberOfThreads: ", p), err) }
  }
  return err
}

func (p *ComputationalResourceSchedulingModel) writeField5(oprot thrift.TProtocol) (err error) {
  if p.IsSetQueueName() {
    if err := oprot.WriteFieldBegin("queueName", thrift.STRING, 5); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:queueName: ", p), err) }
    if err := oprot.WriteString(string(*p.QueueName)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.queueName (5) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 5:queueName: ", p), err) }
  }
  return err
}

func (p *ComputationalResourceSchedulingModel) writeField6(oprot thrift.TProtocol) (err error) {
  if p.IsSetWallTimeLimit() {
    if err := oprot.WriteFieldBegin("wallTimeLimit", thrift.I32, 6); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:wallTimeLimit: ", p), err) }
    if err := oprot.WriteI32(int32(*p.WallTimeLimit)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.wallTimeLimit (6) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 6:wallTimeLimit: ", p), err) }
  }
  return err
}

func (p *ComputationalResourceSchedulingModel) writeField7(oprot thrift.TProtocol) (err error) {
  if p.IsSetTotalPhysicalMemory() {
    if err := oprot.WriteFieldBegin("totalPhysicalMemory", thrift.I32, 7); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:totalPhysicalMemory: ", p), err) }
    if err := oprot.WriteI32(int32(*p.TotalPhysicalMemory)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.totalPhysicalMemory (7) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 7:totalPhysicalMemory: ", p), err) }
  }
  return err
}

func (p *ComputationalResourceSchedulingModel) writeField8(oprot thrift.TProtocol) (err error) {
  if p.IsSetChessisNumber() {
    if err := oprot.WriteFieldBegin("chessisNumber", thrift.STRING, 8); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:chessisNumber: ", p), err) }
    if err := oprot.WriteString(string(*p.ChessisNumber)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.chessisNumber (8) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 8:chessisNumber: ", p), err) }
  }
  return err
}

func (p *ComputationalResourceSchedulingModel) writeField9(oprot thrift.TProtocol) (err error) {
  if p.IsSetStaticWorkingDir() {
    if err := oprot.WriteFieldBegin("staticWorkingDir", thrift.STRING, 9); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:staticWorkingDir: ", p), err) }
    if err := oprot.WriteString(string(*p.StaticWorkingDir)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.staticWorkingDir (9) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 9:staticWorkingDir: ", p), err) }
  }
  return err
}

func (p *ComputationalResourceSchedulingModel) writeField10(oprot thrift.TProtocol) (err error) {
  if p.IsSetOverrideLoginUserName() {
    if err := oprot.WriteFieldBegin("overrideLoginUserName", thrift.STRING, 10); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:overrideLoginUserName: ", p), err) }
    if err := oprot.WriteString(string(*p.OverrideLoginUserName)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.overrideLoginUserName (10) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 10:overrideLoginUserName: ", p), err) }
  }
  return err
}

func (p *ComputationalResourceSchedulingModel) writeField11(oprot thrift.TProtocol) (err error) {
  if p.IsSetOverrideScratchLocation() {
    if err := oprot.WriteFieldBegin("overrideScratchLocation", thrift.STRING, 11); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 11:overrideScratchLocation: ", p), err) }
    if err := oprot.WriteString(string(*p.OverrideScratchLocation)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.overrideScratchLocation (11) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 11:overrideScratchLocation: ", p), err) }
  }
  return err
}

func (p *ComputationalResourceSchedulingModel) writeField12(oprot thrift.TProtocol) (err error) {
  if p.IsSetOverrideAllocationProjectNumber() {
    if err := oprot.WriteFieldBegin("overrideAllocationProjectNumber", thrift.STRING, 12); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 12:overrideAllocationProjectNumber: ", p), err) }
    if err := oprot.WriteString(string(*p.OverrideAllocationProjectNumber)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.overrideAllocationProjectNumber (12) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 12:overrideAllocationProjectNumber: ", p), err) }
  }
  return err
}

func (p *ComputationalResourceSchedulingModel) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ComputationalResourceSchedulingModel(%+v)", *p)
}


/* Autogenerated by Thrift Compiler (0.8.0-dev)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 */
package cdh3u3;

import (
        "github.com/araddon/thrift4go/lib/go/thrift"
        "fmt"
)

//import "thriftlib/fb303"
//import "thriftlib/hive_metastore"
//import "thriftlib/queryplan"



type JobTrackerState int
const (
  INITIALIZING JobTrackerState = 1
  RUNNING JobTrackerState = 2
)
func (p JobTrackerState) String() string {
  switch p {
  case INITIALIZING: return "INITIALIZING"
  case RUNNING: return "RUNNING"
  }
  return ""
}

func FromJobTrackerStateString(s string) JobTrackerState {
  switch s {
  case "INITIALIZING": return INITIALIZING
  case "RUNNING": return RUNNING
  }
  return JobTrackerState(-10000)
}

func (p JobTrackerState) Value() int {
  return int(p)
}

func (p JobTrackerState) IsEnum() bool {
  return true
}

/**
 * Attributes:
 *  - TaskTrackers
 *  - MapTasks
 *  - ReduceTasks
 *  - MaxMapTasks
 *  - MaxReduceTasks
 *  - State
 */
type HiveClusterStatus struct {
  thrift.TStruct
  TaskTrackers int32 "taskTrackers"; // 1
  MapTasks int32 "mapTasks"; // 2
  ReduceTasks int32 "reduceTasks"; // 3
  MaxMapTasks int32 "maxMapTasks"; // 4
  MaxReduceTasks int32 "maxReduceTasks"; // 5
  State JobTrackerState "state"; // 6
}

func NewHiveClusterStatus() *HiveClusterStatus {
  output := &HiveClusterStatus{
    TStruct:thrift.NewTStruct("HiveClusterStatus", []thrift.TField{
    thrift.NewTField("taskTrackers", thrift.I32, 1),
    thrift.NewTField("mapTasks", thrift.I32, 2),
    thrift.NewTField("reduceTasks", thrift.I32, 3),
    thrift.NewTField("maxMapTasks", thrift.I32, 4),
    thrift.NewTField("maxReduceTasks", thrift.I32, 5),
    thrift.NewTField("state", thrift.I32, 6),
    }),
  }
  {
  }
  return output
}

func (p *HiveClusterStatus) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
  _, err = iprot.ReadStructBegin()
  if err != nil { return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err); }
  for {
    fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if fieldId < 0 {
      fieldId = int16(p.FieldIdFromFieldName(fieldName))
    } else if fieldName == "" {
      fieldName = p.FieldNameFromFieldId(int(fieldId))
    }
    if fieldTypeId == thrift.GENERIC {
      fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
    }
    if err != nil {
      return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if fieldId == 1 || fieldName == "taskTrackers" {
      if fieldTypeId == thrift.I32 {
        err = p.ReadField1(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else if fieldTypeId == thrift.VOID {
        err = iprot.Skip(fieldTypeId)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else {
        err = p.ReadField1(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      }
    } else if fieldId == 2 || fieldName == "mapTasks" {
      if fieldTypeId == thrift.I32 {
        err = p.ReadField2(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else if fieldTypeId == thrift.VOID {
        err = iprot.Skip(fieldTypeId)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else {
        err = p.ReadField2(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      }
    } else if fieldId == 3 || fieldName == "reduceTasks" {
      if fieldTypeId == thrift.I32 {
        err = p.ReadField3(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else if fieldTypeId == thrift.VOID {
        err = iprot.Skip(fieldTypeId)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else {
        err = p.ReadField3(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      }
    } else if fieldId == 4 || fieldName == "maxMapTasks" {
      if fieldTypeId == thrift.I32 {
        err = p.ReadField4(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else if fieldTypeId == thrift.VOID {
        err = iprot.Skip(fieldTypeId)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else {
        err = p.ReadField4(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      }
    } else if fieldId == 5 || fieldName == "maxReduceTasks" {
      if fieldTypeId == thrift.I32 {
        err = p.ReadField5(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else if fieldTypeId == thrift.VOID {
        err = iprot.Skip(fieldTypeId)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else {
        err = p.ReadField5(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      }
    } else if fieldId == 6 || fieldName == "state" {
      if fieldTypeId == thrift.I32 {
        err = p.ReadField6(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else if fieldTypeId == thrift.VOID {
        err = iprot.Skip(fieldTypeId)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else {
        err = p.ReadField6(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      }
    } else {
      err = iprot.Skip(fieldTypeId)
      if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
    }
    err = iprot.ReadFieldEnd()
    if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
  }
  err = iprot.ReadStructEnd()
  if err != nil { return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err); }
  return err
}

func (p *HiveClusterStatus) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
  v0, err1 := iprot.ReadI32()
  if err1 != nil { return thrift.NewTProtocolExceptionReadField(1, "taskTrackers", p.ThriftName(), err1); }
  p.TaskTrackers = v0
  return err
}

func (p *HiveClusterStatus) ReadFieldTaskTrackers(iprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.ReadField1(iprot)
}

func (p *HiveClusterStatus) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
  v2, err3 := iprot.ReadI32()
  if err3 != nil { return thrift.NewTProtocolExceptionReadField(2, "mapTasks", p.ThriftName(), err3); }
  p.MapTasks = v2
  return err
}

func (p *HiveClusterStatus) ReadFieldMapTasks(iprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.ReadField2(iprot)
}

func (p *HiveClusterStatus) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
  v4, err5 := iprot.ReadI32()
  if err5 != nil { return thrift.NewTProtocolExceptionReadField(3, "reduceTasks", p.ThriftName(), err5); }
  p.ReduceTasks = v4
  return err
}

func (p *HiveClusterStatus) ReadFieldReduceTasks(iprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.ReadField3(iprot)
}

func (p *HiveClusterStatus) ReadField4(iprot thrift.TProtocol) (err thrift.TProtocolException) {
  v6, err7 := iprot.ReadI32()
  if err7 != nil { return thrift.NewTProtocolExceptionReadField(4, "maxMapTasks", p.ThriftName(), err7); }
  p.MaxMapTasks = v6
  return err
}

func (p *HiveClusterStatus) ReadFieldMaxMapTasks(iprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.ReadField4(iprot)
}

func (p *HiveClusterStatus) ReadField5(iprot thrift.TProtocol) (err thrift.TProtocolException) {
  v8, err9 := iprot.ReadI32()
  if err9 != nil { return thrift.NewTProtocolExceptionReadField(5, "maxReduceTasks", p.ThriftName(), err9); }
  p.MaxReduceTasks = v8
  return err
}

func (p *HiveClusterStatus) ReadFieldMaxReduceTasks(iprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.ReadField5(iprot)
}

func (p *HiveClusterStatus) ReadField6(iprot thrift.TProtocol) (err thrift.TProtocolException) {
  v10, err11 := iprot.ReadI32()
  if err11 != nil { return thrift.NewTProtocolExceptionReadField(6, "state", p.ThriftName(), err11); }
  p.State = JobTrackerState(v10)
  return err
}

func (p *HiveClusterStatus) ReadFieldState(iprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.ReadField6(iprot)
}

func (p *HiveClusterStatus) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
  err = oprot.WriteStructBegin("HiveClusterStatus")
  if err != nil { return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err); }
  err = p.WriteField1(oprot)
  if err != nil { return err }
  err = p.WriteField2(oprot)
  if err != nil { return err }
  err = p.WriteField3(oprot)
  if err != nil { return err }
  err = p.WriteField4(oprot)
  if err != nil { return err }
  err = p.WriteField5(oprot)
  if err != nil { return err }
  err = p.WriteField6(oprot)
  if err != nil { return err }
  err = oprot.WriteFieldStop()
  if err != nil { return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err); }
  err = oprot.WriteStructEnd()
  if err != nil { return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err); }
  return err
}

func (p *HiveClusterStatus) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
  if int64(p.TaskTrackers) == 0 { return nil}
  err = oprot.WriteFieldBegin("taskTrackers", thrift.I32, 1)
  if err != nil { return thrift.NewTProtocolExceptionWriteField(1, "taskTrackers", p.ThriftName(), err); }
  err = oprot.WriteI32(int32(p.TaskTrackers))
  if err != nil { return thrift.NewTProtocolExceptionWriteField(1, "taskTrackers", p.ThriftName(), err); }
  err = oprot.WriteFieldEnd()
  if err != nil { return thrift.NewTProtocolExceptionWriteField(1, "taskTrackers", p.ThriftName(), err); }
  return err
}

func (p *HiveClusterStatus) WriteFieldTaskTrackers(oprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.WriteField1(oprot)
}

func (p *HiveClusterStatus) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
  if int64(p.MapTasks) == 0 { return nil}
  err = oprot.WriteFieldBegin("mapTasks", thrift.I32, 2)
  if err != nil { return thrift.NewTProtocolExceptionWriteField(2, "mapTasks", p.ThriftName(), err); }
  err = oprot.WriteI32(int32(p.MapTasks))
  if err != nil { return thrift.NewTProtocolExceptionWriteField(2, "mapTasks", p.ThriftName(), err); }
  err = oprot.WriteFieldEnd()
  if err != nil { return thrift.NewTProtocolExceptionWriteField(2, "mapTasks", p.ThriftName(), err); }
  return err
}

func (p *HiveClusterStatus) WriteFieldMapTasks(oprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.WriteField2(oprot)
}

func (p *HiveClusterStatus) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
  if int64(p.ReduceTasks) == 0 { return nil}
  err = oprot.WriteFieldBegin("reduceTasks", thrift.I32, 3)
  if err != nil { return thrift.NewTProtocolExceptionWriteField(3, "reduceTasks", p.ThriftName(), err); }
  err = oprot.WriteI32(int32(p.ReduceTasks))
  if err != nil { return thrift.NewTProtocolExceptionWriteField(3, "reduceTasks", p.ThriftName(), err); }
  err = oprot.WriteFieldEnd()
  if err != nil { return thrift.NewTProtocolExceptionWriteField(3, "reduceTasks", p.ThriftName(), err); }
  return err
}

func (p *HiveClusterStatus) WriteFieldReduceTasks(oprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.WriteField3(oprot)
}

func (p *HiveClusterStatus) WriteField4(oprot thrift.TProtocol) (err thrift.TProtocolException) {
  if int64(p.MaxMapTasks) == 0 { return nil}
  err = oprot.WriteFieldBegin("maxMapTasks", thrift.I32, 4)
  if err != nil { return thrift.NewTProtocolExceptionWriteField(4, "maxMapTasks", p.ThriftName(), err); }
  err = oprot.WriteI32(int32(p.MaxMapTasks))
  if err != nil { return thrift.NewTProtocolExceptionWriteField(4, "maxMapTasks", p.ThriftName(), err); }
  err = oprot.WriteFieldEnd()
  if err != nil { return thrift.NewTProtocolExceptionWriteField(4, "maxMapTasks", p.ThriftName(), err); }
  return err
}

func (p *HiveClusterStatus) WriteFieldMaxMapTasks(oprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.WriteField4(oprot)
}

func (p *HiveClusterStatus) WriteField5(oprot thrift.TProtocol) (err thrift.TProtocolException) {
  if int64(p.MaxReduceTasks) == 0 { return nil}
  err = oprot.WriteFieldBegin("maxReduceTasks", thrift.I32, 5)
  if err != nil { return thrift.NewTProtocolExceptionWriteField(5, "maxReduceTasks", p.ThriftName(), err); }
  err = oprot.WriteI32(int32(p.MaxReduceTasks))
  if err != nil { return thrift.NewTProtocolExceptionWriteField(5, "maxReduceTasks", p.ThriftName(), err); }
  err = oprot.WriteFieldEnd()
  if err != nil { return thrift.NewTProtocolExceptionWriteField(5, "maxReduceTasks", p.ThriftName(), err); }
  return err
}

func (p *HiveClusterStatus) WriteFieldMaxReduceTasks(oprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.WriteField5(oprot)
}

func (p *HiveClusterStatus) WriteField6(oprot thrift.TProtocol) (err thrift.TProtocolException) {
  err = oprot.WriteFieldBegin("state", thrift.I32, 6)
  if err != nil { return thrift.NewTProtocolExceptionWriteField(6, "state", p.ThriftName(), err); }
  err = oprot.WriteI32(int32(p.State))
  if err != nil { return thrift.NewTProtocolExceptionWriteField(6, "state", p.ThriftName(), err); }
  err = oprot.WriteFieldEnd()
  if err != nil { return thrift.NewTProtocolExceptionWriteField(6, "state", p.ThriftName(), err); }
  return err
}

func (p *HiveClusterStatus) WriteFieldState(oprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.WriteField6(oprot)
}

func (p *HiveClusterStatus) TStructName() string {
  return "HiveClusterStatus"
}

func (p *HiveClusterStatus) ThriftName() string {
  return "HiveClusterStatus"
}

func (p *HiveClusterStatus) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("HiveClusterStatus(%+v)", *p)
}

func (p *HiveClusterStatus) CompareTo(other interface{}) (int, bool) {
  if other == nil {
    return 1, true
  }
  data, ok := other.(*HiveClusterStatus)
  if !ok {
    return 0, false
  }
  if p.TaskTrackers != data.TaskTrackers {
    if p.TaskTrackers < data.TaskTrackers {
      return -1, true
    }
    return 1, true
  }
  if p.MapTasks != data.MapTasks {
    if p.MapTasks < data.MapTasks {
      return -1, true
    }
    return 1, true
  }
  if p.ReduceTasks != data.ReduceTasks {
    if p.ReduceTasks < data.ReduceTasks {
      return -1, true
    }
    return 1, true
  }
  if p.MaxMapTasks != data.MaxMapTasks {
    if p.MaxMapTasks < data.MaxMapTasks {
      return -1, true
    }
    return 1, true
  }
  if p.MaxReduceTasks != data.MaxReduceTasks {
    if p.MaxReduceTasks < data.MaxReduceTasks {
      return -1, true
    }
    return 1, true
  }
  if p.State != data.State {
    if p.State < data.State {
      return -1, true
    }
    return 1, true
  }
  return 0, true
}

func (p *HiveClusterStatus) AttributeByFieldId(id int) interface{} {
  switch id {
  default: return nil
  case 1: return p.TaskTrackers
  case 2: return p.MapTasks
  case 3: return p.ReduceTasks
  case 4: return p.MaxMapTasks
  case 5: return p.MaxReduceTasks
  case 6: return p.State
  }
  return nil
}

func (p *HiveClusterStatus) TStructFields() thrift.TFieldContainer {
  return thrift.NewTFieldContainer([]thrift.TField{
    thrift.NewTField("taskTrackers", thrift.I32, 1),
    thrift.NewTField("mapTasks", thrift.I32, 2),
    thrift.NewTField("reduceTasks", thrift.I32, 3),
    thrift.NewTField("maxMapTasks", thrift.I32, 4),
    thrift.NewTField("maxReduceTasks", thrift.I32, 5),
    thrift.NewTField("state", thrift.I32, 6),
    })
}

/**
 * Attributes:
 *  - Message
 *  - ErrorCode
 *  - SQLState
 */
type HiveServerException struct {
  thrift.TStruct
  Message string "message"; // 1
  ErrorCode int32 "errorCode"; // 2
  SQLState string "SQLState"; // 3
}

func NewHiveServerException() *HiveServerException {
  output := &HiveServerException{
    TStruct:thrift.NewTStruct("HiveServerException", []thrift.TField{
    thrift.NewTField("message", thrift.STRING, 1),
    thrift.NewTField("errorCode", thrift.I32, 2),
    thrift.NewTField("SQLState", thrift.STRING, 3),
    }),
  }
  {
  }
  return output
}

func (p *HiveServerException) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
  _, err = iprot.ReadStructBegin()
  if err != nil { return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err); }
  for {
    fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if fieldId < 0 {
      fieldId = int16(p.FieldIdFromFieldName(fieldName))
    } else if fieldName == "" {
      fieldName = p.FieldNameFromFieldId(int(fieldId))
    }
    if fieldTypeId == thrift.GENERIC {
      fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
    }
    if err != nil {
      return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if fieldId == 1 || fieldName == "message" {
      if fieldTypeId == thrift.STRING {
        err = p.ReadField1(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else if fieldTypeId == thrift.VOID {
        err = iprot.Skip(fieldTypeId)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else {
        err = p.ReadField1(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      }
    } else if fieldId == 2 || fieldName == "errorCode" {
      if fieldTypeId == thrift.I32 {
        err = p.ReadField2(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else if fieldTypeId == thrift.VOID {
        err = iprot.Skip(fieldTypeId)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else {
        err = p.ReadField2(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      }
    } else if fieldId == 3 || fieldName == "SQLState" {
      if fieldTypeId == thrift.STRING {
        err = p.ReadField3(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else if fieldTypeId == thrift.VOID {
        err = iprot.Skip(fieldTypeId)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      } else {
        err = p.ReadField3(iprot)
        if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
      }
    } else {
      err = iprot.Skip(fieldTypeId)
      if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
    }
    err = iprot.ReadFieldEnd()
    if err != nil { return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err); }
  }
  err = iprot.ReadStructEnd()
  if err != nil { return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err); }
  return err
}

func (p *HiveServerException) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
  v12, err13 := iprot.ReadString()
  if err13 != nil { return thrift.NewTProtocolExceptionReadField(1, "message", p.ThriftName(), err13); }
  p.Message = v12
  return err
}

func (p *HiveServerException) ReadFieldMessage(iprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.ReadField1(iprot)
}

func (p *HiveServerException) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
  v14, err15 := iprot.ReadI32()
  if err15 != nil { return thrift.NewTProtocolExceptionReadField(2, "errorCode", p.ThriftName(), err15); }
  p.ErrorCode = v14
  return err
}

func (p *HiveServerException) ReadFieldErrorCode(iprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.ReadField2(iprot)
}

func (p *HiveServerException) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
  v16, err17 := iprot.ReadString()
  if err17 != nil { return thrift.NewTProtocolExceptionReadField(3, "SQLState", p.ThriftName(), err17); }
  p.SQLState = v16
  return err
}

func (p *HiveServerException) ReadFieldSQLState(iprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.ReadField3(iprot)
}

func (p *HiveServerException) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
  err = oprot.WriteStructBegin("HiveServerException")
  if err != nil { return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err); }
  err = p.WriteField1(oprot)
  if err != nil { return err }
  err = p.WriteField2(oprot)
  if err != nil { return err }
  err = p.WriteField3(oprot)
  if err != nil { return err }
  err = oprot.WriteFieldStop()
  if err != nil { return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err); }
  err = oprot.WriteStructEnd()
  if err != nil { return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err); }
  return err
}

func (p *HiveServerException) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
  if len(p.Message) < 1 { return nil}
  err = oprot.WriteFieldBegin("message", thrift.STRING, 1)
  if err != nil { return thrift.NewTProtocolExceptionWriteField(1, "message", p.ThriftName(), err); }
  err = oprot.WriteString(string(p.Message))
  if err != nil { return thrift.NewTProtocolExceptionWriteField(1, "message", p.ThriftName(), err); }
  err = oprot.WriteFieldEnd()
  if err != nil { return thrift.NewTProtocolExceptionWriteField(1, "message", p.ThriftName(), err); }
  return err
}

func (p *HiveServerException) WriteFieldMessage(oprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.WriteField1(oprot)
}

func (p *HiveServerException) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
  if int64(p.ErrorCode) == 0 { return nil}
  err = oprot.WriteFieldBegin("errorCode", thrift.I32, 2)
  if err != nil { return thrift.NewTProtocolExceptionWriteField(2, "errorCode", p.ThriftName(), err); }
  err = oprot.WriteI32(int32(p.ErrorCode))
  if err != nil { return thrift.NewTProtocolExceptionWriteField(2, "errorCode", p.ThriftName(), err); }
  err = oprot.WriteFieldEnd()
  if err != nil { return thrift.NewTProtocolExceptionWriteField(2, "errorCode", p.ThriftName(), err); }
  return err
}

func (p *HiveServerException) WriteFieldErrorCode(oprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.WriteField2(oprot)
}

func (p *HiveServerException) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
  if len(p.SQLState) < 1 { return nil}
  err = oprot.WriteFieldBegin("SQLState", thrift.STRING, 3)
  if err != nil { return thrift.NewTProtocolExceptionWriteField(3, "SQLState", p.ThriftName(), err); }
  err = oprot.WriteString(string(p.SQLState))
  if err != nil { return thrift.NewTProtocolExceptionWriteField(3, "SQLState", p.ThriftName(), err); }
  err = oprot.WriteFieldEnd()
  if err != nil { return thrift.NewTProtocolExceptionWriteField(3, "SQLState", p.ThriftName(), err); }
  return err
}

func (p *HiveServerException) WriteFieldSQLState(oprot thrift.TProtocol) (thrift.TProtocolException) {
  return p.WriteField3(oprot)
}

func (p *HiveServerException) TStructName() string {
  return "HiveServerException"
}

func (p *HiveServerException) ThriftName() string {
  return "HiveServerException"
}

func (p *HiveServerException) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("HiveServerException(%+v)", *p)
}

func (p *HiveServerException) CompareTo(other interface{}) (int, bool) {
  if other == nil {
    return 1, true
  }
  data, ok := other.(*HiveServerException)
  if !ok {
    return 0, false
  }
  if p.Message != data.Message {
    if p.Message < data.Message {
      return -1, true
    }
    return 1, true
  }
  if p.ErrorCode != data.ErrorCode {
    if p.ErrorCode < data.ErrorCode {
      return -1, true
    }
    return 1, true
  }
  if p.SQLState != data.SQLState {
    if p.SQLState < data.SQLState {
      return -1, true
    }
    return 1, true
  }
  return 0, true
}

func (p *HiveServerException) AttributeByFieldId(id int) interface{} {
  switch id {
  default: return nil
  case 1: return p.Message
  case 2: return p.ErrorCode
  case 3: return p.SQLState
  }
  return nil
}

func (p *HiveServerException) TStructFields() thrift.TFieldContainer {
  return thrift.NewTFieldContainer([]thrift.TField{
    thrift.NewTField("message", thrift.STRING, 1),
    thrift.NewTField("errorCode", thrift.I32, 2),
    thrift.NewTField("SQLState", thrift.STRING, 3),
    })
}

func init() {
}


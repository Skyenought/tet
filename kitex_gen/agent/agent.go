// Code generated by thriftgo (0.3.4). DO NOT EDIT.

package agent

import (
	"fmt"
	"github.com/Skyenought/tet/kitex_gen/model"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type UpdateTasksReq struct {
	Tasks []*model.Task `thrift:"tasks,1" frugal:"1,default,list<model.Task>" json:"tasks"`
}

func NewUpdateTasksReq() *UpdateTasksReq {
	return &UpdateTasksReq{}
}

func (p *UpdateTasksReq) InitDefault() {
	*p = UpdateTasksReq{}
}

func (p *UpdateTasksReq) GetTasks() (v []*model.Task) {
	return p.Tasks
}
func (p *UpdateTasksReq) SetTasks(val []*model.Task) {
	p.Tasks = val
}

var fieldIDToName_UpdateTasksReq = map[int16]string{
	1: "tasks",
}

func (p *UpdateTasksReq) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UpdateTasksReq[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UpdateTasksReq) ReadField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.Tasks = make([]*model.Task, 0, size)
	for i := 0; i < size; i++ {
		_elem := model.NewTask()
		if err := _elem.Read(iprot); err != nil {
			return err
		}

		p.Tasks = append(p.Tasks, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *UpdateTasksReq) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UpdateTasksReq"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UpdateTasksReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("tasks", thrift.LIST, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Tasks)); err != nil {
		return err
	}
	for _, v := range p.Tasks {
		if err := v.Write(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UpdateTasksReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpdateTasksReq(%+v)", *p)

}

func (p *UpdateTasksReq) DeepEqual(ano *UpdateTasksReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Tasks) {
		return false
	}
	return true
}

func (p *UpdateTasksReq) Field1DeepEqual(src []*model.Task) bool {

	if len(p.Tasks) != len(src) {
		return false
	}
	for i, v := range p.Tasks {
		_src := src[i]
		if !v.DeepEqual(_src) {
			return false
		}
	}
	return true
}

type UpdateTasksRes struct {
	Code int32  `thrift:"code,1" frugal:"1,default,i32" json:"code"`
	Msg  string `thrift:"msg,2" frugal:"2,default,string" json:"msg"`
}

func NewUpdateTasksRes() *UpdateTasksRes {
	return &UpdateTasksRes{}
}

func (p *UpdateTasksRes) InitDefault() {
	*p = UpdateTasksRes{}
}

func (p *UpdateTasksRes) GetCode() (v int32) {
	return p.Code
}

func (p *UpdateTasksRes) GetMsg() (v string) {
	return p.Msg
}
func (p *UpdateTasksRes) SetCode(val int32) {
	p.Code = val
}
func (p *UpdateTasksRes) SetMsg(val string) {
	p.Msg = val
}

var fieldIDToName_UpdateTasksRes = map[int16]string{
	1: "code",
	2: "msg",
}

func (p *UpdateTasksRes) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UpdateTasksRes[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UpdateTasksRes) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Code = v
	}
	return nil
}
func (p *UpdateTasksRes) ReadField2(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Msg = v
	}
	return nil
}

func (p *UpdateTasksRes) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UpdateTasksRes"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UpdateTasksRes) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("code", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Code); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UpdateTasksRes) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("msg", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Msg); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *UpdateTasksRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpdateTasksRes(%+v)", *p)

}

func (p *UpdateTasksRes) DeepEqual(ano *UpdateTasksRes) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Code) {
		return false
	}
	if !p.Field2DeepEqual(ano.Msg) {
		return false
	}
	return true
}

func (p *UpdateTasksRes) Field1DeepEqual(src int32) bool {

	if p.Code != src {
		return false
	}
	return true
}
func (p *UpdateTasksRes) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Msg, src) != 0 {
		return false
	}
	return true
}

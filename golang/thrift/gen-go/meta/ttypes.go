// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package meta

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// Attributes:
//  - Ok
//  - Message
type HealthStatus struct {
	Ok      bool    `thrift:"ok,1,required" json:"ok"`
	Message *string `thrift:"message,2" json:"message,omitempty"`
}

func NewHealthStatus() *HealthStatus {
	return &HealthStatus{}
}

func (p *HealthStatus) GetOk() bool {
	return p.Ok
}

var HealthStatus_Message_DEFAULT string

func (p *HealthStatus) GetMessage() string {
	if !p.IsSetMessage() {
		return HealthStatus_Message_DEFAULT
	}
	return *p.Message
}
func (p *HealthStatus) IsSetMessage() bool {
	return p.Message != nil
}

func (p *HealthStatus) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetOk bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetOk = true
		case 2:
			if err := p.readField2(iprot); err != nil {
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
	if !issetOk {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Ok is not set"))
	}
	return nil
}

func (p *HealthStatus) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Ok = v
	}
	return nil
}

func (p *HealthStatus) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Message = &v
	}
	return nil
}

func (p *HealthStatus) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("HealthStatus"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *HealthStatus) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ok", thrift.BOOL, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ok: ", p), err)
	}
	if err := oprot.WriteBool(bool(p.Ok)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ok (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ok: ", p), err)
	}
	return err
}

func (p *HealthStatus) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetMessage() {
		if err := oprot.WriteFieldBegin("message", thrift.STRING, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:message: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Message)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.message (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:message: ", p), err)
		}
	}
	return err
}

func (p *HealthStatus) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HealthStatus(%+v)", *p)
}

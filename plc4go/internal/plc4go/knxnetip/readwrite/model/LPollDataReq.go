//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type LPollDataReq struct {
	Parent *CEMI
}

// The corresponding interface
type ILPollDataReq interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *LPollDataReq) MessageCode() uint8 {
	return 0x13
}

func (m *LPollDataReq) InitializeParent(parent *CEMI) {
}

func NewLPollDataReq() *CEMI {
	child := &LPollDataReq{
		Parent: NewCEMI(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastLPollDataReq(structType interface{}) *LPollDataReq {
	castFunc := func(typ interface{}) *LPollDataReq {
		if casted, ok := typ.(LPollDataReq); ok {
			return &casted
		}
		if casted, ok := typ.(*LPollDataReq); ok {
			return casted
		}
		if casted, ok := typ.(CEMI); ok {
			return CastLPollDataReq(casted.Child)
		}
		if casted, ok := typ.(*CEMI); ok {
			return CastLPollDataReq(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *LPollDataReq) GetTypeName() string {
	return "LPollDataReq"
}

func (m *LPollDataReq) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *LPollDataReq) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *LPollDataReq) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func LPollDataReqParse(readBuffer utils.ReadBuffer) (*CEMI, error) {
	if pullErr := readBuffer.PullContext("LPollDataReq"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("LPollDataReq"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &LPollDataReq{
		Parent: &CEMI{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *LPollDataReq) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LPollDataReq"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("LPollDataReq"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *LPollDataReq) String() string {
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}

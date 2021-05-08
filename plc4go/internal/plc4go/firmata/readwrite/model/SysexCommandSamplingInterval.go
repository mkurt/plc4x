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
type SysexCommandSamplingInterval struct {
	Parent *SysexCommand
}

// The corresponding interface
type ISysexCommandSamplingInterval interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *SysexCommandSamplingInterval) CommandType() uint8 {
	return 0x7A
}

func (m *SysexCommandSamplingInterval) Response() bool {
	return false
}

func (m *SysexCommandSamplingInterval) InitializeParent(parent *SysexCommand) {
}

func NewSysexCommandSamplingInterval() *SysexCommand {
	child := &SysexCommandSamplingInterval{
		Parent: NewSysexCommand(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastSysexCommandSamplingInterval(structType interface{}) *SysexCommandSamplingInterval {
	castFunc := func(typ interface{}) *SysexCommandSamplingInterval {
		if casted, ok := typ.(SysexCommandSamplingInterval); ok {
			return &casted
		}
		if casted, ok := typ.(*SysexCommandSamplingInterval); ok {
			return casted
		}
		if casted, ok := typ.(SysexCommand); ok {
			return CastSysexCommandSamplingInterval(casted.Child)
		}
		if casted, ok := typ.(*SysexCommand); ok {
			return CastSysexCommandSamplingInterval(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *SysexCommandSamplingInterval) GetTypeName() string {
	return "SysexCommandSamplingInterval"
}

func (m *SysexCommandSamplingInterval) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *SysexCommandSamplingInterval) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *SysexCommandSamplingInterval) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func SysexCommandSamplingIntervalParse(readBuffer utils.ReadBuffer) (*SysexCommand, error) {
	if pullErr := readBuffer.PullContext("SysexCommandSamplingInterval"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("SysexCommandSamplingInterval"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SysexCommandSamplingInterval{
		Parent: &SysexCommand{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *SysexCommandSamplingInterval) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandSamplingInterval"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("SysexCommandSamplingInterval"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *SysexCommandSamplingInterval) String() string {
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}

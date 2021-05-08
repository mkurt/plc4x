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
	"github.com/pkg/errors"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type DF1RequestCommand struct {
	Child IDF1RequestCommandChild
}

// The corresponding interface
type IDF1RequestCommand interface {
	FunctionCode() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IDF1RequestCommandParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IDF1RequestCommand, serializeChildFunction func() error) error
	GetTypeName() string
}

type IDF1RequestCommandChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *DF1RequestCommand)
	GetTypeName() string
	IDF1RequestCommand
}

func NewDF1RequestCommand() *DF1RequestCommand {
	return &DF1RequestCommand{}
}

func CastDF1RequestCommand(structType interface{}) *DF1RequestCommand {
	castFunc := func(typ interface{}) *DF1RequestCommand {
		if casted, ok := typ.(DF1RequestCommand); ok {
			return &casted
		}
		if casted, ok := typ.(*DF1RequestCommand); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DF1RequestCommand) GetTypeName() string {
	return "DF1RequestCommand"
}

func (m *DF1RequestCommand) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *DF1RequestCommand) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *DF1RequestCommand) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (functionCode)
	lengthInBits += 8

	return lengthInBits
}

func (m *DF1RequestCommand) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DF1RequestCommandParse(readBuffer utils.ReadBuffer) (*DF1RequestCommand, error) {
	if pullErr := readBuffer.PullContext("DF1RequestCommand"); pullErr != nil {
		return nil, pullErr
	}

	// Discriminator Field (functionCode) (Used as input to a switch field)
	functionCode, _functionCodeErr := readBuffer.ReadUint8("functionCode", 8)
	if _functionCodeErr != nil {
		return nil, errors.Wrap(_functionCodeErr, "Error parsing 'functionCode' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *DF1RequestCommand
	var typeSwitchError error
	switch {
	case functionCode == 0xA2: // DF1RequestProtectedTypedLogicalRead
		_parent, typeSwitchError = DF1RequestProtectedTypedLogicalReadParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("DF1RequestCommand"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *DF1RequestCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *DF1RequestCommand) SerializeParent(writeBuffer utils.WriteBuffer, child IDF1RequestCommand, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("DF1RequestCommand"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (functionCode) (Used as input to a switch field)
	functionCode := uint8(child.FunctionCode())
	_functionCodeErr := writeBuffer.WriteUint8("functionCode", 8, (functionCode))

	if _functionCodeErr != nil {
		return errors.Wrap(_functionCodeErr, "Error serializing 'functionCode' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("DF1RequestCommand"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *DF1RequestCommand) String() string {
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}

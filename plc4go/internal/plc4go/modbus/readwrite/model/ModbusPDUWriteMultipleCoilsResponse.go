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
type ModbusPDUWriteMultipleCoilsResponse struct {
	StartingAddress uint16
	Quantity        uint16
	Parent          *ModbusPDU
}

// The corresponding interface
type IModbusPDUWriteMultipleCoilsResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUWriteMultipleCoilsResponse) ErrorFlag() bool {
	return false
}

func (m *ModbusPDUWriteMultipleCoilsResponse) FunctionFlag() uint8 {
	return 0x0F
}

func (m *ModbusPDUWriteMultipleCoilsResponse) Response() bool {
	return true
}

func (m *ModbusPDUWriteMultipleCoilsResponse) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUWriteMultipleCoilsResponse(startingAddress uint16, quantity uint16) *ModbusPDU {
	child := &ModbusPDUWriteMultipleCoilsResponse{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		Parent:          NewModbusPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastModbusPDUWriteMultipleCoilsResponse(structType interface{}) *ModbusPDUWriteMultipleCoilsResponse {
	castFunc := func(typ interface{}) *ModbusPDUWriteMultipleCoilsResponse {
		if casted, ok := typ.(ModbusPDUWriteMultipleCoilsResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUWriteMultipleCoilsResponse); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUWriteMultipleCoilsResponse(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUWriteMultipleCoilsResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUWriteMultipleCoilsResponse) GetTypeName() string {
	return "ModbusPDUWriteMultipleCoilsResponse"
}

func (m *ModbusPDUWriteMultipleCoilsResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusPDUWriteMultipleCoilsResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUWriteMultipleCoilsResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUWriteMultipleCoilsResponseParse(readBuffer utils.ReadBuffer) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUWriteMultipleCoilsResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (startingAddress)
	startingAddress, _startingAddressErr := readBuffer.ReadUint16("startingAddress", 16)
	if _startingAddressErr != nil {
		return nil, errors.Wrap(_startingAddressErr, "Error parsing 'startingAddress' field")
	}

	// Simple Field (quantity)
	quantity, _quantityErr := readBuffer.ReadUint16("quantity", 16)
	if _quantityErr != nil {
		return nil, errors.Wrap(_quantityErr, "Error parsing 'quantity' field")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteMultipleCoilsResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUWriteMultipleCoilsResponse{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		Parent:          &ModbusPDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ModbusPDUWriteMultipleCoilsResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteMultipleCoilsResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (startingAddress)
		startingAddress := uint16(m.StartingAddress)
		_startingAddressErr := writeBuffer.WriteUint16("startingAddress", 16, (startingAddress))
		if _startingAddressErr != nil {
			return errors.Wrap(_startingAddressErr, "Error serializing 'startingAddress' field")
		}

		// Simple Field (quantity)
		quantity := uint16(m.Quantity)
		_quantityErr := writeBuffer.WriteUint16("quantity", 16, (quantity))
		if _quantityErr != nil {
			return errors.Wrap(_quantityErr, "Error serializing 'quantity' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteMultipleCoilsResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUWriteMultipleCoilsResponse) String() string {
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}

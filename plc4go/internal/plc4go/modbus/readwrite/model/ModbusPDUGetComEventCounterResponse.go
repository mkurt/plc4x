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
type ModbusPDUGetComEventCounterResponse struct {
	Status     uint16
	EventCount uint16
	Parent     *ModbusPDU
}

// The corresponding interface
type IModbusPDUGetComEventCounterResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUGetComEventCounterResponse) ErrorFlag() bool {
	return false
}

func (m *ModbusPDUGetComEventCounterResponse) FunctionFlag() uint8 {
	return 0x0B
}

func (m *ModbusPDUGetComEventCounterResponse) Response() bool {
	return true
}

func (m *ModbusPDUGetComEventCounterResponse) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUGetComEventCounterResponse(status uint16, eventCount uint16) *ModbusPDU {
	child := &ModbusPDUGetComEventCounterResponse{
		Status:     status,
		EventCount: eventCount,
		Parent:     NewModbusPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastModbusPDUGetComEventCounterResponse(structType interface{}) *ModbusPDUGetComEventCounterResponse {
	castFunc := func(typ interface{}) *ModbusPDUGetComEventCounterResponse {
		if casted, ok := typ.(ModbusPDUGetComEventCounterResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUGetComEventCounterResponse); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUGetComEventCounterResponse(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUGetComEventCounterResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUGetComEventCounterResponse) GetTypeName() string {
	return "ModbusPDUGetComEventCounterResponse"
}

func (m *ModbusPDUGetComEventCounterResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusPDUGetComEventCounterResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (status)
	lengthInBits += 16

	// Simple field (eventCount)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUGetComEventCounterResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUGetComEventCounterResponseParse(readBuffer utils.ReadBuffer) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUGetComEventCounterResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (status)
	status, _statusErr := readBuffer.ReadUint16("status", 16)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}

	// Simple Field (eventCount)
	eventCount, _eventCountErr := readBuffer.ReadUint16("eventCount", 16)
	if _eventCountErr != nil {
		return nil, errors.Wrap(_eventCountErr, "Error parsing 'eventCount' field")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUGetComEventCounterResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUGetComEventCounterResponse{
		Status:     status,
		EventCount: eventCount,
		Parent:     &ModbusPDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ModbusPDUGetComEventCounterResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUGetComEventCounterResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (status)
		status := uint16(m.Status)
		_statusErr := writeBuffer.WriteUint16("status", 16, (status))
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		// Simple Field (eventCount)
		eventCount := uint16(m.EventCount)
		_eventCountErr := writeBuffer.WriteUint16("eventCount", 16, (eventCount))
		if _eventCountErr != nil {
			return errors.Wrap(_eventCountErr, "Error serializing 'eventCount' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUGetComEventCounterResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUGetComEventCounterResponse) String() string {
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}

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
type ModbusPDUReadFileRecordResponseItem struct {
	ReferenceType uint8
	Data          []int8
}

// The corresponding interface
type IModbusPDUReadFileRecordResponseItem interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewModbusPDUReadFileRecordResponseItem(referenceType uint8, data []int8) *ModbusPDUReadFileRecordResponseItem {
	return &ModbusPDUReadFileRecordResponseItem{ReferenceType: referenceType, Data: data}
}

func CastModbusPDUReadFileRecordResponseItem(structType interface{}) *ModbusPDUReadFileRecordResponseItem {
	castFunc := func(typ interface{}) *ModbusPDUReadFileRecordResponseItem {
		if casted, ok := typ.(ModbusPDUReadFileRecordResponseItem); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUReadFileRecordResponseItem); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUReadFileRecordResponseItem) GetTypeName() string {
	return "ModbusPDUReadFileRecordResponseItem"
}

func (m *ModbusPDUReadFileRecordResponseItem) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusPDUReadFileRecordResponseItem) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (dataLength)
	lengthInBits += 8

	// Simple field (referenceType)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *ModbusPDUReadFileRecordResponseItem) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUReadFileRecordResponseItemParse(readBuffer utils.ReadBuffer) (*ModbusPDUReadFileRecordResponseItem, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUReadFileRecordResponseItem"); pullErr != nil {
		return nil, pullErr
	}

	// Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	dataLength, _dataLengthErr := readBuffer.ReadUint8("dataLength", 8)
	_ = dataLength
	if _dataLengthErr != nil {
		return nil, errors.Wrap(_dataLengthErr, "Error parsing 'dataLength' field")
	}

	// Simple Field (referenceType)
	referenceType, _referenceTypeErr := readBuffer.ReadUint8("referenceType", 8)
	if _referenceTypeErr != nil {
		return nil, errors.Wrap(_referenceTypeErr, "Error parsing 'referenceType' field")
	}

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	data := make([]int8, 0)
	_dataLength := uint16(dataLength) - uint16(uint16(1))
	_dataEndPos := readBuffer.GetPos() + uint16(_dataLength)
	for readBuffer.GetPos() < _dataEndPos {
		_item, _err := readBuffer.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data = append(data, _item)
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadFileRecordResponseItem"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewModbusPDUReadFileRecordResponseItem(referenceType, data), nil
}

func (m *ModbusPDUReadFileRecordResponseItem) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("ModbusPDUReadFileRecordResponseItem"); pushErr != nil {
		return pushErr
	}

	// Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	dataLength := uint8(uint8(uint8(len(m.Data))) + uint8(uint8(1)))
	_dataLengthErr := writeBuffer.WriteUint8("dataLength", 8, (dataLength))
	if _dataLengthErr != nil {
		return errors.Wrap(_dataLengthErr, "Error serializing 'dataLength' field")
	}

	// Simple Field (referenceType)
	referenceType := uint8(m.ReferenceType)
	_referenceTypeErr := writeBuffer.WriteUint8("referenceType", 8, (referenceType))
	if _referenceTypeErr != nil {
		return errors.Wrap(_referenceTypeErr, "Error serializing 'referenceType' field")
	}

	// Array Field (data)
	if m.Data != nil {
		if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.Data {
			_elementErr := writeBuffer.WriteInt8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'data' field")
			}
		}
		if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	if popErr := writeBuffer.PopContext("ModbusPDUReadFileRecordResponseItem"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *ModbusPDUReadFileRecordResponseItem) String() string {
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}

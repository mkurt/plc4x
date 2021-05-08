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
type ComObjectTableRealisationType6 struct {
	ComObjectDescriptors *GroupObjectDescriptorRealisationType6
	Parent               *ComObjectTable
}

// The corresponding interface
type IComObjectTableRealisationType6 interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ComObjectTableRealisationType6) FirmwareType() FirmwareType {
	return FirmwareType_SYSTEM_300
}

func (m *ComObjectTableRealisationType6) InitializeParent(parent *ComObjectTable) {
}

func NewComObjectTableRealisationType6(comObjectDescriptors *GroupObjectDescriptorRealisationType6) *ComObjectTable {
	child := &ComObjectTableRealisationType6{
		ComObjectDescriptors: comObjectDescriptors,
		Parent:               NewComObjectTable(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastComObjectTableRealisationType6(structType interface{}) *ComObjectTableRealisationType6 {
	castFunc := func(typ interface{}) *ComObjectTableRealisationType6 {
		if casted, ok := typ.(ComObjectTableRealisationType6); ok {
			return &casted
		}
		if casted, ok := typ.(*ComObjectTableRealisationType6); ok {
			return casted
		}
		if casted, ok := typ.(ComObjectTable); ok {
			return CastComObjectTableRealisationType6(casted.Child)
		}
		if casted, ok := typ.(*ComObjectTable); ok {
			return CastComObjectTableRealisationType6(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ComObjectTableRealisationType6) GetTypeName() string {
	return "ComObjectTableRealisationType6"
}

func (m *ComObjectTableRealisationType6) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ComObjectTableRealisationType6) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (comObjectDescriptors)
	lengthInBits += m.ComObjectDescriptors.LengthInBits()

	return lengthInBits
}

func (m *ComObjectTableRealisationType6) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ComObjectTableRealisationType6Parse(readBuffer utils.ReadBuffer) (*ComObjectTable, error) {
	if pullErr := readBuffer.PullContext("ComObjectTableRealisationType6"); pullErr != nil {
		return nil, pullErr
	}

	if pullErr := readBuffer.PullContext("comObjectDescriptors"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (comObjectDescriptors)
	comObjectDescriptors, _comObjectDescriptorsErr := GroupObjectDescriptorRealisationType6Parse(readBuffer)
	if _comObjectDescriptorsErr != nil {
		return nil, errors.Wrap(_comObjectDescriptorsErr, "Error parsing 'comObjectDescriptors' field")
	}
	if closeErr := readBuffer.CloseContext("comObjectDescriptors"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ComObjectTableRealisationType6"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ComObjectTableRealisationType6{
		ComObjectDescriptors: comObjectDescriptors,
		Parent:               &ComObjectTable{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ComObjectTableRealisationType6) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ComObjectTableRealisationType6"); pushErr != nil {
			return pushErr
		}

		// Simple Field (comObjectDescriptors)
		if pushErr := writeBuffer.PushContext("comObjectDescriptors"); pushErr != nil {
			return pushErr
		}
		_comObjectDescriptorsErr := m.ComObjectDescriptors.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("comObjectDescriptors"); popErr != nil {
			return popErr
		}
		if _comObjectDescriptorsErr != nil {
			return errors.Wrap(_comObjectDescriptorsErr, "Error serializing 'comObjectDescriptors' field")
		}

		if popErr := writeBuffer.PopContext("ComObjectTableRealisationType6"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *ComObjectTableRealisationType6) String() string {
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}

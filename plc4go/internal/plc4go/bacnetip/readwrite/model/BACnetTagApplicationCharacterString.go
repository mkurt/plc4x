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
type BACnetTagApplicationCharacterString struct {
	Parent *BACnetTag
}

// The corresponding interface
type IBACnetTagApplicationCharacterString interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetTagApplicationCharacterString) ContextSpecificTag() uint8 {
	return 0
}

func (m *BACnetTagApplicationCharacterString) InitializeParent(parent *BACnetTag, typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8) {
	m.Parent.TypeOrTagNumber = typeOrTagNumber
	m.Parent.LengthValueType = lengthValueType
	m.Parent.ExtTagNumber = extTagNumber
	m.Parent.ExtLength = extLength
}

func NewBACnetTagApplicationCharacterString(typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8) *BACnetTag {
	child := &BACnetTagApplicationCharacterString{
		Parent: NewBACnetTag(typeOrTagNumber, lengthValueType, extTagNumber, extLength),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetTagApplicationCharacterString(structType interface{}) *BACnetTagApplicationCharacterString {
	castFunc := func(typ interface{}) *BACnetTagApplicationCharacterString {
		if casted, ok := typ.(BACnetTagApplicationCharacterString); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetTagApplicationCharacterString); ok {
			return casted
		}
		if casted, ok := typ.(BACnetTag); ok {
			return CastBACnetTagApplicationCharacterString(casted.Child)
		}
		if casted, ok := typ.(*BACnetTag); ok {
			return CastBACnetTagApplicationCharacterString(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetTagApplicationCharacterString) GetTypeName() string {
	return "BACnetTagApplicationCharacterString"
}

func (m *BACnetTagApplicationCharacterString) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetTagApplicationCharacterString) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetTagApplicationCharacterString) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetTagApplicationCharacterStringParse(readBuffer utils.ReadBuffer) (*BACnetTag, error) {
	if pullErr := readBuffer.PullContext("BACnetTagApplicationCharacterString"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BACnetTagApplicationCharacterString"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetTagApplicationCharacterString{
		Parent: &BACnetTag{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetTagApplicationCharacterString) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTagApplicationCharacterString"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetTagApplicationCharacterString"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetTagApplicationCharacterString) String() string {
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}

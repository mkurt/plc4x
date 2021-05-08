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
type BACnetTag struct {
	TypeOrTagNumber uint8
	LengthValueType uint8
	ExtTagNumber    *uint8
	ExtLength       *uint8
	Child           IBACnetTagChild
}

// The corresponding interface
type IBACnetTag interface {
	ContextSpecificTag() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetTagParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetTag, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetTagChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetTag, typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8)
	GetTypeName() string
	IBACnetTag
}

func NewBACnetTag(typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8) *BACnetTag {
	return &BACnetTag{TypeOrTagNumber: typeOrTagNumber, LengthValueType: lengthValueType, ExtTagNumber: extTagNumber, ExtLength: extLength}
}

func CastBACnetTag(structType interface{}) *BACnetTag {
	castFunc := func(typ interface{}) *BACnetTag {
		if casted, ok := typ.(BACnetTag); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetTag); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetTag) GetTypeName() string {
	return "BACnetTag"
}

func (m *BACnetTag) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetTag) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *BACnetTag) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (typeOrTagNumber)
	lengthInBits += 4
	// Discriminator Field (contextSpecificTag)
	lengthInBits += 1

	// Simple field (lengthValueType)
	lengthInBits += 3

	// Optional Field (extTagNumber)
	if m.ExtTagNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (extLength)
	if m.ExtLength != nil {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *BACnetTag) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetTagParse(readBuffer utils.ReadBuffer) (*BACnetTag, error) {
	if pullErr := readBuffer.PullContext("BACnetTag"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (typeOrTagNumber)
	typeOrTagNumber, _typeOrTagNumberErr := readBuffer.ReadUint8("typeOrTagNumber", 4)
	if _typeOrTagNumberErr != nil {
		return nil, errors.Wrap(_typeOrTagNumberErr, "Error parsing 'typeOrTagNumber' field")
	}

	// Discriminator Field (contextSpecificTag) (Used as input to a switch field)
	contextSpecificTag, _contextSpecificTagErr := readBuffer.ReadUint8("contextSpecificTag", 1)
	if _contextSpecificTagErr != nil {
		return nil, errors.Wrap(_contextSpecificTagErr, "Error parsing 'contextSpecificTag' field")
	}

	// Simple Field (lengthValueType)
	lengthValueType, _lengthValueTypeErr := readBuffer.ReadUint8("lengthValueType", 3)
	if _lengthValueTypeErr != nil {
		return nil, errors.Wrap(_lengthValueTypeErr, "Error parsing 'lengthValueType' field")
	}

	// Optional Field (extTagNumber) (Can be skipped, if a given expression evaluates to false)
	var extTagNumber *uint8 = nil
	if bool((typeOrTagNumber) == (15)) {
		_val, _err := readBuffer.ReadUint8("extTagNumber", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extTagNumber' field")
		}
		extTagNumber = &_val
	}

	// Optional Field (extLength) (Can be skipped, if a given expression evaluates to false)
	var extLength *uint8 = nil
	if bool((lengthValueType) == (5)) {
		_val, _err := readBuffer.ReadUint8("extLength", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extLength' field")
		}
		extLength = &_val
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetTag
	var typeSwitchError error
	switch {
	case contextSpecificTag == 0 && typeOrTagNumber == 0x0: // BACnetTagApplicationNull
		_parent, typeSwitchError = BACnetTagApplicationNullParse(readBuffer)
	case contextSpecificTag == 0 && typeOrTagNumber == 0x1: // BACnetTagApplicationBoolean
		_parent, typeSwitchError = BACnetTagApplicationBooleanParse(readBuffer)
	case contextSpecificTag == 0 && typeOrTagNumber == 0x2: // BACnetTagApplicationUnsignedInteger
		_parent, typeSwitchError = BACnetTagApplicationUnsignedIntegerParse(readBuffer, lengthValueType, *extLength)
	case contextSpecificTag == 0 && typeOrTagNumber == 0x3: // BACnetTagApplicationSignedInteger
		_parent, typeSwitchError = BACnetTagApplicationSignedIntegerParse(readBuffer, lengthValueType, *extLength)
	case contextSpecificTag == 0 && typeOrTagNumber == 0x4: // BACnetTagApplicationReal
		_parent, typeSwitchError = BACnetTagApplicationRealParse(readBuffer, lengthValueType, *extLength)
	case contextSpecificTag == 0 && typeOrTagNumber == 0x5: // BACnetTagApplicationDouble
		_parent, typeSwitchError = BACnetTagApplicationDoubleParse(readBuffer, lengthValueType, *extLength)
	case contextSpecificTag == 0 && typeOrTagNumber == 0x6: // BACnetTagApplicationOctetString
		_parent, typeSwitchError = BACnetTagApplicationOctetStringParse(readBuffer)
	case contextSpecificTag == 0 && typeOrTagNumber == 0x7: // BACnetTagApplicationCharacterString
		_parent, typeSwitchError = BACnetTagApplicationCharacterStringParse(readBuffer)
	case contextSpecificTag == 0 && typeOrTagNumber == 0x8: // BACnetTagApplicationBitString
		_parent, typeSwitchError = BACnetTagApplicationBitStringParse(readBuffer, lengthValueType, *extLength)
	case contextSpecificTag == 0 && typeOrTagNumber == 0x9: // BACnetTagApplicationEnumerated
		_parent, typeSwitchError = BACnetTagApplicationEnumeratedParse(readBuffer, lengthValueType, *extLength)
	case contextSpecificTag == 0 && typeOrTagNumber == 0xA: // BACnetTagApplicationDate
		_parent, typeSwitchError = BACnetTagApplicationDateParse(readBuffer)
	case contextSpecificTag == 0 && typeOrTagNumber == 0xB: // BACnetTagApplicationTime
		_parent, typeSwitchError = BACnetTagApplicationTimeParse(readBuffer)
	case contextSpecificTag == 0 && typeOrTagNumber == 0xC: // BACnetTagApplicationObjectIdentifier
		_parent, typeSwitchError = BACnetTagApplicationObjectIdentifierParse(readBuffer)
	case contextSpecificTag == 1: // BACnetTagContext
		_parent, typeSwitchError = BACnetTagContextParse(readBuffer, typeOrTagNumber, *extTagNumber, lengthValueType, *extLength)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetTag"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, typeOrTagNumber, lengthValueType, extTagNumber, extLength)
	return _parent, nil
}

func (m *BACnetTag) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetTag) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetTag, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetTag"); pushErr != nil {
		return pushErr
	}

	// Simple Field (typeOrTagNumber)
	typeOrTagNumber := uint8(m.TypeOrTagNumber)
	_typeOrTagNumberErr := writeBuffer.WriteUint8("typeOrTagNumber", 4, (typeOrTagNumber))
	if _typeOrTagNumberErr != nil {
		return errors.Wrap(_typeOrTagNumberErr, "Error serializing 'typeOrTagNumber' field")
	}

	// Discriminator Field (contextSpecificTag) (Used as input to a switch field)
	contextSpecificTag := uint8(child.ContextSpecificTag())
	_contextSpecificTagErr := writeBuffer.WriteUint8("contextSpecificTag", 1, (contextSpecificTag))

	if _contextSpecificTagErr != nil {
		return errors.Wrap(_contextSpecificTagErr, "Error serializing 'contextSpecificTag' field")
	}

	// Simple Field (lengthValueType)
	lengthValueType := uint8(m.LengthValueType)
	_lengthValueTypeErr := writeBuffer.WriteUint8("lengthValueType", 3, (lengthValueType))
	if _lengthValueTypeErr != nil {
		return errors.Wrap(_lengthValueTypeErr, "Error serializing 'lengthValueType' field")
	}

	// Optional Field (extTagNumber) (Can be skipped, if the value is null)
	var extTagNumber *uint8 = nil
	if m.ExtTagNumber != nil {
		extTagNumber = m.ExtTagNumber
		_extTagNumberErr := writeBuffer.WriteUint8("extTagNumber", 8, *(extTagNumber))
		if _extTagNumberErr != nil {
			return errors.Wrap(_extTagNumberErr, "Error serializing 'extTagNumber' field")
		}
	}

	// Optional Field (extLength) (Can be skipped, if the value is null)
	var extLength *uint8 = nil
	if m.ExtLength != nil {
		extLength = m.ExtLength
		_extLengthErr := writeBuffer.WriteUint8("extLength", 8, *(extLength))
		if _extLengthErr != nil {
			return errors.Wrap(_extLengthErr, "Error serializing 'extLength' field")
		}
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetTag"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetTag) String() string {
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}

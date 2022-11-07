/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestRemoveListElement is the corresponding interface of BACnetConfirmedServiceRequestRemoveListElement
type BACnetConfirmedServiceRequestRemoveListElement interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() BACnetContextTagUnsignedInteger
	// GetListOfElements returns ListOfElements (property field)
	GetListOfElements() BACnetConstructedData
}

// BACnetConfirmedServiceRequestRemoveListElementExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestRemoveListElement.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestRemoveListElementExactly interface {
	BACnetConfirmedServiceRequestRemoveListElement
	isBACnetConfirmedServiceRequestRemoveListElement() bool
}

// _BACnetConfirmedServiceRequestRemoveListElement is the data-structure of this message
type _BACnetConfirmedServiceRequestRemoveListElement struct {
	*_BACnetConfirmedServiceRequest
	ObjectIdentifier   BACnetContextTagObjectIdentifier
	PropertyIdentifier BACnetPropertyIdentifierTagged
	ArrayIndex         BACnetContextTagUnsignedInteger
	ListOfElements     BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestRemoveListElement) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestRemoveListElement) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestRemoveListElement) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestRemoveListElement) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestRemoveListElement) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetConfirmedServiceRequestRemoveListElement) GetArrayIndex() BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *_BACnetConfirmedServiceRequestRemoveListElement) GetListOfElements() BACnetConstructedData {
	return m.ListOfElements
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestRemoveListElement factory function for _BACnetConfirmedServiceRequestRemoveListElement
func NewBACnetConfirmedServiceRequestRemoveListElement(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, arrayIndex BACnetContextTagUnsignedInteger, listOfElements BACnetConstructedData, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestRemoveListElement {
	_result := &_BACnetConfirmedServiceRequestRemoveListElement{
		ObjectIdentifier:               objectIdentifier,
		PropertyIdentifier:             propertyIdentifier,
		ArrayIndex:                     arrayIndex,
		ListOfElements:                 listOfElements,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestRemoveListElement(structType interface{}) BACnetConfirmedServiceRequestRemoveListElement {
	if casted, ok := structType.(BACnetConfirmedServiceRequestRemoveListElement); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestRemoveListElement); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestRemoveListElement) GetTypeName() string {
	return "BACnetConfirmedServiceRequestRemoveListElement"
}

func (m *_BACnetConfirmedServiceRequestRemoveListElement) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestRemoveListElement) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits()

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits()

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += m.ArrayIndex.GetLengthInBits()
	}

	// Optional Field (listOfElements)
	if m.ListOfElements != nil {
		lengthInBits += m.ListOfElements.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestRemoveListElement) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestRemoveListElementParse(theBytes []byte, serviceRequestLength uint32) (BACnetConfirmedServiceRequestRemoveListElement, error) {
	return BACnetConfirmedServiceRequestRemoveListElementParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), serviceRequestLength) // TODO: get endianness from mspec
}

func BACnetConfirmedServiceRequestRemoveListElementParseWithBuffer(readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequestRemoveListElement, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestRemoveListElement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestRemoveListElement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParseWithBuffer(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field of BACnetConfirmedServiceRequestRemoveListElement")
	}
	objectIdentifier := _objectIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
	}

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyIdentifier")
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetPropertyIdentifierTaggedParseWithBuffer(readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field of BACnetConfirmedServiceRequestRemoveListElement")
	}
	propertyIdentifier := _propertyIdentifier.(BACnetPropertyIdentifierTagged)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyIdentifier")
	}

	// Optional Field (arrayIndex) (Can be skipped, if a given expression evaluates to false)
	var arrayIndex BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("arrayIndex"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for arrayIndex")
		}
		_val, _err := BACnetContextTagParseWithBuffer(readBuffer, uint8(2), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'arrayIndex' field of BACnetConfirmedServiceRequestRemoveListElement")
		default:
			arrayIndex = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("arrayIndex"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for arrayIndex")
			}
		}
	}

	// Optional Field (listOfElements) (Can be skipped, if a given expression evaluates to false)
	var listOfElements BACnetConstructedData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("listOfElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for listOfElements")
		}
		_val, _err := BACnetConstructedDataParseWithBuffer(readBuffer, uint8(3), objectIdentifier.GetObjectType(), propertyIdentifier.GetValue(), (CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((arrayIndex) != (nil)), func() interface{} { return CastBACnetTagPayloadUnsignedInteger((arrayIndex).GetPayload()) }, func() interface{} { return CastBACnetTagPayloadUnsignedInteger(nil) }))))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'listOfElements' field of BACnetConfirmedServiceRequestRemoveListElement")
		default:
			listOfElements = _val.(BACnetConstructedData)
			if closeErr := readBuffer.CloseContext("listOfElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for listOfElements")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestRemoveListElement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestRemoveListElement")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestRemoveListElement{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		ObjectIdentifier:   objectIdentifier,
		PropertyIdentifier: propertyIdentifier,
		ArrayIndex:         arrayIndex,
		ListOfElements:     listOfElements,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestRemoveListElement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestRemoveListElement) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestRemoveListElement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestRemoveListElement")
		}

		// Simple Field (objectIdentifier)
		if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
		}
		_objectIdentifierErr := writeBuffer.WriteSerializable(m.GetObjectIdentifier())
		if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectIdentifier")
		}
		if _objectIdentifierErr != nil {
			return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
		}

		// Simple Field (propertyIdentifier)
		if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyIdentifier")
		}
		_propertyIdentifierErr := writeBuffer.WriteSerializable(m.GetPropertyIdentifier())
		if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyIdentifier")
		}
		if _propertyIdentifierErr != nil {
			return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
		}

		// Optional Field (arrayIndex) (Can be skipped, if the value is null)
		var arrayIndex BACnetContextTagUnsignedInteger = nil
		if m.GetArrayIndex() != nil {
			if pushErr := writeBuffer.PushContext("arrayIndex"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for arrayIndex")
			}
			arrayIndex = m.GetArrayIndex()
			_arrayIndexErr := writeBuffer.WriteSerializable(arrayIndex)
			if popErr := writeBuffer.PopContext("arrayIndex"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for arrayIndex")
			}
			if _arrayIndexErr != nil {
				return errors.Wrap(_arrayIndexErr, "Error serializing 'arrayIndex' field")
			}
		}

		// Optional Field (listOfElements) (Can be skipped, if the value is null)
		var listOfElements BACnetConstructedData = nil
		if m.GetListOfElements() != nil {
			if pushErr := writeBuffer.PushContext("listOfElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for listOfElements")
			}
			listOfElements = m.GetListOfElements()
			_listOfElementsErr := writeBuffer.WriteSerializable(listOfElements)
			if popErr := writeBuffer.PopContext("listOfElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for listOfElements")
			}
			if _listOfElementsErr != nil {
				return errors.Wrap(_listOfElementsErr, "Error serializing 'listOfElements' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestRemoveListElement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestRemoveListElement")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestRemoveListElement) isBACnetConfirmedServiceRequestRemoveListElement() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestRemoveListElement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
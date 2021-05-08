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
	"github.com/rs/zerolog/log"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type DeviceConfigurationRequestDataBlock struct {
	CommunicationChannelId uint8
	SequenceCounter        uint8
}

// The corresponding interface
type IDeviceConfigurationRequestDataBlock interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewDeviceConfigurationRequestDataBlock(communicationChannelId uint8, sequenceCounter uint8) *DeviceConfigurationRequestDataBlock {
	return &DeviceConfigurationRequestDataBlock{CommunicationChannelId: communicationChannelId, SequenceCounter: sequenceCounter}
}

func CastDeviceConfigurationRequestDataBlock(structType interface{}) *DeviceConfigurationRequestDataBlock {
	castFunc := func(typ interface{}) *DeviceConfigurationRequestDataBlock {
		if casted, ok := typ.(DeviceConfigurationRequestDataBlock); ok {
			return &casted
		}
		if casted, ok := typ.(*DeviceConfigurationRequestDataBlock); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DeviceConfigurationRequestDataBlock) GetTypeName() string {
	return "DeviceConfigurationRequestDataBlock"
}

func (m *DeviceConfigurationRequestDataBlock) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *DeviceConfigurationRequestDataBlock) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (sequenceCounter)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *DeviceConfigurationRequestDataBlock) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DeviceConfigurationRequestDataBlockParse(readBuffer utils.ReadBuffer) (*DeviceConfigurationRequestDataBlock, error) {
	if pullErr := readBuffer.PullContext("DeviceConfigurationRequestDataBlock"); pullErr != nil {
		return nil, pullErr
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength, _structureLengthErr := readBuffer.ReadUint8("structureLength", 8)
	_ = structureLength
	if _structureLengthErr != nil {
		return nil, errors.Wrap(_structureLengthErr, "Error parsing 'structureLength' field")
	}

	// Simple Field (communicationChannelId)
	communicationChannelId, _communicationChannelIdErr := readBuffer.ReadUint8("communicationChannelId", 8)
	if _communicationChannelIdErr != nil {
		return nil, errors.Wrap(_communicationChannelIdErr, "Error parsing 'communicationChannelId' field")
	}

	// Simple Field (sequenceCounter)
	sequenceCounter, _sequenceCounterErr := readBuffer.ReadUint8("sequenceCounter", 8)
	if _sequenceCounterErr != nil {
		return nil, errors.Wrap(_sequenceCounterErr, "Error parsing 'sequenceCounter' field")
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	if closeErr := readBuffer.CloseContext("DeviceConfigurationRequestDataBlock"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewDeviceConfigurationRequestDataBlock(communicationChannelId, sequenceCounter), nil
}

func (m *DeviceConfigurationRequestDataBlock) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("DeviceConfigurationRequestDataBlock"); pushErr != nil {
		return pushErr
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.LengthInBytes()))
	_structureLengthErr := writeBuffer.WriteUint8("structureLength", 8, (structureLength))
	if _structureLengthErr != nil {
		return errors.Wrap(_structureLengthErr, "Error serializing 'structureLength' field")
	}

	// Simple Field (communicationChannelId)
	communicationChannelId := uint8(m.CommunicationChannelId)
	_communicationChannelIdErr := writeBuffer.WriteUint8("communicationChannelId", 8, (communicationChannelId))
	if _communicationChannelIdErr != nil {
		return errors.Wrap(_communicationChannelIdErr, "Error serializing 'communicationChannelId' field")
	}

	// Simple Field (sequenceCounter)
	sequenceCounter := uint8(m.SequenceCounter)
	_sequenceCounterErr := writeBuffer.WriteUint8("sequenceCounter", 8, (sequenceCounter))
	if _sequenceCounterErr != nil {
		return errors.Wrap(_sequenceCounterErr, "Error serializing 'sequenceCounter' field")
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x00))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	if popErr := writeBuffer.PopContext("DeviceConfigurationRequestDataBlock"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *DeviceConfigurationRequestDataBlock) String() string {
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}

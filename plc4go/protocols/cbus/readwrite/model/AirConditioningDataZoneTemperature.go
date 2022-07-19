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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// AirConditioningDataZoneTemperature is the corresponding interface of AirConditioningDataZoneTemperature
type AirConditioningDataZoneTemperature interface {
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetTemperature returns Temperature (property field)
	GetTemperature() HVACTemperature
	// GetSensorStatus returns SensorStatus (property field)
	GetSensorStatus() HVACSensorStatus
}

// AirConditioningDataZoneTemperatureExactly can be used when we want exactly this type and not a type which fulfills AirConditioningDataZoneTemperature.
// This is useful for switch cases.
type AirConditioningDataZoneTemperatureExactly interface {
	AirConditioningDataZoneTemperature
	isAirConditioningDataZoneTemperature() bool
}

// _AirConditioningDataZoneTemperature is the data-structure of this message
type _AirConditioningDataZoneTemperature struct {
	*_AirConditioningData
	ZoneGroup    byte
	ZoneList     HVACZoneList
	Temperature  HVACTemperature
	SensorStatus HVACSensorStatus
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataZoneTemperature) InitializeParent(parent AirConditioningData, commandTypeContainer AirConditioningCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_AirConditioningDataZoneTemperature) GetParent() AirConditioningData {
	return m._AirConditioningData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataZoneTemperature) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataZoneTemperature) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataZoneTemperature) GetTemperature() HVACTemperature {
	return m.Temperature
}

func (m *_AirConditioningDataZoneTemperature) GetSensorStatus() HVACSensorStatus {
	return m.SensorStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataZoneTemperature factory function for _AirConditioningDataZoneTemperature
func NewAirConditioningDataZoneTemperature(zoneGroup byte, zoneList HVACZoneList, temperature HVACTemperature, sensorStatus HVACSensorStatus, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataZoneTemperature {
	_result := &_AirConditioningDataZoneTemperature{
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		Temperature:          temperature,
		SensorStatus:         sensorStatus,
		_AirConditioningData: NewAirConditioningData(commandTypeContainer),
	}
	_result._AirConditioningData._AirConditioningDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataZoneTemperature(structType interface{}) AirConditioningDataZoneTemperature {
	if casted, ok := structType.(AirConditioningDataZoneTemperature); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataZoneTemperature); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataZoneTemperature) GetTypeName() string {
	return "AirConditioningDataZoneTemperature"
}

func (m *_AirConditioningDataZoneTemperature) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AirConditioningDataZoneTemperature) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits()

	// Simple field (temperature)
	lengthInBits += m.Temperature.GetLengthInBits()

	// Simple field (sensorStatus)
	lengthInBits += 8

	return lengthInBits
}

func (m *_AirConditioningDataZoneTemperature) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AirConditioningDataZoneTemperatureParse(readBuffer utils.ReadBuffer) (AirConditioningDataZoneTemperature, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataZoneTemperature"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataZoneTemperature")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneGroup)
	_zoneGroup, _zoneGroupErr := readBuffer.ReadByte("zoneGroup")
	if _zoneGroupErr != nil {
		return nil, errors.Wrap(_zoneGroupErr, "Error parsing 'zoneGroup' field of AirConditioningDataZoneTemperature")
	}
	zoneGroup := _zoneGroup

	// Simple Field (zoneList)
	if pullErr := readBuffer.PullContext("zoneList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneList")
	}
	_zoneList, _zoneListErr := HVACZoneListParse(readBuffer)
	if _zoneListErr != nil {
		return nil, errors.Wrap(_zoneListErr, "Error parsing 'zoneList' field of AirConditioningDataZoneTemperature")
	}
	zoneList := _zoneList.(HVACZoneList)
	if closeErr := readBuffer.CloseContext("zoneList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneList")
	}

	// Simple Field (temperature)
	if pullErr := readBuffer.PullContext("temperature"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for temperature")
	}
	_temperature, _temperatureErr := HVACTemperatureParse(readBuffer)
	if _temperatureErr != nil {
		return nil, errors.Wrap(_temperatureErr, "Error parsing 'temperature' field of AirConditioningDataZoneTemperature")
	}
	temperature := _temperature.(HVACTemperature)
	if closeErr := readBuffer.CloseContext("temperature"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for temperature")
	}

	// Simple Field (sensorStatus)
	if pullErr := readBuffer.PullContext("sensorStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for sensorStatus")
	}
	_sensorStatus, _sensorStatusErr := HVACSensorStatusParse(readBuffer)
	if _sensorStatusErr != nil {
		return nil, errors.Wrap(_sensorStatusErr, "Error parsing 'sensorStatus' field of AirConditioningDataZoneTemperature")
	}
	sensorStatus := _sensorStatus
	if closeErr := readBuffer.CloseContext("sensorStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for sensorStatus")
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataZoneTemperature"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataZoneTemperature")
	}

	// Create a partially initialized instance
	_child := &_AirConditioningDataZoneTemperature{
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		Temperature:          temperature,
		SensorStatus:         sensorStatus,
		_AirConditioningData: &_AirConditioningData{},
	}
	_child._AirConditioningData._AirConditioningDataChildRequirements = _child
	return _child, nil
}

func (m *_AirConditioningDataZoneTemperature) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataZoneTemperature"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataZoneTemperature")
		}

		// Simple Field (zoneGroup)
		zoneGroup := byte(m.GetZoneGroup())
		_zoneGroupErr := writeBuffer.WriteByte("zoneGroup", (zoneGroup))
		if _zoneGroupErr != nil {
			return errors.Wrap(_zoneGroupErr, "Error serializing 'zoneGroup' field")
		}

		// Simple Field (zoneList)
		if pushErr := writeBuffer.PushContext("zoneList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for zoneList")
		}
		_zoneListErr := writeBuffer.WriteSerializable(m.GetZoneList())
		if popErr := writeBuffer.PopContext("zoneList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for zoneList")
		}
		if _zoneListErr != nil {
			return errors.Wrap(_zoneListErr, "Error serializing 'zoneList' field")
		}

		// Simple Field (temperature)
		if pushErr := writeBuffer.PushContext("temperature"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for temperature")
		}
		_temperatureErr := writeBuffer.WriteSerializable(m.GetTemperature())
		if popErr := writeBuffer.PopContext("temperature"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for temperature")
		}
		if _temperatureErr != nil {
			return errors.Wrap(_temperatureErr, "Error serializing 'temperature' field")
		}

		// Simple Field (sensorStatus)
		if pushErr := writeBuffer.PushContext("sensorStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for sensorStatus")
		}
		_sensorStatusErr := writeBuffer.WriteSerializable(m.GetSensorStatus())
		if popErr := writeBuffer.PopContext("sensorStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for sensorStatus")
		}
		if _sensorStatusErr != nil {
			return errors.Wrap(_sensorStatusErr, "Error serializing 'sensorStatus' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataZoneTemperature"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataZoneTemperature")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_AirConditioningDataZoneTemperature) isAirConditioningDataZoneTemperature() bool {
	return true
}

func (m *_AirConditioningDataZoneTemperature) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

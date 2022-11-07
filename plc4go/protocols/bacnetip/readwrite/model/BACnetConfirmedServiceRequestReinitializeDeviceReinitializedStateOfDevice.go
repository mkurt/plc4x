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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice is an enum
type BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice uint8

type IBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice interface {
	utils.Serializable
}

const (
	BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_COLDSTART                BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice = 0x0
	BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_WARMSTART                BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice = 0x1
	BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ACTIVATE_CHANGES         BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice = 0x2
	BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_STARTBACKUP              BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice = 0x3
	BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ENDBACKUP                BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice = 0x4
	BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_STARTRESTORE             BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice = 0x5
	BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ENDRESTORE               BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice = 0x6
	BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ABORTRESTORE             BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice = 0x7
	BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_VENDOR_PROPRIETARY_VALUE BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice = 0xFF
)

var BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceValues []BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice

func init() {
	_ = errors.New
	BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceValues = []BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice{
		BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_COLDSTART,
		BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_WARMSTART,
		BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ACTIVATE_CHANGES,
		BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_STARTBACKUP,
		BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ENDBACKUP,
		BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_STARTRESTORE,
		BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ENDRESTORE,
		BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ABORTRESTORE,
		BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceByValue(value uint8) (enum BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice, ok bool) {
	switch value {
	case 0x0:
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_COLDSTART, true
	case 0x1:
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_WARMSTART, true
	case 0x2:
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ACTIVATE_CHANGES, true
	case 0x3:
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_STARTBACKUP, true
	case 0x4:
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ENDBACKUP, true
	case 0x5:
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_STARTRESTORE, true
	case 0x6:
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ENDRESTORE, true
	case 0x7:
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ABORTRESTORE, true
	case 0xFF:
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_VENDOR_PROPRIETARY_VALUE, true
	}
	return 0, false
}

func BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceByName(value string) (enum BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice, ok bool) {
	switch value {
	case "COLDSTART":
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_COLDSTART, true
	case "WARMSTART":
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_WARMSTART, true
	case "ACTIVATE_CHANGES":
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ACTIVATE_CHANGES, true
	case "STARTBACKUP":
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_STARTBACKUP, true
	case "ENDBACKUP":
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ENDBACKUP, true
	case "STARTRESTORE":
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_STARTRESTORE, true
	case "ENDRESTORE":
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ENDRESTORE, true
	case "ABORTRESTORE":
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ABORTRESTORE, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_VENDOR_PROPRIETARY_VALUE, true
	}
	return 0, false
}

func BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceKnows(value uint8) bool {
	for _, typeValue := range BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice(structType interface{}) BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice {
	castFunc := func(typ interface{}) BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice {
		if sBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice, ok := typ.(BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice); ok {
			return sBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceParse(theBytes []byte) (BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice, error) {
	return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice, error) {
	val, err := readBuffer.ReadUint8("BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice")
	}
	if enum, ok := BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice) PLC4XEnumName() string {
	switch e {
	case BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_COLDSTART:
		return "COLDSTART"
	case BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_WARMSTART:
		return "WARMSTART"
	case BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ACTIVATE_CHANGES:
		return "ACTIVATE_CHANGES"
	case BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_STARTBACKUP:
		return "STARTBACKUP"
	case BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ENDBACKUP:
		return "ENDBACKUP"
	case BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_STARTRESTORE:
		return "STARTRESTORE"
	case BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ENDRESTORE:
		return "ENDRESTORE"
	case BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_ABORTRESTORE:
		return "ABORTRESTORE"
	case BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	}
	return ""
}

func (e BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice) String() string {
	return e.PLC4XEnumName()
}
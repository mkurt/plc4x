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
type IPAddress struct {
	Addr []int8
}

// The corresponding interface
type IIPAddress interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewIPAddress(addr []int8) *IPAddress {
	return &IPAddress{Addr: addr}
}

func CastIPAddress(structType interface{}) *IPAddress {
	castFunc := func(typ interface{}) *IPAddress {
		if casted, ok := typ.(IPAddress); ok {
			return &casted
		}
		if casted, ok := typ.(*IPAddress); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *IPAddress) GetTypeName() string {
	return "IPAddress"
}

func (m *IPAddress) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *IPAddress) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Addr) > 0 {
		lengthInBits += 8 * uint16(len(m.Addr))
	}

	return lengthInBits
}

func (m *IPAddress) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func IPAddressParse(readBuffer utils.ReadBuffer) (*IPAddress, error) {
	if pullErr := readBuffer.PullContext("IPAddress"); pullErr != nil {
		return nil, pullErr
	}

	// Array field (addr)
	if pullErr := readBuffer.PullContext("addr", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	addr := make([]int8, uint16(4))
	for curItem := uint16(0); curItem < uint16(uint16(4)); curItem++ {
		_item, _err := readBuffer.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'addr' field")
		}
		addr[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("addr", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("IPAddress"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewIPAddress(addr), nil
}

func (m *IPAddress) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("IPAddress"); pushErr != nil {
		return pushErr
	}

	// Array Field (addr)
	if m.Addr != nil {
		if pushErr := writeBuffer.PushContext("addr", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.Addr {
			_elementErr := writeBuffer.WriteInt8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'addr' field")
			}
		}
		if popErr := writeBuffer.PopContext("addr", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	if popErr := writeBuffer.PopContext("IPAddress"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *IPAddress) String() string {
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}

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
type BVLCDistributeBroadcastToNetwork struct {
	Parent *BVLC
}

// The corresponding interface
type IBVLCDistributeBroadcastToNetwork interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BVLCDistributeBroadcastToNetwork) BvlcFunction() uint8 {
	return 0x09
}

func (m *BVLCDistributeBroadcastToNetwork) InitializeParent(parent *BVLC) {
}

func NewBVLCDistributeBroadcastToNetwork() *BVLC {
	child := &BVLCDistributeBroadcastToNetwork{
		Parent: NewBVLC(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBVLCDistributeBroadcastToNetwork(structType interface{}) *BVLCDistributeBroadcastToNetwork {
	castFunc := func(typ interface{}) *BVLCDistributeBroadcastToNetwork {
		if casted, ok := typ.(BVLCDistributeBroadcastToNetwork); ok {
			return &casted
		}
		if casted, ok := typ.(*BVLCDistributeBroadcastToNetwork); ok {
			return casted
		}
		if casted, ok := typ.(BVLC); ok {
			return CastBVLCDistributeBroadcastToNetwork(casted.Child)
		}
		if casted, ok := typ.(*BVLC); ok {
			return CastBVLCDistributeBroadcastToNetwork(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BVLCDistributeBroadcastToNetwork) GetTypeName() string {
	return "BVLCDistributeBroadcastToNetwork"
}

func (m *BVLCDistributeBroadcastToNetwork) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BVLCDistributeBroadcastToNetwork) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *BVLCDistributeBroadcastToNetwork) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCDistributeBroadcastToNetworkParse(readBuffer utils.ReadBuffer) (*BVLC, error) {
	if pullErr := readBuffer.PullContext("BVLCDistributeBroadcastToNetwork"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BVLCDistributeBroadcastToNetwork"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BVLCDistributeBroadcastToNetwork{
		Parent: &BVLC{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BVLCDistributeBroadcastToNetwork) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCDistributeBroadcastToNetwork"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BVLCDistributeBroadcastToNetwork"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BVLCDistributeBroadcastToNetwork) String() string {
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}

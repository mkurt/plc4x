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
package org.apache.plc4x.java.canopen.readwrite;

import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.math.BigInteger;
import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.generation.ByteOrder;
import org.apache.plc4x.java.spi.generation.EvaluationHelper;
import org.apache.plc4x.java.spi.generation.ParseException;
import org.apache.plc4x.java.spi.generation.ReadBuffer;
import org.apache.plc4x.java.spi.generation.SerializationException;
import org.apache.plc4x.java.spi.generation.WriteBuffer;
import org.apache.plc4x.java.spi.values.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

// Code generated by code-generation. DO NOT EDIT.

public class DataItem {

  private static final Logger LOGGER = LoggerFactory.getLogger(DataItem.class);

  public static PlcValue staticParse(ReadBuffer readBuffer, CANOpenDataType dataType, Integer size)
      throws ParseException {
    if (EvaluationHelper.equals(dataType, CANOpenDataType.BOOLEAN)) { // BOOL

      // Simple Field (value)
      Boolean value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readBit("");

      return new PlcBOOL(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED8)) { // USINT

      // Simple Field (value)
      Short value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedShort("", 8);

      return new PlcUSINT(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED16)) { // UINT

      // Simple Field (value)
      Integer value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedInt("", 16);

      return new PlcUINT(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED24)) { // UDINT

      // Simple Field (value)
      Long value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedLong("", 24);

      return new PlcUDINT(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED32)) { // UDINT

      // Simple Field (value)
      Long value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedLong("", 32);

      return new PlcUDINT(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED40)) { // ULINT

      // Simple Field (value)
      BigInteger value = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readUnsignedBigInteger("", 40);

      return new PlcULINT(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED48)) { // ULINT

      // Simple Field (value)
      BigInteger value = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readUnsignedBigInteger("", 48);

      return new PlcULINT(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED56)) { // ULINT

      // Simple Field (value)
      BigInteger value = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readUnsignedBigInteger("", 56);

      return new PlcULINT(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED64)) { // ULINT

      // Simple Field (value)
      BigInteger value = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readUnsignedBigInteger("", 64);

      return new PlcULINT(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER8)) { // SINT

      // Simple Field (value)
      Byte value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readSignedByte("", 8);

      return new PlcSINT(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER16)) { // INT

      // Simple Field (value)
      Short value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readShort("", 16);

      return new PlcINT(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER24)) { // DINT

      // Simple Field (value)
      Integer value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readInt("", 24);

      return new PlcDINT(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER32)) { // DINT

      // Simple Field (value)
      Integer value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readInt("", 32);

      return new PlcDINT(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER40)) { // LINT

      // Simple Field (value)
      Long value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readLong("", 40);

      return new PlcLINT(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER48)) { // LINT

      // Simple Field (value)
      Long value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readLong("", 48);

      return new PlcLINT(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER56)) { // LINT

      // Simple Field (value)
      Long value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readLong("", 56);

      return new PlcLINT(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER64)) { // LINT

      // Simple Field (value)
      Long value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readLong("", 64);

      return new PlcLINT(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.REAL32)) { // REAL

      // Simple Field (value)
      Float value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readFloat("", 32);

      return new PlcREAL(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.REAL64)) { // LREAL

      // Simple Field (value)
      Double value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readDouble("", 64);

      return new PlcLREAL(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.RECORD)) { // List
      // Array field (value)
      // Length array
      long _valueLength = size;
      long valueEndPos = readBuffer.getPos() + _valueLength;
      List<PlcValue> value = new LinkedList<>();
      while (readBuffer.getPos() < valueEndPos) {
        value.add(new PlcSINT(/*TODO: migrate me*//*TODO: migrate me*/ readBuffer.readByte("")));
      }

      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.OCTET_STRING)) { // STRING

      // Simple Field (value)
      String value = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readString("", size, "UTF-8");

      return new PlcSTRING(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.VISIBLE_STRING)) { // STRING

      // Simple Field (value)
      String value = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readString("", size, "UTF-8");

      return new PlcSTRING(value);
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNICODE_STRING)) { // STRING

      // Simple Field (value)
      String value = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readString("", (size) / (8), "UTF-8");

      return new PlcSTRING(value);
    }
    return null;
  }

  public static void staticSerialize(
      WriteBuffer writeBuffer, PlcValue _value, CANOpenDataType dataType, Integer size)
      throws SerializationException {
    staticSerialize(writeBuffer, _value, dataType, size, ByteOrder.BIG_ENDIAN);
  }

  public static void staticSerialize(
      WriteBuffer writeBuffer,
      PlcValue _value,
      CANOpenDataType dataType,
      Integer size,
      ByteOrder byteOrder)
      throws SerializationException {
    if (EvaluationHelper.equals(dataType, CANOpenDataType.BOOLEAN)) { // BOOL
      // Simple Field (value)
      boolean value = (boolean) _value.getBoolean();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeBit("", (boolean) (value));
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED8)) { // USINT
      // Simple Field (value)
      short value = (short) _value.getShort();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedShort("", 8, ((Number) (value)).shortValue());
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED16)) { // UINT
      // Simple Field (value)
      int value = (int) _value.getInt();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedInt("", 16, ((Number) (value)).intValue());
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED24)) { // UDINT
      // Simple Field (value)
      long value = (long) _value.getLong();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedLong("", 24, ((Number) (value)).longValue());
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED32)) { // UDINT
      // Simple Field (value)
      long value = (long) _value.getLong();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedLong("", 32, ((Number) (value)).longValue());
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED40)) { // ULINT
      // Simple Field (value)
      BigInteger value = (BigInteger) _value.getBigInteger();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedBigInteger("", 40, (BigInteger) (value));
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED48)) { // ULINT
      // Simple Field (value)
      BigInteger value = (BigInteger) _value.getBigInteger();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedBigInteger("", 48, (BigInteger) (value));
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED56)) { // ULINT
      // Simple Field (value)
      BigInteger value = (BigInteger) _value.getBigInteger();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedBigInteger("", 56, (BigInteger) (value));
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED64)) { // ULINT
      // Simple Field (value)
      BigInteger value = (BigInteger) _value.getBigInteger();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedBigInteger("", 64, (BigInteger) (value));
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER8)) { // SINT
      // Simple Field (value)
      byte value = (byte) _value.getByte();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeSignedByte("", 8, ((Number) (value)).byteValue());
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER16)) { // INT
      // Simple Field (value)
      short value = (short) _value.getShort();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeShort("", 16, ((Number) (value)).shortValue());
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER24)) { // DINT
      // Simple Field (value)
      int value = (int) _value.getInt();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeInt("", 24, ((Number) (value)).intValue());
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER32)) { // DINT
      // Simple Field (value)
      int value = (int) _value.getInt();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeInt("", 32, ((Number) (value)).intValue());
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER40)) { // LINT
      // Simple Field (value)
      long value = (long) _value.getLong();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeLong("", 40, ((Number) (value)).longValue());
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER48)) { // LINT
      // Simple Field (value)
      long value = (long) _value.getLong();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeLong("", 48, ((Number) (value)).longValue());
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER56)) { // LINT
      // Simple Field (value)
      long value = (long) _value.getLong();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeLong("", 56, ((Number) (value)).longValue());
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER64)) { // LINT
      // Simple Field (value)
      long value = (long) _value.getLong();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeLong("", 64, ((Number) (value)).longValue());
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.REAL32)) { // REAL
      // Simple Field (value)
      float value = (float) _value.getFloat();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeFloat("", 32, (value));
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.REAL64)) { // LREAL
      // Simple Field (value)
      double value = (double) _value.getDouble();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeDouble("", 64, (value));
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.RECORD)) { // List
      PlcList values = (PlcList) _value;

      for (PlcValue val : ((List<PlcValue>) values.getList())) {
        byte[] value = (byte[]) val.getRaw();
        writeBuffer.writeByteArray("", value);
      }

    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.OCTET_STRING)) { // STRING
      // Simple Field (value)
      String value = (String) _value.getString();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeString("", size, "UTF-8", (String) (value));
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.VISIBLE_STRING)) { // STRING
      // Simple Field (value)
      String value = (String) _value.getString();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeString("", size, "UTF-8", (String) (value));
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNICODE_STRING)) { // STRING
      // Simple Field (value)
      String value = (String) _value.getString();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeString("", (size) / (8), "UTF-8", (String) (value));
    }
  }

  public static int getLengthInBytes(PlcValue _value, CANOpenDataType dataType, Integer size) {
    return (int) Math.ceil((float) getLengthInBits(_value, dataType, size) / 8.0);
  }

  public static int getLengthInBits(PlcValue _value, CANOpenDataType dataType, Integer size) {
    int sizeInBits = 0;
    if (EvaluationHelper.equals(dataType, CANOpenDataType.BOOLEAN)) { // BOOL
      // Simple Field (value)
      sizeInBits += 1;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED8)) { // USINT
      // Simple Field (value)
      sizeInBits += 8;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED16)) { // UINT
      // Simple Field (value)
      sizeInBits += 16;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED24)) { // UDINT
      // Simple Field (value)
      sizeInBits += 24;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED32)) { // UDINT
      // Simple Field (value)
      sizeInBits += 32;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED40)) { // ULINT
      // Simple Field (value)
      sizeInBits += 40;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED48)) { // ULINT
      // Simple Field (value)
      sizeInBits += 48;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED56)) { // ULINT
      // Simple Field (value)
      sizeInBits += 56;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNSIGNED64)) { // ULINT
      // Simple Field (value)
      sizeInBits += 64;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER8)) { // SINT
      // Simple Field (value)
      sizeInBits += 8;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER16)) { // INT
      // Simple Field (value)
      sizeInBits += 16;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER24)) { // DINT
      // Simple Field (value)
      sizeInBits += 24;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER32)) { // DINT
      // Simple Field (value)
      sizeInBits += 32;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER40)) { // LINT
      // Simple Field (value)
      sizeInBits += 40;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER48)) { // LINT
      // Simple Field (value)
      sizeInBits += 48;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER56)) { // LINT
      // Simple Field (value)
      sizeInBits += 56;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.INTEGER64)) { // LINT
      // Simple Field (value)
      sizeInBits += 64;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.REAL32)) { // REAL
      // Simple Field (value)
      sizeInBits += 32;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.REAL64)) { // LREAL
      // Simple Field (value)
      sizeInBits += 64;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.RECORD)) { // List
      PlcList values = (PlcList) _value;
      sizeInBits += values.getList().size() * 8;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.OCTET_STRING)) { // STRING
      // Simple Field (value)
      sizeInBits += -1;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.VISIBLE_STRING)) { // STRING
      // Simple Field (value)
      sizeInBits += -1;
    } else if (EvaluationHelper.equals(dataType, CANOpenDataType.UNICODE_STRING)) { // STRING
      // Simple Field (value)
      sizeInBits += -1;
    }
    return sizeInBits;
  }
}
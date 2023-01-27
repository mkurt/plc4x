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
package org.apache.plc4x.java.test.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class ArrayTypeTest implements Message {

  // Properties.
  protected final List<Boolean> bitField;
  protected final byte[] byteField;
  protected final List<Byte> intField;
  protected final List<Short> uintField;
  protected final List<Float> floatField;
  protected final List<Double> doubleField;
  protected final List<String> stringField;

  public ArrayTypeTest(
      List<Boolean> bitField,
      byte[] byteField,
      List<Byte> intField,
      List<Short> uintField,
      List<Float> floatField,
      List<Double> doubleField,
      List<String> stringField) {
    super();
    this.bitField = bitField;
    this.byteField = byteField;
    this.intField = intField;
    this.uintField = uintField;
    this.floatField = floatField;
    this.doubleField = doubleField;
    this.stringField = stringField;
  }

  public List<Boolean> getBitField() {
    return bitField;
  }

  public byte[] getByteField() {
    return byteField;
  }

  public List<Byte> getIntField() {
    return intField;
  }

  public List<Short> getUintField() {
    return uintField;
  }

  public List<Float> getFloatField() {
    return floatField;
  }

  public List<Double> getDoubleField() {
    return doubleField;
  }

  public List<String> getStringField() {
    return stringField;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ArrayTypeTest");

    // Array Field (bitField)
    writeSimpleTypeArrayField("bitField", bitField, writeBoolean(writeBuffer));

    // Array Field (byteField)
    writeByteArrayField("byteField", byteField, writeByteArray(writeBuffer, 8));

    // Array Field (intField)
    writeSimpleTypeArrayField("intField", intField, writeSignedByte(writeBuffer, 8));

    // Array Field (uintField)
    writeSimpleTypeArrayField("uintField", uintField, writeUnsignedShort(writeBuffer, 8));

    // Array Field (floatField)
    writeSimpleTypeArrayField("floatField", floatField, writeFloat(writeBuffer, 32));

    // Array Field (doubleField)
    writeSimpleTypeArrayField("doubleField", doubleField, writeDouble(writeBuffer, 64));

    // Array Field (stringField)
    writeSimpleTypeArrayField("stringField", stringField, writeString(writeBuffer, 8));

    writeBuffer.popContext("ArrayTypeTest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ArrayTypeTest _value = this;

    // Array field
    if (bitField != null) {
      lengthInBits += 1 * bitField.size();
    }

    // Array field
    if (byteField != null) {
      lengthInBits += 8 * byteField.length;
    }

    // Array field
    if (intField != null) {
      lengthInBits += 8 * intField.size();
    }

    // Array field
    if (uintField != null) {
      lengthInBits += 8 * uintField.size();
    }

    // Array field
    if (floatField != null) {
      lengthInBits += 32 * floatField.size();
    }

    // Array field
    if (doubleField != null) {
      lengthInBits += 64 * doubleField.size();
    }

    // Array field
    if (stringField != null) {
      lengthInBits += 8 * stringField.size();
    }

    return lengthInBits;
  }

  public static ArrayTypeTest staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static ArrayTypeTest staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("ArrayTypeTest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    List<Boolean> bitField = readCountArrayField("bitField", readBoolean(readBuffer), 5);

    byte[] byteField = readBuffer.readByteArray("byteField", Math.toIntExact(5));

    List<Byte> intField = readCountArrayField("intField", readSignedByte(readBuffer, 8), 5);

    List<Short> uintField = readCountArrayField("uintField", readUnsignedShort(readBuffer, 8), 5);

    List<Float> floatField = readCountArrayField("floatField", readFloat(readBuffer, 32), 5);

    List<Double> doubleField = readCountArrayField("doubleField", readDouble(readBuffer, 64), 5);

    List<String> stringField = readCountArrayField("stringField", readString(readBuffer, 8), 5);

    readBuffer.closeContext("ArrayTypeTest");
    // Create the instance
    ArrayTypeTest _arrayTypeTest;
    _arrayTypeTest =
        new ArrayTypeTest(
            bitField, byteField, intField, uintField, floatField, doubleField, stringField);
    return _arrayTypeTest;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ArrayTypeTest)) {
      return false;
    }
    ArrayTypeTest that = (ArrayTypeTest) o;
    return (getBitField() == that.getBitField())
        && (getByteField() == that.getByteField())
        && (getIntField() == that.getIntField())
        && (getUintField() == that.getUintField())
        && (getFloatField() == that.getFloatField())
        && (getDoubleField() == that.getDoubleField())
        && (getStringField() == that.getStringField())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getBitField(),
        getByteField(),
        getIntField(),
        getUintField(),
        getFloatField(),
        getDoubleField(),
        getStringField());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
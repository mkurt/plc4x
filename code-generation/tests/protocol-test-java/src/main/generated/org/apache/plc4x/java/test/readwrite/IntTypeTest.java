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

public class IntTypeTest implements Message {

  // Properties.
  protected final byte ThreeField;
  protected final byte ByteField;
  protected final short WordField;
  protected final int WordPlusByteField;
  protected final int DoubleIntField;
  protected final long QuadIntField;

  public IntTypeTest(
      byte ThreeField,
      byte ByteField,
      short WordField,
      int WordPlusByteField,
      int DoubleIntField,
      long QuadIntField) {
    super();
    this.ThreeField = ThreeField;
    this.ByteField = ByteField;
    this.WordField = WordField;
    this.WordPlusByteField = WordPlusByteField;
    this.DoubleIntField = DoubleIntField;
    this.QuadIntField = QuadIntField;
  }

  public byte getThreeField() {
    return ThreeField;
  }

  public byte getByteField() {
    return ByteField;
  }

  public short getWordField() {
    return WordField;
  }

  public int getWordPlusByteField() {
    return WordPlusByteField;
  }

  public int getDoubleIntField() {
    return DoubleIntField;
  }

  public long getQuadIntField() {
    return QuadIntField;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("IntTypeTest");

    // Simple Field (ThreeField)
    writeSimpleField("ThreeField", ThreeField, writeSignedByte(writeBuffer, 3));

    // Simple Field (ByteField)
    writeSimpleField("ByteField", ByteField, writeSignedByte(writeBuffer, 8));

    // Simple Field (WordField)
    writeSimpleField("WordField", WordField, writeSignedShort(writeBuffer, 16));

    // Simple Field (WordPlusByteField)
    writeSimpleField("WordPlusByteField", WordPlusByteField, writeSignedInt(writeBuffer, 24));

    // Simple Field (DoubleIntField)
    writeSimpleField("DoubleIntField", DoubleIntField, writeSignedInt(writeBuffer, 32));

    // Simple Field (QuadIntField)
    writeSimpleField("QuadIntField", QuadIntField, writeSignedLong(writeBuffer, 64));

    writeBuffer.popContext("IntTypeTest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    IntTypeTest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (ThreeField)
    lengthInBits += 3;

    // Simple field (ByteField)
    lengthInBits += 8;

    // Simple field (WordField)
    lengthInBits += 16;

    // Simple field (WordPlusByteField)
    lengthInBits += 24;

    // Simple field (DoubleIntField)
    lengthInBits += 32;

    // Simple field (QuadIntField)
    lengthInBits += 64;

    return lengthInBits;
  }

  public static IntTypeTest staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static IntTypeTest staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("IntTypeTest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte ThreeField = readSimpleField("ThreeField", readSignedByte(readBuffer, 3));

    byte ByteField = readSimpleField("ByteField", readSignedByte(readBuffer, 8));

    short WordField = readSimpleField("WordField", readSignedShort(readBuffer, 16));

    int WordPlusByteField = readSimpleField("WordPlusByteField", readSignedInt(readBuffer, 24));

    int DoubleIntField = readSimpleField("DoubleIntField", readSignedInt(readBuffer, 32));

    long QuadIntField = readSimpleField("QuadIntField", readSignedLong(readBuffer, 64));

    readBuffer.closeContext("IntTypeTest");
    // Create the instance
    IntTypeTest _intTypeTest;
    _intTypeTest =
        new IntTypeTest(
            ThreeField, ByteField, WordField, WordPlusByteField, DoubleIntField, QuadIntField);
    return _intTypeTest;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof IntTypeTest)) {
      return false;
    }
    IntTypeTest that = (IntTypeTest) o;
    return (getThreeField() == that.getThreeField())
        && (getByteField() == that.getByteField())
        && (getWordField() == that.getWordField())
        && (getWordPlusByteField() == that.getWordPlusByteField())
        && (getDoubleIntField() == that.getDoubleIntField())
        && (getQuadIntField() == that.getQuadIntField())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getThreeField(),
        getByteField(),
        getWordField(),
        getWordPlusByteField(),
        getDoubleIntField(),
        getQuadIntField());
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

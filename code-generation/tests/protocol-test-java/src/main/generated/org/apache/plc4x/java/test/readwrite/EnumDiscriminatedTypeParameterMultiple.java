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

public abstract class EnumDiscriminatedTypeParameterMultiple implements Message {

  // Abstract accessors for discriminator values.
  public abstract EnumType getDiscr1();

  public abstract EnumTypeInt getDiscr2();

  public EnumDiscriminatedTypeParameterMultiple() {
    super();
  }

  protected abstract void serializeEnumDiscriminatedTypeParameterMultipleChild(
      WriteBuffer writeBuffer) throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("EnumDiscriminatedTypeParameterMultiple");

    // Switch field (Serialize the sub-type)
    serializeEnumDiscriminatedTypeParameterMultipleChild(writeBuffer);

    writeBuffer.popContext("EnumDiscriminatedTypeParameterMultiple");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    EnumDiscriminatedTypeParameterMultiple _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static EnumDiscriminatedTypeParameterMultiple staticParse(
      ReadBuffer readBuffer, Object... args) throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 2)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 2, but got " + args.length);
    }
    EnumType discr1;
    if (args[0] instanceof EnumType) {
      discr1 = (EnumType) args[0];
    } else if (args[0] instanceof String) {
      discr1 = EnumType.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type EnumType or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    EnumTypeInt discr2;
    if (args[1] instanceof EnumTypeInt) {
      discr2 = (EnumTypeInt) args[1];
    } else if (args[1] instanceof String) {
      discr2 = EnumTypeInt.valueOf((String) args[1]);
    } else {
      throw new PlcRuntimeException(
          "Argument 1 expected to be of type EnumTypeInt or a string which is parseable but was "
              + args[1].getClass().getName());
    }
    return staticParse(readBuffer, discr1, discr2);
  }

  public static EnumDiscriminatedTypeParameterMultiple staticParse(
      ReadBuffer readBuffer, EnumType discr1, EnumTypeInt discr2) throws ParseException {
    readBuffer.pullContext("EnumDiscriminatedTypeParameterMultiple");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    EnumDiscriminatedTypeParameterMultipleBuilder builder = null;
    if (EvaluationHelper.equals(discr1, EnumType.BOOL)
        && EvaluationHelper.equals(discr2, EnumTypeInt.BOOLINT)) {
      builder =
          EnumDiscriminatedTypeParameterMultipleA
              .staticParseEnumDiscriminatedTypeParameterMultipleBuilder(readBuffer, discr1, discr2);
    } else if (EvaluationHelper.equals(discr1, EnumType.UINT)
        && EvaluationHelper.equals(discr2, EnumTypeInt.UINTINT)) {
      builder =
          EnumDiscriminatedTypeParameterMultipleB
              .staticParseEnumDiscriminatedTypeParameterMultipleBuilder(readBuffer, discr1, discr2);
    } else if (EvaluationHelper.equals(discr1, EnumType.INT)
        && EvaluationHelper.equals(discr2, EnumTypeInt.INTINT)) {
      builder =
          EnumDiscriminatedTypeParameterMultipleC
              .staticParseEnumDiscriminatedTypeParameterMultipleBuilder(readBuffer, discr1, discr2);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "discr1="
              + discr1
              + " "
              + "discr2="
              + discr2
              + "]");
    }

    readBuffer.closeContext("EnumDiscriminatedTypeParameterMultiple");
    // Create the instance
    EnumDiscriminatedTypeParameterMultiple _enumDiscriminatedTypeParameterMultiple =
        builder.build();
    return _enumDiscriminatedTypeParameterMultiple;
  }

  public interface EnumDiscriminatedTypeParameterMultipleBuilder {
    EnumDiscriminatedTypeParameterMultiple build();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof EnumDiscriminatedTypeParameterMultiple)) {
      return false;
    }
    EnumDiscriminatedTypeParameterMultiple that = (EnumDiscriminatedTypeParameterMultiple) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
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

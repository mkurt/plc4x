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
package org.apache.plc4x.java.cbus.readwrite;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum AirConditioningCommandTypeContainer {
  AirConditioningCommandSetZoneGroupOff(
      (short) 0x01, (short) 1, AirConditioningCommandType.SET_ZONE_GROUP_OFF),
  AirConditioningCommandZoneHvacPlantStatus(
      (short) 0x05, (short) 5, AirConditioningCommandType.ZONE_HVAC_PLANT_STATUS),
  AirConditioningCommandZoneHumidityPlantStatus(
      (short) 0x0D, (short) 5, AirConditioningCommandType.ZONE_HUMIDITY_PLANT_STATUS),
  AirConditioningCommandZoneTemperature(
      (short) 0x15, (short) 5, AirConditioningCommandType.ZONE_TEMPERATURE),
  AirConditioningCommandZoneHumidity(
      (short) 0x1D, (short) 5, AirConditioningCommandType.ZONE_HUMIDITY),
  AirConditioningCommandRefresh((short) 0x21, (short) 1, AirConditioningCommandType.REFRESH),
  AirConditioningCommandSetZoneHvacMode(
      (short) 0x2F, (short) 7, AirConditioningCommandType.SET_ZONE_HVAC_MODE),
  AirConditioningCommandSetPlantHvacLevel(
      (short) 0x36, (short) 6, AirConditioningCommandType.SET_PLANT_HVAC_LEVEL),
  AirConditioningCommandSetZoneHumidityMode(
      (short) 0x47, (short) 7, AirConditioningCommandType.SET_ZONE_HUMIDITY_MODE),
  AirConditioningCommandSetPlantHumidityLevel(
      (short) 0x4E, (short) 6, AirConditioningCommandType.SET_PLANT_HUMIDITY_LEVEL),
  AirConditioningCommandSetHvacUpperGuardLimit(
      (short) 0x55, (short) 5, AirConditioningCommandType.SET_HVAC_UPPER_GUARD_LIMIT),
  AirConditioningCommandSetHvacLowerGuardLimit(
      (short) 0x5D, (short) 5, AirConditioningCommandType.SET_HVAC_LOWER_GUARD_LIMIT),
  AirConditioningCommandSetHvacSetbackLimit(
      (short) 0x65, (short) 5, AirConditioningCommandType.SET_HVAC_SETBACK_LIMIT),
  AirConditioningCommandSetHumidityUpperGuardLimit(
      (short) 0x6D, (short) 5, AirConditioningCommandType.SET_HUMIDITY_UPPER_GUARD_LIMIT),
  AirConditioningCommandSetHumidityLowerGuardLimit(
      (short) 0x75, (short) 5, AirConditioningCommandType.SET_HUMIDITY_LOWER_GUARD_LIMIT),
  AirConditioningCommandSetZoneGroupOn(
      (short) 0x79, (short) 1, AirConditioningCommandType.SET_ZONE_GROUP_ON),
  AirConditioningCommandSetHumiditySetbackLimit(
      (short) 0x7D, (short) 5, AirConditioningCommandType.SET_HUMIDITY_SETBACK_LIMIT),
  AirConditioningCommandHvacScheduleEntry(
      (short) 0x89, (short) 9, AirConditioningCommandType.HVAC_SCHEDULE_ENTRY),
  AirConditioningCommandHumidityScheduleEntry(
      (short) 0xA9, (short) 9, AirConditioningCommandType.HUMIDITY_SCHEDULE_ENTRY);
  private static final Map<Short, AirConditioningCommandTypeContainer> map;

  static {
    map = new HashMap<>();
    for (AirConditioningCommandTypeContainer value : AirConditioningCommandTypeContainer.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;
  private final short numBytes;
  private final AirConditioningCommandType commandType;

  AirConditioningCommandTypeContainer(
      short value, short numBytes, AirConditioningCommandType commandType) {
    this.value = value;
    this.numBytes = numBytes;
    this.commandType = commandType;
  }

  public short getValue() {
    return value;
  }

  public short getNumBytes() {
    return numBytes;
  }

  public static AirConditioningCommandTypeContainer firstEnumForFieldNumBytes(short fieldValue) {
    for (AirConditioningCommandTypeContainer _val : AirConditioningCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<AirConditioningCommandTypeContainer> enumsForFieldNumBytes(short fieldValue) {
    List<AirConditioningCommandTypeContainer> _values = new ArrayList<>();
    for (AirConditioningCommandTypeContainer _val : AirConditioningCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public AirConditioningCommandType getCommandType() {
    return commandType;
  }

  public static AirConditioningCommandTypeContainer firstEnumForFieldCommandType(
      AirConditioningCommandType fieldValue) {
    for (AirConditioningCommandTypeContainer _val : AirConditioningCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<AirConditioningCommandTypeContainer> enumsForFieldCommandType(
      AirConditioningCommandType fieldValue) {
    List<AirConditioningCommandTypeContainer> _values = new ArrayList<>();
    for (AirConditioningCommandTypeContainer _val : AirConditioningCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static AirConditioningCommandTypeContainer enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}

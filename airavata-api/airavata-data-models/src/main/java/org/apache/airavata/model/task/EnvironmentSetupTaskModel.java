/**
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * Autogenerated by Thrift Compiler (0.9.3)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
package org.apache.airavata.model.task;

import org.apache.thrift.scheme.IScheme;
import org.apache.thrift.scheme.SchemeFactory;
import org.apache.thrift.scheme.StandardScheme;

import org.apache.thrift.scheme.TupleScheme;
import org.apache.thrift.protocol.TTupleProtocol;
import org.apache.thrift.protocol.TProtocolException;
import org.apache.thrift.EncodingUtils;
import org.apache.thrift.TException;
import org.apache.thrift.async.AsyncMethodCallback;
import org.apache.thrift.server.AbstractNonblockingServer.*;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;
import java.util.EnumMap;
import java.util.Set;
import java.util.HashSet;
import java.util.EnumSet;
import java.util.Collections;
import java.util.BitSet;
import java.nio.ByteBuffer;
import java.util.Arrays;
import javax.annotation.Generated;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@SuppressWarnings({"cast", "rawtypes", "serial", "unchecked"})
/**
 * EnvironmentSetupTaskModel: A structure holding the environment creation task details
 * 
 */
@Generated(value = "Autogenerated by Thrift Compiler (0.9.3)", date = "2016-02-01")
public class EnvironmentSetupTaskModel implements org.apache.thrift.TBase<EnvironmentSetupTaskModel, EnvironmentSetupTaskModel._Fields>, java.io.Serializable, Cloneable, Comparable<EnvironmentSetupTaskModel> {
  private static final org.apache.thrift.protocol.TStruct STRUCT_DESC = new org.apache.thrift.protocol.TStruct("EnvironmentSetupTaskModel");

  private static final org.apache.thrift.protocol.TField LOCATION_FIELD_DESC = new org.apache.thrift.protocol.TField("location", org.apache.thrift.protocol.TType.STRING, (short)1);
  private static final org.apache.thrift.protocol.TField PROTOCOL_FIELD_DESC = new org.apache.thrift.protocol.TField("protocol", org.apache.thrift.protocol.TType.I32, (short)2);

  private static final Map<Class<? extends IScheme>, SchemeFactory> schemes = new HashMap<Class<? extends IScheme>, SchemeFactory>();
  static {
    schemes.put(StandardScheme.class, new EnvironmentSetupTaskModelStandardSchemeFactory());
    schemes.put(TupleScheme.class, new EnvironmentSetupTaskModelTupleSchemeFactory());
  }

  private String location; // required
  private org.apache.airavata.model.data.movement.SecurityProtocol protocol; // required

  /** The set of fields this struct contains, along with convenience methods for finding and manipulating them. */
  public enum _Fields implements org.apache.thrift.TFieldIdEnum {
    LOCATION((short)1, "location"),
    /**
     * 
     * @see org.apache.airavata.model.data.movement.SecurityProtocol
     */
    PROTOCOL((short)2, "protocol");

    private static final Map<String, _Fields> byName = new HashMap<String, _Fields>();

    static {
      for (_Fields field : EnumSet.allOf(_Fields.class)) {
        byName.put(field.getFieldName(), field);
      }
    }

    /**
     * Find the _Fields constant that matches fieldId, or null if its not found.
     */
    public static _Fields findByThriftId(int fieldId) {
      switch(fieldId) {
        case 1: // LOCATION
          return LOCATION;
        case 2: // PROTOCOL
          return PROTOCOL;
        default:
          return null;
      }
    }

    /**
     * Find the _Fields constant that matches fieldId, throwing an exception
     * if it is not found.
     */
    public static _Fields findByThriftIdOrThrow(int fieldId) {
      _Fields fields = findByThriftId(fieldId);
      if (fields == null) throw new IllegalArgumentException("Field " + fieldId + " doesn't exist!");
      return fields;
    }

    /**
     * Find the _Fields constant that matches name, or null if its not found.
     */
    public static _Fields findByName(String name) {
      return byName.get(name);
    }

    private final short _thriftId;
    private final String _fieldName;

    _Fields(short thriftId, String fieldName) {
      _thriftId = thriftId;
      _fieldName = fieldName;
    }

    public short getThriftFieldId() {
      return _thriftId;
    }

    public String getFieldName() {
      return _fieldName;
    }
  }

  // isset id assignments
  public static final Map<_Fields, org.apache.thrift.meta_data.FieldMetaData> metaDataMap;
  static {
    Map<_Fields, org.apache.thrift.meta_data.FieldMetaData> tmpMap = new EnumMap<_Fields, org.apache.thrift.meta_data.FieldMetaData>(_Fields.class);
    tmpMap.put(_Fields.LOCATION, new org.apache.thrift.meta_data.FieldMetaData("location", org.apache.thrift.TFieldRequirementType.REQUIRED, 
        new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.STRING)));
    tmpMap.put(_Fields.PROTOCOL, new org.apache.thrift.meta_data.FieldMetaData("protocol", org.apache.thrift.TFieldRequirementType.REQUIRED, 
        new org.apache.thrift.meta_data.EnumMetaData(org.apache.thrift.protocol.TType.ENUM, org.apache.airavata.model.data.movement.SecurityProtocol.class)));
    metaDataMap = Collections.unmodifiableMap(tmpMap);
    org.apache.thrift.meta_data.FieldMetaData.addStructMetaDataMap(EnvironmentSetupTaskModel.class, metaDataMap);
  }

  public EnvironmentSetupTaskModel() {
  }

  public EnvironmentSetupTaskModel(
    String location,
    org.apache.airavata.model.data.movement.SecurityProtocol protocol)
  {
    this();
    this.location = location;
    this.protocol = protocol;
  }

  /**
   * Performs a deep copy on <i>other</i>.
   */
  public EnvironmentSetupTaskModel(EnvironmentSetupTaskModel other) {
    if (other.isSetLocation()) {
      this.location = other.location;
    }
    if (other.isSetProtocol()) {
      this.protocol = other.protocol;
    }
  }

  public EnvironmentSetupTaskModel deepCopy() {
    return new EnvironmentSetupTaskModel(this);
  }

  @Override
  public void clear() {
    this.location = null;
    this.protocol = null;
  }

  public String getLocation() {
    return this.location;
  }

  public void setLocation(String location) {
    this.location = location;
  }

  public void unsetLocation() {
    this.location = null;
  }

  /** Returns true if field location is set (has been assigned a value) and false otherwise */
  public boolean isSetLocation() {
    return this.location != null;
  }

  public void setLocationIsSet(boolean value) {
    if (!value) {
      this.location = null;
    }
  }

  /**
   * 
   * @see org.apache.airavata.model.data.movement.SecurityProtocol
   */
  public org.apache.airavata.model.data.movement.SecurityProtocol getProtocol() {
    return this.protocol;
  }

  /**
   * 
   * @see org.apache.airavata.model.data.movement.SecurityProtocol
   */
  public void setProtocol(org.apache.airavata.model.data.movement.SecurityProtocol protocol) {
    this.protocol = protocol;
  }

  public void unsetProtocol() {
    this.protocol = null;
  }

  /** Returns true if field protocol is set (has been assigned a value) and false otherwise */
  public boolean isSetProtocol() {
    return this.protocol != null;
  }

  public void setProtocolIsSet(boolean value) {
    if (!value) {
      this.protocol = null;
    }
  }

  public void setFieldValue(_Fields field, Object value) {
    switch (field) {
    case LOCATION:
      if (value == null) {
        unsetLocation();
      } else {
        setLocation((String)value);
      }
      break;

    case PROTOCOL:
      if (value == null) {
        unsetProtocol();
      } else {
        setProtocol((org.apache.airavata.model.data.movement.SecurityProtocol)value);
      }
      break;

    }
  }

  public Object getFieldValue(_Fields field) {
    switch (field) {
    case LOCATION:
      return getLocation();

    case PROTOCOL:
      return getProtocol();

    }
    throw new IllegalStateException();
  }

  /** Returns true if field corresponding to fieldID is set (has been assigned a value) and false otherwise */
  public boolean isSet(_Fields field) {
    if (field == null) {
      throw new IllegalArgumentException();
    }

    switch (field) {
    case LOCATION:
      return isSetLocation();
    case PROTOCOL:
      return isSetProtocol();
    }
    throw new IllegalStateException();
  }

  @Override
  public boolean equals(Object that) {
    if (that == null)
      return false;
    if (that instanceof EnvironmentSetupTaskModel)
      return this.equals((EnvironmentSetupTaskModel)that);
    return false;
  }

  public boolean equals(EnvironmentSetupTaskModel that) {
    if (that == null)
      return false;

    boolean this_present_location = true && this.isSetLocation();
    boolean that_present_location = true && that.isSetLocation();
    if (this_present_location || that_present_location) {
      if (!(this_present_location && that_present_location))
        return false;
      if (!this.location.equals(that.location))
        return false;
    }

    boolean this_present_protocol = true && this.isSetProtocol();
    boolean that_present_protocol = true && that.isSetProtocol();
    if (this_present_protocol || that_present_protocol) {
      if (!(this_present_protocol && that_present_protocol))
        return false;
      if (!this.protocol.equals(that.protocol))
        return false;
    }

    return true;
  }

  @Override
  public int hashCode() {
    List<Object> list = new ArrayList<Object>();

    boolean present_location = true && (isSetLocation());
    list.add(present_location);
    if (present_location)
      list.add(location);

    boolean present_protocol = true && (isSetProtocol());
    list.add(present_protocol);
    if (present_protocol)
      list.add(protocol.getValue());

    return list.hashCode();
  }

  @Override
  public int compareTo(EnvironmentSetupTaskModel other) {
    if (!getClass().equals(other.getClass())) {
      return getClass().getName().compareTo(other.getClass().getName());
    }

    int lastComparison = 0;

    lastComparison = Boolean.valueOf(isSetLocation()).compareTo(other.isSetLocation());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetLocation()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.location, other.location);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    lastComparison = Boolean.valueOf(isSetProtocol()).compareTo(other.isSetProtocol());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetProtocol()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.protocol, other.protocol);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    return 0;
  }

  public _Fields fieldForId(int fieldId) {
    return _Fields.findByThriftId(fieldId);
  }

  public void read(org.apache.thrift.protocol.TProtocol iprot) throws org.apache.thrift.TException {
    schemes.get(iprot.getScheme()).getScheme().read(iprot, this);
  }

  public void write(org.apache.thrift.protocol.TProtocol oprot) throws org.apache.thrift.TException {
    schemes.get(oprot.getScheme()).getScheme().write(oprot, this);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder("EnvironmentSetupTaskModel(");
    boolean first = true;

    sb.append("location:");
    if (this.location == null) {
      sb.append("null");
    } else {
      sb.append(this.location);
    }
    first = false;
    if (!first) sb.append(", ");
    sb.append("protocol:");
    if (this.protocol == null) {
      sb.append("null");
    } else {
      sb.append(this.protocol);
    }
    first = false;
    sb.append(")");
    return sb.toString();
  }

  public void validate() throws org.apache.thrift.TException {
    // check for required fields
    if (!isSetLocation()) {
      throw new org.apache.thrift.protocol.TProtocolException("Required field 'location' is unset! Struct:" + toString());
    }

    if (!isSetProtocol()) {
      throw new org.apache.thrift.protocol.TProtocolException("Required field 'protocol' is unset! Struct:" + toString());
    }

    // check for sub-struct validity
  }

  private void writeObject(java.io.ObjectOutputStream out) throws java.io.IOException {
    try {
      write(new org.apache.thrift.protocol.TCompactProtocol(new org.apache.thrift.transport.TIOStreamTransport(out)));
    } catch (org.apache.thrift.TException te) {
      throw new java.io.IOException(te);
    }
  }

  private void readObject(java.io.ObjectInputStream in) throws java.io.IOException, ClassNotFoundException {
    try {
      read(new org.apache.thrift.protocol.TCompactProtocol(new org.apache.thrift.transport.TIOStreamTransport(in)));
    } catch (org.apache.thrift.TException te) {
      throw new java.io.IOException(te);
    }
  }

  private static class EnvironmentSetupTaskModelStandardSchemeFactory implements SchemeFactory {
    public EnvironmentSetupTaskModelStandardScheme getScheme() {
      return new EnvironmentSetupTaskModelStandardScheme();
    }
  }

  private static class EnvironmentSetupTaskModelStandardScheme extends StandardScheme<EnvironmentSetupTaskModel> {

    public void read(org.apache.thrift.protocol.TProtocol iprot, EnvironmentSetupTaskModel struct) throws org.apache.thrift.TException {
      org.apache.thrift.protocol.TField schemeField;
      iprot.readStructBegin();
      while (true)
      {
        schemeField = iprot.readFieldBegin();
        if (schemeField.type == org.apache.thrift.protocol.TType.STOP) { 
          break;
        }
        switch (schemeField.id) {
          case 1: // LOCATION
            if (schemeField.type == org.apache.thrift.protocol.TType.STRING) {
              struct.location = iprot.readString();
              struct.setLocationIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          case 2: // PROTOCOL
            if (schemeField.type == org.apache.thrift.protocol.TType.I32) {
              struct.protocol = org.apache.airavata.model.data.movement.SecurityProtocol.findByValue(iprot.readI32());
              struct.setProtocolIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          default:
            org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
        }
        iprot.readFieldEnd();
      }
      iprot.readStructEnd();
      struct.validate();
    }

    public void write(org.apache.thrift.protocol.TProtocol oprot, EnvironmentSetupTaskModel struct) throws org.apache.thrift.TException {
      struct.validate();

      oprot.writeStructBegin(STRUCT_DESC);
      if (struct.location != null) {
        oprot.writeFieldBegin(LOCATION_FIELD_DESC);
        oprot.writeString(struct.location);
        oprot.writeFieldEnd();
      }
      if (struct.protocol != null) {
        oprot.writeFieldBegin(PROTOCOL_FIELD_DESC);
        oprot.writeI32(struct.protocol.getValue());
        oprot.writeFieldEnd();
      }
      oprot.writeFieldStop();
      oprot.writeStructEnd();
    }

  }

  private static class EnvironmentSetupTaskModelTupleSchemeFactory implements SchemeFactory {
    public EnvironmentSetupTaskModelTupleScheme getScheme() {
      return new EnvironmentSetupTaskModelTupleScheme();
    }
  }

  private static class EnvironmentSetupTaskModelTupleScheme extends TupleScheme<EnvironmentSetupTaskModel> {

    @Override
    public void write(org.apache.thrift.protocol.TProtocol prot, EnvironmentSetupTaskModel struct) throws org.apache.thrift.TException {
      TTupleProtocol oprot = (TTupleProtocol) prot;
      oprot.writeString(struct.location);
      oprot.writeI32(struct.protocol.getValue());
    }

    @Override
    public void read(org.apache.thrift.protocol.TProtocol prot, EnvironmentSetupTaskModel struct) throws org.apache.thrift.TException {
      TTupleProtocol iprot = (TTupleProtocol) prot;
      struct.location = iprot.readString();
      struct.setLocationIsSet(true);
      struct.protocol = org.apache.airavata.model.data.movement.SecurityProtocol.findByValue(iprot.readI32());
      struct.setProtocolIsSet(true);
    }
  }

}


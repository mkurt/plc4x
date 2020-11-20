/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.apache.plc4x.java.opcuaserver;

import java.util.List;
import java.io.File;
import java.io.IOException;
import java.security.SecureRandom;

import org.apache.commons.codec.digest.DigestUtils;
import org.eclipse.jetty.util.security.Password;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonIgnore;
import org.slf4j.LoggerFactory;

public class User {

    @JsonProperty
    private String username;

    //TODO: These need to be hashed
    @JsonProperty
    private String password;

    @JsonProperty
    private String security;

    @JsonProperty
    private byte[] salt;

    @JsonProperty
    private static final SecureRandom randomGen = new SecureRandom();

    public User() {
        if (this.salt == null) {
            this.salt = generateSalt();
        }
    }

    private byte[] generateSalt() {
        byte[] salt = new byte[16];
        randomGen.nextBytes(salt);
        return salt;
    }

    public User(String username, String password, String security) {
        this.username = username;
        this.password = DigestUtils.sha256Hex(this.salt + ":" + password);
        this.security = security;
    }

    public boolean checkPassword(String password) {
        if (this.password.equals((DigestUtils.sha256Hex(this.salt + ":" + password)))) {
            return true;
        }
        return false;
    }

    public String getUsername() {
        return username;
    }

    public String getSecurity() {
        return security;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    @JsonIgnore
    public void setPassword(String password) {
        this.password = this.salt + ":" + DigestUtils.sha256Hex(password);
    }

    public void setSecurity(String security) {
        this.security = security;
    }




}
package com.daemoncoder.java_learning.spring_demo.dao.po;

import java.io.Serializable;

import lombok.Data;

@Data
public class User implements Serializable {

    private Long uid;

    private String uname;

    private String passwdHash;

    private String tel;

    private String email;

    private static final long serialVersionUID = 7341029133221707601L;

}

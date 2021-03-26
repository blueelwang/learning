package com.daemoncoder.java_learning.spring_demo.service;

import com.daemoncoder.java_learning.spring_demo.dao.UserMapper;
import com.daemoncoder.java_learning.spring_demo.dao.po.User;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class UserService {
    
    @Autowired
    private UserMapper userMapper;

    public String hello(String name) {
        return "hello, " + name + ".";
    }

    public User getUserById(Long uid) {
        return userMapper.getById(uid);
    }

}

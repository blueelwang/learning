package com.daemoncoder.java_learning.spring_demo.controller;

import com.daemoncoder.java_learning.spring_demo.dao.po.User;
import com.daemoncoder.java_learning.spring_demo.service.UserService;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping(value = "/user")
public class UserController {
    
    @Autowired
    private UserService userService;


    @GetMapping(value = "/hello")
    public String hello(@RequestParam String name) {
        return userService.hello(name);
    }

    @RequestMapping(value="/hi", method = RequestMethod.GET)
    public String hi(@RequestParam String name) {
        return "hi, " + name + ".";
    }

    @GetMapping(value="/echo/{name}")
    public String echo(@PathVariable String name) {
        return name;
    }

    @GetMapping(value="/id")
    public User getUserName(@RequestParam String id) {
        try {
            Long uid = Long.parseLong(id);
            User user = userService.getUserById(uid);
            return user;
        } catch (Exception e) {
            System.out.println(e);
            return null;
        }
    }

}

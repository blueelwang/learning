package com.daemoncoder.java_learning.spring_demo.dao;

import com.daemoncoder.java_learning.spring_demo.dao.po.User;

import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;


@Mapper
public interface UserMapper {

    @Select("select * from user where uid = #{id,jdbcType=BIGINT}")
    User getById(Long uid);

}

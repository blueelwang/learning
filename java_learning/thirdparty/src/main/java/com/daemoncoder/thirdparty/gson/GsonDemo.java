package com.daemoncoder.thirdparty.gson;

import java.util.HashMap;
import java.util.Map;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonSyntaxException;
import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class GsonDemo {
    
    public static void main(String[] args) {
        try {
            Gson gson = new Gson(); // 线程不安全
            
            // 通过@SerializedName注解映射json映射名字
            // create_time 和 createTime不会匹配
            // graduate_time 和 graduateTime 匹配
            String jsonStr = "{\"name\": \"zhang3\", \"age\": 18, \"weight\": 60.5, \"graduate_time\": 2016, \"company\": {\"name\": \"小米\", \"create_time\": 2010}}"; 
            Person person = gson.fromJson(jsonStr, Person.class);
            System.out.println(person);                 // name:zhang3, age:18, weight:60.5, graduate_time:2016, company_name:小米, company_create_time:0
            System.out.println(gson.toJson(person));    // {"name":"zhang3","age":18,"weight":60.5,"graduate_time":2016,"company":{"name":"小米","createTime":0}}

            // 多的字段被忽略，少的用默认值
            jsonStr = "{\"height\": 13, \"weight\": 60.5, \"company\": {\"name\": \"小米\", \"create_time\": 2010}}";
            person = gson.fromJson(jsonStr, Person.class);
            System.out.println(person);                 // name:null, age:0, weight:60.5, graduate_time:0, company_name:小米, company_create_time:0
            System.out.println(gson.toJson(person));    // {"age":0,"weight":60.5,"graduate_time":0,"company":{"name":"小米","createTime":0}}

            // 任意类型demo
            jsonStr = "{\"height\": 13, \"weight\": 60.5, \"company\": [{\"name\": \"小米\", \"create_time\": 2010}]}";
            Map<String, Object> map = new HashMap<String, Object>();
            map = gson.fromJson(jsonStr, map.getClass());
            // 数值类型都被解析成double（包括int, float）, 对象被解析成 LinkedTreeMap, 数组被解析成 ArrayList
            System.out.println(map);    // {weight=60.5, company={name=小米, create_time=2010.0}, height=13.0}

            exposeDemo();
        } catch (JsonSyntaxException e) {
            // 解析失败抛异常
            System.out.println(e);
        }
        
    }

    public static void exposeDemo() {
        // 加了 Expose 注解，只处理被注解的字段，创建gson实例的时候用下面这个方式，不然不生效
        Gson gson = new GsonBuilder().excludeFieldsWithoutExposeAnnotation().create();
        String jsonStr = "{\"name\": \"zhang3\", \"age\": 18, \"weight\": 60.5, \"company\": {\"name\": \"小米\", \"createTime\": 2010}}";
        Person person = gson.fromJson(jsonStr, Person.class);
        System.out.println(person);                 // name:zhang3, age:0, weight:0.0, graduate_time:0, company_name:null, company_create_time:2010
        System.out.println(gson.toJson(person));    // {"name":"zhang3","company":{"createTime":2010}}
    }

}

class Person {
    @Expose 
    String name;
    int age;
    float weight;
    @SerializedName("graduate_time")
    int graduateTime;
    @Expose
    Company company;

    public String toString() {
        return "name:" + name + ", age:" + age + ", weight:" + weight + ", graduate_time:" + graduateTime + ", company_name:" + company.name + ", company_create_time:" + company.createTime;
    }
}

class Company {
    String name;
    @Expose
    int createTime;
}

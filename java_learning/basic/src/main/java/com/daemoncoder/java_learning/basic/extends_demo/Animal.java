package com.daemoncoder.java_learning.basic.extends_demo;


public class Animal {
    public String type;

    public static void main(String[] args) {
        try {
            Animal animal = new Animal();
            Dog dog;
            Cat cat;

            System.out.println(animal instanceof Animal);   // true
            System.out.println(animal instanceof Dog);      // false
            System.out.println(animal instanceof Cat);      // false
            System.out.println("--------------------------");

            // Dog dog = (Dog)animal;  // java.lang.ClassCastException: com.daemoncoder.basic.extends_demo.Animal cannot be cast to com.daemoncoder.basic.extends_demo.Dog
            System.out.println("--------------------------");
    

            animal = new Dog();
            dog = (Dog)animal;
            System.out.println(dog instanceof Animal);      // true
            System.out.println(dog instanceof Dog);         // true
            System.out.println("--------------------------");
    
            cat = (Cat)animal;      // java.lang.ClassCastException: com.daemoncoder.basic.extends_demo.Dog cannot be cast to com.daemoncoder.basic.extends_demo.Cat
            System.out.println(cat instanceof Animal);
            System.out.println(cat instanceof Cat);
            System.out.println("--------------------------");
        } catch (Exception e) {
            System.out.println("！！！！！！！！！！！！！！！");
            System.out.println(e);
        }
        
    }
}

class Dog extends Animal {
    public String type = "dog";
    public String wang = "wang";
}

class Cat extends Animal {
    public String type = "cat";
    public String miao = "miao";
}

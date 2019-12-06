package basic;

public class Test {

    public static void main(String[] args) {
        System.out.println(add(1, 2));                  // 3
        System.out.println(add("daemoncoder", ".com")); // daemoncoder.com
    }

    public static int add(int a, int b) {
        return a + b;
    }
    
    public static String add(String a, String b) {
        return a + b;
    }

}

package com.company;

import iotdb.*;

import java.sql.SQLException;

public class Main {

    public static void main(String[] args) throws SQLException {
        System.out.println("链接参数");
        System.out.println(args[0]);
        System.out.println(args[1]);
        System.out.println(args[2]);
        System.out.println(args[3]);
        System.out.println(args[4]);
        System.out.println("----------------------------");
        IOTDB ioTDB = new IOTDB();
        String ip = args[0];
        Integer port = Integer.valueOf(args[1]);
        String username = args[2];
        String password = args[3];
        String path = args[4];
        ioTDB.connect(ip,port,username,password);
        try {
            ioTDB.executefile(path);
        }catch (SQLException e){
            System.out.println("失败了");
            System.out.println(e.getMessage());
            ioTDB.disconnect();
        }
	// write your code here
    }
}

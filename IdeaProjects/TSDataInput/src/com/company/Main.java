package com.company;

import iotdb.*;

import java.sql.SQLException;

public class Main {

    public static void main(String[] args) throws SQLException {
        IOTDB ioTDB = new IOTDB();
        String ip = "127.0.0.1";
        Integer port = 6667;
        ioTDB.connect(ip,port,"root","root");
        try {
            ioTDB.executefile("./iotdb.sql");
        }catch (SQLException e){
            System.out.println(e.getMessage());
            ioTDB.disconnect();
        }
	// write your code here
    }
}

package iotdb;

import java.io.*;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.sql.Statement;

/**
 * Created by hailan on 18/8/1.
 */
public class IOTDB {
    public Connection connection = null;
    public void connect(){
        try {
            Class.forName("cn.edu.tsinghua.iotdb.jdbc.TsfileDriver");
        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        }
        try {

            connection =  DriverManager.getConnection("jdbc:tsfile://127.0.0.1:6667/",
                    "root","root");
        } catch (SQLException e) {
            e.printStackTrace();
        }
    }

    public void disconnect() throws SQLException {
        connection.close();
    }

    public void connect(String ip,int port,String user,String password){
        try {
            Class.forName("cn.edu.tsinghua.iotdb.jdbc.TsfileDriver");
        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        }
        try {

            connection =  DriverManager.getConnection("jdbc:tsfile://"+ip+":"+port+"/",
                    "root","root");
        } catch (SQLException e) {
            e.printStackTrace();
        }
    }

    public void execute(String sql) throws SQLException {
        Statement statement = null;
        statement = connection.createStatement();
        statement.execute(sql);
    }
    public void executefile(String path) throws SQLException {

        File readFileSpice = new File(path);
        BufferedReader bReader = null;
        try {
            FileReader reader = new FileReader(readFileSpice);
            bReader = new BufferedReader(reader);
            String line = bReader.readLine();

            int count =0;
            while(line != null){
                System.out.println(line);
                count++;
                try{

                    execute(line);
                }catch (SQLException e){

                   System.out.println(e.getMessage());
                    System.out.println("SQL语句失败");
                }
                finally {

                    System.out.println("Have executed "+ Integer.toString(count)+" lines");
                }
                line = bReader.readLine();
            }
            System.out.println(count);
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }catch(IOException e){
            e.printStackTrace();
        }  finally {
            try {
                if (bReader != null) {
                    bReader.close();
                }
            } catch (IOException e) {
                e.printStackTrace();
            }
        }


    }
}

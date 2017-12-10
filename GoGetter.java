import java.io.*;
import java.util.*;

public class GoGetter {
  public static void main(String args[]) throws Exception {
    BufferedReader read = new BufferedReader(new FileReader("sample.txt"));

    String line;

    while((line=read.readLine())!=null) {
      if(line.contains("cannot find package")) {
        String goPackage = line.substring(line.indexOf("\"")+1, line.lastIndexOf("\""));

        System.out.println("go get "+goPackage);      
      }

    }
  }
}

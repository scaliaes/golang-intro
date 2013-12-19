package es.scalia.example;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;

import com.crittercism.app.Crittercism;

public class Example {
  // STARTREADFILE OMIT
  public static final String readFile(String filename) {
    StringBuilder sb = new StringBuilder();
    try {
      BufferedReader br = new BufferedReader(new FileReader(filename)); // HL

      int read = 0;
      char[] buffer = new char[1024];
      while( -1 != (read = br.read(buffer, 0, buffer.length)) ) { // HL
        sb.append(buffer, 0, read);
      }
      br.close();
    } catch (IOException e) { // HL
      Log.e("HTTP", "readStream", e);
      Crittercism.logHandledException(e);
      return null;
    }

    return sb.toString();
  }
  // STOPREADFILE OMIT
}

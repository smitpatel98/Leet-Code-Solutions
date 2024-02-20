import java.io.*;
import java.util.*;

class Result {

    /*
     * Complete the 'timeConversion' function below.
     *
     * The function is expected to return a STRING.
     * The function accepts STRING s as parameter.
     */

    public static String timeConversion(String s) {
        String newFormat = "";
        String meridian = s.substring(s.length() - 2); // Extract AM/PM
        String timeWithoutMeridian = s.substring(0, s.length() - 2); // Extract time without AM/PM

        String[] parts = timeWithoutMeridian.split(":"); // Split hours, minutes, seconds

        int hours = Integer.parseInt(parts[0]);

        // Handle cases for AM and PM
        if (meridian.equals("AM")) {
            if (hours == 12)
                hours = 0;
        } else { // PM
            if (hours != 12)
                hours += 12;
        }

        // Format hours part (ensure two digits)
        newFormat += String.format("%02d", hours) + ":";
        newFormat += parts[1] + ":"; // Minutes part
        newFormat += parts[2]; // Seconds part

        return newFormat;
    }

}

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String s = bufferedReader.readLine();

        String result = Result.timeConversion(s);

        bufferedWriter.write(result);
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}

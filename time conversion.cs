using System;
using System.IO;

class Result
{
    /*
     * Complete the 'timeConversion' function below.
     *
     * The function is expected to return a STRING.
     * The function accepts STRING s as parameter.
     */

    public static string timeConversion(string s)
    {
        string newFormat = "";
        string meridian = s.Substring(s.Length - 2); // Extract AM/PM
        string timeWithoutMeridian = s.Substring(0, s.Length - 2); // Extract time without AM/PM
        
        string[] parts = timeWithoutMeridian.Split(':'); // Split hours, minutes, seconds
        
        int hours = int.Parse(parts[0]);
        
        // Handle cases for AM and PM
        if (meridian == "AM")
        {
            if (hours == 12)
                hours = 0;
        }
        else // PM
        {
            if (hours != 12)
                hours += 12;
        }
        
        // Format hours part (ensure two digits)
        newFormat += hours.ToString("00") + ":";
        newFormat += parts[1] + ":"; // Minutes part
        newFormat += parts[2]; // Seconds part

        return newFormat;
    }
}

class Solution
{
    public static void Main(string[] args)
    {
        TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

        string s = Console.ReadLine();

        string result = Result.timeConversion(s);

        textWriter.WriteLine(result);

        textWriter.Flush();
        textWriter.Close();
    }
}

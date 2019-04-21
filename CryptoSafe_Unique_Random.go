package strings

import (
        "crypto/rand"
        "encoding/base32"
        "strconv"
        "time"
)

func CryptoSafe_Unique_Random_32 () (cryptoSafe_String string, err error) { /* This function generates a string that is unique and cryptographically safe. Any two strings generated by this function should never be the same. Strings it generates are 32 characters long.

        ** On success nil error is returned, else the error that occured is returned. */

        // This part generates a cryptosafe random string of 13 characters. { ...
        cryptoSafe_Bytes := make ([]byte, 13)
        _, err1 := rand.Read (cryptoSafe_Bytes)
        if err1 != nil {
                return cryptoSafe_String, err1
        }
        cryptoSafe_SubString := base32.StdEncoding.EncodeToString (cryptoSafe_Bytes)
        cryptoSafe_String += cryptoSafe_SubString[0:13]
        // ... }

        // This part generates a 19-character string that represents the current time. String format: SSNNNNNNNNNDDMMYYYY. SS refers to the seconds, NNNNNNNNN refers to the nano seconds, DD refers to the day, MM refers to the month, and YYYY refers to the year. { ...

                // Adding of the seconds part { ...
                current_Time := time.Now ()
                sec := current_Time.Second ()
                sec_String := Prepend (strconv.Itoa (sec), "0", 2)
                cryptoSafe_String += sec_String
                // ... }

                // Adding of the nano seconds part { ...
                nanoSec := current_Time.Nanosecond ()
                nanoSec_String := Prepend (strconv.Itoa (nanoSec), "0", 9)
                cryptoSafe_String += nanoSec_String
                // ... }

                // Adding of the day part { ...
                day := current_Time.Day ()
                day_String := Prepend (strconv.Itoa (day), "0", 2)
                cryptoSafe_String += day_String
                // ... }

                // Adding of the months part { ...
                month := Month_String_To_Int (current_Time.Month ().String ())
                month_String := Prepend (strconv.Itoa (month), "0", 2)
                cryptoSafe_String += month_String
                // ... }

                // Adding of the year part { ...
                year := current_Time.Year ()
                year_String := Prepend (strconv.Itoa (year), "0", 4)
                cryptoSafe_String += year_String
                // ... }
        // ... }

        return cryptoSafe_String, err
}

func CryptoSafe_Unique_Random_64 () (string, error) { // This function is the 64-character version of CryptoSafe_Unique_Random_32 ().

        part1, err1 := CryptoSafe_Unique_Random_32 ()
        if err1 != nil {
                return "", err1
        }

        part2, err2 := CryptoSafe_Unique_Random_32 ()
        if err2 != nil {
                return "", err2
        }

        return part1 + part2, nil
}

func CryptoSafe_Unique_Random_128 () (string, error) { // This function is the 128-character version of CryptoSafe_Unique_Random_32 ().
        part1, err1 := CryptoSafe_Unique_Random_64 ()
        if err1 != nil {
                return "", err1
        }

        part2, err2 := CryptoSafe_Unique_Random_64 ()
        if err2 != nil {
                return "", err2
        }

        return part1 + part2, nil
}

func CryptoSafe_Unique_Random_256 () (string, error) { // This function is the 256-character version of CryptoSafe_Unique_Random_32 ().
        part1, err1 := CryptoSafe_Unique_Random_128 ()
        if err1 != nil {
                return "", err1
        }

        part2, err2 := CryptoSafe_Unique_Random_128 ()
        if err2 != nil {
                return "", err2
        }

        return part1 + part2, nil
}

func Prepend (input string, prepend string, length int) (string) { // This function prepends variable "prepend", to variable "input", until variable "input" has a length greater than or equals to variable "length". In simpler terms, the function repeatedly prepends a string to an object string until the length of the object string is greater or equal to the object string.
        for {
                if len (input) >= length {
                        break
                }
                input = prepend + input
        }
        return input
}

func Month_String_To_Int (month string) (int) { // This function returns the numerical value of months. If the string provided is not a valid month name, 0 is returned.
        switch month {
                case "January":   return 1
                case "February":  return 2
                case "March":     return 3
                case "April":     return 4
                case "May":       return 5
                case "June":      return 6
                case "July":      return 7
                case "August":    return 8
                case "September": return 9
                case "October":   return 10
                case "November":  return 11
                case "December":  return 12
        }
        return 0
}
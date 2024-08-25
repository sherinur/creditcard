# creditcard
In this project, I created a tool called creditcard to: 
  1. Validate credit card numbers.
  2. Generate possible card numbers.
  3. Get information about card brands and issuers.
  4. Issue new card numbers.
     
## Context
Credit cards are used to pay for goods and services. Each card has a unique number that helps identify the cardholder and the issuing bank. Credit card numbers are long to ensure each card is unique. 
For example:
   1. Visa uses 13- and 16-digit numbers.
   2. MasterCard uses 16-digit numbers.
   3. American Express uses 15-digit numbers.

These numbers are not random. They follow specific patterns:
   1. Visa numbers start with 4.
   2. MasterCard numbers start with 51, 52, 53, 54, or 55.
   3. American Express numbers start with 34 or 37.

Credit card numbers also include a "checksum" that helps detect errors. This is done using Luhn's Algorithm, a simple math formula that checks if the number is valid.

### Validate
The validate feature checks if a credit card number is valid using Luhn's Algorithm.
  Requirements:
    The number must be at least 13 digits long.
    If valid, print OK to stdout and exit with status 0.
    If invalid, print INCORRECT to stderr and exit with status 1.
    Support passing multiple entries.
    Support --stdin flag to pass number from stdin.

    $ ./creditcard validate "4400430180300003"
    OK
    $ ./creditcard validate "4400430180300002"
    INCORRECT
    $ ./creditcard validate "4400430180300003" "4400430180300011"
    OK
    OK
    $ echo "4400430180300003" | ./creditcard validate --stdin
    OK
    $ echo "4400430180300003" "4400430180300011" | ./creditcard validate --stdin
    OK
    OK

### Generate
The generate feature creates possible credit card numbers by replacing asterisks (*) with digits.
  Requirements:
    Replace up to 4 asterisks (*) with digits. If more - it's an error. Asterisks should be at the end of the given credit card number.
    Print the generated numbers to stdout.
    Numbers must be printed in ascending order.
    Exit with status 1 if there is any error.
    Support --pick flag to randomly pick a single entry.

    $ ./creditcard generate "440043018030****"
    4400430180300003
    4400430180300011
    4400430180300029
    ...
    4400430180309988
    4400430180309996
    $ ./creditcard generate --pick "440043018030****"
    4400430180304385

    In case of an error:
    $ ./creditcard generate --pick "44004301803*****"
    $ echo $?
    1

### Information
The information feature provides details about the card based on data in brands.txt and issuers.txt.
  Requirements:
    Output the card number, validity, brand, and issuer.
    Support --stdin flag to pass number from stdin.
    Support passing multiple entries.

    $ ./creditcard information --brands=brands.txt --issuers=issuers.txt "4400430180300003"
    4400430180300003
    Correct: yes
    Card Brand: VISA
    Card Issuer: Kaspi Gold
    
### Issue
The issue feature generates a random valid credit card number for a specified brand and issuer.
    Requirements:
      Pick a random number for the specified brand and issuer.
      Exit with status 1 if there is any error.

      $ ./creditcard issue --brands=brands.txt --issuers=issuers.txt --brand=VISA --issuer="Kaspi Gold"
      4400430180300003

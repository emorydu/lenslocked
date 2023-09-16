Session tokens prevent cookie tampering through obfuscation.

user_id | session token
==================================
      1 | AUniqueRandomString123abc
      2 | AnotherString789xyz

Cookies can be changes, but attackers won't be able to predict or easily guess
a valid session token

We are going to make our session tokens unpredictable by generating them with
sufficient randomness.
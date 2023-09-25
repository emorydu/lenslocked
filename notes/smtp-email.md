# SMTP Email Example

```
MIME-Version: 1.0
Date: Mon, 25 Sep 2023 15:04:48 +0800
From: test@gmail.com
Subject: This is test email
To: test@qq.com
Content-Type: text/plain; charset=UTF-8
Content-Transfer-Encoding: quoted-printable

This is the body of the email
```




```
MIME-Version: 1.0
Date: Mon, 25 Sep 2023 15:06:53 +0800
To: test@qq.com
From: test@gmail.com
Subject: This is test email
Content-Type: multipart/alternative;
 boundary=19b2d25b69ea2a701c93e73984a72963f4aa9461f40fb997504ae3e2145d

--19b2d25b69ea2a701c93e73984a72963f4aa9461f40fb997504ae3e2145d
Content-Transfer-Encoding: quoted-printable
Content-Type: text/plain; charset=UTF-8

This is the body of the email
--19b2d25b69ea2a701c93e73984a72963f4aa9461f40fb997504ae3e2145d
Content-Transfer-Encoding: quoted-printable
Content-Type: text/html; charset=UTF-8

<h1>Hello there buddy!</h1><p>This is the email</p><p>Hope you enjoy it</p>
--19b2d25b69ea2a701c93e73984a72963f4aa9461f40fb997504ae3e2145d--


```
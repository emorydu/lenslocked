# Password Reset Process

1. User forgets their PW
2. User visits a page like `/forgets-pw` to request a new PW. This is a form 
that asks for their address.
3. User submits the form, our server creates a reset token, and emails it to
the user.
4. The user checks their email. In it they have a link allowing them to reset
their password. This link embeds the password reset token so the user doesn't
have to think about it.
5. When the user follows the password reset link, we parse the token from the
URL and generate an HTML form for the user to update their password. This form
has the password reset token as hidden. The user enters a new password in 
this form and hits submit.
6. Our server process the form by:
    1. Verifying the reset token is valid and looking up the user associated to
    it.
    2. Once verified, the reset token is deleted.
    3. We then update that user's password.
    4. Next we sign the user in
    5. Finally, we redirect them to the dashboard. (`/users/me` for our app.)

Magic sign in links use this same idea, but don't ask for a new password.

Putting this all together, we will need to code the following:

Views (templates):
    1. An HTML page with a form to type your email address and request a password
    reset token.
    2. An HTML page that tells the user to go check their email.
    3. An HTML page with a form to update a password.

Models (service):
    Reset tokens:
        1. A migration adding a database table for reset tokens.
        2. A reset type mapped to the DB table.
        3. A models service to create & consume reset tokens.

    Emails:
        1. An email server to manage sending the forget password email and other
        future email.

    Controllers: (handlers)
        1. An HTTP handler to render and process the reset pw form.
        2. An HTTP handler to render process the forgotten password form.
    
    Main:
        1. Setup routes.
        2. Setup templates processing.

This is a lot, so I'll break this into two sections.
In the first we will work on the emails, then we will look at the code to 
generate reset tokens.

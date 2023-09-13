# Exercises

1. Add an FAQ page.

We want visit <127.0.0.1:3000/faq> and have it render an FAQ page with
questions and answer like:

```
Q: Is there a free version?
A: Yes! We offer a free trial for 30 days on any paid plans.

Q: What are your support hours?
A: We have support staff answering emails 24/7, though response times may be a
bit slower on weekends.

Q: How do I contact support?
A: Email us - support@orangeduxiaocheng@gmail.com
```


2. Add a URL Parameter

Read the docs and see if you can add a URL parameter to one of your router,
retrieve it in your handler, and output it to the resulting HTML.

*Hint: See [these docs](https://github.com/go-chi/chi#url-parameters) if you
need some guidance. You shouldn't need to use context, just the `URLParam`
method.*


3. Experiment with builtin middleware

Chi provides quite a few builtin middleware. One is the Logger middleware,
which will track how long each request is taking. Try to add it to your
application, then to only a single route.
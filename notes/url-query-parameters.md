# URL Query Parameters

GET requests DO NOT have a body, so they can't include POST form values.
Instead, data must be passed as URL query parameters.

Query params are key/value pairs appended to the URL after the `?`

Example:

```
GET https://example.com/widgets?page=3
                                 |   |
                                key  |
                                     |
                                   value    
``` 

Extra ones can be added with `&`

```
GET https://example.com/widgets?page=3&color=green
                                 |   |   |     |
                                key  |   |     |
                                     |   |     |
                                   value |     |
                                        key    |
                                              value
```

These can be incredibly useful for a number of reasons. One we saw here is
pagination, but another use is to prefill forms.
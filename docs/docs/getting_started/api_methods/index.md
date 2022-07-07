# API Methods

  - [Token Methods](token_methods.md)
  - [Info Methods](info_methods.md)
  - [Scan Methods](scan_methods.md)
  - [Ban Methods](ban_methods.md)
  - [Polling Methods](polling_methods.md)

<hr/>

In this section, everything about how to send a specified request to the API and their necessary parameters are explained.

We will be using curl to display how to send a request to the API to make it more simpler. Later on you can use online tools to convert a curl request to your desired programming language.

Here is a sample request, all of the requests explained in this section will be in the following format:

## **RequestName**

Request explanation.

### **Params**

|           Param               |   Type      |   Required    |    Permission    |
|         :----------:          | :---------: |  :---------:  |    :---------:   |
|        user-id                |   int64     |      No       |     Enforcer     |
|        reason                 |   string    |      Yes      |     Enforcer     |
|        sample-param           |    bool     |      No       |     Enforcer     |
|        something              |   string    |      Yes      |     Inspector    |

- reason: the reason.
- user-id: the user-id.
- sample-param: a sample param which doesn't really need to be passed to the api.
- something: a param. please do notice that this param can only be used by inspectors.

### **Example**

```sh
curl 'https://PsychoPass.AnimeKaizoku.com/requestName' -H 'user-id: 12345' -H 'reason: the reason' \
      -H 'simple-param: true' -H 'something: ok'
```

# How to run test files?

1- Make sure you have the file `config.ini` in the `tests` directory. Its content should be in the same format as the `config.ini` file in the root directory of the project.

2- Make sure you have the file `baseUrl.ini` in the `tests` directory. Its content should be only your base url. If you are going to run the application on your server and run the tests on your local machine, put your server url in there. it should look like this: `http://host_url:8080/` or `http://1.2.3.4:50100/`. Don't use anything else, since the file won't be parsed via a library, all of its contents will be read as the base url, so comments are not allowed.

3- If you are running the application on your server and you want to run the test files in your local, please make sure to create the file `owner.token` in the `tests` directory. It should contain a token with `owner` permission.

4- If you are running into an error, or something is acting weird, please try to reproduce the error more than once. Please remove `*.db` files from `tests` directory once and then try again. If the error occurs again, please try to fix it in the code. If you don't have any idea why it's happening, please share it with the team.

5- If you are on linux platform, simply execute command `./run.sh test` for executing all test files. If not, please try `go test -v .\...` command.

# Docs

## How to generate

To generate document pages with highlighted JSON, run these commands:
```sh
pip3 install -r requirements.txt 
python3 format.py
```

<hr/>

## Tips

- If you are in an Unix-based OS, you can simply run `./build.sh`. Please notice that `./run.sh` itself will run this file, so you no need to run it twice.

- When serving this page CSP and Referrer-Policy should be set via headers and canonical meta tags should be set

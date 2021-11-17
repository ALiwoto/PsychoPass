import os
import shutil
from bs4 import BeautifulSoup
from pygments import highlight
from pygments.lexers import JsonLexer
from pygments.formatters import HtmlFormatter

json_lexer = JsonLexer()
html_formatter = HtmlFormatter(wrapcode=True, style='rrt')
for dirname, _, files in os.walk('in'):
    out_dirname = 'out' + dirname[2:]
    os.makedirs(out_dirname, exist_ok=True)
    for file in files:
        if not (file.endswith('.html') or file.endswith('.htm')):
            shutil.copyfile(os.path.join(dirname, file), os.path.join(out_dirname, file))
            continue
        with open(os.path.join(dirname, file)) as f:
            soup = BeautifulSoup(f.read(), 'lxml')
        for i in soup.find_all('div', {'class': 'highlightjson'}):
            highlight_soup = BeautifulSoup(highlight(i.string, json_lexer, html_formatter), 'lxml')
            i.replace_with(highlight_soup.div)
        with open(os.path.join(out_dirname, file), 'w+') as file:
            file.write(soup.prettify())
with open('out/pygments.css', 'w+') as file:
    file.write('// This is an auto-generated file, do not edit\n\n')
    file.write(html_formatter.get_style_defs('.highlight'))

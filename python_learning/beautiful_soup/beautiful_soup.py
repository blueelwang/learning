from bs4 import BeautifulSoup

with open('sample.html', 'r') as fd:
    htmlContent = fd.read()
# soup = BeautifulSoup(htmlContent, 'html.parser')
soup = BeautifulSoup(htmlContent, 'lxml')
print(type(soup.body))
print(soup.body)        # soup.body是Tag类型
print(soup.body.name)   # 标签名: body

print(soup.body.div.attrs['class'])     # 取div的属性, 没有这个属性时会抛KeyError: 'class'
print(soup.body.div['class'])           # 取div的属性
print(soup.body.div.string)             # 取div里的内容

print(soup.find_all('div', 'left'))     # 找出所有 class=left 的 div 标签



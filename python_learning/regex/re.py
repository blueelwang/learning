import re

pattern = re.compile('[a-zA-z]+')
match = re.findall(pattern, '<div>test string</div>')
print(match)

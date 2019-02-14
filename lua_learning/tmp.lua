---
--- @author      jianwei18 <jianwei18@staff.weibo.com>
--- @date        2018-11-01 21:28:55
--- @copyright   Copyright (c) 2009-2018 Weibo.com All Rights Reserved.
---

print "Hello World"

local image_data = 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAcQAAALQCA'
print(image_data)

image_data = string.gsub(image_data, "data:image/png;base64,", "")
print(image_data)

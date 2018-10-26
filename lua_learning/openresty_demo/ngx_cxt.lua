---
--- @author      jianwei18@staff.weibo.com
--- @version     2018-10-26 18:55:00
---

if nil == ngx.cxt then
    ngx.cxt = {}
end
ngx.cxt.foo = 'bar'
ngx.say(ngx.cxt.foo)

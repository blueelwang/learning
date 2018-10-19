

-- https://www.runoob.com/lua/lua-iterators.html


t = {10, 20, 30, 40, 50}
function iterator(t)
    local i = 0
    return function()
        i = i + 1
        return i, t[i]
    end
end
for k, v in iterator(t) do
    if (nil == v) then
        break;
    end
    print(k, v)
end
--10
--20
--30
--40
--50

---ipairs()
function iter (a, i)
    i = i + 1
    local v = a[i]
    if v then
        return i, v
    end
end

function ipairs (a)
    return iter, a, 0
end
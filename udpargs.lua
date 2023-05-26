local socket = require("socket")
local udp = assert(socket.udp())
local data

udp:settimeout(1)
assert(udp:setsockname("*",0))
assert(udp:setpeername("localhost",3000))

for i = 0, 2, 1 do
  assert(udp:send(arg[1]))
  data = udp:receive()
  if data then
    break
  end
end


if data == nil then
  print("timeout")
else
  print(data)
end
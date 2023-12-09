local io = require("io")
local string = require("string")
local table = require("table")

local args = { ... }
local file = assert(io.open(args[1], "r"))
local input = file:read("*a")
file:close()

local dirs_r, els_r = string.match(input, "([LR]+)\n\n(.*)")

local dirs = {}
for dir in string.gmatch(dirs_r, "([LR])") do
  table.insert(dirs, dir)
end

local els = {}
for name, a, b in string.gmatch(els_r, "(%a+) = %((%a+), (%a+)%)") do
  els[name] = { a = a, b = b }
end

local steps = 0
local node = "AAA"
while true
do
  for _, dir in ipairs(dirs) do
    steps = steps + 1

    if dir == "L" then
      node = els[node].a
    else
      node = els[node].b
    end

    if node == "ZZZ" then
      print(steps)
      return
    end
  end
end

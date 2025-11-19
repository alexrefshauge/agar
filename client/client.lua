local socket = require "socket"
local json = require "json"

local M = {
	connected = false,
	client = socket.tcp(),
	id = -1,
}



function M:connect(address, port)
	self.client:settimeout(10)
	local ok, err = self.client:connect(address, port)
	if not ok then
		print("client failed to connect", err)
	end

	local data, id, size
	local world = require "world"
	data, err = self.client:receive("*l")
	if err then
		print(err)
	elseif data then
		print(data)
		local packet = {}
		for value in data:gmatch("[^;]+") do
			table.insert(packet, value)
		end
		id = tonumber(packet[1])
		size = tonumber(packet[2])
		self.id = id
		print("received id: " .. id)
		world:update({ size = size })
	end

	-- disable blocking
	self.client:settimeout(0)
	self.connected = ok
	return self.id, world
end

function M:updateWorld(world)
	if not self.connected then
		return world
	end

	local data, err = self.client:receive("*l")
	if err and not err == "timeout" then
		print(err)
		return world
	end

	if data then
		world:update(json.decode(data))
	end

	return world
end

local lastDirection = 0
function M:sendInput()
	if not self.connected then
		print("Unable to send input (client is not connected")
		return
	end
	local ww, wh, flags = love.window.getMode()
	local mx, my = love.mouse.getPosition()
	local cx, cy = ww / 2, wh / 2
	local direction = math.atan2(my - cy, mx - cx)

	if math.abs(lastDirection - direction) > 0.01 then
		-- client:send(tostring(direction))
		local _, err = self.client:send(tostring(direction))
		print(tostring(direction))
		if err then
			print("failed to send direction: " .. err)
		end
		lastDirection = direction
	end
end

return M

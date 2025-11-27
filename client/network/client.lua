local socket = require "socket"
local M = {
	tcp = socket.tcp4(),
	udp = socket.udp4(),
	clientId = -1,
	playerId = -1,
}

---connect to game server
---@param address string
function M:connect(address, port)
	print("connecting...")
	local ok = self.tcp:connect(address, port)
	if not ok == 1 then
		print("failed to connect to game server")
	end
	print("connection to game server established!")
	--TODO: connect udp as well

	self.tcp:settimeout(0)
	self.udp:settimeout(0)
end

local headerSize = 1
local payloadSize = {
	welcome = 4 + 4,
	stateHeader = 4 + 4,
	statePlayer = 92,
	stateBlob = 16,
}
function M:handleReceive(world)
	local header, error = self.tcp:receive(headerSize)
	if header then
		local packetType = string.byte(header, 1)
		if packetType == 1 then self:handleWelcome() end
		if packetType == 2 then self:handleState(world) end
		if packetType == 3 then self:handleDeltaState() end
	end

	if error and not error == "timeout" then
		print("failed to receive")
		print(error)
	end
	self:flush()
end

function M:handleWelcome()
	print("received: WELCOME")
	local payload, error = self.tcp:receive(payloadSize.welcome)
	if error then print(error) end
	self.clientId, self.playerId = love.data.unpack(">I4I4", payload)
end

function M:handleState(world)
	print("received: STATE")
	local header, error = self.tcp:receive(payloadSize.stateHeader)
	local playerCount, blobCount = love.data.unpack(">I4I4", header)
	local objects = {}
	for i = 0, playerCount do
		if i == playerCount then break end
		local data, error = self.tcp:receive(payloadSize.statePlayer)
		local id, size, px, py, vx, vy, dir, name = love.data.unpack(">I4I4fffffc16", data)
		print(vx, vy, dir)
		table.insert(objects,
			{ type = "player", id = id, size = size, pos = { x = px, y = py }, vel = { x = vx, y = vy }, dir = dir, name = name })
	end
	for i = 0, blobCount do
		if i == blobCount then break end
		local data, error = self.tcp:receive(payloadSize.stateBlob)
		local id, size, x, y = love.data.unpack(">I4I4ff", data)
		table.insert(objects, { type = "blob", id = id, size = size, pos = { x = x, y = y } })
	end
	world:putObjects(objects)
end

function M:handleDeltaState()
	print("received: DELTA_STATE")
	local payload, error = self.tcp:receive(payloadSize.welcome)
end

function M:flush()
	while true do
		local _, error = self.tcp:receive(1)
		if error == "timeout" then
			return
		end
	end
end

return M

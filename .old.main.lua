local world = require "world"
local json = require "json"

playerId = -1

local client
local socket = require "socket"
local status = "no status"
local message = "no message"

local screenCenter = { x = 0, y = 0 }
local mousePosition = { x = 0, y = 0 }

local direction = 0

function love.load()
	print("starting")
	world:initialize()
	client = socket.tcp()
	client:settimeout(5)
	local ok, err = client:connect("localhost", 42069)
	if ok then
		status = "connected to server!"
		client:settimeout(0)
	else
		status = "failed to connect to server :("
	end
end

local lastDirection = 0
function love.update()
	local ww, wh, flags = love.window.getMode()
	local mx, my = love.mouse.getPosition()
	screenCenter = { x = ww / 2, y = wh / 2 }
	mousePosition = { x = mx, y = my }

	direction = math.atan2(
		my - screenCenter.y,
		mx - screenCenter.x
	)

	if not client then
		return
	end

	local success = love.window.setFullscreen(true)
	local data, err = client:receive("*l")
	message = data or message

	if data then
		world = json.decode(data)
	end

	if math.abs(lastDirection - direction) > 0.1 then
		-- client:send(tostring(direction))
		lastDirection = direction
	end
end

function love.draw()
	love.graphics.circle("line", 0, 0, world.size)
	love.graphics.print(status, 300, 100, 0, 5, 5)
	love.graphics.print(message, 300, 300, 0, 3, 3)
	love.graphics.print(direction, 300, 500, 0, 3, 3)

	for i, b in ipairs(world.blobs) do
		love.graphics.circle("line", b.pos.x * 1000, b.pos.y * 1000, 10)
	end

	love.graphics.line(screenCenter.x, screenCenter.y, mousePosition.x, mousePosition.y)
end

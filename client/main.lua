local camera = require "camera"
local client = require "client"
local address, port = "130.225.37.50", 42069
address = "localhost"
local clientId = -1
local world

_G.debug = false


function love.load()
	clientId, world = client:connect(address, port)
	if client.connected then
		print("connected as client " .. clientId)
		love.window.setFullscreen(true)
	end
end

function love.update(dt)
	world = client:updateWorld(world)
	if not world then return end
	if not client.connected then return end

	local player = world.objects[clientId]
	if player then camera.target = player.pos end

	world:step(dt)
	camera:update()
	client:sendInput()
end

function love.draw()
	if not world then return end
	if not client.connected then
		love.graphics.print("connecting to server...", 100, 100, 0, 2)
		return
	end


	camera:applyTransform()


	local spacing = 200
	local w, h = love.graphics.getDimensions()

	-- find where the first visible grid line is
	local startX = math.floor(camera.center.x / spacing) * spacing
	local startY = math.floor(camera.center.y / spacing) * spacing

	love.graphics.setColor(0.1, 0.1, 0.1, 1)

	-- vertical lines
	for x = startX, startX + w + spacing, spacing do
		love.graphics.line(x, camera.center.x, x, camera.center.y + h)
	end

	-- horizontal lines
	for y = startY, startY + h + spacing, spacing do
		love.graphics.line(camera.center.x, y, camera.center.y + w, y)
	end

	love.graphics.setColor(1, 1, 1)

	world:draw()
	camera:clearTransform()
end

love = require "love"
local client = require "network.client"
local world = require "world"
local camera = require "camera"

---@diagnostic disable-next-line: duplicate-set-field
function love.load()
	client:connect("localhost", 42069)
end

---@diagnostic disable-next-line: duplicate-set-field
function love.update(dt)
	client:handleReceive(world)
	if world.playerId > 0 then
		camera.target = world.objects[world.playerId].pos
	end
	camera:update()
end

---@diagnostic disable-next-line: duplicate-set-field
function love.draw() 
	camera:applyTransform()
	world:draw()
	camera:clearTransform()
	end

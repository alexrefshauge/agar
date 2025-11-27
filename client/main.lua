love = require "love"
local client = require "network.client"
local world = require "world"
local camera = require "camera"

function love.load()
	client:connect("localhost", 42069)
	camera.target = { x = 0, y = 0 }
end

function love.update()
	client:handleReceive(world)
end

function love.draw() world:draw() end

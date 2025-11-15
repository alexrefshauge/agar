local love = love
local world = require "world"

function love.load()
	world:initialize()
end

function love.update()
	local success = love.window.setFullscreen(true)
end

function love.draw()
	world:draw()
end

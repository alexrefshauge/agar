require "love"

function love.keypressed(key)
	if key == "r" then love.event.quit "restart" end
	if key == "escape" then love.event.quit() end
end

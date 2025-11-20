M = {
	size = 100,
	players = {},
	blobs = {},
	objects = {}
}

local debug = _G.debug
function M:update(data)
	self.size = data.size or self.size
	local unloadIds = {}
	local eatIds = data.eat or {}

	for _, id in ipairs(unloadIds) do
		print("- unloading game object: " .. id)
		table.remove(self.objects, id)
	end

	for _, id in ipairs(eatIds) do
		print(id .. " has been eaten!")
		table.remove(self.objects, id)
	end

	local count = 0
	-- load players
	for _, o in ipairs(data.players or {}) do
		if self.objects[o.id] then -- update object
			if debug then
				print("= updating game object [player]: " .. o.id)
			end
			self.objects[o.id].pos = o.pos
			self.objects[o.id].dir = o.dir
			self.objects[o.id].vel = o.vel
			self.objects[o.id].name = o.name
			self.objects[o.id].size = o.size
		else -- load as new object
			if debug then
				print("+ loading game object [player]: " .. o.id)
			end
			self.objects[o.id] = o
			self.objects[o.id].type = "player"
			count = count + 1
		end
	end

	-- load blobs
	for _, o in ipairs(data.blobs or {}) do
		if self.objects[o.id] then -- update object
			if debug then
				print("= updating game object [blob]: " .. o.id)
			end
			self.objects[o.id].pos = o.pos
			self.objects[o.id].size = o.size
		else -- load as new object
			if debug then
				print("+ loading game object [blob]: " .. o.id)
			end
			self.objects[o.id] = o
			self.objects[o.id].type = "blob"
			count = count + 1
		end
	end


	-- print(tostring(count) .. " objects loaded")
end

function M:step(dt)
	for _, object in pairs(self.objects) do
		local type = object.type
		if type == "player" then
			local vx, vy = object.vel.x, object.vel.y
			local ox, oy = object.pos.x + vx * dt, object.pos.y + vy * dt

			-- gpt
			local px, py = ox, oy
			local dist = math.sqrt(px * px + py * py)
			-- how far the player sticks outside the world boundary
			local distanceOutside = (dist + object.size) - self.size
			-- inside the world: no correction
			if distanceOutside > 0 then
				-- normalize (px, py)
				local nx, ny = px / dist, py / dist
				-- move player back inside
				ox = ox - nx * distanceOutside
				oy = oy - ny * distanceOutside
			end
			-- gpt end

			object.pos = { x = ox, y = oy }
			self.objects[object.id] = object
		end
	end
end

function M:draw()
	for id, o in pairs(self.objects) do
		local t = o.type or "none"
		if t == "player" then
			love.graphics.circle("line", o.pos.x, o.pos.y, o.size)
			local xx, yy =
					o.pos.x + math.cos(o.dir) * 20,
					o.pos.y + math.sin(o.dir) * 20
			love.graphics.line(o.pos.x, o.pos.y, xx, yy)
		elseif t == "blob" then
			love.graphics.circle("fill", o.pos.x, o.pos.y, o.size)
		end
	end

	love.graphics.setColor(1, 0, 0)
	love.graphics.circle("line", 0, 0, self.size or 10)
	love.graphics.setColor(1, 1, 1)
end

return M

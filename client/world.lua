---@class Vec
---@field x number
---@field y number

---@class Object
---@field id integer
---@field pos Vec
---@field type "blob" | "player"

---@class Blob : Object

---@class Player : Object
---@field name string
---@field size integer
---@field vel Vec
---@field dir number

---@class World
---@field size number
---@field players table[]
---@field blobs table[]
---@field objects table<integer, Object>
local M = {
	size = 100,
	players = {},
	blobs = {},
	objects = {}
}

local inspect = require "inspect"
local debug = _G.debug

---update world based on UpdateData
---@param data UpdateData
function M:update(data)
	self.size = data.size or self.size
	local unloadIds = {}
	local eatIds = data.eat or {}

	for _, id in ipairs(unloadIds) do
		print("- unloading game object: " .. id)
		self.objects[id] = nil
	end

	for _, id in ipairs(eatIds) do
		print(id .. " has been eaten!")
		self.objects[id] = nil
	end

	-- load players
	for _, o in ipairs(data.players or {}) do
		if self.objects[o.id] then -- update object
			if debug then
				print("= updating game object [player]: " .. o.id)
			end
			self.objects[o.id] = nil
			self.objects[o.id] = o
			self.objects[o.id].type = "player"
		else -- load as new object
			if debug then
				print("+ loading game object [player]: " .. o.id)
			end
			self.objects[o.id] = o
			self.objects[o.id].type = "player"

		end
	end

	-- load blobs
	for _, o in ipairs(data.blobs or {}) do
		if self.objects[o.id] then -- update object
			if debug then
				print("= updating game object [blob]: " .. o.id)
			end
			self.objects[o.id] = nil
			self.objects[o.id] = o
			self.objects[o.id].type = "blob"
		else -- load as new object
			if debug then
				print("+ loading game object [blob]: " .. o.id)
			end
			self.objects[o.id] = o
			self.objects[o.id].type = "blob"
		end
	end
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
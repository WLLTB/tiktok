package constant

const TokenBucketLimiterTryAcquireRedisScript = `
-- ARGV[1]: 容量
-- ARGV[2]: 发放令牌速率/秒
-- ARGV[3]: 当前时间（秒）

local capacity = tonumber(ARGV[1])
local rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])

local lastTime = tonumber(redis.call("hget", KEYS[1], "lastTime"))
local currentTokens = tonumber(redis.call("hget", KEYS[1], "currentTokens"))
-- 初始化
if lastTime == nil then
   lastTime = now
   currentTokens = capacity
   redis.call("hmset", KEYS[1], "currentTokens", currentTokens, "lastTime", lastTime)
end

-- 尝试发放令牌
-- 距离上次发放令牌的时间
local interval = now - lastTime
if interval > 0 then
   -- 当前令牌数量+距离上次发放令牌的时间(秒)*发放令牌速率
   local newTokens = currentTokens + interval * rate
   if newTokens > capacity then
      newTokens = capacity
   end
   currentTokens = newTokens
   redis.call("hmset", KEYS[1], "currentTokens", newTokens, "lastTime", now)
end

-- 如果没有令牌，请求失败
if currentTokens == 0 then
   return 0
end
-- 果有令牌，当前令牌-1，请求成功
redis.call("hincrby", KEYS[1], "currentTokens", -1)
redis.call("expire", KEYS[1], capacity / rate)
return 1
`

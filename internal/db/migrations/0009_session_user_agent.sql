-- The device a session was last used from (OS · browser), saved as its page
-- closes, so the GM's device list can name an offline device. Empty until
-- the device first disconnects.
ALTER TABLE sessions ADD COLUMN user_agent TEXT NOT NULL DEFAULT '';

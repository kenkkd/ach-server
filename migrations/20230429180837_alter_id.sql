-- modify "responses" table
ALTER TABLE `responses` MODIFY COLUMN `id` varchar(255) NOT NULL, MODIFY COLUMN `thread_id` varchar(255) NOT NULL;
-- modify "threads" table
ALTER TABLE `threads` MODIFY COLUMN `id` varchar(255) NOT NULL, ADD COLUMN `title` varchar(255) NOT NULL;

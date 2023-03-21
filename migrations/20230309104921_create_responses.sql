-- create "responses" table
CREATE TABLE `responses` (`id` bigint NOT NULL AUTO_INCREMENT, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `deleted_at` timestamp NULL, `name` varchar(255) NOT NULL, `number` bigint NOT NULL, `content` varchar(255) NOT NULL, `thread_id` bigint NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `responses_threads_responses` FOREIGN KEY (`thread_id`) REFERENCES `threads` (`id`) ON DELETE NO ACTION) CHARSET utf8mb4 COLLATE utf8mb4_bin;
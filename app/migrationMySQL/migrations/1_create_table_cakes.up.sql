CREATE TABLE IF NOT EXISTS `cakes` (
  `id` bigint(20) NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` TEXT NOT NULL,
  `rating` TEXT NOT NULL,
  `image` text NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
);

ALTER TABLE `cakes`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `cakes`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;
COMMIT;
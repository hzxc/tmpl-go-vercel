-- CreateTable
CREATE TABLE `Project` (
    `id` INTEGER NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(191) NOT NULL,
    `pin` BOOLEAN NOT NULL DEFAULT false,
    `personId` INTEGER NOT NULL,
    `organization` VARCHAR(191) NOT NULL,
    `description` VARCHAR(191) NULL,
    `created` DATETIME(3) NOT NULL,

    UNIQUE INDEX `Project_name_key`(`name`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

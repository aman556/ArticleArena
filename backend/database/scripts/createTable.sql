CREATE DATABASE IF NOT EXISTS `articlearenadatabase` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `articlearenadatabase`;

CREATE TABLE `UserHandles` (
`ArticleArenaHandle` char(32) NOT NULL PRIMARY KEY,
`GeeksforgeeksHandle`  char(32),
`MediumHandle` char(32),
`TutorialpointHandle` char(32)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `UserInfo` (
`UserName` char(32),
`ArticleArenaHandle` char(32) NOT NULL PRIMARY KEY,
`Email`  char(32),
`GithubUrl` char(32),
`LinkedinUrl` char(32)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 17, 2021 at 08:06 AM
-- Server version: 10.4.14-MariaDB
-- PHP Version: 7.4.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `pokemondb`
--

-- --------------------------------------------------------

--
-- Table structure for table `deletedusers`
--

CREATE TABLE `deletedusers` (
  `userId` int(11) NOT NULL,
  `email` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `role` int(11) NOT NULL,
  `token` longtext DEFAULT NULL,
  `name` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `deletedusers`
--

INSERT INTO `deletedusers` (`userId`, `email`, `password`, `role`, `token`, `name`) VALUES
(4, 'test@email.com', 'password', 0, NULL, 'Test User Deleted');

-- --------------------------------------------------------

--
-- Table structure for table `pokemons`
--

CREATE TABLE `pokemons` (
  `pokeId` int(11) NOT NULL,
  `pokeName` varchar(250) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `pokemons`
--

INSERT INTO `pokemons` (`pokeId`, `pokeName`) VALUES
(1, 'Pikachu'),
(2, 'Charmander'),
(6, 'Eevee');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `userId` int(11) NOT NULL,
  `email` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `role` int(11) NOT NULL,
  `token` longtext DEFAULT NULL,
  `name` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`userId`, `email`, `password`, `role`, `token`, `name`) VALUES
(0, 'admin@admin.com', 'admin123', 1, NULL, 'Admin'),
(1, 'user@user.com', 'user123', 0, NULL, 'User'),
(5, 'testedit@newuser.com', 'password', 0, NULL, 'New User');

-- --------------------------------------------------------

--
-- Table structure for table `__efmigrationshistory`
--

CREATE TABLE `__efmigrationshistory` (
  `MigrationId` varchar(95) NOT NULL,
  `ProductVersion` varchar(32) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `__efmigrationshistory`
--

INSERT INTO `__efmigrationshistory` (`MigrationId`, `ProductVersion`) VALUES
('20210225041727_InitialMigration', '5.0.1'),
('20210325135534_InitialMigrations', '5.0.1'),
('20210329041231_InitialMigrations', '5.0.1'),
('20210329041531_InitialMigrations', '5.0.1');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `deletedusers`
--
ALTER TABLE `deletedusers`
  ADD PRIMARY KEY (`userId`);

--
-- Indexes for table `pokemons`
--
ALTER TABLE `pokemons`
  ADD PRIMARY KEY (`pokeId`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`userId`);

--
-- Indexes for table `__efmigrationshistory`
--
ALTER TABLE `__efmigrationshistory`
  ADD PRIMARY KEY (`MigrationId`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `pokemons`
--
ALTER TABLE `pokemons`
  MODIFY `pokeId` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `userId` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

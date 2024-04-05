-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 05, 2024 at 09:43 AM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_tugas_explorasi_3_pbp`
--

-- --------------------------------------------------------

--
-- Table structure for table `tasks`
--

CREATE TABLE `tasks` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `start_task` timestamp NOT NULL DEFAULT current_timestamp(),
  `due_date` timestamp NOT NULL DEFAULT current_timestamp(),
  `details` text DEFAULT NULL,
  `notified` int(11) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `tasks`
--

INSERT INTO `tasks` (`id`, `user_id`, `title`, `start_task`, `due_date`, `details`, `notified`) VALUES
(1, 4, 'mengerjakan tugas pbp', '2024-03-30 04:19:54', '2024-04-01 03:00:00', 'frontend, backend (go mail)', 0),
(2, 1, 'ngerjain metpen', '2024-04-04 10:38:11', '2024-04-04 11:29:00', 'Jangan lupa ngerjain metpen Bab 1 - 3', 1),
(6, 1, 'siapin interview', '2024-04-05 02:36:36', '2024-04-05 02:40:00', 'siapin 3 minute presentation, jangan lupa print CV', 2),
(7, 1, 'panasin motor', '2024-04-05 02:38:31', '2024-04-05 02:44:00', 'berangkat ke kampus jgn lupa', 2),
(8, 1, 'jgn lupa mandi', '2024-04-05 02:44:20', '2024-04-05 02:50:00', 'berangkat ke kampus jgn lupa mandi dluuuu', 2),
(9, 4, 'ngerjain tugas', '2024-04-05 06:21:40', '2024-04-05 02:50:00', 'tugas abc', 0),
(10, 4, 'ngerjain tugas', '2024-04-05 06:22:13', '2024-04-05 06:25:00', 'tugas abc', 0),
(11, 4, 'ngerjain tugas', '2024-04-05 06:28:26', '2024-04-05 06:35:00', 'tugas abc', 1),
(12, 4, 'ngerjain tugas', '2024-04-05 06:39:57', '2024-04-05 06:49:00', 'tugas bcde', 2);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `email`, `password`) VALUES
(1, 'jason', 'jasonenrico79@gmail.com', 'jason123'),
(2, 'alex', '', 'alex123'),
(3, 'marcel', '', 'marcel123'),
(4, 'nico', 'nicojonathan69@gmail.com', 'nico123');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tasks`
--
ALTER TABLE `tasks`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `tasks`
--
ALTER TABLE `tasks`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `tasks`
--
ALTER TABLE `tasks`
  ADD CONSTRAINT `tasks_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

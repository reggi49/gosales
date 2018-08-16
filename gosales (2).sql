-- phpMyAdmin SQL Dump
-- version 4.7.9
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Aug 15, 2018 at 06:38 PM
-- Server version: 10.1.31-MariaDB
-- PHP Version: 5.6.34

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gosales`
--

-- --------------------------------------------------------

--
-- Table structure for table `penjualan`
--

CREATE TABLE `penjualan` (
  `id_penjualan` int(11) NOT NULL,
  `tgl_keluar` varchar(100) NOT NULL,
  `kode_bon` varchar(100) NOT NULL,
  `merk` varchar(100) NOT NULL,
  `nopol` varchar(100) NOT NULL,
  `suplier` varchar(100) NOT NULL,
  `harga_modal` varchar(100) NOT NULL,
  `harga_jual` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `penjualan`
--

INSERT INTO `penjualan` (`id_penjualan`, `tgl_keluar`, `kode_bon`, `merk`, `nopol`, `suplier`, `harga_modal`, `harga_jual`) VALUES
(1, '14/08/2018', '123', 'honda', 'f123', 'aaa', '12', '13'),
(2, '08/15/2018', '11212', 'Honda', 'f12212', 'freen', '12', '13'),
(3, '08/14/2018', '11122', 'Honda', 'Honda', 'tiga bisa edit', '1', '2');

-- --------------------------------------------------------

--
-- Table structure for table `persediaan`
--

CREATE TABLE `persediaan` (
  `id_barang` int(11) NOT NULL,
  `nopol` varchar(50) NOT NULL,
  `tgl_masuk` varchar(50) NOT NULL,
  `merk` varchar(50) NOT NULL,
  `tahun` varchar(50) NOT NULL,
  `harga_modal` varchar(50) NOT NULL,
  `harga_jual` varchar(50) NOT NULL,
  `create_date` varchar(50) NOT NULL,
  `suplier` varchar(50) NOT NULL,
  `telp` varchar(50) NOT NULL,
  `alamat` varchar(50) NOT NULL,
  `deskripsi` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `persediaan`
--

INSERT INTO `persediaan` (`id_barang`, `nopol`, `tgl_masuk`, `merk`, `tahun`, `harga_modal`, `harga_jual`, `create_date`, `suplier`, `telp`, `alamat`, `deskripsi`) VALUES
(1, 'F117', '2018-08-02', 'Honda', '2018', '1000', '12000', '2017-01-01', '', '085782057092', 'jakarta', ''),
(2, 'ad', '19/09/12', 'a', 'd', '0', '0', '2017-01-01', '', '2017-01-01', '2017-01-01', '2017-01-01'),
(3, 'Honda', '08/15/2018', 'Honda', '', '1', '2', '', 'tiga bisa edit', '', '', ''),
(4, 'f321', '2017-01-01', 'honda', '2017', '1000', '2000', '2017-01-01', '', '3232', 'jakarta', 'bogogoog'),
(5, 'sdsa', 'sds', 'ds', 'ads', '12', '11', '2017-01-01', 'ads', 'asds', 'sdsa', 'sadas'),
(6, 'f1231', '2017-01-01', 'honda', '2018', '12', '13', '2017-01-01', 'honda', '123', 'jakarta', 'bogoroor'),
(7, 'f1235', '2017-01-01', 'honda', '2017', '14', '15', '2017-01-01', 'honda', '123', 'daa', 'add'),
(8, 'f12346', '2017-01-01', 'hond', '2018', '12', '14', '2017-01-01', 'aa', '1212', 'sdsd', 'sdsd'),
(9, '', '', '', '', '0', '0', '', '', '', '', ''),
(10, 'f9090', '08/13/2018', 'Honda', '2018', '1', '2', '2017-01-01', 'honda', '085782057092', 'bogor, Bogor', 'deskripsi'),
(11, 'f12340', '', 'Honda', '', '12', '13', '', 'ads', '', '', '');

-- --------------------------------------------------------

--
-- Table structure for table `person`
--

CREATE TABLE `person` (
  `id` int(11) NOT NULL,
  `first_name` varchar(40) DEFAULT NULL,
  `last_name` varchar(40) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `person`
--

INSERT INTO `person` (`id`, `first_name`, `last_name`) VALUES
(1, 'satu edit', 'satuu'),
(2, 'dua', 'duaa'),
(3, 'tiga', 'tigaa\n');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `username` varchar(255) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` int(11) NOT NULL DEFAULT '2'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `username`, `email`, `password`, `role`) VALUES
(1, 'admin', 'admin123', 'admin@gmai.com', '$2a$10$jRGs7WDZ0G75fQhjF28eS.IsXokDDdsmXztwGgiJG8JAjF2WkqSqO', 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `penjualan`
--
ALTER TABLE `penjualan`
  ADD PRIMARY KEY (`id_penjualan`);

--
-- Indexes for table `persediaan`
--
ALTER TABLE `persediaan`
  ADD PRIMARY KEY (`id_barang`);

--
-- Indexes for table `person`
--
ALTER TABLE `person`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `penjualan`
--
ALTER TABLE `penjualan`
  MODIFY `id_penjualan` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `persediaan`
--
ALTER TABLE `persediaan`
  MODIFY `id_barang` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `person`
--
ALTER TABLE `person`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

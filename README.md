# API Shell Executor Documentation

---

## Overview
API ini menyediakan HTTP endpoint untuk mengelola layanan VPN/proxy, termasuk **SSH, Shadowsocks (SS), VMess, VLESS, Trojan, Noobz, Zivpn**, dan lainnya. Semua endpoint diakses menggunakan **Bearer Token** dan mengembalikan output dalam format **JSON**.

---

## Table of Contents
- [Authentication](#authentication)
- [Base URL](#base-url)
- [General Usage](#general-usage)
- [API by Protocol](#api-by-protocol)
  - [SSH](docs/ssh.md)
  - [Shadowsocks (SS)](docs/shadowsocks.md)
  - [XRAY (VMess, VLESS, Trojan)](docs/xray.md)
  - [NoobzVPN](docs/noobzvpn.md)
  - [UDP Zivpn](docs/udpzivpn.md)
  - [Wireguard WARP](docs/wireguard-warp.md)
- [Error Handling](#error-handling)
- [Notes](#notes)

---

## Authentication
Semua endpoint memerlukan **Bearer Token** di header `Authorization`: Bearer your-token

Token ini adalah kunci akses Anda ke API. Pastikan untuk menjaganya tetap aman.
---

## Base URL
DIRECT: `http://server.com:9000/code`

API HTTP: `http://server.com/vps/code`

API HTTP(S): `https://server.com/vps/code`
---

## General Usage
- **Endpoint:** `/NAMA_SCRIPT` (tanpa ekstensi `.sh`)
- **HTTP Method:** Tergantung pada fungsi (GET, POST, DELETE, dll.).
- **Input:** Biasanya JSON body jika menggunakan metode POST/DELETE.
- **Output:** Selalu dalam format JSON (menunjukkan `success` atau `error`).

---

## API by Protocol
Untuk detail lengkap setiap protokol, termasuk endpoint spesifik, parameter input, dan contoh output, silakan merujuk ke dokumentasi terpisah:

- **[SSH Documentation](docs/ssh.md)**
- **[Shadowsocks (SS) Documentation](docs/shadowsocks.md)**
- **[XRAY (VMess, VLESS, Trojan) Documentation](docs/xray.md)**
- **[NoobzVPN Documentation](docs/noobzvpn.md)**
- **[UDP Zivpn Documentation](docs/udpzivpn.md)**
- **[Wireguard WARP Documentation](docs/wireguard-warp.md)**

---

## Error Handling
Untuk informasi tentang penanganan error dan respons yang mungkin Anda terima, silakan lihat:

- **[Error Handling Documentation](docs/error-handling.md)**

---

## Notes
Beberapa catatan penting mengenai penggunaan API:

- **[Important Notes](docs/notes.md)**

---

Untuk pertanyaan lebih lanjut atau bantuan teknis, silakan hubungi tim dukungan kami.

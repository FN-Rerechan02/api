# API Shell Executor Documentation

## Overview
API ini menyediakan HTTP endpoint untuk mengeksekusi 51+ shell script backend layanan VPN/proxy: SSH, Shadowsocks (SS), VMess, VLESS, Trojan, Noobz, Zivpn, dsb. Semua endpoint diakses via Bearer Token dan output dalam format JSON.

---

## Table of Contents
- [Authentication](#authentication)
- [Base URL](#base-url)
- [General Usage](#general-usage)
- [API by Protocol](#api-by-protocol)
  - [SSH](#ssh)
  - [Shadowsocks (SS)](#shadowsocks-ss)
  - [VMess](#vmess)
  - [VLESS](#vless)
  - [Trojan](#trojan)
  - [Noobz](#noobz)
  - [Zivpn](#zivpn)
- [API by Feature](#api-by-feature)
  - [Create User](#create-user)
  - [Delete User](#delete-user)
  - [Extend User](#extend-user)
  - [Trial User](#trial-user)
  - [Check Login](#check-login)
- [Error Handling](#error-handling)
- [Example curl Commands](#example-curl-commands)
- [Notes](#notes)

---

## Authentication
Semua endpoint membutuhkan Bearer Token di header:
```
Authorization: Bearer <TOKEN>
```
Token disimpan di `/api/server/key`.

---

## Base URL

```
http://localhost/api/
```

---

## General Usage
- Endpoint: `/NAMA_SCRIPT` (tanpa .sh)
- HTTP method: GET, POST, DELETE, dsb (lihat fitur)
- Input: JSON body jika POST/DELETE
- Output: JSON (success/error)

---

## API by Protocol

### SSH
- **Create:** `add-ssh` (POST)
- **Delete:** `delete-ssh` (DELETE)
- **Extend:** `extend-ssh` (POST)
- **Trial:** `trial-ssh` (POST)
- **Check Login:** `cek-login-ssh` (GET)

### Shadowsocks (SS)
- **Create:** `add-ss` (POST)
- **Delete:** `delete-ss` (DELETE)
- **Extend:** `extend-ss` (POST)
- **Check Login:** `cek-ss` (GET)

  - **Create Output Example:**
    ```json
    {
      "status": "success",
      "message": "Shadowsocks account created successfully",
      "account_details": {
        "hostname": "yourdomain.com",
        "username": "ssuser1",
        "password": "ssuser1",
        "method": "aes-256-cfb",
        "tls_port": "2443",
        "http_port": "3443",
        "expired_on": "2025-06-08",
        "link_tls": "ss://...#ssuser1-TLS",
        "link_http": "ss://...#ssuser1-HTTP"
      }
    }
    ```
  - **Delete Output Example:**
    ```json
    {
      "status": "success",
      "message": "Shadowsocks account deleted successfully.",
      "username": "ssuser1",
      "expired_on": "2025-06-08",
      "tls_port": "2443",
      "http_port": "3443"
    }
    ```
  - **Extend Output Example:**
    ```json
    {
      "status": "success",
      "message": "Shadowsocks account renewed successfully.",
      "username": "ssuser1",
      "old_expired_on": "2025-06-08",
      "new_expired_on": "2025-06-13",
      "days_added": "5"
    }
    ```
  - **Check Login Output Example:**
    ```json
    {
      "status": "success",
      "timestamp": "2025-06-05T10:00:00Z",
      "server_info": {
        "domain": "yourdomain.com",
        "ipvps": "192.168.1.2"
      },
      "login_status": {
        "total_online_users": 2,
        "users": [
          {"username": "ssuser1", "active_ips": 1, "ips": ["1.2.3.4"]},
          {"username": "ssuser2", "active_ips": 2, "ips": ["5.6.7.8", "9.10.11.12"]}
        ]
      }
    }
    ```

### VMess
- **Create:** `add-vmess-ws`, `add-vmess-grpc`, `add-vmess-xhttp`, `add-vmess-hu` (POST)
- **Delete:** `delete-xray-ws`, `delete-xray-grpc`, `delete-xray-xhttp`, `delete-xray-hu` (DELETE)
- **Extend:** `extend-xray-ws`, `extend-xray-grpc`, `extend-xray-xhttp`, `extend-xray-hu` (POST)
- **Trial:** `add-trial-vmess-ws`, `add-trial-vmess-grpc`, `add-trial-vmess-xhttp`, `add-trial-vmess-hu` (POST)
- **Check Login:** `cek-login-xray-ws`, `cek-login-xray-grpc`, `cek-login-xray-xhttp`, `cek-login-xray-hu` (GET)

### VLESS
- **Create:** `add-vless-ws`, `add-vless-grpc`, `add-vless-xhttp`, `add-vless-hu` (POST)
- **Delete:** `delete-xray-ws`, `delete-xray-grpc`, `delete-xray-xhttp`, `delete-xray-hu` (DELETE)
- **Extend:** `extend-xray-ws`, `extend-xray-grpc`, `extend-xray-xhttp`, `extend-xray-hu` (POST)
- **Trial:** `add-trial-vless-ws`, `add-trial-vless-grpc`, `add-trial-vless-xhttp`, `add-trial-vless-hu` (POST)
- **Check Login:** `cek-login-xray-ws`, `cek-login-xray-grpc`, `cek-login-xray-xhttp`, `cek-login-xray-hu` (GET)

### Trojan
- **Create:** `add-trojan-ws`, `add-trojan-grpc`, `add-trojan-xhttp`, `add-trojan-hu` (POST)
- **Delete:** `delete-xray-ws`, `delete-xray-grpc`, `delete-xray-xhttp`, `delete-xray-hu` (DELETE)
- **Extend:** `extend-xray-ws`, `extend-xray-grpc`, `extend-xray-xhttp`, `extend-xray-hu` (POST)
- **Trial:** `add-trial-trojan-ws`, `add-trial-trojan-grpc`, `add-trial-trojan-xhttp`, `add-trial-trojan-hu` (POST)
- **Check Login:** `cek-login-xray-ws`, `cek-login-xray-grpc`, `cek-login-xray-xhttp`, `cek-login-xray-hu` (GET)

### Noobz
- **Create:** `add-noobz` (POST)
- **Delete:** `delete-noobz` (DELETE)
- **Extend:** `extend-noobz` (POST)
- **Check Login:** `cek-login-noobz` (GET)

### Zivpn
- **Create:** `add-zivpn` (POST)
- **Delete:** `delete-zivpn` (DELETE)
- **Extend:** `extend-zivpn` (POST)

---

## API by Feature

### Create User
- SSH: `add-ssh`
- SS: `add-ss`
- VMess: `add-vmess-ws`, `add-vmess-grpc`, `add-vmess-xhttp`, `add-vmess-hu`
- VLESS: `add-vless-ws`, `add-vless-grpc`, `add-vless-xhttp`, `add-vless-hu`
- Trojan: `add-trojan-ws`, `add-trojan-grpc`, `add-trojan-xhttp`, `add-trojan-hu`
- Noobz: `add-noobz`
- Zivpn: `add-zivpn`

### Delete User
- SSH: `delete-ssh`
- SS: `delete-ss`
- VMess: `delete-xray-ws`, `delete-xray-grpc`, `delete-xray-xhttp`, `delete-xray-hu`
- VLESS: `delete-xray-ws`, `delete-xray-grpc`, `delete-xray-xhttp`, `delete-xray-hu`
- Trojan: `delete-xray-ws`, `delete-xray-grpc`, `delete-xray-xhttp`, `delete-xray-hu`
- Noobz: `delete-zivpn`
- Zivpn: `delete-zivpn`

### Extend User
- SSH: `extend-ssh`
- SS: `extend-ss`
- VMess: `extend-xray-ws`, `extend-xray-grpc`, `extend-xray-xhttp`, `extend-xray-hu`
- VLESS: `extend-xray-ws`, `extend-xray-grpc`, `extend-xray-xhttp`, `extend-xray-hu`
- Trojan: `extend-xray-ws`, `extend-xray-grpc`, `extend-xray-xhttp`, `extend-xray-hu`
- Noobz: `extend-noobz`
- Zivpn: `extend-zivpn`

### Trial User
- SSH: `trial-ssh`
- VMess: `add-trial-vmess-ws`, `add-trial-vmess-grpc`, `add-trial-vmess-xhttp`, `add-trial-vmess-hu`
- VLESS: `add-trial-vless-ws`, `add-trial-vless-grpc`, `add-trial-vless-xhttp`, `add-trial-vless-hu`
- Trojan: `add-trial-trojan-ws`, `add-trial-trojan-grpc`, `add-trial-trojan-xhttp`, `add-trial-trojan-hu`

### Check Login
- SSH: `cek-login-ssh`
- SS: `cek-ss`
- VMess: `cek-login-xray-ws`, `cek-login-xray-grpc`, `cek-login-xray-xhttp`, `cek-login-xray-hu`
- VLESS: `cek-login-xray-ws`, `cek-login-xray-grpc`, `cek-login-xray-xhttp`, `cek-login-xray-hu`
- Trojan: `cek-login-xray-ws`, `cek-login-xray-grpc`, `cek-login-xray-xhttp`, `cek-login-xray-hu`
- Noobz: `cek-login-noobz`

---

## Error Handling
- **401 Unauthorized:**
  ```json
  {"message": "Unauthorized: Missing or invalid Bearer token"}
  ```
- **404 Not Found (script):**
  ```json
  Script not found
  ```
- **500 Internal Server Error:**
  ```json
  {"error": "<error message>"}
  ```

---

## Example curl Commands

### Create SSH User
```powershell
curl -X POST http://localhost/api/add-ssh `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"username":"userbaru","password":"passku","device":2,"expired":3}'
```

### Delete SSH User
```powershell
curl -X DELETE http://localhost/api/delete-ssh `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"username":"userbaru"}'
```

### Create Trial SSH User
```powershell
curl -X POST http://localhost/api/trial-ssh `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"expired":60}'
```

### Check SSH Login
```powershell
curl -X GET http://localhost/api/cek-login-ssh `
  -H "Authorization: Bearer <TOKEN>"
```

### Create Shadowsocks User
```powershell
curl -X POST http://localhost/api/add-ss `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"email":"ssuser1","expired":3}'
```

### Delete Shadowsocks User
```powershell
curl -X DELETE http://localhost/api/delete-ss `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"emaill":"ssuser1"}'
```

### Check Shadowsocks Login Status
```powershell
curl -X GET http://localhost/api/cek-ss `
  -H "Authorization: Bearer <TOKEN>"
```

### Create VMess WebSocket User
```powershell
curl -X POST http://localhost/api/add-vmess-ws `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"email":"user1@domain.com","quota":10,"device":2,"expired":7}'
```

### Delete VMess WebSocket User
```powershell
curl -X DELETE http://localhost/api/delete-xray-ws `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"email":"user1@domain.com"}'
```

### Check VMess WebSocket Login Status
```powershell
curl -X GET http://localhost/api/cek-login-xray-ws `
  -H "Authorization: Bearer <TOKEN>"
```

### Create VLESS User
```powershell
curl -X POST http://localhost/api/add-vless-ws `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"email":"user1@domain.com","quota":10,"device":2,"expired":7}'
```

### Delete VLESS User
```powershell
curl -X DELETE http://localhost/api/delete-xray-ws `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"email":"user1@domain.com"}'
```

### Check VLESS Login Status
```powershell
curl -X GET http://localhost/api/cek-login-xray-ws `
  -H "Authorization: Bearer <TOKEN>"
```

### Create Trojan User
```powershell
curl -X POST http://localhost/api/add-trojan-ws `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"email":"user1@domain.com","quota":10,"device":2,"expired":7}'
```

### Delete Trojan User
```powershell
curl -X DELETE http://localhost/api/delete-xray-ws `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"email":"user1@domain.com"}'
```

### Check Trojan Login Status
```powershell
curl -X GET http://localhost/api/cek-login-xray-ws `
  -H "Authorization: Bearer <TOKEN>"
```

### Create Noobz User
```powershell
curl -X POST http://localhost/api/add-noobz `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"username":"noobz1","expired":3}'
```

### Delete Noobz User
```powershell
curl -X DELETE http://localhost/api/delete-zivpn `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"username":"noobz1"}'
```

### Check Noobz Login Status
```powershell
curl -X GET http://localhost/api/cek-login-noobz `
  -H "Authorization: Bearer <TOKEN>"
```

### Extend SSH User
```powershell
curl -X POST http://localhost/api/extend-ssh `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"email":"user1@domain.com","expired":5}'
```

### Extend Shadowsocks User
```powershell
curl -X POST http://localhost/api/extend-ss `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"email":"ssuser1","expired":5}'
```

### Extend VMess WebSocket User
```powershell
curl -X POST http://localhost/api/extend-xray-ws `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"email":"user1@domain.com","expired":5}'
```

### Extend VLESS User
```powershell
curl -X POST http://localhost/api/extend-xray-ws `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"email":"user1@domain.com","expired":5}'
```

### Extend Trojan User
```powershell
curl -X POST http://localhost/api/extend-xray-ws `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"email":"user1@domain.com","expired":5}'
```

### Extend Noobz User
```powershell
curl -X POST http://localhost/api/extend-noobz `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"username":"noobz1","expired":5}'
```

### Create Trial VMess WebSocket User
```powershell
curl -X POST http://localhost/api/add-trial-vmess-ws `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"expired":60}'
```

### Create Trial VLESS WebSocket User
```powershell
curl -X POST http://localhost/api/add-trial-vless-ws `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"expired":60}'
```

### Create Trial Trojan WebSocket User
```powershell
curl -X POST http://localhost/api/add-trial-trojan-ws `
  -H "Authorization: Bearer <TOKEN>" `
  -H "Content-Type: application/json" `
  -d '{"expired":60}'
```

---

## Notes
- Semua script dieksekusi dari `/api/script/` di server. Pastikan script ada dan permission eksekusi.
- Format output dan parameter tergantung script. Cek script terkait untuk detail.
- Server API harus punya permission eksekusi system command tanpa password.
- Batasi akses API dan simpan token dengan aman.

---

Untuk detail lebih lanjut, cek source code setiap script di folder `api/`.

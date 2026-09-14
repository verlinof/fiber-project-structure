# Example Module Template

Folder ini berfungsi sebagai **blueprint / referensi template** untuk membuat modul baru pada arsitektur ini.

## Struktur Direktori

```
example/
├── http/
│   ├── handler.go      # HTTP handlers / controller logic
│   ├── type.go         # Struct handler & dependency injection
│   └── route/
│       └── main.go     # Route registration untuk Fiber Router
├── model/
│   ├── entity.go       # Model entitas database (GORM)
│   ├── request.go      # Struct request DTO & validation tags
│   └── response.go     # Struct response DTO & serialization
├── service/
│   ├── service.go      # Business logic implementation
│   └── type.go         # Struct service & dependency injection
└── README.md
```

## Langkah Membuat Modul Baru

1. **Duplikasi Folder**:
   Copy folder `internal/modules/example` dan ubah namanya menjadi nama modul baru Anda (contoh: `internal/modules/product`).

2. **Sesuaikan Package Name**:
   - `example_model` -> `product_model`
   - `example_service` -> `product_service`
   - `example_http` -> `product_http`
   - `example_route` -> `product_route`

3. **Definisikan Model & DTO** (`model/`):
   - Buat entitas database di `entity.go`.
   - Buat DTO request & validation di `request.go`.
   - Buat DTO response di `response.go`.

4. **Implementasikan Business Logic** (`service/`):
   - Tambahkan fungsi business logic di `service.go` atau buat file baru per fitur (misal: `create.go`, `get.go`).

5. **Implementasikan Handler HTTP** (`http/`):
   - Tangani parsing request, validasi, pemanggilan service, dan format response.

6. **Daftarkan Route** (`http/route/` & `internal/routes/main.go`):
   - Tentukan path dan HTTP method di `http/route/main.go`.
   - Inisialisasi service, handler, dan route pada `internal/routes/main.go`.

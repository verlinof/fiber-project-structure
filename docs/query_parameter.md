# Dokumentasi API: Filtering & Ordering

Sistem API ini menyediakan fungsionalitas filtering (penyaringan) dan ordering (pengurutan) yang fleksibel dan aman untuk mengambil data. Dokumen ini menjelaskan cara menggunakan parameter-parameter tersebut.

## 1. Cara Kerja

Sistem ini bekerja dengan mengubah parameter pada URL (query parameters) menjadi kondisi query ke database. Untuk menjaga keamanan dan performa, hanya field dan operator yang telah di-whitelist di backend yang dapat digunakan.

- **Filtering**: Menggunakan awalan `q_` pada query parameter.
- **Ordering**: Menggunakan satu query parameter bernama `order_by`.

## 2. Format Query Parameter

### Filtering

Format umum untuk filtering adalah:

```
q_{fieldName}_{operator}={value}
```

- `{fieldName}`: Nama alias dari field yang ingin difilter (lihat daftar di bawah).
- `{operator}`: Operator perbandingan yang ingin digunakan.
- `{value}`: Nilai untuk perbandingan.

Contoh:

- Menyaring material dengan nama yang mengandung "pasir": `?q_name_ilike=pasir`
- Menyaring material dengan harga lebih dari atau sama dengan 50000: `?q_price_gte=50000`
- Menyaring material berdasarkan beberapa ID kategori: `?q_categoryId_in=1,2`

### Operator yang Didukung

| Operator | Deskripsi         | Contoh Penggunaan     |
| :------- | :---------------- | :-------------------- |
| `eq`     | Sama dengan (`=`) | `q_status_eq=shipped` |

neq Tidak sama dengan (!=) q_status_neq=draft
gt Lebih besar dari (>) q_price_gt=100000
gte Lebih besar atau sama dengan (>=) q_price_gte=50000
lt Lebih kecil dari (<) q_price_lt=200000
lte Lebih kecil atau sama dengan (<=) q_price_lte=150000
like Pencarian teks (case-sensitive) q_name_like=Pasir
ilike Pencarian teks (case-insensitive) q_name_ilike=pasir
| `neq` | Tidak sama dengan (`!=`) | `q_status_neq=draft` |
<br>
| `gt` | Lebih besar dari (`>`) | `q_price_gt=100000` |
<br>
| `gte` | Lebih besar atau sama dengan (`>=`) | `q_price_gte=50000` |
<br>
| `lt` | Lebih kecil dari (`<`) | `q_price_lt=200000` |
<br>
| `lte` | Lebih kecil atau sama dengan (`<=`) | `q_price_lte=150000` |
<br>
| `like` | Pencarian teks (case-sensitive) | `q_name_like=Pasir` |
<br>
| `ilike` | Pencarian teks (case-insensitive) | `q_name_ilike=pasir` |
<br>
| `in` | Nilai ada di dalam daftar (dipisah koma) | `q_categoryId_in=1,2,5` |

### Ordering

Format umum untuk ordering adalah:

```
order_by={field1}:{direction},{field2}:{direction}
```

- `{field}`: Nama alias dari field yang ingin diurutkan (lihat daftar di bawah).
- `{direction}`: `asc` (ascending/menaik) atau `desc` (descending/menurun). Jika dikosongkan, default-nya adalah `asc`.

Contoh:

- Mengurutkan berdasarkan harga termurah: `?order_by=price:asc` atau `?order_by=price`
- Mengurutkan berdasarkan penjualan terbanyak, lalu yang terbaru: `?order_by=totalSales:desc,createdAt:desc`

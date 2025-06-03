# Tubes-Alpro-Semester-2
Judul: Program Manajemen Data Pegawai

Bahasa: Go (Golang)

Deskripsi Singkat:

File tubes.go adalah sebuah program berbasis terminal yang ditulis dalam bahasa Go untuk mengelola data pegawai. Program ini menyediakan menu interaktif untuk melakukan berbagai operasi terhadap data pegawai yang disimpan dalam array statis, seperti:

### Fitur Utama:

1. *Tambah Data Pegawai:* Menambahkan informasi pegawai baru.
2. *Tampilkan Data Pegawai:* Menampilkan seluruh data pegawai dalam bentuk tabel ASCII.
3. *Edit Data Pegawai:* Mengubah data pegawai berdasarkan ID menggunakan Sequential Search.
4. *Hapus Data Pegawai:* Menghapus data pegawai berdasarkan ID.
5. *Urutkan Gaji Tertinggi:* Mengurutkan data pegawai berdasarkan gaji secara menurun (descending) menggunakan algoritma Selection Sort.
6. *Urutkan Gaji Terendah:* Mengurutkan data pegawai berdasarkan gaji secara naik (ascending) menggunakan algoritma Insertion Sort.
7. *Total Gaji:* Menghitung total seluruh gaji pegawai.
8. *Cari Pegawai (Binary Search):* Mencari pegawai berdasarkan nama setelah data diurutkan secara alfabetis.
9. *Keluar Program:* Menampilkan ASCII art sebagai ucapan terima kasih sebelum keluar.

### Struktur Data:

* **Pegawai struct**: Menyimpan ID, nama, shift, dan gaji pegawai.
* **Array pegawaiList**: Menyimpan hingga 100 data pegawai.
* **Variabel jumlahData**: Menyimpan jumlah data pegawai yang saat ini aktif.

### Aspek Tambahan:

* Fungsi clearScreen() untuk membersihkan tampilan terminal agar lebih rapi.
* Tampilan antarmuka dengan padding agar semua output berada di tengah layar.

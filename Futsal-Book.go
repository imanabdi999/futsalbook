package main
import "fmt"
const nmax int = 999
type lapangan struct {
	id, jenis string
	tarif float64
}
type penyewa struct{
	id, nama, notelp string
}
type jadwal struct {
	id string
	tanggal, bulan int
	jamMulai int
	jamSelesai int
	tersedia bool
	idLapangan, idPenyewa string
	profit float64
}

type tabLapangan [nmax-1]lapangan
type tabPenyewa [nmax-1]penyewa
type tabJadwal [nmax-1]jadwal

func main() {
	var dataPenyewa tabPenyewa
	var dataLapangan tabLapangan
	var dataJadwal tabJadwal

	var nPenyewa int = 0
	var nLapangan int = 0
	var nJadwal int = 0
	var pilihan int
	var jalan bool = true

	for jalan {
		fmt.Println("\n=======================================")
		fmt.Println("       APLIKASI FUTSAL BOOK")
		fmt.Println("=======================================")
		fmt.Println("[ MANAJEMEN PENYEWA ]")
		fmt.Println(" 1. Tambah Penyewa      3. Ubah Penyewa")
		fmt.Println(" 2. List Penyewa      4. Hapus Penyewa")
		fmt.Println(" 5. Cari Penyewa (Nama)")
		fmt.Println(" 6. Cari Penyewa (No. Telp - Binary Search)")

		fmt.Println("\n[ MANAJEMEN LAPANGAN ]")
		fmt.Println(" 7. Tambah Lapangan     9. Ubah Lapangan")
		fmt.Println(" 8. List Lapangan    10. Hapus Lapangan")

		fmt.Println("\n[ TRANSAKSI & JADWAL ]")
		fmt.Println("11. Tambah Jadwal Kosong")
		fmt.Println("12. Tampilkan Jadwal Tersedia")
		fmt.Println("13. Booking Lapangan")
		fmt.Println("14. Tampilkan Jadwal Ter-booking")

		fmt.Println("\n[ PENGURUTAN JADWAL ]")
		fmt.Println("15. Urutkan Waktu (Selection Sort)")
		fmt.Println("16. Urutkan Tarif (Insertion Sort)")

		fmt.Println("\n[ STATISTIK ]")
		fmt.Println("17. Total Pendapatan Bulanan")
		fmt.Println("18. Jam Favorit Pelanggan")

		fmt.Println(" 0. KELUAR PROGRAM")
		fmt.Println("=======================================")
		fmt.Print("Pilih menu (0-18): ")
		fmt.Scan(&pilihan)

		fmt.Println()

		if pilihan == 1 {
			tambahPenyewa(&dataPenyewa, &nPenyewa)
		} else if pilihan == 2 {
			showPenyewa(dataPenyewa, nPenyewa)
		} else if pilihan == 3 {
			ubahPenyewa(&dataPenyewa, nPenyewa)
		} else if pilihan == 4 {
			hapusPenyewa(&dataPenyewa, &nPenyewa)
		} else if pilihan == 5 {
			searchNamaPenyewa(dataPenyewa, nPenyewa)
		} else if pilihan == 6 {
			searchNotelpPenyewa(dataPenyewa, nPenyewa)
		} else if pilihan == 7 {
			tambahLapangan(&dataLapangan, &nLapangan)
		} else if pilihan == 8 {
			showLapangan(dataLapangan, nLapangan)
		} else if pilihan == 9 {
			ubahLapangan(&dataLapangan, nLapangan)
		} else if pilihan == 10 {
			hapusLapangan(&dataLapangan, &nLapangan)
		} else if pilihan == 11 {
			tambahJadwal(&dataJadwal, &nJadwal)
		} else if pilihan == 12 {
			jadwalKosong(dataJadwal, nJadwal, dataLapangan, nLapangan)
		} else if pilihan == 13 {
			booking(&dataJadwal, nJadwal, nPenyewa, nLapangan, dataPenyewa, dataLapangan)
		} else if pilihan == 14 {
			jadwalBooking(dataJadwal, nJadwal)
		} else if pilihan == 15 {
			selectionSortJadwalMulai(&dataJadwal, nJadwal)
			jadwalKosong(dataJadwal, nJadwal, dataLapangan, nLapangan)
		} else if pilihan == 16 {
			insertionSortTarif(&dataJadwal, nJadwal, dataLapangan, nLapangan)
			jadwalKosong(dataJadwal, nJadwal, dataLapangan, nLapangan)
		} else if pilihan == 17 {
			totalPendapatanBulanan(dataJadwal, nJadwal)
		} else if pilihan == 18 {
			jamFavorit(dataJadwal, nJadwal)
		} else if pilihan == 0 {
			jalan = false
		} else {
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func tambahPenyewa(x *tabPenyewa, n *int){
	var in string
	fmt.Print("Input data penyewa (id, nama, nomor telepon)\nex: P01 Budi 62812345678\nKetik 'STOP' untuk berhenti.\n")
	fmt.Scan(&in)
	for in!="STOP" {
		x[*n].id = in 
		fmt.Scan(&x[*n].nama, &x[*n].notelp)
		*n++	
		fmt.Println("=> Jadwal berhasil ditambahkan!")
		fmt.Scan(&in)
	}
	fmt.Printf("Banyak data penyewa yang sudah masuk: %d\n", *n)
}

func ubahPenyewa(x *tabPenyewa, n int) {
	var target string
	var i, idx int
	idx = -1
	fmt.Print("Penyewa yang ingin diubah (ID): ")
	fmt.Scan(&target)
	for i<n && idx==-1 {
		if x[i].id==target{
			idx = i
		}
		i++
	}
	if idx==-1 {
		fmt.Println("Data tidak ditemukan.")
	} else {
		fmt.Println("Data ditemukan, silahkan input data baru (nama, nomor telepon)")
		fmt.Scan(&x[idx].nama, &x[idx].notelp)
	}
}

func hapusPenyewa(x *tabPenyewa, n *int) {
	var target string
	var i, idx int
	idx = -1
	showPenyewa(*x, *n)
	fmt.Print("Penyewa yang ingin dihapus (ID): ")
	fmt.Scan(&target)
	for i<*n && idx==-1 {
		if x[i].id==target{
			idx = i
		}
		i++
	}
	if idx==-1 {
		fmt.Println("Data tidak ditemukan.")
	} else {
		fmt.Println("Data ditemukan dan telah dihapus.")
		for i=idx; i<*n-1; i++ {
			x[i] = x[i+1]
		}
		*n = *n-1
	}
}

func showPenyewa(x tabPenyewa, n int){
	var i int
	if n==0 {
		fmt.Println("Data masih kosong")
	} else {
		fmt.Printf("%-4s %-10s %-10s\n", "ID", "Nama", "No Telp")
		for i=0; i<n; i++ {
			fmt.Printf("%-4s %-10s %-10s\n", x[i].id, x[i].nama, x[i].notelp)
		}
	}
}

func tambahLapangan(x *tabLapangan, n *int){
    var in string
    fmt.Print("Input data lapangan (id, jenis, tarif)\nex: L01 Sintetis 100000\nKetik 'STOP' untuk berhenti.\n")
    fmt.Scan(&in)
    for in != "STOP" {
        x[*n].id = in 
        fmt.Scan(&x[*n].jenis, &x[*n].tarif)
        *n++
		fmt.Println("=> Jadwal berhasil ditambahkan!")
        fmt.Scan(&in)
    }
    fmt.Printf("Banyak data lapangan yang sudah masuk: %d\n", *n)
}

func ubahLapangan(x *tabLapangan, n int) {
    var target string
    var i, idx int
    idx = -1
    fmt.Print("Lapangan yang ingin diubah (ID): ")
    fmt.Scan(&target)
    for i < n && idx == -1 {
        if x[i].id == target {
            idx = i
        }
        i++
    }
    if idx == -1 {
        fmt.Println("Data tidak ditemukan.")
    } else {
        fmt.Println("Data ditemukan, silahkan input data baru (jenis, tarif):")
        fmt.Scan(&x[idx].jenis, &x[idx].tarif)
    }
}

func hapusLapangan(x *tabLapangan, n *int) {
    var target string
    var i, idx int
    idx = -1
    fmt.Print("Lapangan yang ingin dihapus (ID): ") 
    fmt.Scan(&target)
    for i < *n && idx == -1 {
        if x[i].id == target {
            idx = i
        }
        i++
    }
    if idx == -1 {
        fmt.Println("Data tidak ditemukan.")
    } else {
        fmt.Println("Data ditemukan dan telah dihapus.")
        for i = idx; i < *n-1; i++ {
            x[i] = x[i+1]
        }
        *n = *n - 1
    }
}

func showLapangan(x tabLapangan, n int){
    var i int
	if n==0 {
		fmt.Println("Data masih kosong")
	} else {
		fmt.Printf("%-4s %-10s %-10s\n", "ID", "Jenis", "Tarif/Jam")
		for i=0; i<n; i++ {
			fmt.Printf("%-4s %-10s %-10.2f\n", x[i].id, x[i].jenis, x[i].tarif)
		}
	}
}

func tambahJadwal(x *tabJadwal, n *int) {
	var in, idL string
	var tgl, bulan, mulai, akhir int
	var valid, bentrok bool
	fmt.Print("Input data jadwal kosong (id, tanggal, bulan, jamMulai, jamSelesai, idLapangan)\nex: J01 5 12 19 21 L01\nKetik 'STOP' untuk berhenti.\n")
	fmt.Scan(&in)
	
	for in != "STOP" {
		fmt.Scan(&tgl, &bulan, &mulai, &akhir, &idL)
		valid = true
		//error handling
		if bulan < 1 || bulan > 12 {
			fmt.Println("ERROR: Bulan tidak valid. harus antara 1 - 12.")
			valid = false
		}
		if tgl < 1 || tgl > 31 {
			fmt.Println("ERROR: Tanggal tidak valid. harus antara 1 - 31.")
			valid = false
		}
		if mulai < 0 || mulai > 23 || akhir < 0 || akhir > 23 {
			fmt.Println("ERROR: Format jam tidak valid. harus antara 0 - 23.")
			valid = false
		} else if mulai >= akhir {
			fmt.Println("ERROR: Jam selesai harus lebih besar dari jam mulai")
			valid = false
		}

		if valid { // mengecek bentrok
			bentrok = cekBentrok(*x, *n, tgl, bulan, mulai, akhir, idL)
			if bentrok {
				valid = false
				fmt.Println("ERROR: Ada jadwal yang bertabrakan")
			}
		}

		if valid {
			x[*n].id = in
			x[*n].tanggal = tgl
			x[*n].bulan = bulan
			x[*n].jamMulai = mulai
			x[*n].jamSelesai = akhir
			x[*n].idLapangan = idL
			x[*n].tersedia = true
			*n++
			fmt.Println("Jadwal berhasil ditambahkan")
		} else {
			fmt.Println("Jadwal gagal ditambahkan")
		}
		fmt.Scan(&in)
	}
	fmt.Printf("Banyak slot jadwal operasional yang dibuat: %d\n", *n)
}

//fungsi khusus untuk tambahJadwal
func cekBentrok(x tabJadwal, n int, tgl int, bln int, mulai int, akhir int, idL string) bool {
	var i int
	for i = 0; i < n; i++ {
		if x[i].tanggal == tgl && x[i].bulan == bln && x[i].idLapangan == idL {
			if mulai < x[i].jamSelesai && akhir > x[i].jamMulai {
				return true
			}
		}
	}
	return false 
}

func booking(x *tabJadwal, nJadwal, nPenyewa, nLapangan int, y tabPenyewa, z tabLapangan){
	var penyewa, jadwal string
	var found bool
	var i, idxJ, idxL int
	fmt.Println("List penyewa:")
	showPenyewa(y, nPenyewa)
	fmt.Print("Masukkan ID penyewa: ")
	fmt.Scan(&penyewa)
	i=0
	for found==false && i<nPenyewa {
		if y[i].id == penyewa {
			found = true
		}
		i++
	}
	if found {
		found = false
		i=0
		fmt.Println("Jadwal yang tersedia: ")
		jadwalKosong(*x, nJadwal, z, nLapangan)
		fmt.Print("Masukkan ID Jadwal: ")
		fmt.Scan(&jadwal)
		for found==false && i<nJadwal {
			if x[i].id == jadwal {
				found = true
				idxJ = i
			}
			i++
		}
		if found {
			//mencari tarif
			for i=0; i<nLapangan; i++ {
				if z[i].id == x[idxJ].idLapangan {
					idxL = i
			}
			x[idxJ].idPenyewa = penyewa
			x[idxJ].tersedia = false
			x[idxJ].profit = x[idxJ].profit + (float64(x[idxJ].jamSelesai)-float64(x[idxJ].jamMulai))*z[idxL].tarif
			fmt.Println("Jadwal berhasil dibooking.")
			}
		}
	} else {
		fmt.Println("Data penyewa tidak ditemukan")
	}
}

func jadwalKosong(x tabJadwal, nJadwal int, z tabLapangan, nLap int) {
	var i int
	var tarif float64
	fmt.Printf("%-5s %-9s %-6s %-10s %-11s %-9s %-10s\n", "ID", "Tanggal", "Bulan", "Jam Mulai", "Jam Selesai", "Lapangan", "Tarif/Jam")
	
	for i = 0; i < nJadwal; i++ {
		if x[i].tersedia == true {
			tarif = getTarifLapangan(x[i].idLapangan, z, nLap)
			fmt.Printf("%-5s %-9d %-6d %-10d %-11d %-9s Rp%-10.0f\n", x[i].id, x[i].tanggal, x[i].bulan, x[i].jamMulai, x[i].jamSelesai, x[i].idLapangan, tarif)
		}
	}
}

func jadwalBooking(x tabJadwal, nJadwal int) {
	var i int
	fmt.Printf("%-5s %-9s %-6s %-10s %-11s %-9s %-9s %-12s\n", "ID", "Tanggal", "Bulan", "Jam Mulai", "Jam Selesai", "Lapangan", "Penyewa", "Profit")
	
	for i = 0; i < nJadwal; i++ {
		if x[i].tersedia == false {
			fmt.Printf("%-5s %-9d %-6d %-10d %-11d %-9s %-9s Rp%-12.0f\n", x[i].id, x[i].tanggal, x[i].bulan, x[i].jamMulai, x[i].jamSelesai, x[i].idLapangan, x[i].idPenyewa, x[i].profit)
		}
	}
}

func searchNamaPenyewa(x tabPenyewa, n int){
	//sequential search
	var i, idx int
	var target string
	fmt.Print("Cari data penyewa berdasarkan nama: ")
	fmt.Scan(&target)
	idx = -1
	for i<n && idx==-1{
		if x[i].nama==target {
			idx = i	
		}
		i++
	}
	if idx==-1 {
		fmt.Println("Data penyewa tidak ditemukan")
	} else {
		fmt.Printf("Data ditemukan di urutan ke-%d\n", idx+1)
		fmt.Printf("%-4s %-10s %-10s\n", "ID", "Nama", "No Telp")
		fmt.Printf("%-4s %-10s %-10s\n", x[idx].id, x[idx].nama, x[idx].notelp)
	}
}

func searchNotelpPenyewa(x tabPenyewa, n int){
	//binary search
	var left, right, mid, idx int
	var target string
	insertionSortPenyewaNotelp(&x, n)
	fmt.Print("Cari data penyewa berdasarkan No Telp: ")
	fmt.Scan(&target)
	left = 0
	right = n-1
	mid = (left+right)/2
	idx = -1
	for left <= right && idx==-1 {
		if target < x[mid].notelp {
			right = mid-1
		} else if target > x[mid].notelp {
			left = mid+1
		} else {
			idx = mid
		}
		mid = (left+right)/2
	}
	if idx==-1 {
		fmt.Println("Data penyewa tidak ditemukan")
	} else {
		fmt.Printf("Data ditemukan di urutan ke-%d\n", idx+1)
		fmt.Printf("%-4s %-10s %-10s\n", "ID", "Nama", "No Telp")
		fmt.Printf("%-4s %-10s %-10s\n", x[idx].id, x[idx].nama, x[idx].notelp)
	}	
}

//procedure khusus untuk procedure diatas karena memerlukan data yg terurut
func insertionSortPenyewaNotelp(x *tabPenyewa, n int) {
	var pass, i int
	var temp penyewa

	pass = 1
	for pass < n {
		temp = x[pass]
		i = pass
		for i > 0 && x[i-1].notelp > temp.notelp {
			x[i] = x[i-1]
			i--
		}

		x[i] = temp
		pass++
	}
} 

func selectionSortJadwalMulai(x *tabJadwal, n int) {
	var pass, i, idx int
	var temp jadwal

	pass = 1
	for pass < n {
		idx = pass-1
		i = pass
		for i < n {
			if x[i].bulan < x[idx].bulan {
				idx = i
			} else if x[i].bulan == x[idx].bulan {
				if x[i].tanggal < x[idx].tanggal {
					idx = i
				} else if x[i].tanggal == x[idx].tanggal {
					if x[i].jamMulai < x[idx].jamMulai {
						idx = i
					}
				}
			}
			i++
		}
		
		temp = x[pass-1]
		x[pass-1] = x[idx]
		x[idx] = temp
		pass++
	}
	fmt.Println("Data telah diurutkan berdasarkan jam mulai menggunakan")
}

func getTarifLapangan(idLap string, z tabLapangan, nLap int) float64 {
	var i int
	for i = 0; i < nLap; i++ {
		if z[i].id == idLap {
			return z[i].tarif
		}
	}
	return 0
}

func insertionSortTarif(x *tabJadwal, nJadwal int, z tabLapangan, nLap int) {
	var pass, i int
	var temp jadwal

	pass = 1
	for pass < nJadwal {
		temp = x[pass]
		i = pass

		for i > 0 && getTarifLapangan(x[i-1].idLapangan, z, nLap) > getTarifLapangan(temp.idLapangan, z, nLap) {
			x[i] = x[i-1]
			i--
		}

		x[i] = temp
		pass++
	}
	fmt.Println("Data telah diurutkan berdasarkan harga sewa")
}

func totalPendapatanBulanan(x tabJadwal, njadwal int){
	var i int
	var target int
	var total float64
	fmt.Print("Masukkan bulan (contoh: 5): ")
	fmt.Scan(&target)
	for i=0; i<njadwal; i++ {
		if x[i].tersedia==false && x[i].bulan==target {
			total = total + x[i].profit
		} 
	}
	fmt.Printf("Total pendapatan pada bulan %d: Rp%.2f\n", target, total)
}

func jamFavorit(x tabJadwal, nJadwal int){
	var freq [24]int
	var i, max, idx int
	idx = -1
	for i = 0; i<nJadwal; i++ {
		if x[i].tersedia == false {
			freq[x[i].jamMulai]++
		}
	}

	max = 0
	for i=0; i<24; i++ {
		if freq[i] > max {
			max = freq[i]
			idx = i
		}
	}

	if idx==-1{
		fmt.Println("Belum ada jadwal yang telah dibooking")
	} else {
		fmt.Printf("Jam yang paling sering dipesan adalah jam %d:00", idx)
	}
}
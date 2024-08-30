package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// pengguna bisa melakukan registrasi
// admin bisa melakukan yes/no registrasi akun dan mencetak daftar
// chat antar pengguna
// membuat grup chatting dan add anggota
// chat di grup
// bisa lihat anggota grup

const NMAX int = 100
const GMAX int = 1000

type user struct {
	data [NMAX]regist
	m    int
}

type regist struct {
	nama     string
	umur     int
	username string
	password string
}

// struct untuk pesan

type tabpesan struct {
	message [NMAX]pesan
	nMessage int
}

type pesan struct{
	penerima	string
	pengirim	string
	text		string
}

// struct untuk grup

type tabgroupchat struct {
	dataGrup	[GMAX]groupchat 
	nGrup		int
}

type groupchat struct { // isi dari data grupchat
	namaGrup     string
	anggotagrup  [NMAX]string // insert user.data[].username anggota
	chatGrup     [NMAX]string // nomor array disingkronkan dengan pengirimChat
	pengirimChat [NMAX]string // nomor array disingkronkan dengan chatGrup
	angootaGrupN int
	chatGrupN	 int
}
func main() {
	var banyakakun int
	var akunawal, akunacc user
	var chattingan tabpesan
	var group tabgroupchat
	
	dummyData(&akunawal, &akunacc, &banyakakun, &chattingan, &group)

	cetakawal()
	aturan()
	masukakun(&akunawal, &akunacc, &banyakakun, &chattingan, &group)
}


func dummyData(A, B *user, n *int, chattingan *tabpesan, group *tabgroupchat) {
	
	B.data[0] = regist{"Alice", 25, "alice", "123"}
	B.data[1] = regist{"Bob", 30, "bob", "123"}
	B.data[2] = regist{"Charlie", 35, "charlie", "123"}
	B.data[3] = regist{"David", 28, "david", "123"}
	B.data[4] = regist{"Eve", 22, "eve", "123"}
	B.data[5] = regist{"Frank", 27, "frank", "123"}
	B.data[6] = regist{"Grace", 29, "grace", "123"}
	B.data[7] = regist{"Hank", 33, "hank", "123"}
	B.data[8] = regist{"Ivy", 31, "ivy", "123"}
	B.data[9] = regist{"Jack", 24, "jack", "123"}
	B.m = 10	
	
	// Menambahkan data dummy pesan
	chattingan.message[0] = pesan{"alice", "bob", "Hi Bob, how are you?"}
	chattingan.message[1] = pesan{"bob", "alice", "I'm good, thank you!"}
	chattingan.message[2] = pesan{"charlie", "david", "Hello David!"}
	chattingan.message[3] = pesan{"david", "charlie", "Hi Charlie!"}
	chattingan.message[4] = pesan{"eve", "frank", "Hey Frank!"}
	chattingan.message[5] = pesan{"frank", "eve", "Hello Eve!"}
	chattingan.message[6] = pesan{"grace", "hank", "Hi Hank!"}
	chattingan.message[7] = pesan{"hank", "grace", "Hello Grace!"}
	chattingan.message[8] = pesan{"ivy", "jack", "Hey Jack!"}
	chattingan.message[9] = pesan{"jack", "ivy", "Hi Ivy!"}
	chattingan.nMessage = 10
	
	// Menambahkan data dummy grup
	group.dataGrup[0] = groupchat{
		namaGrup: "Group1",
		anggotagrup: [NMAX]string{"alice", "bob", "charlie"},
		chatGrup: [NMAX]string{"Welcome to the group", "Hello everyone", "Hi all!"},
		pengirimChat: [NMAX]string{"alice", "bob", "charlie"},
		angootaGrupN: 3,
		chatGrupN: 3,
	}
	group.dataGrup[1] = groupchat{
		namaGrup: "Group2",
		anggotagrup: [NMAX]string{"david", "eve", "frank"},
		chatGrup: [NMAX]string{"Welcome!", "Hi there!", "Good to see you!"},
		pengirimChat: [NMAX]string{"david", "eve", "frank"},
		angootaGrupN: 3,
		chatGrupN: 3,
	}
	group.dataGrup[2] = groupchat{
		namaGrup: "Group3",
		anggotagrup: [NMAX]string{"grace", "hank", "ivy", "jack"},
		chatGrup: [NMAX]string{"Hello group!", "Good morning!", "How's everyone?", "Hi!"},
		pengirimChat: [NMAX]string{"grace", "hank", "ivy", "jack"},
		angootaGrupN: 4,
		chatGrupN: 4,
	}
	group.nGrup = 3
}

func cetakawal() {
	//sapaan pembuka aplikasi chatting
	clearScreen()
	fmt.Println(strings.Repeat("-", 50))
	texta := "Wangsaf"
	textb := "Aplikasi Sejuta Pengguna"
	width := 50
	padding := (width - len(texta)) / 2
	paddingb := (width - len(textb)) / 2
	fmt.Println(strings.Repeat(" ", padding) + texta + strings.Repeat(" ", padding))
	fmt.Println(strings.Repeat(" ", paddingb) + textb + strings.Repeat(" ", paddingb))
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Loading...")
	fmt.Println(strings.Repeat("-", 50))
}

func aturan() {
	texta := "Cara Pemakaian"
	width := 50
	padding := (width - len(texta)) / 2
	fmt.Println(strings.Repeat(" ", padding) + texta + strings.Repeat(" ", padding))
	fmt.Println("1. Buat akun terlebih dahulu")
	fmt.Println("2. Akun bisa dibuat lebih dari satu")
	fmt.Println("3. Buat akun juga sebagai akun teman")
	fmt.Println("4. Persetujuan akun dilakukan dengan mode admin")
	fmt.Println("5. Mode admin hanya bisa diakses dengan password")
	fmt.Println("6. Untuk fitur yang keren, gunakan mode login")
	fmt.Println(strings.Repeat("-", 50))
}

func masukakun(A *user, B *user, n *int, chattingan *tabpesan, group *tabgroupchat) {
	//untuk mengetahui admin atau pengguna
	var akun int
	fmt.Println("Pilih mode:")
	fmt.Println("1. Buat Akun")
	fmt.Println("2. Login")
	fmt.Println("3. Admin")
	fmt.Println("4. Exit")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Print("Pilih 1/2/3/4: ")
	fmt.Scan(&akun)
	fmt.Println(strings.Repeat("-", 50))
	switch akun {
	case 1:
		jadipengguna(&*A, &*B, &*n, &*chattingan, &*group)
	case 2:
		login(&*A, &*B, &*n, &*chattingan, &*group)
	case 3:
		jadiadmin(&*A, &*B, &*n, &*chattingan, &*group)
	case 4:
		exit()
	}
}

//punyanya admin ===============================

func jadiadmin(A *user, B *user, n *int, chattingan *tabpesan, group *tabgroupchat) {
	//kalau jadi admin 
	fmt.Println("Selamat datang di mode admin")
	fmt.Println("Silakan masuk")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Loading...")
	fmt.Println(strings.Repeat("-", 50))
	passwordnya(&*A, &*B, &*n, &*chattingan, &*group)
}

func passwordnya(A *user, B *user, n *int, chattingan *tabpesan, group *tabgroupchat) {
	var pass string

	fmt.Print("Silakan masukkan password: ")
	fmt.Scan(&pass)
	if pass == "password" {
		fmt.Println("Password benar, silakan masuk")
		fmt.Println(strings.Repeat("-", 50))
		fituradmin(&*A, &*B, &*n, &*chattingan, &*group)
	} else {
		fmt.Println("Password salah, silakan kembali ke menu awal")
		masukakun(&*A, &*B, &*n, &*chattingan, &*group)
	}
}

func fituradmin(A *user, B *user, n *int, chattingan *tabpesan, group *tabgroupchat) {
	var akun int
	fmt.Println("Pilih mode:")
	fmt.Println("1. Persetujuan Pengguna")
	fmt.Println("2. Cetak Pengguna")
	fmt.Println("3. Home")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Print("Pilih 1/2/3: ")
	fmt.Scan(&akun)
	fmt.Println(strings.Repeat("-", 50))
	switch akun {
	case 1:
		cekakun(&*A, &*B, &*n, &*chattingan, &*group) //array dimasukkan ke dalam array baru yang memenuhi syarat
	case 2:
		cetakakun(&*A, &*B, &*n, &*chattingan, &*group)
	case 3:
		masukakun(&*A, &*B, &*n, &*chattingan, &*group)
	}

}

func cekakun(A *user, B *user, n *int, chattingan *tabpesan, group *tabgroupchat) {
	var i int
	var yatidak string
	var temp user

	fmt.Println("Persetujuan Pengguna")
	fmt.Println(strings.Repeat("-", 50))
	for i = 0; i < *n; i++ {
		fmt.Println("Nama:", A.data[i].nama)
		fmt.Println("Umur:", A.data[i].umur)
		fmt.Println("Username:", A.data[i].username)
		fmt.Print("Setuju untuk membuat akun (Ya/Tidak): ")
		fmt.Scan(&yatidak)
		fmt.Println(strings.Repeat("-", 50))
		if yatidak == "Ya" || yatidak == "YA" || yatidak == "ya" || yatidak == "yA" {
			fmt.Println("Akun disetujui")
			fmt.Println(strings.Repeat("-", 50))
			(*B).data[(*B).m] = (*A).data[i]
			(*B).m++
		}
	}
	*n = 0
	*A = temp
	fituradmin(&*A, &*B, &*n, &*chattingan, &*group)
}

func cetakakun(A *user, B *user, n *int, chattingan *tabpesan, group *tabgroupchat) {
	var akun int

	fmt.Println("Pilih mode:")
	fmt.Println("1. Cetak sesuai pendaftaran")
	fmt.Println("2. Cetak sesuai urutan umur")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Print("Pilih 1/2: ")
	fmt.Scan(&akun)
	fmt.Println(strings.Repeat("-", 50))
	switch akun {
	case 1:
		sesuaipendaf(&*A, &*B, &*n) //array dimasukkan ke dalam array baru yang memenuhi syarat
	case 2:
		sesuaiumur(&*A, &*B, &*n)
	}
}

func sesuaipendaf(A *user, B *user, n *int) {
	var i int
	var chattingan tabpesan
	var group tabgroupchat


	fmt.Println("+----+------------+------------+-------------+")
	fmt.Println("| No |    Nama    |    Umur    |   Username  |")
	fmt.Println("+----+------------+------------+-------------+")
	for i = 0; i < B.m; i++ {
		fmt.Printf("| %-2d | %-10s | %-10d | %-11s |\n", i+1, B.data[i].nama, B.data[i].umur, B.data[i].username)
	}
	fmt.Println("+----+------------+------------+-------------+")
	fmt.Println(strings.Repeat("-", 50))
	fituradmin(&*A, &*B, &*n, &chattingan, &group)
}

func sesuaiumur(A *user, B *user, n *int) {
	var akun int

	fmt.Println("Pilih mode:")
	fmt.Println("1. Urutkan umur naik")
	fmt.Println("2. Urutkan umur turun")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Print("Pilih 1/2: ")
	fmt.Scan(&akun)
	fmt.Println(strings.Repeat("-", 50))
	switch akun {
	case 1:
		umurnaik(&*A, &*B, &*n) //array dimasukkan ke dalam array baru yang memenuhi syarat
	case 2:
		umurturun(&*A, &*B, &*n)
	}

}

func umurnaik(A *user, B *user, n *int) {
	var i int
	var C user
	var pass int
	var temp user

	C = *B
	pass = 1
	for pass < C.m {
		i = pass
		temp.data[0] = C.data[pass]
		for i > 0 && temp.data[0].umur < C.data[i-1].umur {
			C.data[i] = C.data[i-1]
			i--
		}
		C.data[i] = temp.data[0]
		pass++
	}

	fmt.Println("+----+------------+------------+-------------+")
	fmt.Println("| No |    Nama    |    Umur    |   Username  |")
	fmt.Println("+----+------------+------------+-------------+")
	for i = 0; i < C.m; i++ {
		fmt.Printf("| %-2d | %-10s | %-10d | %-11s |\n", i+1, C.data[i].nama, C.data[i].umur, C.data[i].username)
	}
	fmt.Println("+----+------------+------------+-------------+")
	fmt.Println(strings.Repeat("-", 50))
	var chattingan tabpesan
	var group tabgroupchat
	fituradmin(&*A, &*B, &*n, &chattingan, &group)
}

func umurturun(A *user, B *user, n *int) {
	var i, idx int
	var C user
	var pass int
	var temp user

	C = *B
	pass = 1
	for pass <= C.m-1 {
		idx = pass - 1
		i = pass
		for i < C.m {
			if C.data[idx].umur < C.data[i].umur {
				idx = i
			}
			i++
		}
		temp.data[0] = C.data[pass-1]
		C.data[pass-1] = C.data[idx]
		C.data[idx] = temp.data[0]
		pass++
	}

	fmt.Println("+----+------------+------------+-------------+")
	fmt.Println("| No |    Nama    |    Umur    |   Username  |")
	fmt.Println("+----+------------+------------+-------------+")
	for i = 0; i < C.m; i++ {
		fmt.Printf("| %-2d | %-10s | %-10d | %-11s |\n", i+1, C.data[i].nama, C.data[i].umur, C.data[i].username)
	}
	fmt.Println("+----+------------+------------+-------------+")
	fmt.Println(strings.Repeat("-", 50))
	var chattingan tabpesan
	var group tabgroupchat
	fituradmin(&*A, &*B, &*n, &chattingan, &group)
}

//pengguna ================================================================

// bagian regist user ==================================

func jadipengguna(A *user, B *user, n *int, chattingan *tabpesan, group *tabgroupchat) {
	//kalau jadi pengguna
	fmt.Println("Selamat datang di mode pengguna")
	fmt.Println("Silakan mendafkarkan akun")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Loading...")
	fmt.Println(strings.Repeat("-", 50))
	daftarakun(&*A, &*B, &*n, &*chattingan, &*group)

}

func daftarakun(A *user, B *user, n *int, chattingan *tabpesan, group *tabgroupchat) {
	var akun regist

	fmt.Println("Isi data berikut:")
	fmt.Print("Nama: ")
	fmt.Scan(&akun.nama)
	fmt.Print("Umur: ")
	fmt.Scan(&akun.umur)
	fmt.Print("Username: ")
	fmt.Scan(&akun.username)
	fmt.Print("Password: ")
	fmt.Scan(&akun.password)
	fmt.Println(strings.Repeat("-", 50))

	if cekusername(akun, *A, *B, *n) == true {
		fmt.Println("Username telah digunakan")
		fmt.Println("Silakan ulangi pendaftaran")
		fmt.Println(strings.Repeat("-", 50))
		daftarakun(&*A, &*B, &*n, &*chattingan, &*group)
	} else {
		cekdata(akun, &*A, &*B, &*n)
	}

}

func cekusername(akun regist, A user, B user, n int) bool {
	var i int
	n = n + 1
	for i = 0; i < n; i++ {
		if akun.username == A.data[i].username {
			return true
		}
	}

	for i = 0; i < B.m+1; i++ {
		if akun.username == B.data[i].username {
			return true
		}
	}
	return false
}

func cekdata(akun regist, A *user, B *user, n *int) {
	//cek data
	var cek string
	var benar bool = false
	var chattingan tabpesan
	var group tabgroupchat

	fmt.Println("Nama:", akun.nama)
	fmt.Println("Umur:", akun.umur)
	fmt.Println("Username:", akun.username)
	fmt.Println("Apakah benar?")
	fmt.Print("Ketik 1 jika benar: ")
	fmt.Scan(&cek)
	if cek == "1" {
		fmt.Println(strings.Repeat("-", 50))
		fmt.Println("Silakan verifikasi")
		i := 1
		for i <= 5 && benar == false {
			kodeverifikasi(&benar)
		}
		if benar == true {
			A.data[*n] = akun
			*n++
			masukakun(&*A, &*B, &*n, &chattingan, &group)
		} else {
			fmt.Println("Pembuatan akun gagal")
			masukakun(&*A, &*B, &*n, &chattingan, &group)
		}
	} else {
		fmt.Println("Silakan mengisi ulang data")
		fmt.Println(strings.Repeat("-", 50))
		daftarakun(&*A, &*B, &*n, &chattingan, &group)
	}

}

func kodeverifikasi(cek *bool) {
	//kode verifikasi nih
	var angka, kode int

	rand.Seed(time.Now().UnixNano())

	angka = rand.Intn(99999 - 10000)
	fmt.Println("Kode verifikasi:", angka)
	fmt.Print("Masukkan kode verifikasi: ")
	fmt.Scan(&kode)
	if kode == angka {
		fmt.Println("Verifikasi Berhasil")
		fmt.Println(strings.Repeat("-", 50))
		*cek = true
	} else {
		fmt.Println("Kode salah")
		fmt.Println("Coba lagi")
		fmt.Println(strings.Repeat("-", 50))
		*cek = false
	}
}

// bagian login user ==================================

func login(A *user, B *user, n *int, chattingan *tabpesan, group *tabgroupchat) {
	//kalau jadi pengguna, tapi udah regist
	var akun regist
	var noAkun int

	fmt.Println("Selamat datang di mode pengguna")
	fmt.Println("Silakan masukan username dan password anda untuk masuk")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Loading...")
	fmt.Println(strings.Repeat("-", 50))
	
	fmt.Print("Username : ")
	fmt.Scan(&akun.username)
	fmt.Print("Password : ")
	fmt.Scan(&akun.password)
	
	if cekakunLogin(akun, *B, *n, &noAkun) {
		fiturlogin(&*A, &*B, &*n, noAkun, chattingan, group)
	} else {
		fmt.Println("Username atau Password yang anda masukan salah")
		fmt.Println("Silakan masukan kembali username dan password anda untuk masuk\n")
		login(&*A, &*B, &*n, chattingan, group)
	}
}

func cekakunLogin(akun regist, B user, n int, x *int) bool {
	n += 1
	for i := 0; i < B.m; i++ {
		if akun.username == B.data[i].username && akun.password == B.data[i].password {
			*x = i
			return true
		}
	}
	return false
}

func fiturlogin(A, B *user, n *int, x int, chattingan *tabpesan, group *tabgroupchat) {
	var pilih int

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Pilih mode:")
	fmt.Println("1. Kirim Pesan Pribadi")
	fmt.Println("2. Baca Pesan Pribadi")
	fmt.Println("3. Akses Grup")
	fmt.Println("4. Logout")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Print("Pilih 1/2/3/4: ")
	fmt.Scan(&pilih)
	fmt.Println(strings.Repeat("-", 50))
	switch pilih {
	case 1:
		kirimpesan(A, B, n, x, chattingan, group)
	case 2:
		bacapesan(A, B, n, x, chattingan, group)
	case 3:
		grup(&*A, &*B, &*n, x, &*chattingan, &*group)
	case 4:
		masukakun(&*A, &*B, &*n, &*chattingan, &*group)
	}
}

// fungsi pesan +++++++++++++++++++++++++++++++++++

func kirimpesan(A *user, B *user, n *int, x int, chattingan *tabpesan, group *tabgroupchat) { // diasumsikan bahwa 1 pesan dikirimkan dengan isi 1 kalimat panjang yang diakhiri oleh tanda "."
	var chitChat pesan
	var akun regist
	var pilihan, word string
	
	fmt.Print("Penerima : ")
	fmt.Scan(&akun.username)
	fmt.Println(strings.Repeat("-", 50))


	if cekusernameLogin(akun, *B, *n){
		deIdPesan := cariPesan(*A, *B, *n, x, *chattingan, *group)
		message := ""
		fmt.Println("Pesan harus diakhiri dengan ' .'")
		fmt.Print("Pesan    : ")
		fmt.Scan(&word)
		for string(word[len(word)-1]) != "." {
			message += word + " "
			fmt.Scan(&word)
		}
		chitChat.text = message

		fmt.Println(strings.Repeat("=", 50))
		fmt.Println(chitChat.text)
		fmt.Println(strings.Repeat("=", 50))
	
		fmt.Print("Kirim pesan? y/n : ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case "y":
			chattingan.message[deIdPesan].penerima	= akun.username
			chattingan.message[deIdPesan].pengirim	= B.data[x].username
			chattingan.message[deIdPesan].text		= chitChat.text
			fiturlogin(&*A, &*B, &*n, x, chattingan, group)
		case "n":
			kirimpesan(&*A, &*B, &*n, x, chattingan, group)
		}
	} else {
		fmt.Println("Penerima tidak ditemukan")
		fmt.Println("Silakan masukan kembali Username Penerima")
		fmt.Println(strings.Repeat("-", 50))
		kirimpesan(A, &*B, &*n, x, chattingan, group)
	} 
}

func bacapesan(A *user, B *user, n *int, x int, chattingan *tabpesan, group *tabgroupchat) {
	fmt.Println("Anda mendapatkan pesan baru dari : ")
	for i := 0; i < chattingan.nMessage; i++ {
		if chattingan.message[i].penerima == B.data[x].username {
			fmt.Print(i+1, ".", chattingan.message[i].pengirim, "\n")
			fmt.Print("Pesan: ")
			fmt.Println(chattingan.message[i].text)
		}
	}
	fmt.Println()
	fiturlogin(&*A, &*B, &*n, x, &*chattingan, &*group)
}

func cariPesan(A user, B user, n int, x int, chattingan tabpesan, group tabgroupchat) int {
	var found int = 0
	for i := 0; i < chattingan.nMessage+1; i++ {
		if (chattingan.message[i].penerima == B.data[x].username) || (chattingan.message[i].pengirim == B.data[x].username) {
			found = i
		}
	}
	if found == 0 {
		found = chattingan.nMessage
		chattingan.nMessage++
	}
	return found
}

func cekusernameLogin(akun regist, B user, n int) bool {
	var i int

	for i = 0; i < B.m+1; i++ {
		if akun.username == B.data[i].username {
			return true
		}
	}
	return false
}

// fungsi grup +++++++++++++++++++++++++++++++++++

func grup(A *user, B *user, n *int, x int, chattingan *tabpesan, group *tabgroupchat) {
	var pilih int

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("1. Buat Grup")
	fmt.Println("2. Buka Grup")
	fmt.Println("3. Hapus Grup")
	fmt.Println("4. Kembali ke menu chat")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Print("Pilih 1/2/3/4 : ")
	fmt.Scan(&pilih)
	fmt.Println(strings.Repeat("-", 50))
	switch pilih {
	case 1:
		buatGrup(&*A, &*B, &*n, x, &*chattingan, &*group)
	case 2:
		bukaGrup(&*A, &*B, &*n, x, &*chattingan, &*group)
	case 3:
		hapusGrup(&*A, &*B, &*n, x, &*chattingan, &*group)
	case 4:
		fiturlogin(&*A, &*B, &*n, x, &*chattingan, &*group)
	}
}

func buatGrup(A *user, B *user, n *int, x int, chattingan *tabpesan, group *tabgroupchat) {
	var groupTemp groupchat
	var pilih string
	var nGrup = group.nGrup //menyederhanakan penamaan variabel untuk berapa banyak grup yang ada

	fmt.Print("Silakan Masukan nama Grup yang akan dibuat : ")
	fmt.Scan(&groupTemp.namaGrup)
	if cekNamaGrup(groupTemp, *group, x) {
		groupTemp.anggotagrup[0] = B.data[x].username
		fmt.Print("Tambahkan anggota grup? Y/N: ")
		fmt.Scan(&pilih)
		i := 1
		for pilih == "Y" || pilih == "y" {
			fmt.Print("Masukkan username: ")
			fmt.Scan(&groupTemp.anggotagrup[i])
			
			groupTemp.angootaGrupN = i
			
			fmt.Print("Tambahkan lagi anggota grup? Y/N: ")
			fmt.Scan(&pilih)
			i++
			groupTemp.angootaGrupN = i
		}
		fmt.Println(strings.Repeat("-", 50))

		group.dataGrup[nGrup] = groupTemp // mengisi data grup
		nGrup++
		group.nGrup = nGrup // memperbaharui banyaknya grup
	}

	grup(&*A, &*B, &*n, x, &*chattingan, &*group)
}

func cekNamaGrup(groupTemp groupchat, group tabgroupchat, x int) bool {
	var i int
	var n = group.nGrup
	
	for i = 0; i < n; i++ {
		if groupTemp.namaGrup == group.dataGrup[i].namaGrup {
			return false
		}
	}

	return true
}

func bukaGrup(A *user, B *user, n *int, x int, chattingan *tabpesan, group *tabgroupchat) {
	var idgroup int
	
	cetakGrup(*B, x,*group)
	fmt.Println()
	fmt.Print("Pilih ID grup yang ingin dibuka\n(ketik 123 jika tidak ingin masuk ke grup) : ")
	fmt.Scan(&idgroup)
	fmt.Println(strings.Repeat("-", 50))
	if idgroup >= 0 {
		fiturGrup(&*A, &*B, &*n, x, &*chattingan, &*group, idgroup)
	} else if idgroup == 123 {
		grup(&*A, &*B, &*n, x, &*chattingan, &*group)
		fmt.Println()
	}
}

func fiturGrup(A *user, B *user, n *int, x int, chattingan *tabpesan, group *tabgroupchat, idgroup int) {
	// z berbeda dengan y, z mengartikan no array grup yang sedang dibuka
	var pilih int
	
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Di grup mau ngapain nich?")
	fmt.Println("1. Akses pesan dalam grup")
	fmt.Println("2. Lihat semua anggota grup")
	fmt.Println("3. Tambahkan/Hapus anggota")
	fmt.Println("4. Kembali ke menu login")
	fmt.Println(strings.Repeat("-", 50))

	fmt.Print("Pilih 1/2/3/4: ")
	fmt.Scan(&pilih)
	fmt.Println(strings.Repeat("-", 50))

	switch pilih {
	case 1:
		pesanGrup(A, B, n, x, chattingan, group, idgroup)
	case 2:
		cetakAnggotaGrup(*B, x, *group, idgroup)
	case 3:
		editAnggotaGrup(A, B, n, x, chattingan, group, idgroup)
	case 4:
		grup(A, B, n, x, chattingan, group)
	}
}

func pesanGrup(A *user, B *user, n *int, x int, chattingan *tabpesan, group *tabgroupchat, idgroup int) {
	var word string
	var pilih int
	var nchat = group.dataGrup[idgroup].chatGrupN // simplifikasi
	
	for i := 0; i < group.dataGrup[idgroup].chatGrupN; i++ {
		fmt.Println(group.dataGrup[idgroup].pengirimChat[i], ":", group.dataGrup[idgroup].chatGrup[i])
	}
	fmt.Println()
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("1. Kirim pesan")
	fmt.Println("2. Nggak ngirim pesan")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Print("Pilih 1/2: ")
	fmt.Scan(&pilih)
	fmt.Println(strings.Repeat("-", 50))

	switch pilih {
	case 1:
		message := ""
		fmt.Println("Pesan harus diakhiri dengan ' .'")
		fmt.Print("Kirim Pesan    : ")
		fmt.Scan(&word)
		for string(word[len(word)-1]) != "." { // diasumsikan bahwa 1 pesan dikirimkan dengan isi 1 kalimat panjang yang diakhiri oleh tanda "."
			message += word + " "
			fmt.Scan(&word)
		}

		fmt.Println(group.dataGrup[idgroup].chatGrup[nchat])
		
		group.dataGrup[idgroup].pengirimChat[nchat] = B.data[x].username
		group.dataGrup[idgroup].chatGrup[nchat] 	= message
		
		nchat++
		group.dataGrup[idgroup].chatGrupN = nchat
		pesanGrup(A, B, n, x, chattingan, group, idgroup)
	case 2:
		fiturGrup(A, B, n, x, chattingan, group, idgroup)
	}
}

func cetakAnggotaGrup(B user, x int, group tabgroupchat, idgroup int) {
	var A user
	var n int
	var chattingan tabpesan
	if group.dataGrup[idgroup].angootaGrupN > 0 {
		fmt.Println("Anggota Grup :")
		for i := 0; i < group.dataGrup[idgroup].angootaGrupN; i++ {
			fmt.Print(i+1, ". ", group.dataGrup[idgroup].anggotagrup[i], "\n")
		}
		fmt.Println()
		fiturGrup(&A, &B, &n, x, &chattingan, &group, idgroup)
	} else {
		fmt.Println("Tidak ada siapapun di grup ini")
		fmt.Println(strings.Repeat("-", 50))
		fmt.Println()
	}
}

func editAnggotaGrup(A *user, B *user, n *int, x int, chattingan *tabpesan, group *tabgroupchat, idgroup int) {
	var pilih int

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Pilih tindakan terhadap anggota: ")
	fmt.Println("1. Tambah Anggota")
	fmt.Println("2. Hapus Anggota")
	fmt.Println("3. Kembali ke menu grup")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Print("Pilih 1/2/3: ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		tambahAnggota(A, B, n, x, chattingan, group, idgroup)
	case 2:
		hapusAnggota(A, B, n, x, chattingan, group, idgroup)
	case 3:
		fiturGrup(A, B, n, x, chattingan, group, idgroup)
	}
	
	fmt.Println("Keanggotaan grup telah diubah")
	fmt.Println(strings.Repeat("-", 50))

	fiturGrup(A, B, n, x, chattingan, group, idgroup)
}

func tambahAnggota(A *user, B *user, n *int, x int, chattingan *tabpesan, group *tabgroupchat, idgroup int) { // add anggota
	var pilih string

	fmt.Println(strings.Repeat("-", 50))
	fmt.Print("Tambah anggota grup ? Y/N : ")
	fmt.Scan(&pilih)
	i := group.dataGrup[idgroup].angootaGrupN
	for pilih == "Y" || pilih == "y" {
		fmt.Print("Masukkan nama anggota yang akan ditambahkan:")
		fmt.Scan(&group.dataGrup[idgroup].anggotagrup[i])
		i++
		group.dataGrup[idgroup].angootaGrupN = i

		fmt.Print("Tambah lagi anggota grup? Y/N: ")
		fmt.Scan(&pilih)
	}
	fmt.Println(strings.Repeat("-", 50))
	fiturGrup(A, B, n, x, chattingan, group, idgroup)
}

func hapusAnggota(A *user, B *user, n *int, x int, chattingan *tabpesan, group *tabgroupchat, idgroup int) { // kick anggota
	var i, y, pilih int
	y = group.dataGrup[idgroup].angootaGrupN
	if group.dataGrup[idgroup].angootaGrupN > 0 {
		fmt.Println("Anggota Grup :")
		for i := 0; i < group.dataGrup[idgroup].angootaGrupN; i++ {
			fmt.Print(i+1, ". ", group.dataGrup[idgroup].anggotagrup[i], "\n")
		}
		fmt.Print("Pilih nomor anggota yang akan dihapus:")
		fmt.Scan(&pilih)
		for i = pilih-1; i < y-1; i++ {
			group.dataGrup[idgroup].anggotagrup[i] = group.dataGrup[idgroup].anggotagrup[i+1]
		}
		group.dataGrup[idgroup].angootaGrupN = group.dataGrup[idgroup].angootaGrupN - 1
	
		fmt.Println("Anggota telah berhasil dihapus")
		fmt.Println(strings.Repeat("-", 50))
	} else {
		fmt.Println("Tidak ada anggota di dalam grup")
	}
	fiturGrup(&*A, &*B, &*n, x, &*chattingan, &*group, idgroup)

}

func hapusGrup(A *user, B *user, n *int, x int, chattingan *tabpesan, group *tabgroupchat) {
	var i, pilih int
	nGrup := group.nGrup // penyederhanaan variabel terhadap banyak nya grup yg dimiliki user
	
	cetakGrup(*B, x, *group)
	fmt.Println(strings.Repeat("-", 50))
	fmt.Print("Pilih ID grup yang ingin dihapus: ")
	fmt.Scan(&pilih)

	for i = pilih; i < nGrup; i++ {
		group.dataGrup[i] = group.dataGrup[i+1]
	}
	group.nGrup = group.nGrup - 1

	fmt.Println("Grup telah berhasil dihapus")
	fmt.Println(strings.Repeat("-", 50))
	fiturlogin(A, B, n, x, chattingan, group)
}

func cetakGrup(B user, x int, group tabgroupchat) {
	var found = false 
	n := group.nGrup
	for i := 0; i < n; i++ {
		m := group.dataGrup[i].angootaGrupN
		for j := 0; j < m; j++ {
			if B.data[x].username == group.dataGrup[i].anggotagrup[j] {
				found = true
			}
		}
	}

	if found == true {
		fmt.Println("+----+----+-----------------+")
		fmt.Println("| No | ID |    Nama Grup    |")
		fmt.Println("+----+----+-----------------+")
		n := group.nGrup
		noGrup := 0
		for i := 0; i < n; i++ {
			m := group.dataGrup[i].angootaGrupN
			for j := 0; j < m; j++ {
				if B.data[x].username == group.dataGrup[i].anggotagrup[j] {
					fmt.Printf("| %-2d | %-2d | %-15s |\n", noGrup+1, i ,  group.dataGrup[i].namaGrup)
				}
			}
		}
		fmt.Println("+----+----+-----------------+")
	} else {
		fmt.Println("Anda tidak memasuki grup manapun")
		fmt.Println(strings.Repeat("-", 50))
	}
}
//exit ===================================================================================================


func exit() {
	fmt.Println(" ")
	fmt.Println(strings.Repeat("-", 50))
	texta := "Terima Kasih"
	textb := "Semoga harimu menyenangkan"
	textc := "Wangsaf"
	textd := "Created by Abil n Gabriel"
	width := 50
	padding := (width - len(texta)) / 2
	paddingb := (width - len(textb)) / 2
	paddingc := (width - len(textc)) / 2
	paddingd := (width - len(textd)) / 2
	fmt.Println(strings.Repeat(" ", padding) + texta + strings.Repeat(" ", padding))
	fmt.Println(strings.Repeat(" ", paddingb) + textb + strings.Repeat(" ", paddingb))
	fmt.Println(strings.Repeat(" ", paddingc) + textc + strings.Repeat(" ", paddingc))
	fmt.Println(strings.Repeat(" ", paddingd) + textd + strings.Repeat(" ", paddingd))
	fmt.Println(strings.Repeat("-", 50))
}

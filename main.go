package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// }
func main() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	clock := time.Now().In(loc).Format("15:04:05")
	// fmt.Println(clock)
	// menu(clock)
	recursif(clock)

}
func serve() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		loc, _ := time.LoadLocation("Asia/Jakarta")
		time := time.Now().In(loc).Format("15:04:05")

		c.JSON(http.StatusOK, gin.H{
			"message": time,
		})

	})

	r.Run()
}

func recursif(clock string) string {

	if clock == "00:00:00" {
		loc, _ := time.LoadLocation("Asia/Jakarta")
		clock = time.Now().In(loc).Format("15:04:05")
		didntReport(clock)
		changedDay(clock)
		log.Println(clock)
		fmt.Println(clock)
		duration := time.Second
		time.Sleep(duration)
		return recursif(clock)
	}
	if clock == "20:00:00" {
		loc, _ := time.LoadLocation("Asia/Jakarta")
		clock = time.Now().In(loc).Format("15:04:05")
		menu(clock)
		fmt.Println(clock)
		duration := time.Second
		time.Sleep(duration)
		return recursif(clock)
	}
	if clock == "08:15:00" {
		loc, _ := time.LoadLocation("Asia/Jakarta")
		clock = time.Now().In(loc).Format("15:04:05")
		changedDay(clock)
		fmt.Println(clock)
		duration := time.Second
		time.Sleep(duration)
		return recursif(clock)
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	clock = time.Now().In(loc).Format("15:04:05")
	fmt.Println(clock)

	duration := time.Second
	time.Sleep(duration)
	return recursif(clock)
}
func didntReport(time string) {
	sa := option.WithCredentialsFile("./ServiceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	dsnap := client.Collection("user").Where("role", "==", "user").Documents(context.Background())

	for {
		doc, err := dsnap.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		data := doc.Data()
		hari := int(data["hari"].(int64))

		id := data["id"].(string)
		// schedule, err := client.Collection("user").Doc(id).Collection("day").Doc("day" + strconv.Itoa(hari)).Get(context.Background())
		if err != nil {
			fmt.Println(err)
		}
		schedule, err := client.Collection("user").Doc(id).Collection("day").Doc("day" + strconv.Itoa(hari)).Get(context.Background())
		if err != nil {
			fmt.Println(err)
		}

		deadline := schedule.Data()
		if deadline["status"].(string) == "belum" {
			client.Collection("user").Doc(id).Collection("day").Doc("day"+strconv.Itoa(hari)).Set(context.Background(), map[string]interface{}{
				"status": "tidak",
			}, firestore.MergeAll)
			fmt.Println("success fill didnt report day")
		}

	}
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
}
func changedDay(time string) {
	sa := option.WithCredentialsFile("./ServiceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	dsnap := client.Collection("user").Where("role", "==", "user").Documents(context.Background())

	for {
		doc, err := dsnap.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		data := doc.Data()
		hari := int(data["hari"].(int64))

		id := data["id"].(string)
		// schedule, err := client.Collection("user").Doc(id).Collection("day").Doc("day" + strconv.Itoa(hari)).Get(context.Background())
		if err != nil {
			fmt.Println(err)
		}

		if hari < 30 {
			client.Collection("user").Doc(id).Set(context.Background(), map[string]interface{}{
				"hari": hari + 1,
			}, firestore.MergeAll)
			fmt.Println("success changed day")
		}

		// client.Collection("user").Doc(id).Collection("day").Doc("day"+strconv.Itoa(hari)).Set(context.Background(), map[string]interface{}{
		// 	"status": "tidak",
		// }, firestore.MergeAll)
		fmt.Println(hari)

	}
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
}
func menu(time string) {
	//Use Credential Firebase
	tips := map[string]interface{}{
		"1":  "Yth.Ibu Jangan lupa melakukan pemeriksaan kehamilan di Bidan setiap bulan atau jika mengalami tanda bahaya selama kehamilan.",
		"2":  "Tanda bahaya dalam kehamilan: keluar darah dari jalan lahir, ketuban pecah sebelum tanda persalinan, bengkak seluruh tubuh, sakit kepala hebat, kejang, demam lebih dari 2 hari, berat badan tidak naik.",
		"3":  "Selamat malam ibu. Jangan lupa minum obat tambah darah untuk mencegah anemia agar terhindar dari perdarahan banyak setelah melahirkan. Obat tambah darah tidak boleh diminum bersamaan dengan teh, kopi, susu karena dapat merusak khasiat obat.",
		"4":  "Selamat pagi ibu. Semoga ibu dan bayi dalam keadaan sehat. Selama hamil makanlah apa yang ibu sukai, jangan dipantang agar ibu dan bayi dalam kandungan sehat.",
		"5":  "Makanlah makanan dengan menu seimbang: nasi, sayur, lauk (ikan, tempe, telur, ayam), buah-buah, susu dan air putih minimal 8 gelas dalam sehari.",
		"6":  "Selamat malam ibu. Sudahkah ibu minum obat tambah darahnya malam ini? Kalau belum, diminum ya bu agar ibu dan bayi tetap sehat dan terhindar dari perdarahan banyak setelah melahirkan. Apabila di pagi hari BAB ibu berwarna hitam, ibu jangan khawatir.",
		"7":  "Tahukah ibu? Anemia pada kehamilan dapat menyebabkan keguguran, persalinan prematur, serta gangguan pertumbuhan janin.",
		"8":  "Perlu dikenali beberapa tanda dan gejala anemia pada ibu hamil agar dapat dideteksi secara mandiri dan sedini mungkin: wajah, terutama kelopak mata dan bibir tampak pucat. Mata berkunang-kunang, serta sering mengalami pusing. Jangan disepelakan ya bu. Segera kungjungi pusat kesehatan terkait jka mengalami salah satu gejala diatas.",
		"9":  "Perlu diketahui bu, bahwa dalam masa pandemi, konsumsi Tablet Tambah Darah  sangat penting, untuk mencegah anemia sekaligus meningkatkan kekebalan tubuh terhadap virus corona.",
		"10": "Bagaimana sih Biar Tidak Anemia? Makan Makanan BERGIZI SEIMBANG, Terutama Tinggi Protein, Kaya zat besi Minum tablet tambah darah secara teratur.",
		"11": "Ingat ya bu, JANGAN minum tablet tambah darah dengan teh, kopi, atau susu, karena akan MENGHAMBAT penyerapan zat besi.",
		"12": "Efek samping yang mungkin dialami ibu hamil dari mengonsumsi tablet tambah darah yaitu nyeri/perih di ulu hati, mual serta tinja berwarna kehitaman (yang berasal dari sisa zat besi yang dikeluarkan oleh tubuh melalui feses). Efek samping tidak sama dialami oleh setiap orang dan akan hilang dengan sendirinya.",
		"13": "Jangan lupa bu! Tes laboratorium yang wajib dilakukan adalah tes hemoglobin darah (Hb) untuk mengetahui apakah ibu hamil menderita anemia, golongan darah untuk mempersiapkan donor bila diperlukan kelak.",
		"14": "Sebaiknya ibu tidur pakai kelambu, jangan memakai anti nyamuk (bakar atau semprot).",
		"15": "Selamat pagi ibu. Semoga ibu dan bayi tetap dalam keadaan sehat. Melahirkanlah ke tenaga kesehatan: bidan praktik, puskesmas dan Rumah Sakit untuk menghindari terjadinya masalah saat melahirkan. Terima Kasih.",
		"16": "Selamat pagi ibu. Semoga ibu dan bayi yang dikandung selalu dalam keadaan sehat. Sudah berapa bulan usia kehamilannya sekarang bu? Marilah kita mengenali tanda-tanda bahwa ibu akan melahirkan, yaitu: Sakit perut sampai ke pinggang, keluar lendir bercampur darah, keluar air ketuba, perut kencang, jalan lahir terbuka.",
		"17": "Ibu hamil setiap kali makan harus mengonsumsi makanan yang mengandung protein, karbohidrat dan zat gizi mikro (vitamin dan mineral) agar terhindar dari resiko anemia karena pola makan yang kurang beragam dan tidak bergizi seimbang merupakan faktor utama penyebab anemia dalam kehamilan.",
		"18": "Pada masa pandemi covid-19, ibu hamil perlu memperhatikan: Konsumsi makanan bergizi seimbang (isi priringku), Kebiasaan selalu mencuci tangan pakai sabun dan air mengalir (Perilaku Hidup Bersih dan Sehat, Aktivitas fisik ringan sesuai untuk ibu hamil, Minum air putih minimal 8 gelas sehari",
		"19": "Bila perut ibu hamil terasa perih karena meminum tablet tambah darah, mual serta tinja berwarna kehitaman, tidak perlu khawatir karena tubuh akan menyesuaikan. Untuk meminimalkan efek samping tersebut, jangan minum tablet tambah darah  dalam kondisi perut kosong.",
		"20": "Sebaiknya tablet tambah darah diminum setelah makan malam agar konsumsi tablet tambah darah dapat lebih efektif untuk mencegah anemia. Tablet tambah darah dapat dikonsumsi bersama makanan atau minuman yang mengandung Vitamin C seperti buah segar, sayuran dan jus buah, agar penyerapan zat besi didalam tubuh lebih baik.",
		"21": "Yang harus dipersiapkan sebelum melahirkan: tahu tanggal melahirkan, siapa yang menolong melahirkan, siapkan uang, siapkan kendaraan, siapkan orang yang akan mendonorkan darah dan semua keperluan melahirkan.",
		"22": "Tanda-tanda melahirkan: Sakit perut sampai ke pinggang, keluar lendir bercampur darah, keluar air ketuban, perut kencang, jalan lahir tebuka.",
		"23": "Tanda-tanda bahaya waktu melahirkan:  bayi tidak keluar dalam 12 jam dari perut sakit, perdarahan banyak dari jalan lahir, tali pusat dan tangan bayi keluar dan jalan lahir, ibu tidak kuat mengedan atau mengalami kejang, air ketuban keruh atau busuk.",
		"24": "Tanda bahaya bayi baru lahir: bayi tidak mau menyusu, kejang, lemah, sesak nafas, pusat kemerahan, demam tinggi atau badan terasa dingin.",
		"25": "Yang harus dihindari selama kehamilan: Kerja berat, Merokok dan terpapar asap rokok selama kehamilan karena akan mengganggu pertumbuhan janin, Mengonsumsi minuman yang mengandung soda, alkohol dan lain lain, Tidur terlentang pada hamil tua, Mengonsumsi obat tanpa resep dokter apabila ada keluhan",
		"26": "Persiapkan pendonor yang memiliki golongan darah yang sama dengan ibu bersalin, upayakan pendonor berasal dari orang terdekat (keluarga atau tetangga).",
		"27": "Bersama dengan suami lakukan stimulasi janin dengan cara: Stimulasi suara, berbicara dengan janin sejak hamil muda dengan kata-kata yang lemah lembut dan positif. Stimulasi raba, lakukan sentuhan dengan cara mengusap perut ibu hamil sesering mungkin.",
		"28": "Selamat pagi ibu. Jangan lupa ke posyandu untuk mengetahui kesehatan ibu dan bayi di dalam kandungan.",
		"29": "Simpan tablet tambah darah di tempat yang kering, terhindar dari sinar matahari langsung, jauh dari jangkauan anak dan setelah dibuka harus ditutup kembali dengan rapat tablet Tambah darah yang sudah berubah warna sebaiknya tidak di minum (warna asli: merah darah).",
		"30": "Selamat pagi ibu. Tablet tambah darah tidak menyebabkan tekanan darah tinggi atau kelebihan darah.",
	}
	a := ""
	sa := option.WithCredentialsFile("./ServiceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	dsnap := client.Collection("user").Where("role", "==", "user").Documents(context.Background())

	for {
		doc, err := dsnap.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		data := doc.Data()
		nomorhp := data["nomorhp"]
		nomorhp2 := nomorhp.(string)
		hari := int(data["hari"].(int64))

		// id := data["id"].(string)
		// schedule, err := client.Collection("user").Doc(id).Collection("day").Doc("day" + strconv.Itoa(hari)).Get(context.Background())
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// deadline := schedule.Data()
		// if time == "07:11:00" {
		// 	client.Collection("user").Doc(id).Set(context.Background(), map[string]interface{}{
		// 		"hari": hari + 1,
		// 	}, firestore.MergeAll)
		// 	fmt.Println("success changed day")
		// }

		// client.Collection("user").Doc(id).Collection("day").Doc("day"+strconv.Itoa(hari)).Set(context.Background(), map[string]interface{}{
		// 	"status": "tidak",
		// }, firestore.MergeAll)
		a = "day" + strconv.Itoa(hari)
		fmt.Println(hari)
		// if time == "06:35:00" {
		// 	if deadline["status"].(string) == "belum" {
		// 		client.Collection("user").Doc(id).Collection("day").Doc("day"+strconv.Itoa(hari)).Set(context.Background(), map[string]interface{}{
		// 			"status": "tidak",
		// 		}, firestore.MergeAll)
		// 		a = "a"
		// 		fmt.Println(a)

		// 	}
		// 	client.Collection("user").Doc(id).Collection("day").Doc("day"+strconv.Itoa(hari)).Set(context.Background(), map[string]interface{}{
		// 		"status": "tidak",
		// 	}, firestore.MergeAll)
		// 	a = "b"
		// 	fmt.Println(a)
		// 	if hari < 30 {
		// 		client.Collection("user").Doc(id).Set(context.Background(), map[string]interface{}{
		// 			"hari": hari + 1,
		// 		}, firestore.MergeAll)
		// 		a = "Report and day changed success"
		// 		fmt.Println(a)
		// 	}

		// }

		test("Selamat malam ibu..\nIngin mengingatkan sekarang sudah memasuki jam 20.00 waktunya ibu untuk meminum obat. Segera melakukan konfirmasi melalui aplikasi Mamiana dengan mengupload obat yang akan diminum hari ini\nTerimakasih", nomorhp.(string))
		test("Tips : \n\n"+tips[strconv.Itoa(hari)].(string), nomorhp2)

	}
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(a)
	defer client.Close()
}

func test(message, phone string) {
	headers := map[string][]string{
		"Content-Type": []string{"application/json"},
		"Accept":       []string{"application/json"},
		"APIKey":       []string{"56E4E650EF96812F367FB6096B806DE7"},
	}
	payload, _ := json.Marshal(map[string]string{
		// Sender is optional
		// "sender": "YOUR_SENDER",
		"destination": phone,
		"message":     message,
	})
	url := "https://api.nusasms.com/nusasms_api/1.0/whatsapp/message"
	// For testing
	// url := "https://dev.nusasms.com/nusasms_api/1.0/whatsapp/message"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header = headers

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	// ...
}

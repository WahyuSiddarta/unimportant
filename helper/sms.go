package helper

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	zenzivaUserKey = "yxwebs"
	zenzivaPassKey = "upszkk1nct"
)

//ZenzivaDomainSMSOTP :
func SendSMSOTP(phone, code string) int {
	phone1 := strings.ReplaceAll(phone, ".", "")
	msg := "RuangNyaman \n<< " + code + " >>\njangan berikan kode ini kepada siapapun."
	hurl := "https://console.zenziva.net/reguler/api/sendsms/"
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 15 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 30,
		Transport: netTransport,
	}
	response, _ := netClient.PostForm(hurl, url.Values{"userkey": {zenzivaUserKey}, "passkey": {zenzivaPassKey}, "to": {phone1}, "message": {msg}})
	fmt.Printf("%+v", response)
	return response.StatusCode
}

func SendSmsBooking(phone string, booking_id string) int {
	phone1 := strings.ReplaceAll(phone, ".", "")
	msg := "Pemesanan Anda di ruangnyaman.id terkonfirmasi dengan booking ID " + booking_id + ". Segera lakukan pembayaran sebelum batas waktu berakhir. CS : +62818228828"
	hurl := "https://console.zenziva.net/reguler/api/sendsms/"
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 15 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 30,
		Transport: netTransport,
	}
	response, _ := netClient.PostForm(hurl, url.Values{"userkey": {zenzivaUserKey}, "passkey": {zenzivaPassKey}, "to": {phone1}, "message": {msg}})
	fmt.Printf("%+v", response)
	return response.StatusCode
}

func SendSmsPaymentSuccess(phone string, booking_id string) int {
	phone1 := strings.ReplaceAll(phone, ".", "")
	msg := "Terima kasih atas kepercayaan anda memilih ruangnyaman.id,\nPemesanan kamu berhasil diverifikasi dengan Booking ID " + booking_id + ". CS : +62818228828"
	hurl := "https://console.zenziva.net/reguler/api/sendsms/"
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 15 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 30,
		Transport: netTransport,
	}
	response, _ := netClient.PostForm(hurl, url.Values{"userkey": {zenzivaUserKey}, "passkey": {zenzivaPassKey}, "to": {phone1}, "message": {msg}})
	fmt.Printf("%+v", response)
	return response.StatusCode
}

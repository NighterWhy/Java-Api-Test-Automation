package com.odev;

import static io.restassured.RestAssured.*;
import static org.hamcrest.Matchers.*;
import org.junit.Test;

public class AppTest {

        /// LOGİN TESTİ

        @Test
        public void loginBasariliTesti() {
                baseURI = "http://localhost:8080";

                String loginData = "{\"email\": \"serhat@test.com\", \"password\": \"123456\"}";

                given()
                                .contentType("application/json")
                                .body(loginData)
                                .when()
                                .post("/users/login")
                                .then()
                                .statusCode(200)
                                .body("message", equalTo("Giriş Başarılı!"));
        }

        /// KAYIT TESTİ - NAME KISMINDA SORUN VAR
        @Test
        public void kayitTesti() {
                baseURI = "http://localhost:8080";

                String requestBody = "{" +
                                "\"name\": \"mehmet\"," +
                                "\"email\": \"mehmet_v2" + System.currentTimeMillis() + "@test.com\"," +
                                "\"password\": \"123456\"" +
                                "}";

                given()
                                .header("Content-Type", "application/json")
                                .body(requestBody)
                                .when()
                                .post("/users/register")
                                .then()
                                .statusCode(201) // Go'dan gelen 201'i onaylıyoruz
                                .body("message", equalTo("Kullanici Veri Tabanina Kaydedildi..")); // Go'daki tam mesaj
        }

        @Test
        public void sahaListelemeVeVeriKontrolu() {
                baseURI = "http://localhost:8080";

                given()
                                .when()
                                .get("/sahalar/")
                                .then()
                                .statusCode(200)
                                .body("[0].price", greaterThan(0));
        }

        @Test
        public void sahalariListelemeTesti() {
                baseURI = "http://localhost:8080";

                given()
                                .when()
                                .get("/sahalar/")
                                .then()
                                .statusCode(200)
                                .body("size()", greaterThanOrEqualTo(0))
                                .time(lessThan(1500L));
        }

        @Test
        public void yeniSahaOlusturmaTesti() {
                baseURI = "http://localhost:8080";

                String sahaBody = "{" +
                                "\"name\": \"Şampiyon Halı Saha v3\"," +
                                "\"location\": \"İstanbul\"," +
                                "\"price\": 50" +
                                "}";

                given()
                                .header("Content-Type", "application/json")
                                .body(sahaBody)
                                .when()
                                .post("/sahalar/")
                                .then()
                                .statusCode(201)
                                .body("name", equalTo("Şampiyon Halı Saha v3"))
                                .time(lessThan(2000L));
        }

        @Test
        public void rezervasyonOlusturmaTesti() {
                baseURI = "http://localhost:8080";

                
                String rezervasyonPayload = "{" +
                                "\"user_id\": 1," +
                                "\"saha_id\": 1," +
                                "\"tarih\": \"2025-12-20T20:00:00Z\"," +
                                "\"saat\": \"20:00\"" +
                                "}";

                given()
                                .header("Content-Type", "application/json")
                                .body(rezervasyonPayload)
                                .when()
                                .post("/rezervasyonlar/")
                                .then()
                                .statusCode(201) // Kaydedildi kodu
                                .body("message", equalTo("Rezervasyon Başarıyla Oluşturuldu"))
                                .time(lessThan(2000L));
        }

        @Test
        public void rezervasyonCakismaTesti() {
                baseURI = "http://localhost:8080";

                String ayniRandevu = "{" +
                                "\"user_id\": 1," +
                                "\"saha_id\": 1," +
                                "\"tarih\": \"2025-12-20T20:00:00Z\"," +
                                "\"saat\": \"20:00\"" +
                                "}";

                // Birinciyi gönderiyoruz (Sistem kabul edecek - 201)
                given().contentType("application/json").body(ayniRandevu).post("/rezervasyonlar/");

                // İkinciyi gönderiyoruz (Sistem çakışma diyecek - 409)
                given()
                                .contentType("application/json")
                                .body(ayniRandevu)
                                .when()
                                .post("/rezervasyonlar/")
                                .then()
                                .statusCode(409) // http.StatusConflict
                                .body("error", equalTo("bu saat için zaten rezervasyon bulunuyor"));
        }
}
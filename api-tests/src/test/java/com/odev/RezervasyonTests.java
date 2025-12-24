package com.odev;

import static io.restassured.RestAssured.*;
import static org.hamcrest.Matchers.*;
import org.junit.Test;

public class RezervasyonTests extends BaseTest {

    @Test
    public void rezervasyonOlusturmaTesti() {
        String payload = "{\"user_id\": 1, \"saha_id\": 1, \"tarih\": \"2025-12-20T20:00:00Z\", \"saat\": \"17:00\"}";

        given()
            .contentType("application/json")
            .body(payload)
        .when()
            .post("/rezervasyonlar/")
        .then()
            .statusCode(201)
            .body("message", equalTo("Rezervasyon Başarıyla Oluşturuldu"))
            .time(lessThan(2000L));
    }
    
    
}
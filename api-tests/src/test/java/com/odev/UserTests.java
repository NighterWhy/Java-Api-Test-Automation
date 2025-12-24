package com.odev;

import static io.restassured.RestAssured.*;
import static org.hamcrest.Matchers.*;
import org.junit.Test;

public class UserTests extends BaseTest {

    @Test
    public void loginBasariliTesti() {
        String loginData = "{\"email\": \"serhat@test.com\", \"password\": \"123456\"}";

        given()
            .contentType("application/json")
            .body(loginData)
        .when()
            .post("/users/login")
        .then()
            .statusCode(200)
            .body("message", equalTo("Giriş Başarılı!"))
            .time(lessThan(2000L));
    }

    @Test
    public void kayitTesti() {
        String requestBody = "{" +
                "\"name\": \"mehmet\"," +
                "\"email\": \"mehmet_v3" + System.currentTimeMillis() + "@test.com\"," +
                "\"password\": \"123456\"" +
                "}";
        given()
            .contentType("application/json")
            .body(requestBody)
        .when()
            .post("/users/register")
        .then()
            .statusCode(201) 
            .body("message", equalTo("Kullanici Veri Tabanina Kaydedildi.."))
            .time(lessThan(2000L));
    }
}
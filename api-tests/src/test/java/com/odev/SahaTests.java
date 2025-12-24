package com.odev;

import static io.restassured.RestAssured.*;
import static org.hamcrest.Matchers.*;
import org.junit.Test;

public class SahaTests extends BaseTest {

    @Test
    public void sahalariListelemeTesti() {
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
        String sahaBody = "{\"name\": \"Zaim Halı Saha\",\"location\": \"İstanbul\",\"price\": 50}";

        given()
            .contentType("application/json")
            .body(sahaBody)
        .when()
            .post("/sahalar/")
        .then()
            .statusCode(201)
            .body("name", equalTo("Zaim Halı Saha"))
            .time(lessThan(2000L));
    }
}
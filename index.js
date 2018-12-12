"use strict";

Object.defineProperty(exports, "__esModule", { value: true });
var grpc_web_client_1 = require("grpc-web-client");
var readinglist_pb_service_1 = require("./proto-types/readinglist_pb_service");
var readinglist_pb_1 = require("./proto-types/readinglist_pb");

var host = "https://localhost:9090";

console.log('hitting server at: ', host);

function addBook() {
    var magazine = new readinglist_pb_1.Book();
    magazine.setTitle("To Kill a Mockingbird");
    magazine.setAuthor("Harper Lee");
    grpc_web_client_1.grpc.unary(readinglist_pb_service_1.ReadingList.AddBook, {
        request: magazine,
        host: host,
        onEnd: function (res) {
            console.log("onEnd", res);
        }
    });
}



setInterval(function () {
  // add a book every 2 seconds
  addBook();
}, 2000);

var ele = document.getElementById("app");
if (ele && ele.innerHTML) {
    ele.innerHTML = "Check the console";
}

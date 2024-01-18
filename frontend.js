// frontend.js
var app = angular.module('myApp', []);

app.controller('myCtrl', function ($scope, $http) {
    $http.get('/getEndpointNames')
        .then(function (response) {
            $scope.endpoints = response.data.endpoints;
        });

    $scope.callEndpoint = function (endpoint) {
        $http.get('/' + endpoint)
            .then(function (response) {
                console.log(response.data);
                // Handle the response as needed
            });
    };
});

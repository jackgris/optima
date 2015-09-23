angular
   .module('optima', [
       'satellizer',
       'ngRoute',
   ])
    .config(function($authProvider, $routeProvider){
        // Params to user authentication
        $authProvider.loginUrl = "http://localhost:8080/auth/login";
        $authProvider.signupUrl ="http://localhost:8080/auth/signup";
        $authProvider.tokenName = "token";
        $authProvider.tokenPrefix = "optima";          
            
        $routeProvider.
            when('/', {
                templateUrl: '/static/partials/home.html',
                controller: 'HomeController'
            }).
            when('/login', {
                templateUrl: '/static/partials/home-login.html',
                controller: 'LoginController'
            }).     
            when('/registro', {
                templateUrl: '/static/partials/home-register.html',
                controller: 'SignUpController'
            }).
            when('/logout', {
                templateUrl: null,
                controller: 'LogoutController'
            }).
            otherwise({
                redirectTo: '/'
            });
});    

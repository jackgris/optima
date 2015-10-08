angular
   .module('optima', [
       'satellizer',
       'ngRoute',
       'ngResource',
       'smart-table',
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
                controller: 'HomeController',
                resolve: {
                    isLoginListAdvertiser: isLoginListAdvertiser
                }
            }).
            when('/login', {
                templateUrl: '/static/partials/home-login.html',
                controller: 'LoginController',
                resolve: {
                    isLoginListAdvertiser: isLoginListAdvertiser
                }
            }).     
            when('/registro', {
                templateUrl: '/static/partials/home-register.html',
                controller: 'SignUpController',
                resolve: {
                    isLoginListAdvertiser: isLoginListAdvertiser
                }
            }).
            when('/logout', {
                templateUrl: null,
                controller: 'LogoutController'
            }).
            when('/listadvertisers', {
                templateUrl: '/static/partials/listadvertisers.html',
                controller: 'ListAdvertiserController',
                resolve: {
                    loginRequired: loginRequired
                }
            }).
            when('/addadvertiser', {
                templateUrl: '/static/partials/add-advertiser.html',
                controller: 'AddAdvertiserController',
                resolve: {
                    loginRequired: loginRequired
                }
            }).
            otherwise({
                redirectTo: '/'
            });
  })
  .run(function($rootScope, $auth, $location) {
      $rootScope.logout = function(){
            $auth.logout()
                .then(function() {
                    // Desconectamos al usuario y lo redirijimos
                    $location.path("/")
                });
        }

});

// Redirect authenticated users from section of register or authentication
function isLoginListAdvertiser($q, $location, $auth){
    var deferred = $q.defer();
    if ($auth.isAuthenticated()) {
        $location.path('/listadvertisers');
    } else {
        deferred.resolve();
    }
    return deferred.promise;
}

// Redirect unauthenticated users to the login state
function loginRequired($q, $location, $auth){
    var deferred = $q.defer();
    if ($auth.isAuthenticated()) {
        deferred.resolve();
    } else {
        $location.path('/login');
    }
    return deferred.promise;
}

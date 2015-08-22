var optima = angular.module('optima', [
    'ngRoute', 
    'homeControllers'
]);

optima.config(['$routeProvider',
  function($routeProvider){
      $routeProvider.
          when('/', {
              templateUrl: '/static/partials/home.html',
              controller: 'OptimaHome'
          }).
          when('/login', {
              templateUrl: '/static/partials/home-login.html',
              controller: 'OptimaLogin'
          }).     
          when('/registro', {
              templateUrl: '/static/partials/home-register.html',
              controller: 'OptimaRegister'
          }).
          otherwise({
              redirectTo: '/'
          });
  }]);

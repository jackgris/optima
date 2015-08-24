'use strict';

/* Home Controllers */
var homeControllers = angular.module('homeControllers', []);

homeControllers.controller('OptimaHome', ['$scope',
  function($scope){
}]);

homeControllers.controller('OptimaLogin', ['$scope', 'Login',
  function($scope, Login){
}]);

homeControllers.controller('OptimaRegister', ['$scope', 'Registro',
  function($scope, Registro){
      $scope.registerUser = function(){
        Registro.newRegister($scope.user);
      }
}]);


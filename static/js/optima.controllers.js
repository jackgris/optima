angular  
    .module('optima')
    .controller('HomeController', HomeController)
    .controller('SignUpController', SignUpController)
    .controller('LoginController', LoginController)
    .controller('LogoutController', LogoutController)
    .controller('PrivateController', PrivateController);

function HomeController($log) {  
    $log.info('Estamos en el home');
}

function PrivateController($auth, $scope, $log, $location){
    $log.info('Estamos en una seccion privada, '
                             + 'solo debe poder ingresar alguien autenticado');
    $scope.logout = function(){
        $auth.logout()
            .then(function() {
                // Desconectamos al usuario y lo redirijimos
                $location.path("/")
            });
    }
}

function SignUpController($auth, $location, $scope, $log) {  

    $scope.signup = function() {
        $auth.signup({
            email: $scope.signup.email,
            password: $scope.signup.password,
            name: $scope.signup.nombre
        })
        .then(function() {
            $log.info('Se realizo el registro correctamente');
            // Si se ha registrado correctamente,
            // Podemos redirigirle a otra parte
            $location.path("/private");
        })
        .catch(function(response) {
            // Si ha habido errores, llegaremos a esta función
            $log.info('Hubo un error en el registro');
        });
    }
}

function LoginController($log, $auth, $location, $scope) {  
    $log.debug('Se carga el controller de login');
    $scope.login = function(){
        $auth.login({
            email: $scope.login.email,
            password: $scope.login.password
        })
        .then(function(){
            // Si se ha logueado correctamente, lo tratamos aquí.
            // Podemos también redirigirle a una ruta
            $log.info('Se realizo el login correctamente');
            $location.path("/private")
        })
        .catch(function(response){
            // Si ha habido errores llegamos a esta parte
            $log.info('Hubo un error en el login');
            $location.path("/registro")
        });
    }
}

function LogoutController($auth, $location) {  
    $auth.logout()
        .then(function() {
            // Desconectamos al usuario y lo redirijimos
            $location.path("/")
        });
}

angular  
    .module('optima')
    .controller('HomeController', HomeController)
    .controller('SignUpController', SignUpController)
    .controller('LoginController', LoginController)
    .controller('LogoutController', LogoutController)
    .controller('ListAdvertiserController', ListAdvertiserController)
    .controller('AddAdvertiserController', AddAdvertiserController);

function HomeController($log, $auth) {  
    $log.info('We are at home');
    $log.info($auth.getToken())
    $log.info($auth.getPayload())
}

function AddAdvertiserController($log, $scope, AddAdvertiser){
    $log.info('We will add an advertiser')
    $scope.addadvertiser = function(){
        AddAdvertiser.put({
            name : $scope.addadvertiser.name,
            age : $scope.addadvertiser.age,
            sex : $scope.addadvertiser.sex,
            nse : $scope.addadvertiser.nse,
            coverage : $scope.addadvertiser.coverage,
            interets : '',
            category : $scope.addadvertiser.category,
            budget : $scope.addadvertiser.budget,
            objetives : $scope.addadvertiser.objetives,
        })

    }
    $scope.addadvertiser.nse = {
        selectnse: null,
        availableOptions: [
            {id: '1', name: 'Option E'},
            {id: '2', name: 'Option D'},
            {id: '3', name: 'Option D+'},
            {id: '4', name: 'Option C'},
            {id: '5', name: 'Option C+'},
            {id: '6', name: 'Option A/B'}
        ],
    }
}

function ListAdvertiserController($auth, $scope, $log, $location, Advertiser){

    $scope.GoAddAdvertiser = function(){
            $location.path("/addadvertiser");
    }
    $scope.advertisers = Advertiser.query();
    $scope.orderAdvertiser = 'age';
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
            $location.path("/listadvertisers");
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
            $location.path("/listadvertisers")
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

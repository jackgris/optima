<!DOCTYPE html>
<html lang="en" ng-app="optima">
  <head>
    <title>Optima</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="/static/css/bootstrap.css">
    <link rel="stylesheet" href="/static/css/app.css">
    <script src="/static/js/angular/angular.js"></script>
    <script src="/static/js/angular/angular-route.js"></script>
    <script src="/static/js/home.js"></script>
    <script src="/static/js/homeController.js"></script>
  </head>

  <body>
    <header>
      <h1>Optima</h1>
      <div class="description">
        Optima es una aplicación que nos dará la posibilidad de planear nuestras campañas publicitarias de una forma simple.
      </div>
    </header>

    <div class="view-container">
      <div ng-view class="view-frame"></div>
    </div>

    <footer>
      <div class="author">
        Official website:
        <a href="http://{{.Website}}">{{.Website}}</a> /
        Contact me:
        <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
      </div>
    </footer>
    <div class="backdrop"></div>
  </body>
</html>

<html>
  <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Login</title>
        <link rel="stylesheet" href="/static/css/registration.css">
        <link href='http://fonts.googleapis.com/css?family=Nunito:400,300' rel='stylesheet' type='text/css'>
    </head>
    <body>

      <form action="login" method="post">
      
        <h1>Login</h1>
        
        <fieldset>          
          <label for="mail">Email:</label>
          <input type="email" id="mail" name="user_email">
          
          <label for="password">Password:</label>
          <input type="password" id="password" name="user_password">

        </fieldset>

        <fieldset>  

          <label>Type:</label>
          <input type="radio" id="type_user" value="user" name="type_user"><label class="light" for="development">User</label><br>
          <input type="radio" id="type_publicist" value="publicist" name="type_user"><label class="light" for="design">Publicist</label><br>
        
        </fieldset>
        <button type="submit">Login</button>
      </form>
      
    </body>
</html>

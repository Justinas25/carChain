<!DOCTYPE html>
<html lang="en">
<head>
  <title>Owner</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
</head>
<body>

<div class="container">
  <h2>Owner</h2>
  <form action="/" method="POST">
    <div class="form-group">
      <label for="Name">Name:</label>
      <input class="form-control" id="name" name="name" placeholder="Enter name">
    </div>
    <div class="form-group">
      <label for="surname">Surname:</label>
      <input class="form-control" id="surname" name="surname" placeholder="Enter surname">
    </div>
    <div class="form-group">
      <label for="birth">Birth day:</label>
      <input type="date" class="form-control" id="birth" name="birth" placeholder="Enter date of birth">
    </div>
    <div class="form-group">
      <label for="surname">Nationality:</label>
      <input class="form-control" id="nationality" name="nationality" placeholder="Enter nationality">
    </div>
    <button type="submit" class="btn btn-default">Save</button>
  </form>
</div>

</body>
</html>

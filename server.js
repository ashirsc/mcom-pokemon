
const serverless = require( 'serverless-http')
const express = require('express')
const path = require('path');
const app = express();



app.use('/pokemon',express.static(path.join(__dirname, 'build')));

app.get('/pokemon', function (req, res) {
  res.sendFile(path.join(__dirname, 'build', 'index.html'));
});


app.listen(process.env.PORT || 8080);
console.log('started server on port', process.env.PORT || 8080);

module.exports.handler = serverless(app);


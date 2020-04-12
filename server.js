
const express = require('express');
const path = require('path');
const app = express();
const fs = require('fs')
const axios = require('axios').default





app.use(express.static(path.join(__dirname, 'build')));

app.get('/', function (req, res) {
  res.sendFile(path.join(__dirname, 'build', 'index.html'));
});

app.get('/pokemon/:pokeId/download', async (request, response) => {

const pokeRes = await axios.get('https://pokeapi.co/api/v2/pokemon/'+ request.params.pokeId)
console.log(pokeRes.data)
fs.writeFileSync(path.join(__dirname, 'pokemon.json'), JSON.stringify(pokeRes.data))
  response.download(path.join(__dirname, 'pokemon.json'))
})

app.listen(process.env.PORT || 8080);
console.log('started server on port', process.env.PORT || 8080);

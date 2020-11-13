const axios = require('axios').default

var pokemon;
let charmander
let background

var dots;


var leftDirection = false;
var rightDirection = false;
var upDirection = true;
var downDirection = false;
var inGame = true;

const DOT_SIZE = 1;
const ALL_DOTS = 900;
const MAX_RAND = 29;
const DELAY = 1000/60;
const C_HEIGHT = 1300;
const C_WIDTH = 800;

const LEFT_KEY = 37;
const RIGHT_KEY = 39;
const UP_KEY = 38;
const DOWN_KEY = 40;

var x = 150
var y = 700

export async function init(ctx) {
    console.log('Setting up game')
    console.log('ctx', ctx)

    await loadImages()
   
    console.log('Game starting')
    setTimeout(() => {
        gameCycle(ctx)
    } , DELAY);
}


async function loadImages() {
    
    pokemon = new Image();
    const pokeData = await (await axios.get('https://pokeapi.co/api/v2/pokemon/1')).data
    // pokemon.src = pokeData.sprites.front_default;    
    pokemon.src = pokeData.sprites.back_default;    
    console.log('poke img src', pokemon.src)

    background = new Image()
    // background.src = './background.png'
    background.src = require(`./assets/background.png`)
    console.log('background', background)
   
}

// function drawSnakePart(ctx, x,y) {
//     ctx.fillStyle = "lightgreen";
//     ctx.strokestyle = "darkgreen";
//     ctx.fillRect(x, y, 10, 10);
//     ctx.strokeRect(x, y, 300, 10);
// }


function doDrawing(ctx) {

    ctx.clearRect(0, 0, C_WIDTH, C_HEIGHT);
    ctx.drawImage(background, 0, 0)
    ctx.drawImage(pokemon, x, y)

   
}



function move() {

    for (var z = dots; z > 0; z--) {

        x[z] = x[(z - 1)];
        y[z] = y[(z - 1)];
    }

    if (leftDirection) {

        x -= DOT_SIZE;
    }

    if (rightDirection) {

        x += DOT_SIZE;
    }

    if (upDirection) {

        y -= DOT_SIZE;
    }

    if (downDirection) {

        y += DOT_SIZE;
    }
}





function gameCycle(ctx) {
    console.log('game cycle');
    
    if (inGame) {
        move();
        doDrawing(ctx);
        setTimeout(()=>{gameCycle(ctx)}, DELAY);
    }
}

onkeydown = function (e) {
    
    e.preventDefault()
    console.log(e)
    var key = e.keyCode;

    if ((key == LEFT_KEY) && (!rightDirection)) {

        leftDirection = true;
        upDirection = false;
        downDirection = false;
    }

    if ((key == RIGHT_KEY) && (!leftDirection)) {

        rightDirection = true;
        upDirection = false;
        downDirection = false;
    }

    if ((key == UP_KEY) && (!downDirection)) {

        upDirection = true;
        rightDirection = false;
        leftDirection = false;
    }

    if ((key == DOWN_KEY) && (!upDirection)) {

        downDirection = true;
        rightDirection = false;
        leftDirection = false;
    }
};  
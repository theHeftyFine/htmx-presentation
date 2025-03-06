const can = document.createElement('canvas')
can.style.position = 'absolute';
can.style.pointerEvents = 'none';
can.style.top = '0px';
can.width = window.innerWidth;
can.height = window.innerHeight;

const ct = can.getContext('2d');

let mouse = ({
    x: 0,
    y: 0
});

const createStalker = () => ({
    x: 20,
    y: 20,
    radius: 5,
    color: '#ddd',
    speed: 2
})

const drawStalker = stalker => {
    ct.beginPath();
    ct.arc(stalker.x, stalker.y, stalker.radius, 0, Math.PI * 2);
    ct.fillStyle = stalker.color;
    ct.fill();
    ct.closePath();
}

const updateStalker = stalker => {
    deltax = mouse.x - stalker.x
    deltay = mouse.y - stalker.y
    mag = Math.sqrt(deltax * deltax + deltay * deltay)
    console.log(deltax, deltay, mag)

    stalker.x += stalker.speed * (deltax / mag);
    stalker.y += stalker.speed * (deltay / mag);
}

const calcDir = val => {
    return val > 0 ? 1 : val < 0 ? -1 : 0;
}

const stalker = createStalker();

const animateStalker = () => {
    ct.clearRect(0, 0, canvas.width, canvas.height);
    updateStalker(stalker);
    drawStalker(stalker);

    requestAnimationFrame(animateStalker)
}

document.onmousemove = event => {
    mouse.x = event.pageX;
    mouse.y = event.pageY;
}

document.body.appendChild(can);
animateStalker();

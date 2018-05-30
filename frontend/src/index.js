import Matchmaker from './matchmaker.js';

function canvas() {
  var canv = document.createElement('canvas');
  return canv;
}

var canv = canvas();
document.body.appendChild(canv);

canv.width = window.innerWidth;
canv.height = window.innerHeight;

var ctx = canv.getContext('2d');

ctx.beginPath();
ctx.arc(40, 40, 40, 0, Math.PI);
ctx.stroke();


//function resizeCanvas() {
//  canv.style.width = window.innerWidth + "px";
//  // artifical delay so innerHeight is correct
//  setTimeout(function() {
//    canv.style.height = window.innerHeight + "px";
//  }, 0)j
//
//  drawStuff();
//}
//
//window.onresize = resizeCanvas;
//
//resizeCanvas();
//drawStuff();
//
//function drawStuff() {
//	var ctx = canv.getContext('2d');
//
//    var devicePixelRatio = window.devicePixelRatio || 1;
//
//    var backingStoreRatio = ctx.webkitBackingStorePixelRatio ||
//                        ctx.mozBackingStorePixelRatio ||
//                        ctx.msBackingStorePixelRatio ||
//                        ctx.oBackingStorePixelRatio ||
//                        ctx.backingStorePixelRatio || 1;
//
//    var ratio = devicePixelRatio / backingStoreRatio;
//
//    if (devicePixelRatio !== backingStoreRatio)
//    {
//        var oldWidth = canvas.width;
//        var oldHeight = canvas.height;
//
//        canvas.width = oldWidth * ratio;
//        canvas.height = oldHeight * ratio;
//
//        canvas.style.width = oldWidth + 'px';
//        canvas.style.height = oldHeight + 'px';
//
//        context.scale(ratio, ratio);
//    }
//
//	console.log(ratio);
//
//	ctx.beginPath();
//	ctx.arc(canv.width / 2, canv / 2, 40, 0, Math.PI);
//	ctx.stroke();
//}


// var matchmaker = new Matchmaker("ws://127.0.0.1:8001/ws", canv.getContext('2d'));
// matchmaker.joinQueue();

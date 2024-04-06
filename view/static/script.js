



/*                           projects               */
/*                        header          */

let menuLsit = document.getElementById('menu-list');
menuLsit.style.maxHeight = '0px';
document.querySelector(".menu-icon").addEventListener("click", () => {
    if(menuLsit.style.maxHeight == '0px') {
        menuLsit.style.maxHeight = '370px';
    } else {
        menuLsit.style.maxHeight = '0px';
    }
});






const typed = new Typed('#multiple', {
    strings: ['Golang programmer', 'Database engineer', 'Web developer'],
    typeSpeed: 70,
    backSpeed: 70,
    backDelay: 100,
    loop: true,
});






// function toggleMenu() {
//
// }

///////////////////////////////////////////////////

lightGallery(document.querySelector('.gallery'));
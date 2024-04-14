"use strict";


lightGallery(document.querySelector('.gallery'));



// image uploader

const inputFile = document.querySelector('#myFile');
const uploadP = document.querySelector('.upload-p');


inputFile.addEventListener('change', uploadImage);

function uploadImage() {
    const file = this.files[0];
    uploadP.innerHTML = file.name;
    
}









/*                        header          */

let menuLsit = document.querySelector('#menu-list');
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







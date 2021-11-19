const signUpButton = document.getElementById('signUp');
const signInButton = document.getElementById('signIn');
const container = document.getElementById('container');

var password = document.getElementById("password"), confirm_password = document.getElementById("confirm_password");
var password_1 = document.getElementById("password_1"), confirm_password_1 = document.getElementById("confirm_password_1");

signUpButton.addEventListener('click', () => {
    container.classList.add("right-panel-active");
});

signInButton.addEventListener('click', () => {
    container.classList.remove("right-panel-active");
});

function validatePassword(){
    if(password.value != confirm_password.value) {
        confirm_password.setCustomValidity("Passwords Don't Match");
    } else {
        confirm_password.setCustomValidity('');
    }
}

function validatePassword_1(){
    if(password_1.value != confirm_password_1.value) {
        confirm_password_1.setCustomValidity("Passwords Don't Match");
    } else {
        confirm_password_1.setCustomValidity('');
    }
}

    password.onchange = validatePassword;
    confirm_password.onkeyup = validatePassword;

    password_1.onchange = validatePassword_1;
    confirm_password_1.onkeyup = validatePassword_1;
const signUpButton = document.getElementById('signUp');
const signInButton = document.getElementById('signIn');
const container = document.getElementById('login-container');
const closeButton = document.getElementById('close')

signUpButton.addEventListener('click', () => {
    container.classList.add("right-panel-active");
    closeButton.classList.add("left")
});

signInButton.addEventListener('click', () => {
    container.classList.remove("right-panel-active");
    closeButton.classList.remove("left")

});

function validatePassword(){
    if(password.value !== confirm_password.value) {
        confirm_password.setCustomValidity("Passwords Don't Match");
    } else {
        confirm_password.setCustomValidity('');
    }
}
password.onchange = validatePassword;
confirm_password.onkeyup = validatePassword;



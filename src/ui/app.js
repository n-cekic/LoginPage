let confirmPasswordVisible = false;

async function login() {
    console.log("Submitted!");
    const username = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    // Perform your login logic here
    // Check if input is correct
    const data = {
        username: username,
        password: password,
    };

    sendRequest(data, "login");
}

async function signin() {
    console.log("sign in clicked.");
   
    if (!validateSigninPassword()) {
        moveConfirmPasswordCover();
        return;
    }

    moveConfirmPasswordCover();

    const password = document.getElementById("password").value;
    const username = document.getElementById("email").value;

    sendRequest({ username: username, password: password }, "signup");
}

async function sendRequest(data, resource) {
    const resp = await fetch("http://localhost:8080/" + resource, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    });

    console.log(await resp.text());
}

function validateSigninPassword() {
    const password = document.getElementById("password");
    const confPswd = document.getElementById("confirm-password");
    const pswd = password.value;
    const conf = confPswd.value;

    if (!confirmPasswordVisible) {
        return false;
    }

    if (pswd === "") {
        console.log("empty password")
        password.style= {};
        confPswd.style = {};
        return false
    }else if (pswd !== conf) {
        console.log("passwords don't match")
        password.style.borderColor = "red";
        confPswd.style.borderColor = "red";
        return false;
    } else {
        console.log("password matched")
        password.style.borderColor = "#66ff99";
        confPswd.style.borderColor = "#66ff99";
        return true;
    }
}

function moveConfirmPasswordCover() {
    const buttonContainer = document.getElementById("button-container");

    if (!confirmPasswordVisible) {
        revalConfirmPassword(buttonContainer);
    } else {
        hideConfirmPassword(buttonContainer);
    }
}

function revalConfirmPassword(buttons) {
    buttons.style.transform = "translateY(55px)";
    confirmPasswordVisible = true;
    console.log("confirm password revealed.");
}

function hideConfirmPassword(buttons) {
    buttons.style.transform = "translateY(0px)";
    confirmPasswordVisible = false;
    console.log("confirm password hidden.");
}

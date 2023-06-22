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

    const resp = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    });

    console.log(await resp.text());
}

function passwordsMatch() {
    const password = document.getElementById("password");
    const confPswd = document.getElementById("confirm-password");
    const pswd = password.value;
    const conf = confPswd.value;

    if (!confirmPasswordVisible) {
        password.style.borderColor = "black";
        return;
    }

    if (pswd !== conf) {
        password.style.borderColor = "red";
        confPswd.style.borderColor = "red";
        return false;
    } else {
        password.style.borderColor = "#66ff99";
        confPswd.style.borderColor = "#66ff99";
        return true;
    }
}

function moveConfirmPasswordCover() {
    const buttonContainer = document.getElementById("button-container");
    const confirmPassworInput = document.getElementById("confirm-password");

    if (!confirmPasswordVisible) {
        revalConfirmPassword(buttonContainer);
    } else {
        if (!confirmPassworInput.value) {
            hideConfirmPassword(buttonContainer);
            return;
        }
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

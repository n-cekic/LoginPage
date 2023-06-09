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
    mode: "no-cors",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  console.log(resp);
}

let confirmPasswordVisible = false;

function signin() {
  console.log("sign in clicked.");

  const confPswd = document.getElementById("confirm-password");
  const buttons = document.getElementById("button-container");

  if (!confirmPasswordVisible) {
    revalConfirmPassword(buttons);
  } else {
    if (!confPswd.value) {
      hideConfirmPassword(buttons);
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

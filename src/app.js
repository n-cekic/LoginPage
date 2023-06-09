async function submit() {
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

function revealConfirmPassword() {
  console.log("sign in clicked.");

  const confPswd = document.getElementById("confirm-password");
  const buttons = document.getElementById("button-container");

  if (!confirmPasswordVisible) {
    buttons.style.transform = "translateY(55px)";
    confirmPasswordVisible = true;
  } else {
    if (!confPswd.value) {
      buttons.style.transform = "translateY(0px)";
      confirmPasswordVisible = false;
    }
  }

  console.log("confirm password revealed.");
}

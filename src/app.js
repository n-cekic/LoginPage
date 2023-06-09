let clickedOnSignIn = false;

function submit() {
  console.log("Submitted!");
  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;

  // Perform your login logic here
  // Check if input is correct
  const data = {
    username: username,
    password: password
  };


  const resp = await fetch('http://localhost:8080/login', {
    method: 'POST',
    mode: "no-cors",
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
  });
  
  console.log(resp);
}

function revealConfirmPassword() {
  console.log("sign in clicked.");
  const confPswd = document.getElementById("confirm-password");
  const buttons = document.getElementById("button-container");
  buttons.style.transform = clickedOnSignIn
    ? "translateY(0px)"
    : "translateY(55px)";
  clickedOnSignIn = !clickedOnSignIn;
  //confPswd.style.opacity = confPswd.style.opacity === "1" ? "0" : "1";

  //confPswd.classList.toggle("open");
  console.log("confirm password revealed.");
}

let clickedOnSignIn = false;

function submit() {
  console.log("Submitted!");
  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;

  // Perform your login logic here
  // For example, make an AJAX request to the server to validate the credentials

  // Prevent the form from submitting and the page from refreshing
  event.preventDefault();
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

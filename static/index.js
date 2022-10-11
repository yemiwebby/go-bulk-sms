const MAX_BADGES_PER_ROW = 5;

const possibleBadgeColours = [
  "red",
  "orange",
  "yellow",
  "green",
  "teal",
  "blue",
  "purple",
  "indigo",
  "pink",
  "gray",
];

const recipients = [];

const input = document.querySelector("input");
const textArea = document.querySelector("textarea");
const phoneNumberGrid = document.getElementById("phoneNumbers");
const alert = document.getElementById("alert");

input.addEventListener("keyup", (event) => {
  if (event.key === "Enter") {
    const recipient = event.target.value;
    if (recipient !== "") {
      recipients.push(recipient);
      input.value = "";
      let lastRow = phoneNumberGrid.lastElementChild;
      if (
        lastRow === null ||
        lastRow.childElementCount === MAX_BADGES_PER_ROW
      ) {
        lastRow = getNewRow();
        phoneNumberGrid.appendChild(lastRow);
      }
      const phoneNumberBadge = getNewPhoneNumberBadge(recipient);
      lastRow.appendChild(phoneNumberBadge);
    }
  }
});

const getNewRow = () => {
  const row = document.createElement("div");
  row.classList.add("flex", "space-x2");
  return row;
};

const getRandomBadgeColour = () => {
  const randomIndex = Math.floor(
    Math.random() * (possibleBadgeColours.length - 1)
  );
  return possibleBadgeColours[randomIndex];
};

const getNewPhoneNumberBadge = (phoneNumber) => {
  const badgeColour = getRandomBadgeColour();
  const badge = document.createElement("div");
  badge.classList.add(
    "badge",
    "text-sm",
    "px-3",
    "mx-2",
    `bg-${badgeColour}-200`,
    `text-${badgeColour}-800`,
    "rounded-full"
  );
  badge.textContent = phoneNumber;
  return badge;
};

const showButton = (buttonId) => {
  const button = document.getElementById(buttonId);
  button.classList.remove("hidden");
};

const hideButton = (buttonId) => {
  const button = document.getElementById(buttonId);
  button.classList.add("hidden");
};

const reset = () => {
  recipients.splice(0, recipients.length);
  phoneNumberGrid.replaceChildren();
  textArea.value = "";
  showButton("submit");
  hideButton("loading");
};

const handleSubmit = async () => {
  if (recipients.length === 0) {
    showAlert("Please provide at least one phone number");
    return;
  }

  const message = textArea.value.trim();
  if (message === "") {
    showAlert("Please provide a message for the recipients");
    return;
  }

  hideButton("submit");
  showButton("loading");

  const response = await fetch("/messages", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ recipients, message }),
  });

  const { message: alertMessage } = await response.json();
  showAlert(alertMessage, response.ok);
  reset();
};

const hideAlert = () => {
  alert.classList.add("hidden");
};

const showAlert = (message, isSuccess = false) => {
  const alertColour = isSuccess ? "green" : "red";

  alert.classList.add(
    `bg-${alertColour}-100`,
    "border",
    `border-${alertColour}-400`,
    `text-${alertColour}-700`
  );

  const alertTitle = document.getElementById("alertTitle");
  const alertMessage = document.getElementById("alertMessage");
  const closeAlertButton = document.getElementById("closeAlert");

  closeAlertButton.classList.add(`text-${alertColour}-500`);
  alertTitle.innerText = isSuccess ? "Success! \n" : "Error! \n";
  alertMessage.innerText = message;
  alert.classList.remove("hidden");

  setTimeout(() => {
    alert.classList.add("hidden");
  }, 5000);
};

document.getElementById("calcForm").addEventListener("submit", function (event) {
  event.preventDefault();
  let formData = new FormData(this);
  let jsonData = {};

  formData.forEach((value, key) => {
      jsonData[key] = parseFloat(value.replace(",", "."));
  });

  fetch("/calculator2", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(jsonData)
  }).then(response => response.json())
    .then(data => {
        document.getElementById("result").innerHTML = `
            <p>Hp: ${data.Hp.toFixed(2)}</p>
            <p>Cp: ${data.Cp.toFixed(2)}</p>
            <p>Sp: ${data.Sp.toFixed(2)}</p>
            <p>Op: ${data.Op.toFixed(2)}</p>
            <p>Vp: ${data.Vp.toFixed(2)}</p>
            <p>Ap: ${data.Ap.toFixed(2)}</p>
            <p>Qri: ${data.Qri.toFixed(2)}</p>
        `;
    });
});
document.getElementById("calcForm").addEventListener("submit", function (event) {
  event.preventDefault();
  let formData = new FormData(this);
  let jsonData = {};

  formData.forEach((value, key) => {
      jsonData[key] = parseFloat(value.replace(",", "."));
  });

  fetch("/calculator1", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(jsonData)
  }).then(response => response.json())
    .then(data => {
        document.getElementById("result").innerHTML = `
            <p>Kpc: ${data.Kpc.toFixed(2)}</p>
            <p>Kpg: ${data.Kpg.toFixed(2)}</p>
            <p>Hc: ${data.Hc.toFixed(2)}</p>
            <p>Cc: ${data.Cc.toFixed(2)}</p>
            <p>Sc: ${data.Sc.toFixed(2)}</p>
            <p>Nc: ${data.Nc.toFixed(2)}</p>
            <p>Oc: ${data.Oc.toFixed(2)}</p>
            <p>Ac: ${data.Ac.toFixed(2)}</p>
            <p>Hr: ${data.Hr.toFixed(2)}</p>
            <p>Cr: ${data.Cr.toFixed(2)}</p>
            <p>Sr: ${data.Sr.toFixed(2)}</p>
            <p>Nr: ${data.Nr.toFixed(2)}</p>
            <p>Or: ${data.Or.toFixed(2)}</p>
            <p>Qph: ${data.Qph.toFixed(2)}</p>
            <p>Qch: ${data.Qch.toFixed(2)}</p>
            <p>Qrh: ${data.Qrh.toFixed(2)}</p>
        `;
    });
});

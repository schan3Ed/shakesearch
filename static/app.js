const Controller = {
  search: (ev) => {
    ev.preventDefault();
    const form = document.getElementById("form");
    const data = Object.fromEntries(new FormData(form));
    const response = fetch(`/search?q=${data.query}`).then((response) => {
      response.json().then((results) => {
        Controller.updateMeta(results);
        Controller.updateTable(results);
      });
    });
  },
  updateMeta: (results) => {
    const table = document.getElementById("meta-body");
    const rows = [];
    rows.push("How many times this word show up? " + results.Occur + " times");
    rows.push("<br>");
    rows.push("How many books have this word? " + Object.keys(results.Results).length + " book");
    table.innerHTML = rows;
  },
  updateTable: (results) => {
    const table = document.getElementById("table-body");
    const head = document.getElementById("table-head");
    const rows = [];
    const header = [];

    header.push(`<tr</tr>`)
    header.push(`<th>Book</th>`)
    header.push(`<th>Content</th>`)
    header.push(`<tr></tr>`)

    for (book in results.Results) {
      for (let lines of results.Results[book]) {
        const l = lines.replace(/(\r\n|\n|\r)/gm, "<br>")
        rows.push(`<tr><tr/>`);
        rows.push(`<td>${book}<td/>`);
        rows.push(`<td>${l}<td/>`);
        rows.push(`<tr><tr/>`);
      }
    }

    table.innerHTML = rows;
    head.innerHTML = header;
  },
};

const form = document.getElementById("form");
form.addEventListener("submit", Controller.search);

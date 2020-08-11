fetch('http://localhost:8080/fontsApi')
.then( async (jsonData) => {
    const res = await jsonData.json()
    var node = document.getElementById("fontsSelect");
    const fonts = res.font;
    fonts.map(font => {
        node.options.add(new Option(font, font));
    });
});
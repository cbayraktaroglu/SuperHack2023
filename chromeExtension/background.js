chrome.omnibox.onInputEntered.addListener(function (text) {
    console.log(text);
    console.log("egegegegegegegegegegegegegeg");
    var _a = text.split('/'), value1 = _a[0], value2 = _a[1];
    console.log(value1); // Outputs: 0x1234
    console.log(value2); // Outputs: Abc
    // URL of the API
    var apiUrl = 'http://localhost:3000/decodePublic/' + text;
    // Configure the fetch POST request
    fetch(apiUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(function (response) { return response.json(); })
        .then(function (responseData) {
        console.log('API response:', responseData);
        // Create a new tab after the API request
        var url = "https://www.wikipedia.org";
        setTimeout(function () {
            chrome.tabs.create({ url: url });
        }, 10000); // Delay in milliseconds (e.g., 10000 ms = 10 seconds)
    })
        .catch(function (error) {
        console.error('API request failed:', error);
    });
});
/*
chrome.omnibox.onInputEntered.addListener((text: string) => {
  console.log(text);
  console.log("egegegegegegegegegegegegegeg");

  const url = "https://www.wikipedia.org";
  const [scAddress, chainName] = text.split('/');

  setTimeout(() => {
    chrome.tabs.create({ url });
  }, 10000);

});
*/ 

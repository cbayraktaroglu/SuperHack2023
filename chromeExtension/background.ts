chrome.omnibox.onInputEntered.addListener((text: string) => {
  console.log(text);
  console.log("egegegegegegegegegegegegegeg");

  const [value1, value2] = text.split('/');

  console.log(value1); // Outputs: 0x1234
  console.log(value2); // Outputs: Abc

  // URL of the API
  const apiUrl = 'http://localhost:3000/decodePublic/'+ text;

  // Configure the fetch POST request
  fetch(apiUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    }
  })
  .then(response => response.json())
  .then(responseData => {
    console.log('API response:', responseData);

    // Create a new tab after the API request
    const url = "https://www.wikipedia.org";
    setTimeout(() => {
      chrome.tabs.create({ url });
    }, 10000); // Delay in milliseconds (e.g., 10000 ms = 10 seconds)
  })
  .catch(error => {
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
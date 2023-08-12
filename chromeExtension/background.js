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
        .then(function (response) { return response.arrayBuffer(); })
        .then(function (responseData) {
        var responseDataStr = new TextDecoder().decode(responseData);
        // Parse the JSON data
        var parsedData = JSON.parse(responseDataStr);
        var return_buffer = parsedData.return_buffer, file_name = parsedData.file_name;
        console.log('API response:', responseData);
        // Map file extension to MIME type
        var mimeTypes = {
            '.gif': 'image/gif',
            '.jpg': 'image/jpeg',
            '.jpeg': 'image/jpeg',
            '.png': 'image/png',
            '.svg': 'image/svg+xml',
            '.pdf': 'application/pdf',
            '.txt': 'text/plain',
            '.json': 'application/json',
            '.html': 'text/html' // HTML file
            // Add more mappings as needed
        };
        var fileEx = file_name.substring(file_name.lastIndexOf('.'));
        var fileType = mimeTypes[fileEx];
        //Create the blob from buffer
        var buffer = new Uint8Array(return_buffer);
        var blob = new Blob([buffer], { type: fileType });
        var blobUrl = URL.createObjectURL(blob);
        //Create the new tab for blob
        chrome.tabs.create({ url: blobUrl });
        // Create a link element for downloading
        var downloadLink = document.createElement('a');
        downloadLink.href = blobUrl;
        downloadLink.download = file_name;
        // Add the link to the document
        document.body.appendChild(downloadLink);
        // Click the link to trigger the download
        downloadLink.click();
        // Clean up the Object URL and remove the link from the document
        URL.revokeObjectURL(blobUrl);
        downloadLink.remove();
        /*
        setTimeout(() => {
          chrome.tabs.create({ url });
        }, 10000); // Delay in milliseconds (e.g., 10000 ms = 10 seconds)
        */
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

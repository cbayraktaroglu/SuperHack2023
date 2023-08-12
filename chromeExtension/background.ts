chrome.omnibox.onInputEntered.addListener((text: string) => {
  console.log(text);
  console.log("egegegegegegegegegegegegegeg");

  const [value1, value2] = text.split('/');

  console.log(value1); // Outputs: 0x1234
  console.log(value2); // Outputs: Abc

  // URL of the API
  const apiUrl = 'http://localhost:3000/decodePublic/' + text;

  // Configure the fetch POST request
  fetch(apiUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(response => response.arrayBuffer())
    .then(responseData => {
      const responseDataStr = new TextDecoder().decode(responseData);

      // Parse the JSON data
      const parsedData = JSON.parse(responseDataStr) as {
        return_buffer: number[]; // Change to the actual type of return_buffer
        file_name: string;
      };

      const { return_buffer, file_name } = parsedData;
      
      console.log('API response:', responseData);

      // Map file extension to MIME type
      const mimeTypes: Record<string, string> = {
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
      
      const fileEx = file_name.substring(file_name.lastIndexOf('.'));
      const fileType = mimeTypes[fileEx];

      //Create the blob from buffer
      const buffer = new Uint8Array(return_buffer);
      const blob = new Blob([buffer], { type: fileType });
      const blobUrl = URL.createObjectURL(blob);

      //Create the new tab for blob
      chrome.tabs.create({ url: blobUrl });

      // Create a link element for downloading
      const downloadLink = document.createElement('a');
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
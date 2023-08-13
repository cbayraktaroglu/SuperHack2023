document.addEventListener('DOMContentLoaded', function () {
  const fileInput = document.getElementById('fileInput') as HTMLInputElement;
  const uploadPrivate = document.getElementById('uploadPrivate') as HTMLButtonElement;
  const privateInputs = document.getElementById('privateInputs') as HTMLInputElement;

  const publicView = document.getElementById('publicView') as HTMLInputElement;
  const privateView = document.getElementById('privateView') as HTMLInputElement;

  const publicInputs = document.getElementById('publicInputs') as HTMLInputElement;
  const publicDropdown = document.getElementById('publicDropdown') as HTMLInputElement;
  const submitPublicButton = document.getElementById('submitPublic') as HTMLInputElement;

  //Public view handler
  submitPublicButton.addEventListener('click', async () => {
    const contractAddressInput = document.getElementById('contractAddress') as HTMLInputElement;
    const publicDropdownValue = publicDropdown.value;

    console.log('Public Input:', contractAddressInput.value);
    console.log('Public Dropdown:', publicDropdownValue);

    const query = contractAddressInput.value + "/" + publicDropdownValue

    try {
      const response = await fetch('http://localhost:3000/decodePublic/' + query, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },

      });

      if (response.ok) {
        const jsonResponse = await response.json();

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

        // Access the values from the JSON response
        const returnBuffer = jsonResponse.Data.return_buffer;
        const returnBufferBytes = new Uint8Array(atob(returnBuffer).split('').map(char => char.charCodeAt(0)));

        const fileName = jsonResponse.Data.file_name;
        const fileExtension = fileName.substring(fileName.lastIndexOf('.'));
        const status = jsonResponse.Status;

        // Extract file name and type from the uploaded file
        const fileType = mimeTypes[fileExtension];

        // Create a new Uint8Array from the buffer
        //const uint8Array = new Uint8Array(buffer);
        //console.log(uint8Array);

        //const blob = new Blob([buffer], { type: 'image/gif' });
        const blob = new Blob([returnBufferBytes], { type: fileType });
        const blobUrl = URL.createObjectURL(blob);


        console.log(blob);

        // Display the preview image
        //await sleep(100);

        chrome.tabs.create({ url: blobUrl });

        // Create a link element for downloading
        const downloadLink = document.createElement('a');
        downloadLink.href = blobUrl;
        downloadLink.download = fileName;

        // Add the link to the document
        document.body.appendChild(downloadLink);

        // Click the link to trigger the download
        downloadLink.click();

        // Clean up the Object URL and remove the link from the document
        URL.revokeObjectURL(blobUrl);
        downloadLink.remove();
      } else {
        console.error('API request failed:', response.status, response.statusText);
      }
    } catch (error) {
      console.error('Error:', error);
    }
  });


  publicView.addEventListener('click', function () {

    publicInputs.style.display = 'block';
    privateInputs.style.display = 'none';
    submitPublicButton.style.display = 'block';
    uploadPrivate.style.display = 'none';
  });

  privateView.addEventListener('click', function () {
    publicInputs.style.display = 'none';
    privateInputs.style.display = 'block';
    submitPublicButton.style.display = 'none';
    uploadPrivate.style.display = 'block';
  });

  uploadPrivate.addEventListener('click', async () => {
    const file = fileInput.files?.[0];

    if (file) {
      const reader = new FileReader();
      reader.onload = async (event) => {
        const fileContents = event.target?.result as string;
        const uploadedData = JSON.parse(fileContents) as {
          file_name: string;
          file_type: string;
        };

        try {
          const response = await fetch('http://localhost:3000/decodePrivate', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json'
            },
            body: fileContents
          });

          if (response.ok) {
            const buffer = await response.arrayBuffer();

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

            // Extract file name and type from the uploaded file
            const fileEx = uploadedData.file_type;
            const fileType = mimeTypes[fileEx];

            // Create a new Uint8Array from the buffer
            //const uint8Array = new Uint8Array(buffer);
            //console.log(uint8Array);

            //const blob = new Blob([buffer], { type: 'image/gif' });
            const blob = new Blob([buffer], { type: fileType });
            const blobUrl = URL.createObjectURL(blob);


            console.log(blob);

            // Display the preview image
            //await sleep(100);

            chrome.tabs.create({ url: blobUrl });

            // Create a link element for downloading
            const downloadLink = document.createElement('a');
            downloadLink.href = blobUrl;
            downloadLink.download = uploadedData.file_name;

            // Add the link to the document
            document.body.appendChild(downloadLink);

            // Click the link to trigger the download
            downloadLink.click();

            // Clean up the Object URL and remove the link from the document
            URL.revokeObjectURL(blobUrl);
            downloadLink.remove();
          } else {
            console.error('API request failed:', response.status, response.statusText);
          }
        } catch (error) {
          console.error('Error:', error);
        }
      };
      reader.readAsText(file);
    }
  });
});

function sleep(ms: number): Promise<void> {
  return new Promise((resolve) => {
    setTimeout(resolve, ms);
  });
}

/*
document.addEventListener('DOMContentLoaded', function() {
    const fileInput = document.getElementById('fileInput') as HTMLInputElement;
    const uploadButton = document.getElementById('uploadButton') as HTMLButtonElement;
  
    uploadButton.addEventListener('click', async () => {
      const file = fileInput.files?.[0];
  
      if (file) {
        const reader = new FileReader();
        reader.onload = async (event) => {
          const fileContents = event.target?.result as string;
          
          try {
            const response = await fetch('http://localhost:3000/decodePrivate', {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json' // Set the Content-Type to JSON
              },
              body: fileContents // Send the JSON data as the request body
            });
  
            if (response.ok) {
              const buffer = await response.arrayBuffer();
              const blob = new Blob([buffer]);
              const url = URL.createObjectURL(blob);
  
              chrome.tabs.create({ url });
            } else {
              console.error('API request failed:', response.status, response.statusText);
            }
          } catch (error) {
            console.error('Error:', error);
          }
        };
        reader.readAsText(file); // Read the file contents as text
      }
    });
  });
 */ 
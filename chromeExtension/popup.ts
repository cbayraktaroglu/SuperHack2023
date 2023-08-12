document.addEventListener('DOMContentLoaded', function () {
    const fileInput = document.getElementById('fileInput') as HTMLInputElement;
    const uploadButton = document.getElementById('uploadButton') as HTMLButtonElement;
    const previewImage = document.getElementById('previewImage') as HTMLImageElement;

    uploadButton.addEventListener('click', async () => {
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
                            // Add more mappings as needed
                        };

                        // Extract file name and type from the uploaded file
                        const fileEx = uploadedData.file_type;
                        const fileType = mimeTypes[fileEx];

                        // Create a new Uint8Array from the buffer
                        //const uint8Array = new Uint8Array(buffer);
                        //console.log(uint8Array);

                        //const blob = new Blob([buffer], { type: 'image/gif' });
                        const blob = new Blob([buffer], { type: fileType});
                        const blobUrl = URL.createObjectURL(blob);
                        console.log(blob);

                        // Display the preview image
                        await sleep(100);

                        chrome.tabs.create({ url: blobUrl });
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
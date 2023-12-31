var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
document.addEventListener('DOMContentLoaded', function () {
    var _this = this;
    var fileInput = document.getElementById('fileInput');
    var uploadPrivate = document.getElementById('uploadPrivate');
    var privateInputs = document.getElementById('privateInputs');
    var publicView = document.getElementById('publicView');
    var privateView = document.getElementById('privateView');
    var publicInputs = document.getElementById('publicInputs');
    var publicDropdown = document.getElementById('publicDropdown');
    var submitPublicButton = document.getElementById('submitPublic');
    //Public view handler
    submitPublicButton.addEventListener('click', function () { return __awaiter(_this, void 0, void 0, function () {
        var contractAddressInput, publicDropdownValue, query, response, jsonResponse, mimeTypes, returnBuffer, returnBufferBytes, fileName, fileExtension, status_1, fileType, blob, blobUrl, downloadLink, error_1;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    contractAddressInput = document.getElementById('contractAddress');
                    publicDropdownValue = publicDropdown.value;
                    console.log('Public Input:', contractAddressInput.value);
                    console.log('Public Dropdown:', publicDropdownValue);
                    query = contractAddressInput.value + "/" + publicDropdownValue;
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 6, , 7]);
                    return [4 /*yield*/, fetch('http://localhost:3000/decodePublic/' + query, {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json'
                            },
                        })];
                case 2:
                    response = _a.sent();
                    if (!response.ok) return [3 /*break*/, 4];
                    return [4 /*yield*/, response.json()];
                case 3:
                    jsonResponse = _a.sent();
                    mimeTypes = {
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
                    returnBuffer = jsonResponse.Data.return_buffer;
                    returnBufferBytes = new Uint8Array(atob(returnBuffer).split('').map(function (char) { return char.charCodeAt(0); }));
                    fileName = jsonResponse.Data.file_name;
                    fileExtension = fileName.substring(fileName.lastIndexOf('.'));
                    status_1 = jsonResponse.Status;
                    fileType = mimeTypes[fileExtension];
                    blob = new Blob([returnBufferBytes], { type: fileType });
                    blobUrl = URL.createObjectURL(blob);
                    console.log(blob);
                    // Display the preview image
                    //await sleep(100);
                    chrome.tabs.create({ url: blobUrl });
                    downloadLink = document.createElement('a');
                    downloadLink.href = blobUrl;
                    downloadLink.download = fileName;
                    // Add the link to the document
                    document.body.appendChild(downloadLink);
                    // Click the link to trigger the download
                    downloadLink.click();
                    // Clean up the Object URL and remove the link from the document
                    URL.revokeObjectURL(blobUrl);
                    downloadLink.remove();
                    return [3 /*break*/, 5];
                case 4:
                    console.error('API request failed:', response.status, response.statusText);
                    _a.label = 5;
                case 5: return [3 /*break*/, 7];
                case 6:
                    error_1 = _a.sent();
                    console.error('Error:', error_1);
                    return [3 /*break*/, 7];
                case 7: return [2 /*return*/];
            }
        });
    }); });
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
    uploadPrivate.addEventListener('click', function () { return __awaiter(_this, void 0, void 0, function () {
        var file, reader;
        var _this = this;
        var _a;
        return __generator(this, function (_b) {
            file = (_a = fileInput.files) === null || _a === void 0 ? void 0 : _a[0];
            if (file) {
                reader = new FileReader();
                reader.onload = function (event) { return __awaiter(_this, void 0, void 0, function () {
                    var fileContents, uploadedData, response, buffer, mimeTypes, fileEx, fileType, blob, blobUrl, downloadLink, error_2;
                    var _a;
                    return __generator(this, function (_b) {
                        switch (_b.label) {
                            case 0:
                                fileContents = (_a = event.target) === null || _a === void 0 ? void 0 : _a.result;
                                uploadedData = JSON.parse(fileContents);
                                _b.label = 1;
                            case 1:
                                _b.trys.push([1, 6, , 7]);
                                return [4 /*yield*/, fetch('http://localhost:3000/decodePrivate', {
                                        method: 'POST',
                                        headers: {
                                            'Content-Type': 'application/json'
                                        },
                                        body: fileContents
                                    })];
                            case 2:
                                response = _b.sent();
                                if (!response.ok) return [3 /*break*/, 4];
                                return [4 /*yield*/, response.arrayBuffer()];
                            case 3:
                                buffer = _b.sent();
                                mimeTypes = {
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
                                fileEx = uploadedData.file_type;
                                fileType = mimeTypes[fileEx];
                                blob = new Blob([buffer], { type: fileType });
                                blobUrl = URL.createObjectURL(blob);
                                console.log(blob);
                                // Display the preview image
                                //await sleep(100);
                                chrome.tabs.create({ url: blobUrl });
                                downloadLink = document.createElement('a');
                                downloadLink.href = blobUrl;
                                downloadLink.download = uploadedData.file_name;
                                // Add the link to the document
                                document.body.appendChild(downloadLink);
                                // Click the link to trigger the download
                                downloadLink.click();
                                // Clean up the Object URL and remove the link from the document
                                URL.revokeObjectURL(blobUrl);
                                downloadLink.remove();
                                return [3 /*break*/, 5];
                            case 4:
                                console.error('API request failed:', response.status, response.statusText);
                                _b.label = 5;
                            case 5: return [3 /*break*/, 7];
                            case 6:
                                error_2 = _b.sent();
                                console.error('Error:', error_2);
                                return [3 /*break*/, 7];
                            case 7: return [2 /*return*/];
                        }
                    });
                }); };
                reader.readAsText(file);
            }
            return [2 /*return*/];
        });
    }); });
});
function sleep(ms) {
    return new Promise(function (resolve) {
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

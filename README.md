# SuperHack2023
Repository for Superhack2023

![Logo](https://github.com/cbayraktaroglu/SuperHack2023/blob/readme/banner.jpg)

# keepIt: Revolutionizing Blockchain File Storage

Introducing keepIt: a cutting-edge file storage protocol that empowers users to seamlessly upload and download files directly on the blockchain.

## How it Works

1. **Effortless Upload Process**: Divide your files into multiple parts as needed, tailoring the process to your preferences.

2. **Seamless Transmission**: Each file part is encapsulated within a transaction, automatically routed to the designated smart contract address: `0x00000000000000000000000000000000002023Ee` (default).

3. **User-Friendly Approval**: Every transaction requires explicit approval through your wallet for security and user control.

4. **Flexibility in Privacy**: After safely transmitting each file part, you have two options:

   a. **keepIt Private**: Maintain maximum privacy by encoding the transaction order into a JSON file, known only to you.
   
   b. **keepIt Public**: Opt for transparency by deploying a dedicated smart contract, publicly revealing the transaction order.

5. **User Verification**: Users can choose between two verification options:
   
   a. **WorldCoin Verification**: Verify yourself through the WorldCoin.
   
   b. **Attention ID**: Retrieve an Attention ID from us for verification through Ethereum Attestation Service (EAS).

6. **Access on Your Terms**: When it's time to access uploaded files:

   a. Retrieve your files using the keepIt Private JSON file, retaining exclusive control over your transaction order.
   
   b. Share the keepIt Public smart contract address and specified blockchain details to let others witness the transaction sequence.

## Deployed Contracts

The following contracts have been deployed for different networks:

- **Base Goerli**: `0xfbBc9950cB64912EDd88bCf47f6D85957C2aBcd0`
  - Status: Full ready for use.
 
- **Mode Sepholia**: `0xA4575B1d61AA4fE963373e8FD535427205B91135`
  - Status: Ready, but without Ethereum Attestation Service (EAS) since EAS is not supported on this chain.

- **Optimism Goerli**: `0xcf91A26C978c33fCe244412cBcb602C63F749A8b`
  - Status: Full ready for use.
    
- **Polygon Mumbai**: `0xce2dA00922faf10dd5bE5229666691eB28FcB75D`
  - Status: Ready, but without Ethereum Attestation Service (EAS) since EAS is not supported on this chain.

- **Zora Goerli**: `0xfbBc9950cB64912EDd88bCf47f6D85957C2aBcd0`
  - Status: Ready, but without any verification. Worldcoin and Ethereum Attestation Service (EAS) are not supported on this chain.


## Project Structure

The keepIt project is organized as follows:

- `frontend/`: The SvelteKit frontend for the application. Run with the command `npm run dev`.
- `backend/server/`: Contains the Golang backend service for decoding procedures needed by the Chrome extension.
   To run the backend service, navigate to the `backend/server/` directory and use the command `go run main.go`.
- `chrome-extension/`: This directory contains the Chrome extension for viewing and downloading files from keepIt. To use the extension, simply enable "Developer mode" in Chrome's Extensions settings and load the `chrome-extension/` directory as an unpacked extension.

## Embrace the Future with keepIt!

Elevate your file storage experience with the revolutionary keepIt protocol. Step into a new era of secure and transparent file management. Ready to get started? Let's embark on your keepIt journey today!

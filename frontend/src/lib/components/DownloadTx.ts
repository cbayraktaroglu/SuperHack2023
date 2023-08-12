import type { TxInfoFileJson } from "$lib/types/SmartContractBridgeData";

export function downloadTxInfoAsJsonFile(jsonData: TxInfoFileJson, fileName: string): void {
  const jsonContent = JSON.stringify(jsonData);
  const blob = new Blob([jsonContent], { type: 'application/json' });
  const url = URL.createObjectURL(blob);

  const link = document.createElement('a');
  link.href = url;
  link.download = fileName;
  link.click();

  // Clean up by revoking the object URL
  URL.revokeObjectURL(url);
}
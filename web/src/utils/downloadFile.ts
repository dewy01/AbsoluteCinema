export const downloadFileFromPath = async (url: string, filename: string) => {
  try {
    const response = await fetch(url);
    const blob = await response.blob();

    const downloadUrl = URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = downloadUrl;
    link.download = filename;
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    URL.revokeObjectURL(downloadUrl);
  } catch (error) {
    console.error('Failed to download file:', error);
    throw error;
  }
};
